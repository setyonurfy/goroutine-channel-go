# goroutine-go

Catatan belajar **concurrency di Go** (goroutine, channel, select, mutex, atomic, context) yang dipelajari lewat **unit test**, bukan lewat `main.go`. Setiap konsep punya fungsi kecil dan test yang membuktikan cara kerjanya.

## Prasyarat

- Go 1.21 atau lebih baru (`go version`)
- Tidak ada dependensi eksternal, hanya library standar

## Cara menjalankan

Project ini berupa **package pustaka** (`package belajar`) tanpa `func main()`, jadi dijalankan lewat `go test`, bukan `go run`.

```bash
# jalankan semua test, dengan detektor race condition
go test -race -v

# jalankan satu topik saja
go test -race -v -run Unduh
go test -race -v -run "Kuadrat|Genap"
go test -race -v -run WorkerPool
go test -race -v -run "Tercepat|Timeout|CobaAmbil"
go test -race -v -run "SafeCounter|SafeStok"
go test -race -v -run "Atomic|CompareAndSwap"
go test -race -v -run "KerjaLama|Pencacah"

# cek masalah umum (misalnya lupa memanggil cancel)
go vet ./...
```

> Beberapa test mengukur waktu dan memakai `time.Sleep`, jadi seluruh suite butuh beberapa detik.

## Struktur project

```
goroutine-go/
├── go.mod
├── unduh.go      unduh_test.go      goroutine dan WaitGroup
├── channel.go    channel_test.go    channel
├── pool.go       pool_test.go       worker pool
├── select.go     select_test.go     select, timeout, non-blocking
├── counter.go                       SafeCounter (mutex)
├── stok.go       stok_test.go       SafeStok (mutex + map)
├── mutex_test.go                    race condition dan test SafeCounter
├── atomic.go     atomic_test.go     atomic
└── context.go    context_test.go    context
```

## Materi

| # | Topik | Fungsi / tipe | Konsep yang dipelajari |
|---|---|---|---|
| 1 | Goroutine dan `WaitGroup` | `Unduh`, `UnduhSemua` | `go`, `wg.Add/Done/Wait`, menulis ke indeks berbeda, jebakan variabel loop |
| 2 | Channel | `Kuadrat`, `Genap` | unbuffered vs buffered, `close`, `range`, channel searah `<-chan` |
| 3 | Worker pool | `WorkerPool` | banyak pekerja berbagi satu antrean, hasil berurutan lewat indeks |
| 4 | `select` | `Tercepat`, `TungguDenganTimeout`, `CobaAmbil`, `TercepatDenganTimeout` | menunggu banyak channel, timeout dengan `time.After`, `default` |
| 5 | Mutex | `SafeCounter`, `SafeStok` | race condition, `Lock/Unlock`, mutex di dalam struct, cek-lalu-ubah |
| 6 | Atomic | `AtomicCounter` | `atomic.Int64`, `Add/Load`, `CompareAndSwap` |
| 7 | Context | `KerjaLama`, `Pencacah` | `WithCancel`, `WithTimeout`, `ctx.Done()`, `ctx.Err()`, mencegah goroutine leak |

Selain materi di atas, ada juga pembahasan `RWMutex` (kunci baca untuk data yang jarang berubah), tapi default yang dipakai tetap `sync.Mutex`.

## Pola yang dipakai

**Menunggu semua goroutine**

```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
	wg.Add(1)           // sebelum go, di luar goroutine
	go func(id int) {
		defer wg.Done()
		// kerja...
	}(i)                // kirim i sebagai argumen
}
wg.Wait()
```

**Generator dengan channel**

```go
func Kuadrat(angka []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // pengirim yang menutup
		for _, n := range angka {
			out <- n * n
		}
	}()
	return out
}
```

**Menunggu dengan batas waktu atau pembatalan**

```go
select {
case hasil := <-ch:
	return hasil, nil
case <-ctx.Done():
	return "", ctx.Err()
}
```

## Mutex atau atomic?

| Yang dilindungi | Pakai |
|---|---|
| Satu angka atau satu flag | `atomic` |
| `map`, slice, struct | `sync.Mutex` |
| Dua variabel atau lebih yang berubah bersamaan | `sync.Mutex` |
| Cek dulu lalu ubah (misalnya cek stok lalu kurangi) | `sync.Mutex` |
| Ragu | `sync.Mutex` |

## Pelajaran penting

- **Selalu jalankan test dengan `-race`.** Detektor ini menangkap akses data tanpa kunci, bahkan saat hasil hitungan kebetulan benar.
- **`-race` tidak menangkap kesalahan logika penguncian.** Cek-lalu-ubah yang kuncinya dipecah dua tidak memicu `DATA RACE`, tapi hasilnya salah. Test yang memeriksa hasil akhir yang menangkapnya.
- **Jangan mengandalkan urutan eksekusi goroutine.** Urutkan hasil dulu (`sort`) atau tulis ke indeks yang pasti.
- **Test yang bergantung pada waktu butuh margin.** Pasang batas di antara waktu ideal dan waktu kalau kodenya salah, dan hati-hati dengan `Millisecond` vs `Microsecond`.
- **`wg.Add` sebelum `go`, `defer wg.Done()` di dalam goroutine.**
- **Yang mengirim yang menutup channel.** Mengirim ke channel tertutup akan panic.
- **Jangan menyalin `Mutex`, `WaitGroup`, atau tipe `atomic`.** Pakai pointer receiver atau pointer argumen.
- **Channel hasil yang mungkin tidak dibaca sebaiknya buffered**, supaya pengirimnya tidak macet selamanya.
- **`cancel()` tidak memaksa goroutine berhenti.** Goroutine harus memeriksa `ctx.Done()` sendiri. Selalu `defer cancel()`, dan letakkan `ctx` sebagai parameter pertama.

## Catatan

- `TestCounterTanpaMutex` di `mutex_test.go` adalah **demo race condition yang sengaja bermasalah**. Dengan `-race` ia akan gagal. Hapus atau jadikan komentar setelah paham supaya hasil `go test -race` keseluruhan bisa hijau.
- Contoh eksperimen (menghapus `go`, `close`, `Lock`, atau `case <-ctx.Done()` untuk melihat test gagal) berguna untuk memahami apa yang dijaga oleh tiap baris kode.

## Latihan lanjutan

- `TungguPesan(ctx, ch)`: versi `TungguDenganTimeout` yang memakai context
- Ubah `WorkerPool` agar menerima `ctx` dan berhenti saat dibatalkan
- Topik lanjutan: pipeline, fan-in/fan-out, `errgroup`, `sync.Once`
- Terapkan ke mini project: proses pesanan toko online (worker pool + mutex untuk stok + context untuk timeout)