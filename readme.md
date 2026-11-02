# Goroutine Go

Project latihan belajar **goroutine dan channel di Golang**. Setiap contoh ditulis sebagai fungsi test, sehingga bisa langsung dijalankan dengan `go test`.

## Struktur Project

```
goroutine-go/
├── go.mod
├── goroutine_test.go   # Contoh goroutine
└── channel_test.go     # Contoh channel
```

Project ini tidak memakai dependensi eksternal, hanya library bawaan Go (`fmt`, `testing`, `time`).

## Prasyarat

- [Go](https://go.dev/dl/) versi 1.21.5 atau lebih baru

## Menjalankan

Masuk ke folder project:

```bash
cd goroutine-go
```

Jalankan semua test dengan output rinci. Opsi `-v` penting supaya teks dari `fmt.Println` terlihat:

```bash
go test -v
```

Jalankan satu test tertentu:

```bash
go test -v -run TestGoroutine
go test -v -run TestManyGoroutine
go test -v -run TestChannel
```

Setiap test memakai `time.Sleep` untuk menunggu goroutine selesai, jadi total waktu berjalan sekitar 13 detik.

## Isi Materi

### Goroutine (`goroutine_test.go`)

| Test | Penjelasan |
| --- | --- |
| `TestGoroutine` | Menjalankan fungsi `RunGoroutine()` dengan kata kunci `go`, sehingga berjalan bersamaan dengan fungsi utama. Urutan output tidak dijamin. |
| `TestManyGoroutine` | Membuat 100.000 goroutine yang masing-masing memanggil `DisplayNumber(i)`. Menunjukkan bahwa goroutine sangat ringan dibanding thread biasa. |

Kedua test memakai `time.Sleep` agar program menunggu goroutine selesai sebelum test berakhir. Tanpa jeda ini, test bisa selesai lebih dulu dan goroutine belum sempat mencetak apa pun.

### Channel (`channel_test.go`)

| Test | Penjelasan |
| --- | --- |
| `TestChannel` | Membuat channel bertipe `string` (tanpa buffer). Sebuah goroutine mengirim pesan setelah 2 detik, dan fungsi utama menerimanya dengan `<-channel`. Penerimaan akan menunggu sampai ada data yang dikirim. |

Channel ditutup dengan `defer close(channel)` di akhir fungsi.

## Contoh Output

`TestChannel` kira-kira menghasilkan:

```
Goroutine finished sending value
Hello from goroutine!
```

Urutan dua baris itu bisa tertukar, karena goroutine pengirim dan fungsi penerima berjalan bersamaan.

## Catatan

- `time.Sleep` dipakai di sini hanya untuk keperluan belajar. Pada kode sungguhan, gunakan `sync.WaitGroup` atau channel untuk menunggu goroutine selesai.
- Mencetak 100.000 baris di `TestManyGoroutine` membuat output terminal cukup panjang. Kalau mau lebih ringkas, jalankan test lain saja dengan `-run`.