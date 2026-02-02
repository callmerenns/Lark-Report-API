# Lark Webhook API

> **Private Repository – Internal / Client Project**

Lark Webhook API adalah backend service berbasis **Golang (Gin)** yang dirancang untuk kebutuhan **production-grade webhook processing**, autentikasi JWT, rate limiting, dan integrasi dengan Redis serta Supabase.

Aplikasi ini **tidak ditujukan untuk konsumsi publik** dan hanya digunakan oleh sistem internal / klien yang telah terotorisasi.

---

## 🔒 Repository Status

⚠️ **PRIVATE REPOSITORY**

* Akses terbatas
* Tidak untuk distribusi publik
* Dokumentasi ditujukan untuk **developer internal & client technical team**

---

## 📌 Badges

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go\&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker\&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-Enabled-DC382D?logo=redis\&logoColor=white)
![Supabase](https://img.shields.io/badge/Supabase-Postgres-3ECF8E?logo=supabase\&logoColor=white)
![JWT](https://img.shields.io/badge/Auth-JWT-orange)
![License](https://img.shields.io/badge/License-Proprietary-red)

---

## 🧱 Tech Stack

| Layer              | Technology              |
| ------------------ | ----------------------- |
| Language           | Golang                  |
| Web Framework      | Gin                     |
| Auth               | JWT (HS256, 1h expiry)  |
| Cache / Rate Limit | Redis                   |
| Database           | Supabase (PostgreSQL)   |
| API Docs           | Swagger (swaggo)        |
| Container          | Docker & Docker Compose |

---

## 🏗️ Architecture Diagram

```
┌──────────────┐
│ Client / Bot │
└──────┬───────┘
       │ HTTP (JSON)
       ▼
┌──────────────────────┐
│   Gin API Server     │
│──────────────────────│
│ - JWT Middleware     │
│ - Webhook Secret     │
│ - Rate Limiter       │
│ - Validation Layer   │
└──────┬───────────────┘
       │
       ├── Redis (Rate Limit, Token State)
       │
       └── Supabase (PostgreSQL)

```

---

## 🔐 Authentication & Security

### JWT Token

* **Expiry:** 1 jam
* **Single Active Token**

  * Token lama otomatis invalid saat token baru dibuat
* **Header:**

  ```http
  Authorization: Bearer <token>
  ```

### Webhook Secret

Digunakan untuk endpoint tertentu (misalnya generate token):

```http
X-Webhook-Secret: <static-secret>
```

### Rate Limiting

* Redis-based
* Per IP / Identifier
* Response otomatis:

  * HTTP `429 Too Many Requests`
  * Header `Retry-After`

---

## 📦 Environment Configuration

Gunakan file `.env` berdasarkan template berikut:

```env
PORT=8080
APP_ENV=dev

DATABASE_URL=

JWT_SECRET=
WEBHOOK_SECRET=

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

⚠️ **Jangan commit `.env` ke repository**

---

## ▶️ Running the Application

### Local Development

```bash
go mod tidy
go run main.go
```

Akses:

* API: `http://localhost:8080`
* Swagger: `http://localhost:8080/swagger/index.html`

---

### Docker Compose (Recommended)

```bash
docker compose up --build
```

Service yang dijalankan:

* `api` (Gin App)
* `redis`

---

## 📄 API Documentation (Swagger)

Generate swagger docs:

```bash
swag init
```

Endpoint Swagger:

```
GET /swagger/index.html
```

Dokumentasi sudah mencakup:

* Success response
* Error response (400, 401, 403, 429, 500)
* Example payload

---

## ✅ Health Check

```http
GET /health
```

Response:

```json
{
  "success": true,
  "message": "Welcome to API Lark",
  "data": {
    "name": "Lark Webhook API",
    "version": "v1",
    "status": "running"
  },
  "error": null
}
```

---

## 🚨 Error Response Standard

Semua error mengikuti format berikut:

```json
{
  "success": false,
  "message": "Unauthorized",
  "error": {
    "code": "INVALID_JWT",
    "details": "Invalid or expired token"
  }
}
```

---

## 🧪 Production Notes

* Gunakan **HTTPS** di production
* Simpan secret di **Vault / Secret Manager**
* Redis **wajib aktif** untuk rate limit
* JWT Secret **harus strong random**
* Enable logging & monitoring

---

## 👥 Intended Audience

* Internal Backend Engineers
* Client Technical Team
* DevOps / Infra Team

---

## 📜 License

© 2026 Adamata

**Proprietary Software** — All rights reserved.
Unauthorized copying or distribution is prohibited.
