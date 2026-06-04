package main

import "fmt"

const Max int = 100

type Perangkat struct {
	perangkat string
	ruangan   string
	watt      int
	durasi    int
}

type arrPerangkat [Max]Perangkat

func addData(arr *arrPerangkat, banyakdata *int) int
func showData(arr arrPerangkat, banyakdata int)
func changeData(arr *arrPerangkat, banyakdata int)
func deleteData(arr *arrPerangkat, banyakdata *int) int
func searchData(arr arrPerangkat, banyakdata int) int
func sortData(arr *arrPerangkat, banyakdata int)

// fungsi untuk menentukan total watt dari keseluruhan perangkat dan menentukan perangkat yang paling boros energi.
func analysisData(arr arrPerangkat, banyakdata int) {
	var total int
	var idxmax int

	if banyakdata == 0 {
		fmt.Println("Data Kosong! Silakan tambah data terlebih dahulu.")
		return
	}

	for i := 0; i < banyakdata; i++ {
		total += arr[i].watt

		if arr[i].watt > arr[idxmax].watt {
			idxmax = i
		}
	}
	fmt.Println("Total konsumsi listrik ", total)
	fmt.Println("Perangkat paling boros energi ", arr[idxmax].perangkat)
	fmt.Println("Daya Listrik : ", arr[idxmax].watt)
}

func main() {
	var arr arrPerangkat
	var banyakdata int

	var pilihan string

	fmt.Println(" ————— Aplikasi Pencatatan Konsumsi Listrik Perangkat (PowerLog) ————— ")

	for {
		fmt.Println()
		fmt.Println("1. Tambah Perangkat")
		fmt.Println("2. Tampilkan Perangkat")
		fmt.Println("3. Edit Perangkat")
		fmt.Println("4. Hapus Perangkat")
		fmt.Println("5. Cari Perangkat")
		fmt.Println("6. Urutkan Konsumsi")
		fmt.Println("7. Analisis Penggunaan Listrik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")

		fmt.Scan(&pilihan)

		if banyakdata == 0 && (pilihan == "2" || pilihan == "3" || pilihan == "4" || pilihan == "5" || pilihan == "6" || pilihan == "7") {
			fmt.Println(" Data masih kosong! Silakan pilih menu 1 untuk menambah data.")
			continue
		}

		switch pilihan {

		case "1":
			if banyakdata >= len(arr) {
				fmt.Println("ERROR: Kapasitas penyimpanan penuh! Tidak bisa menambah data.")
			} else {
				addData(&arr, &banyakdata)
			}

		case "2":
			showData(arr, banyakdata)

		case "3":
			changeData(&arr, banyakdata)

		case "4":
			deleteData(&arr, &banyakdata)

		case "5":
			idx := searchData(arr, banyakdata)

			if idx != -1 {
				fmt.Println("Data ditemukan!")
				fmt.Println("Nama Perangkat :", arr[idx].perangkat)
				fmt.Println("Durasi         :", arr[idx].durasi)
				fmt.Println("Watt           :", arr[idx].watt)
				fmt.Println("Nama Ruangan   :", arr[idx].ruangan)
			} else {
				fmt.Println("Data tidak ditemukan!")
			}

		case "6":
			sortData(&arr, banyakdata)

		case "7":
			analysisData(arr, banyakdata)

		case "0":
			fmt.Println("Terima kasih telah menggunakan PowerLog!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Pilih angka dari 0 sampai 7.")
		}
	}
}
