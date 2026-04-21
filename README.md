# Port Checker

Tool sederhana berbasis Go untuk mengecek status port TCP pada suatu host.

## Fungsi

Melakukan scan terhadap port-port umum berikut:

| Port | Layanan    |
|------|------------|
| 22   | SSH        |
| 80   | HTTP       |
| 443  | HTTPS      |
| 3306 | MySQL      |
| 5432 | PostgreSQL |
| 6379 | Redis      |

Setiap port akan ditampilkan statusnya: `OPEN` jika dapat dikoneksi, atau `CLOSED` jika tidak.

## Cara Penggunaan

### Build

```bash
go build -o port-checker
```

### Jalankan

Scan `localhost` (default):
```bash
./port-checker
```

Scan host tertentu (IP atau domain):
```bash
./port-checker 192.168.1.1
./port-checker example.com
```

Scan dengan IPv6:
```bash
./port-checker ::1
```

### Contoh Output

```
Scanning 192.168.1.1...

Port 22     CLOSED
Port 80     OPEN  
Port 443    OPEN  
Port 3306   CLOSED
Port 5432   CLOSED
Port 6379   CLOSED

2 port terbuka dari 6 yang dicek
```

## Requirement

- Go 1.18 atau lebih baru
