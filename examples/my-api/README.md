# my-api — Contoh API minimal (GET /my-api)

API HTTP minimal: **GET /my-api** mengembalikan JSON (status, uptime, board). Tanpa dependency eksternal (hanya std), cocok untuk cross-compile ke **armv7-unknown-freebsd** dan deploy ke Orange Pi Zero LTS.

## Build dan jalankan

```bash
# Native (di PC atau di board)
cargo build --release
./target/release/my-api
```

```bash
# Cross-compile ke FreeBSD armv7 (di PC, butuh toolchain)
cargo build --release --target armv7-unknown-freebsd
scp target/armv7-unknown-freebsd/release/my-api user@opi:/usr/local/bin/
# Di board: my-api
```

## Penggunaan

- Listen: `0.0.0.0:8080`
- **GET /my-api** → `200 OK`, body JSON: `{"status":"ok","uptime_sec":N,"board":"Orange Pi Zero LTS","endpoint":"/my-api"}`
- Lainnya → `404 Not Found`

```bash
curl http://<board-ip>:8080/my-api
```

## Referensi

- [FreeBSD 15 - Orange Pi Zero LTS (contoh layanan my-api)](../../docs/FreeBSD-15-Orange-Pi-Zero-LTS.md#contoh-layanan-my-api)
- [Rancangan Agile/Scrum (Sprint 4)](../../docs/Agile-Scrum-Penelitian-FreeBSD-Orange-Pi-Zero-LTS.md)
