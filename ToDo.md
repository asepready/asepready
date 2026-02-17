# TODO List
## Phase 1: Infrastructure & Planning (Fondasi)
Ini adalah tahap "Design Document" sebelum coding.
- Pilih Domain Name
  1. Usahakan namaanda.com, namaanda.net, atau namaanda.id.
  2. Hindari domain gratisan (.tk, .ml) agar terlihat profesional.
- Pilih Tech Stack (Static Site)
  1. Rekomendasi: Hugo, Astro, atau Jekyll.
  2. Alasan: Cepat, aman (tidak ada database untuk di-hack), mudah di-version control.
- Pilih Hosting
  1. Rekomendasi: GitHub Pages, Netlify, atau Cloudflare Pages.
  2. Alasan: Gratis tier cukup mumpuni, SSL otomatis, CDN global.
- Setup Version Control
  1. Buat repository private/public di GitHub/GitLab.
  2. Setup .gitignore yang proper.

## Phase 2: Core Features (Frontend & UX)
Fitur wajib agar pengunjung paham siapa Anda dalam 5 detik.
- Hero Section yang Jelas
  1. Headline: Role spesifik (misal: "Linux System Administrator").
  2. Sub-headline: Value proposition (misal: "Automating Infrastructure & Ensuring Security").
  3. Foto Profil: Profesional (bisa casual tapi rapi, atau avatar teknis).
- Navigasi Simpel
  1. Menu: Home, About, Skills, Projects, Blog, Contact.
  2. Sticky header atau mudah diakses.
- Responsive Design
  1. Wajib rapi di Mobile (Admin sering cek link dari HP).
  2. Wajib rapi di Desktop.
- Dark Mode Toggle
  1. Fitur wajib untuk audiens teknis (mengurangi eye strain).
- Fast Loading Performance
  1. Target Google Lighthouse Score > 90.
  2. Minimalkan JavaScript eksternal.
  3. Gunakan format gambar modern (WebP/AVIF).

## Phase 3: Content & Portfolio (Bukti Kompetensi)
Bagian ini menjual skill Anda.
- Halaman Skills (Tech Stack)
  1. Kategorikan skill: OS, Network, Cloud, Automation, Monitoring.
  2. Jangan pakai progress bar (misal: "90% Linux" itu tidak relevan), cukup list atau level (Fundamental, Advanced, Expert).
- Halaman Projects (Case Studies)
  1. Minimal 3 studi kasus nyata.
  2. Format: Masalah -> Solusi -> Tools -> Hasil (Quantifiable, misal: "Uptime naik ke 99.9%").
  3. Fitur Khusus: Sertakan Diagram Topologi (buat di Draw.io/Excalidraw) untuk proyek jaringan.
- Halaman Blog / Write-ups
  1. Artikel teknis (Tutorial, Troubleshooting, Post-Mortem).
  2. Fitur Code Syntax Highlighting (warna-warni untuk script).
- Downloadable CV
  1. Tombol download PDF (Update terbaru).
  2. Pastikan PDF bersih dan tidak perlu password.
- Certifications Badge
  1. Logo sertifikat (RHCE, CCNA, AWS, dll) dengan link validasi jika ada.

## Phase 4: Security & Privacy (Kredibilitas Admin)
Sebagai admin keamanan/jaringan, website Anda harus aman.
- HTTPS/SSL Enforcement
  1. Wajib HTTPS (HSTS enabled).
- Security Headers
  1. Implementasi CSP (Content Security Policy).
  2. Implementasi X-Frame-Options, X-Content-Type-Options.
- Kontak Aman
  1. Jangan pakai Contact Form PHP (rawan spam & exploit).
  2. Gunakan mailto: link biasa ATAU layanan form statis (Formspree/Netlify Forms).
  3. Fitur Pro: Tampilkan PGP/GPG Public Key untuk enkripsi email.
- Privacy Focused Analytics
  1. Jangan pakai Google Analytics (berat & privacy issue).
  2. Gunakan Plausible, Umami, atau Matomo (self-hosted).
- Hide Server Info
  1. Pastikan header HTTP tidak membocorkan versi server (misal: X-Powered-By harus di-disable).

## Phase 5: "Dogfooding" & Monitoring (Nilai Plus)
Menunjukkan Anda mempraktikkan apa yang Anda preach.
- Uptime Status Page
  1. Buat halaman kecil /status yang menampilkan uptime website Anda sendiri (menggunakan Uptime Kuma atau sejenisnya).
  2. Ini menunjukkan Anda peduli pada monitoring.
- CI/CD Pipeline
  1. Setup GitHub Actions untuk auto-deploy saat ada push ke branch main.
  2. Setup auto-build check sebelum merge.
- Backup Strategy
  1. Pastikan repository memiliki backup remote.
- DNS Configuration
  1. Setup SPF, DKIM, DMARC untuk domain email Anda (menunjukkan pemahaman email security).

## Phase 6: Easter Eggs (Sentuhan Personal)
Fitur opsional untuk menunjukkan kreativitas & humor teknis.
- CLI Easter Egg
  1. Jika user klik kanan -> Inspect Element, berikan pesan rahasia di Console.
  2. Atau buat halaman /ssh yang mensimulasikan terminal browser (menggunakan library seperti xterm.js).
- Robots.txt & Sitemap.xml
  1. Pastikan SEO dasar terpenuhi agar mudah ditemukan recruiter.
- 404 Page Custom
  1. Buat halaman 404 yang unik (misal: "Error 404: Route Not Found" dengan gaya pesan error Cisco/Linux).

## Prioritas Eksekusi (Saran Saya)
- Minggu 1: Setup Domain, GitHub Repo, & Pilih Theme Hugo/Astro. (Phase 1 & 2)
- Minggu 2: Isi Konten (Skills, Projects, CV). Jangan perfeksionis, yang penting ada isinya. (Phase 3)
- Minggu 3: Hardening Security (Headers, SSL, GPG) & Setup Analytics. (Phase 4)
- Minggu 4: Polish, Testing Mobile, & Launch. (Phase 5 & 6) 