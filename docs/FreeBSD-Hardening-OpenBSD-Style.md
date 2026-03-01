# Pengamanan FreeBSD (setara semangat OpenBSD)

Dokumentasi ini berisi panduan mengamankan FreeBSD agar mendekati prinsip **secure by default** ala OpenBSD: firewall default-deny, layanan minimal, hardening kernel, SSH ketat, update rutin, dan isolasi (jails, Capsicum). Cocok untuk server maupun board (mis. Orange Pi Zero LTS).

**Referensi resmi:** [FreeBSD Handbook – Security](https://docs.freebsd.org/en/books/handbook/security/).

---

## Daftar isi

1. [Firewall (pf)](#1-firewall-pf)
2. [Layanan minimal](#2-layanan-minimal)
3. [SSH](#3-ssh)
4. [Kernel hardening (sysctl)](#4-kernel-hardening-sysctl)
5. [Securelevel (opsional)](#5-securelevel-opsional)
6. [Update dan audit](#6-update-dan-audit)
7. [Jails](#7-jails)
8. [Capsicum](#8-capsicum)
9. [Filesystem dan mount](#9-filesystem-dan-mount)
10. [TLS dan protokol](#10-tls-dan-protokol)
11. [Log dan monitoring](#11-log-dan-monitoring)
12. [Checklist ringkas](#12-checklist-ringkas)

---

## 1. Firewall (pf)

FreeBSD memakai **pf** (port dari OpenBSD). Kebijakan disarankan: **default-deny**, lalu buka hanya layanan yang dipakai.

### 1.1 Aktifkan pf

Tambahkan di `/etc/rc.conf`:

```bash
pf_enable="YES"
pf_flags=""
pf_rules="/etc/pf.conf"
pflog_enable="YES"
```

### 1.2 Contoh `/etc/pf.conf` minimal

Ganti `em0` dengan interface eksternal Anda (cek dengan `ifconfig`). Contoh ini mengizinkan hanya SSH (port 22) dari luar; lalu lintas keluar di-state.

```text
# Interface (sesuaikan nama)
ext_if = "em0"

# Kebijakan default
set block policy return
set skip on lo0

# Blok semua masuk; izinkan keluar
block in all
pass out on $ext_if inet keep state

# Izinkan SSH dari luar (sesuaikan jika pakai port lain)
pass in on $ext_if inet proto tcp to ($ext_if) port 22 keep state

# Untuk layanan lain (mis. HTTP 80, HTTPS 443), tambahkan baris serupa:
# pass in on $ext_if inet proto tcp to ($ext_if) port { 80 443 } keep state
```

### 1.3 Terapkan dan cek

```bash
service pf start
pfctl -s rules
pfctl -s info
```

---

## 2. Layanan minimal

Hanya nyalakan layanan yang benar-benar dipakai. Kurangi permukaan serangan dan penggunaan resource.

### 2.1 `/etc/rc.conf` — contoh kepala

```bash
# Jaringan
ifconfig_DEFAULT="DHCP"

# Hanya SSH jika perlu akses remote
sshd_enable="YES"

# Waktu (penting untuk TLS dan log)
ntpdate_enable="YES"
ntpd_enable="YES"

# Matikan sendmail jika tidak dipakai
sendmail_enable="NONE"
sendmail_submit_enable="NO"
sendmail_outbound_enable="NO"
sendmail_msp_queue_enable="NO"

# Layanan lain hanya jika dipakai (uncomment jika perlu)
# nfs_client_enable="YES"
# mountd_enable="YES"
```

### 2.2 Cek layanan yang listen

```bash
sockstat -l -4
netstat -an | grep LISTEN
```

Nonaktifkan atau uninstall layanan yang tidak dipakai.

---

## 3. SSH

Konfigurasi SSH yang ketat mengurangi risiko brute-force dan akses tidak sah.

### 3.1 `/etc/ssh/sshd_config` — rekomendasi

Edit dengan `ee` atau `vi`; sesuaikan sesuai kebutuhan.

```text
# Tidak login sebagai root lewat SSH
PermitRootLogin no

# Hanya autentikasi publik key (nonaktifkan password)
PasswordAuthentication no
PubkeyAuthentication yes
ChallengeResponseAuthentication no

# Batasi user yang boleh login (ganti youruser)
AllowUsers youruser

# Port standar 22; atau ganti ke port lain dan buka di pf
# Port 2222

# Batas percobaan login (opsional, tergantung versi sshd)
# MaxAuthTries 3
```

### 3.2 Pastikan user punya kunci sebelum menonaktifkan password

Di **klien** (bukan server):

```bash
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ""
ssh-copy-id -i ~/.ssh/id_ed25519.pub youruser@server
```

Tes login dengan kunci, lalu set `PasswordAuthentication no` dan restart `sshd`.

```bash
service sshd restart
```

---

## 4. Kernel hardening (sysctl)

Parameter berikut membatasi informasi dan perilaku yang bisa dimanfaatkan penyerang. Beberapa hanya bisa di-set saat boot (lihat [FreeBSD Handbook](https://docs.freebsd.org/en/books/handbook/security/)).

### 4.1 `/etc/sysctl.conf` — contoh

```text
# Random PID (sulit prediksi)
kern.randompid=1

# Batasi visibilitas proses/user lain
security.bsd.see_other_uids=0
security.bsd.see_other_gids=0
security.bsd.unprivileged_read_msgbuf=0
security.bsd.unprivileged_proc_debug=0

# Log koneksi TCP/UDP yang ditolak (debug firewall)
net.inet.tcp.log_in_vain=1
net.inet.udp.log_in_vain=1
```

Parameter yang **read-only** (hanya bisa di-set dari loader) harus masuk `/boot/loader.conf`, bukan `sysctl.conf`. Contoh (cek dulu apakah didukung di arsitektur Anda):

```text
# Di /boot/loader.conf (efek setelah reboot)
# security.bsd.see_other_uids=0
# security.bsd.see_other_gids=0
```

Terapkan sysctl yang bisa diubah saat jalan:

```bash
sysctl -f /etc/sysctl.conf
```

---

## 5. Securelevel (opsional)

**Securelevel** membatasi aksi tertentu setelah boot (mis. load kernel module, ubah waktu, tulis ke device). Dokumentasi FreeBSD menyarankan pemakaian hanya jika Anda paham dampaknya.

- **Level 1:** Batasan tertentu pada kernel dan device.
- **Level 2:** Lebih ketat; antara lain tidak bisa load module.

Lihat `man 7 securelevel` dan `man 8 init`. Untuk kebanyakan sistem, **sysctl + pf + layanan minimal** sudah cukup tanpa securelevel.

---

## 6. Update dan audit

### 6.1 Base system (freebsd-update)

```bash
freebsd-update fetch
freebsd-update install
# Jika ada kernel update, reboot setelah install
```

Jadwalkan secara rutin (cron) atau lakukan manual secara berkala.

### 6.2 Ports/paket (pkg audit)

```bash
pkg audit -F
```

Perbaiki vulnerability yang dilaporkan (upgrade paket atau nonaktifkan layanan). Contoh jadwal cron (root):

```text
0 3 * * 0 /usr/sbin/freebsd-update fetch install 2>/dev/null
0 4 * * 0 /usr/sbin/pkg audit -F 2>&1 | mail -s "pkg audit" admin@local
```

---

## 7. Jails

**Jails** mengisolasi layanan dalam lingkungan terpisah (filesystem, proses, sering juga jaringan). Cocok untuk menjalankan web server, database, atau API tanpa memberi akses penuh ke host.

### 7.1 Persiapan dasar

- Buat filesystem untuk jail (atau dataset ZFS).
- Install world ke direktori jail atau pakai `pkg -j` untuk mengisi jail.
- Konfigurasi jail di `/etc/jail.conf` atau `/etc/rc.conf` (jail_*).

### 7.2 Contoh minimal `/etc/jail.conf`

```text
myjail {
    host.hostname = "myjail.local";
    path = "/usr/jails/myjail";
    ip4.addr = "192.168.1.100";
    interface = "em0";
    exec.start = "/bin/sh /etc/rc";
    exec.stop = "/bin/sh /etc/rc.shutdown jail";
}
```

Aktifkan:

```bash
service jail start myjail
jls
```

Panduan lengkap: [FreeBSD Handbook – Jails](https://docs.freebsd.org/en/books/handbook/jails/).

---

## 8. Capsicum

**Capsicum** adalah mekanisme capability-based security: proses dapat dibatasi hanya mengakses file, socket, atau operasi yang diizinkan. Berguna untuk layanan yang Anda develop atau paket yang sudah mendukung Capsicum.

- Program harus secara eksplisit memasuki capability mode (`cap_enter()` atau wrapper).
- Setelah masuk, kemampuan proses dibatasi oleh file descriptor yang diizinkan.

Untuk layanan pihak ketiga, cek dokumentasi atau opsi konfigurasi yang mengaktifkan Capsicum. Untuk program sendiri, gunakan API Capsicum (lihat `man 2 cap_enter`, `man 3 libcapsicum`).

---

## 9. Filesystem dan mount

### 9.1 Opsi noexec

Gunakan **noexec** pada mount point yang tidak perlu menjalankan binary (mis. `/tmp`, `/var/tmp`), untuk mengurangi risiko eksekusi dari direktori yang world-writable.

Di `/etc/fstab`, tambahkan opsi `noexec` pada partisi yang sesuai. Contoh (sesuaikan device):

```text
/dev/ada0p3  /tmp  ufs  rw,noexec,nosuid  0  0
```

Perhatikan: beberapa program mengharapkan bisa menjalankan binary dari `/tmp`; uji setelah perubahan.

### 9.2 Permission

- Hindari file atau direktori world-writable kecuali yang memang diperlukan (mis. `/tmp` dengan sticky bit).
- Batasi akses ke file konfigurasi (mis. `chmod 600` untuk file yang berisi rahasia).

---

## 10. TLS dan protokol

- Layanan yang terbuka ke jaringan (web, API, mail) sebaiknya memakai **TLS** (HTTPS, SMTP dengan STARTTLS, dll.).
- Nonaktifkan protokol dan cipher lama (SSLv3, TLS 1.0 jika tidak diperlukan) di konfigurasi layanan (nginx, Apache, Postfix, dll.).
- Gunakan sertifikat yang valid (mis. Let's Encrypt) dan perbarui sebelum kedaluwarsa.

---

## 11. Log dan monitoring

### 11.1 Syslog

Pastikan syslog mencatat auth, daemon, dan kernel. Konfigurasi di `/etc/syslog.conf` (atau konfigurasi modern `syslogd`/`rc.conf`). Rotasi log dengan **newsyslog** (biasanya sudah dijadwalkan).

### 11.2 Log SSH dan auth

- Failed login dan event auth tercatat (biasanya di `/var/log/auth.log` atau output dari `sshd`).
- Pantau secara berkala: `grep Failed /var/log/auth.log` atau gunakan tool analisis log.

### 11.3 Firewall log

Jika di `pf.conf` Anda memakai `log` pada rule tertentu:

```text
block in log all
```

Log pf biasanya ditangani oleh `pflog`; pastikan `pflog_enable="YES"` dan periksa di mana log pf ditulis (sering ke `pflog0` atau file yang dikonfigurasi).

---

## 12. Checklist ringkas

| Aspek | Tindakan |
|-------|----------|
| Firewall | pf default-deny; hanya buka port yang dipakai (SSH, HTTP/HTTPS, dll.). |
| Layanan | Hanya yang diperlukan; matikan sendmail dan layanan tidak dipakai. |
| SSH | `PermitRootLogin no`; autentikasi key-only; `AllowUsers` jika memungkinkan. |
| Kernel | sysctl hardening (see_other_uids/gids, unprivileged_*, log_in_vain). |
| Securelevel | Opsional; hanya jika paham dampaknya. |
| Update | Rutin: `freebsd-update`, `pkg audit -F`. |
| Isolasi | Jails untuk layanan; Capsicum untuk program yang mendukung. |
| Filesystem | noexec di tempat yang tepat; permission ketat. |
| Protokol | TLS untuk layanan jaringan; nonaktifkan protokol lama. |
| Log | Syslog dan auth log aktif; pantau failed login dan akses mencurigakan. |

---

## Referensi

- [FreeBSD Handbook – Security](https://docs.freebsd.org/en/books/handbook/security/)
- [FreeBSD Handbook – Firewalls (pf)](https://docs.freebsd.org/en/books/handbook/firewalls/)
- [FreeBSD Handbook – Jails](https://docs.freebsd.org/en/books/handbook/jails/)
- [FreeBSD Documentation Portal](https://docs.freebsd.org/en/)
- `man 7 securelevel` — securelevel
- `man 5 pf.conf` — sintaks pf
- `man 2 cap_enter` — Capsicum

---

Dokumentasi ini melengkapi [FreeBSD-15-Orange-Pi-Zero-LTS.md](FreeBSD-15-Orange-Pi-Zero-LTS.md). Setelah instalasi dan optimalisasi board, terapkan langkah di atas sesuai kebutuhan (firewall, SSH, dan update minimal untuk production).
