# Superbank Assessment

This is a backend API for the Superbank Assessment project. It provides features for managing users, accounts, pockets, and transactions. It uses **Go**, **PostgreSQL**, **JWT Authentication**, and **Docker** for containerization.

## Instalasi

### Clone Repository

```bash
git clone https://github.com/username-anda/superbank_assessment.git
cd superbank_assessment/backend
```

### Uncomment example.env menjadi .env

## Menjalankan Aplikasi

### Dengan Docker

```bash
cd /path/to/superbank_assessment/backend
docker-compose up --build
```

### Tanpa Docker

```bash
go mod tidy
cd /path/to/superbank_assessment/backend
go run main.go
```

## Dokumentasi API

### Login: POST /login

Request

```json
{
  "username": "user1",
  "password": "kata_sandi_anda"
}
```

Response

```json
{
  "token": "jwt_token_anda"
}
```

Gunakan token untuk mengakses route yang dilindungi dengan menyetelnya di header Authorization sebagai token Bearer.

### Endpoint User

**Buat Pengguna**: `POST /users`
**Dapatkan Pengguna**: `GET /users/{id}`
**Perbarui Pengguna**: `PUT /users/{id}`
**Hapus Pengguna**: `DELETE /users/{id}`

### Endpoint Pocket

**Buat Pocket**: `POST /pockets`
**Dapatkan Pocket**: `GET /pockets/{id}`
**Perbarui Pocket**: `PUT /pockets/{id}`
**Hapus Pocket**: `DELETE /pockets/{id}`
