# SysAdmin Personal Branding Website - Setup Guide

Panduan lengkap command-line untuk membuat website personal branding SysAdmin dengan stack:
- **Backend**: Golang + MariaDB
- **Frontend**: Vue.js + TailwindCSS  
- **Deployment**: Podman (Alpine-based containers)

---

## 📋 Daftar Isi

1. [Persiapan Environment](#1-persiapan-environment)
2. [Inisialisasi Backend (Golang)](#2-inisialisasi-backend-golang)
3. [Inisialisasi Frontend (Vue.js)](#3-inisialisasi-frontend-vuejs)
4. [Setup Database (MariaDB)](#4-setup-database-mariadb)
5. [Konfigurasi Podman & Container](#5-konfigurasi-podman--container)
6. [Deployment](#6-deployment)

---

## 1. Persiapan Environment

### 1.1 Verifikasi Tools

```powershell
# Cek versi Golang (minimal 1.21)
go version

# Cek versi Node.js (minimal 18.x)
node --version
npm --version

# Cek versi Podman
podman --version

# Cek versi MariaDB client (opsional untuk testing)
mysql --version
```

### 1.2 Buat Struktur Proyek

```powershell
# Pindah ke workspace
cd c:\laragon\www\personal

# Buat folder proyek utama
mkdir sysadmin-portfolio
cd sysadmin-portfolio

# Buat struktur folder
mkdir backend, frontend, deployment, docs
```

---

## 2. Inisialisasi Backend (Golang)

### 2.1 Setup Go Module & Struktur Folder

```powershell
# Masuk ke direktori backend
cd backend

# Inisialisasi Go module (ganti username dengan GitHub username Anda)
go mod init github.com/yourusername/sysadmin-portfolio-backend

# Buat struktur folder sesuai best practice
mkdir cmd\api, internal\handlers, internal\models, internal\database, internal\middleware, pkg\utils, configs, migrations
```

**Struktur folder backend:**
```
backend/
├── cmd/api/main.go              # Entry point
├── internal/
│   ├── handlers/                # HTTP handlers
│   ├── models/                  # Data models
│   ├── database/                # DB connection
│   └── middleware/              # Middleware
├── pkg/utils/                   # Utilities
├── configs/config.yaml          # Config
├── migrations/                  # SQL migrations
├── go.mod
└── .env.example
```

### 2.2 Install Dependencies

```powershell
# Web framework - Fiber (minimalis & cepat)
go get -u github.com/gofiber/fiber/v2

# Database driver MariaDB/MySQL
go get -u github.com/go-sql-driver/mysql

# Environment variables
go get -u github.com/joho/godotenv

# CORS middleware
go get -u github.com/gofiber/fiber/v2/middleware/cors

# JWT authentication
go get -u github.com/golang-jwt/jwt/v5

# Database migration tool
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Air untuk hot reload (development)
go install github.com/air-verse/air@latest
```

### 2.3 Buat File Konfigurasi

```powershell
# Buat .env.example
@"
# Server Configuration
PORT=8080
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=sysadmin
DB_PASSWORD=your_password
DB_NAME=sysadmin_portfolio

# JWT Configuration
JWT_SECRET=your-super-secret-key-change-this
JWT_EXPIRE=24h

# CORS
ALLOWED_ORIGINS=http://localhost:5173
"@ | Out-File -FilePath .env.example -Encoding UTF8

# Copy ke .env untuk development
Copy-Item .env.example .env
```

---

## 3. Inisialisasi Frontend (Vue.js)

### 3.1 Setup Vue.js dengan Vite

```powershell
# Kembali ke root proyek
cd ..

# Buat project Vue.js dengan Vite
npm create vite@latest frontend -- --template vue-ts

# Masuk ke direktori frontend
cd frontend

# Install dependencies
npm install
```

### 3.2 Install Dependencies Frontend

```powershell
# Vue Router
npm install vue-router@4

# Pinia (state management)
npm install pinia

# Axios (HTTP client)
npm install axios

# TailwindCSS (styling minimalis)
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p

# VueUse (utility functions)
npm install @vueuse/core

# Heroicons (icon library minimalis)
npm install @heroicons/vue

# Form validation
npm install vee-validate yup

# Date formatting
npm install date-fns
```

### 3.3 Konfigurasi TailwindCSS

```powershell
# Edit tailwind.config.js
@"
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        primary: '#0f172a',
        secondary: '#1e293b',
        accent: '#3b82f6',
      },
    },
  },
  plugins: [],
}
"@ | Out-File -FilePath tailwind.config.js -Encoding UTF8

# Buat file CSS utama
@"
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  body {
    @apply bg-gray-50 text-gray-900;
  }
}
"@ | Out-File -FilePath src\assets\main.css -Encoding UTF8
```

### 3.4 Buat Struktur Folder Frontend

```powershell
# Buat folder struktur
mkdir src\components\common, src\components\sections, src\views, src\router, src\stores, src\services, src\types, src\utils
```

**Struktur folder frontend:**
```
frontend/
├── src/
│   ├── components/
│   │   ├── common/              # Reusable components
│   │   └── sections/            # Page sections
│   ├── views/                   # Pages
│   ├── router/                  # Vue Router
│   ├── stores/                  # Pinia stores
│   ├── services/                # API services
│   ├── types/                   # TypeScript types
│   └── utils/                   # Utilities
├── vite.config.ts
└── tailwind.config.js
```

### 3.5 Konfigurasi Environment

```powershell
# Buat .env.development
@"
VITE_API_BASE_URL=http://localhost:8080/api
VITE_APP_TITLE=SysAdmin Portfolio
"@ | Out-File -FilePath .env.development -Encoding UTF8

# Buat .env.production
@"
VITE_API_BASE_URL=/api
VITE_APP_TITLE=SysAdmin Portfolio
"@ | Out-File -FilePath .env.production -Encoding UTF8
```

---

## 4. Setup Database (MariaDB)

### 4.1 Buat Migration Files

```powershell
# Kembali ke backend
cd ..\backend

# Buat migration untuk tabel users
@"
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    bio TEXT,
    avatar_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
"@ | Out-File -FilePath migrations\001_create_users.up.sql -Encoding UTF8

@"
DROP TABLE IF EXISTS users;
"@ | Out-File -FilePath migrations\001_create_users.down.sql -Encoding UTF8

# Buat migration untuk tabel projects
@"
CREATE TABLE IF NOT EXISTS projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    technologies JSON,
    github_url VARCHAR(255),
    demo_url VARCHAR(255),
    image_url VARCHAR(255),
    is_featured BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_featured (is_featured)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
"@ | Out-File -FilePath migrations\002_create_projects.up.sql -Encoding UTF8

@"
DROP TABLE IF EXISTS projects;
"@ | Out-File -FilePath migrations\002_create_projects.down.sql -Encoding UTF8
```

### 4.2 Buat Init Script Database

```powershell
# Kembali ke root dan buat folder init
cd ..
mkdir deployment\mariadb\init

# Buat init script
@"
-- Create database
CREATE DATABASE IF NOT EXISTS sysadmin_portfolio CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create user
CREATE USER IF NOT EXISTS 'sysadmin'@'%' IDENTIFIED BY 'sysadmin_password';

-- Grant privileges
GRANT ALL PRIVILEGES ON sysadmin_portfolio.* TO 'sysadmin'@'%';

-- Flush privileges
FLUSH PRIVILEGES;

-- Use database
USE sysadmin_portfolio;
"@ | Out-File -FilePath deployment\mariadb\init\01-init.sql -Encoding UTF8
```

---

## 5. Konfigurasi Podman & Container

### 5.1 Buat Containerfile Backend

```powershell
cd deployment

# Buat Containerfile untuk backend
@"
# Build stage
FROM docker.io/library/golang:1.21-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Final stage
FROM docker.io/library/alpine:latest

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -g 1000 appuser && adduser -D -u 1000 -G appuser appuser

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ['./main']
"@ | Out-File -FilePath Containerfile.backend -Encoding UTF8
```

### 5.2 Buat Containerfile Frontend

```powershell
# Buat Containerfile untuk frontend
@"
# Build stage
FROM docker.io/library/node:18-alpine AS builder

WORKDIR /app

COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build

# Production stage
FROM docker.io/library/nginx:alpine

COPY deployment/nginx/nginx.conf /etc/nginx/nginx.conf
COPY deployment/nginx/default.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist /usr/share/nginx/html

RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /usr/share/nginx/html && \
    chown -R appuser:appuser /var/cache/nginx && \
    chown -R appuser:appuser /var/log/nginx && \
    touch /var/run/nginx.pid && \
    chown -R appuser:appuser /var/run/nginx.pid

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080 || exit 1

CMD ['nginx', '-g', 'daemon off;']
"@ | Out-File -FilePath Containerfile.frontend -Encoding UTF8
```

### 5.3 Buat Konfigurasi Nginx

```powershell
mkdir nginx

# nginx.conf
@"
user appuser;
worker_processes auto;
error_log /var/log/nginx/error.log warn;
pid /var/run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include /etc/nginx/mime.types;
    default_type application/octet-stream;

    log_format main '\$remote_addr - \$remote_user [\$time_local] "\$request" '
                    '\$status \$body_bytes_sent "\$http_referer" '
                    '"\$http_user_agent" "\$http_x_forwarded_for"';

    access_log /var/log/nginx/access.log main;

    sendfile on;
    tcp_nopush on;
    keepalive_timeout 65;
    gzip on;

    include /etc/nginx/conf.d/*.conf;
}
"@ | Out-File -FilePath nginx\nginx.conf -Encoding UTF8

# default.conf
@"
server {
    listen 8080;
    server_name localhost;
    root /usr/share/nginx/html;
    index index.html;

    add_header X-Frame-Options 'SAMEORIGIN' always;
    add_header X-Content-Type-Options 'nosniff' always;
    add_header X-XSS-Protection '1; mode=block' always;

    location / {
        try_files \$uri \$uri/ /index.html;
    }

    location /api {
        proxy_pass http://backend:8080;
        proxy_http_version 1.1;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    }

    location ~* \.(jpg|jpeg|png|gif|ico|css|js|svg|woff|woff2)$ {
        expires 1y;
        add_header Cache-Control 'public, immutable';
    }
}
"@ | Out-File -FilePath nginx\default.conf -Encoding UTF8
```

### 5.4 Buat Podman Compose File

```powershell
cd ..

# Buat podman-compose.yml
@"
version: '3.8'

services:
  database:
    image: docker.io/library/mariadb:11-alpine
    container_name: sysadmin-db
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: root_password
      MYSQL_DATABASE: sysadmin_portfolio
      MYSQL_USER: sysadmin
      MYSQL_PASSWORD: sysadmin_password
      TZ: Asia/Jakarta
    volumes:
      - mariadb_data:/var/lib/mysql
      - ./deployment/mariadb/init:/docker-entrypoint-initdb.d:ro
    ports:
      - '3306:3306'
    networks:
      - sysadmin-network
    healthcheck:
      test: ['CMD', 'healthcheck.sh', '--connect', '--innodb_initialized']
      interval: 10s
      timeout: 5s
      retries: 5
      start_period: 30s

  backend:
    build:
      context: .
      dockerfile: deployment/Containerfile.backend
    container_name: sysadmin-backend
    restart: unless-stopped
    environment:
      PORT: 8080
      ENV: production
      DB_HOST: database
      DB_PORT: 3306
      DB_USER: sysadmin
      DB_PASSWORD: sysadmin_password
      DB_NAME: sysadmin_portfolio
      JWT_SECRET: your-super-secret-key-change-this
      JWT_EXPIRE: 24h
      ALLOWED_ORIGINS: http://localhost:3000
    ports:
      - '8080:8080'
    depends_on:
      database:
        condition: service_healthy
    networks:
      - sysadmin-network

  frontend:
    build:
      context: .
      dockerfile: deployment/Containerfile.frontend
    container_name: sysadmin-frontend
    restart: unless-stopped
    ports:
      - '3000:8080'
    depends_on:
      - backend
    networks:
      - sysadmin-network

volumes:
  mariadb_data:
    driver: local

networks:
  sysadmin-network:
    driver: bridge
"@ | Out-File -FilePath podman-compose.yml -Encoding UTF8
```

### 5.5 Buat Makefile Helper

```powershell
# Buat Makefile (untuk PowerShell, kita buat script .ps1)
@"
# Helper commands untuk Podman

function Build {
    podman-compose build
}

function Up {
    podman-compose up -d
}

function Down {
    podman-compose down
}

function Logs {
    podman-compose logs -f
}

function Clean {
    podman-compose down -v
    podman system prune -f
}

function Restart {
    podman-compose restart
}

function Backend-Shell {
    podman exec -it sysadmin-backend sh
}

function Frontend-Shell {
    podman exec -it sysadmin-frontend sh
}

function DB-Shell {
    podman exec -it sysadmin-db mysql -u sysadmin -psysadmin_password sysadmin_portfolio
}

function Migrate-Up {
    cd backend
    migrate -path migrations -database 'mysql://sysadmin:sysadmin_password@tcp(localhost:3306)/sysadmin_portfolio' up
    cd ..
}

function Dev-Backend {
    cd backend
    air
}

function Dev-Frontend {
    cd frontend
    npm run dev
}

# Export functions
Export-ModuleMember -Function *
"@ | Out-File -FilePath podman-helpers.ps1 -Encoding UTF8
```

---

## 6. Deployment

### 6.1 Build & Run dengan Podman

```powershell
# Build semua container
podman-compose build

# Start semua services
podman-compose up -d

# Cek status
podman ps

# Lihat logs
podman-compose logs -f

# Stop services
podman-compose down
```

### 6.2 Development Mode (Tanpa Container)

```powershell
# Terminal 1 - Database
podman run -d `
  --name sysadmin-db-dev `
  -e MYSQL_ROOT_PASSWORD=root `
  -e MYSQL_DATABASE=sysadmin_portfolio `
  -e MYSQL_USER=sysadmin `
  -e MYSQL_PASSWORD=sysadmin_password `
  -p 3306:3306 `
  docker.io/library/mariadb:11-alpine

# Terminal 2 - Backend
cd backend
air

# Terminal 3 - Frontend
cd frontend
npm run dev
```

### 6.3 Akses Aplikasi

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Database**: localhost:3306

---

## 📝 Checklist Deployment

### Pre-Production
- [ ] Ganti semua password di `.env` dan `podman-compose.yml`
- [ ] Ganti `JWT_SECRET` dengan nilai random yang aman
- [ ] Update `ALLOWED_ORIGINS` dengan domain production
- [ ] Test semua endpoint API
- [ ] Test responsive design (mobile & desktop)

### Production
- [ ] Setup HTTPS/SSL
- [ ] Setup domain name
- [ ] Configure firewall
- [ ] Setup backup database
- [ ] Enable monitoring & logging
- [ ] Setup CI/CD pipeline

---

## 🔗 Referensi

- [Golang Docs](https://go.dev/doc/)
- [Fiber Framework](https://docs.gofiber.io/)
- [Vue.js 3](https://vuejs.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [Podman Docs](https://docs.podman.io/)
- [MariaDB Docs](https://mariadb.org/documentation/)

---

**Dibuat**: 2026-02-17  
**Versi**: 1.0.0
