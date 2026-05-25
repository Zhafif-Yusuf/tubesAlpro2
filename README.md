# PowerLog

Aplikasi pencatatan konsumsi listrik perangkat elektronik berbasis terminal menggunakan bahasa Go (Golang).

## Deskripsi

PowerLog merupakan aplikasi yang dibuat untuk membantu pengguna dalam mencatat dan memantau konsumsi listrik perangkat elektronik. Aplikasi ini dikembangkan sebagai Tugas Besar (TuBes) mata kuliah Algoritma dan Pemrograman.

Melalui aplikasi ini, pengguna dapat:

* Menambahkan data perangkat elektronik
* Menampilkan daftar perangkat
* Menghitung total konsumsi listrik
* Melakukan pencarian data perangkat
* Mengurutkan data perangkat

## Fitur

* CRUD data perangkat elektronik
* Perhitungan konsumsi daya listrik
* Pencarian data perangkat
* Pengurutan data perangkat
* Interface berbasis terminal/CLI

## Teknologi yang Digunakan

* Bahasa Pemrograman Go (Golang)
* Visual Studio Code

## Struktur Data

```go
type Perangkat struct {
    perangkat string
    ruangan   string
    watt      int
    durasi    int
}
```

## Algoritma yang Digunakan

Beberapa konsep algoritma dan pemrograman yang diterapkan:

* Array
* Struct
* Function/Procedure
* Sequential Search
* Sorting Algorithm
* Perulangan
* Percabangan

## Anggota Kelompok

Kelompok 11

* Dadi Maulana Muhammad
* Hanan Fahri Abiyyu
* Zhafif Yusuf Al Amin

## Referensi

1. Donovan, A. A., & Kernighan, B. *The Go Programming Language*. Addison-Wesley, 2015.
2. Cormen, Thomas H., dkk. *Introduction to Algorithms Third Edition*. MIT Press, 2009.
3. Sedgewick, Robert & Wayne Kevin. *Algorithms Fourth Edition*. Addison-Wesley.
4. Materi Mata Kuliah Algoritma dan Pemrograman.
5. Dokumentasi Resmi Go Programming Language.

## Lisensi

Project ini dibuat untuk keperluan akademik dan pembelajaran. Silahkan gunakan source code kami sebijak mungkin.
