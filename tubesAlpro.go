	// Aplikasi Pencatatan Konsumsi Listrik Perangkat (PowerLog)
	// Kelompok 11
	// S1IF-13-04

	package main

	import (
		"bufio"
		"fmt"
		"os"
		"strings"
	)

	const Max int = 100

	type Perangkat struct {
		perangkat string
		ruangan   string
		watt      int
		durasi    int
	}

	type arrPerangkat [Max]Perangkat

	func addData(arr *arrPerangkat, banyakdata *int) int {
		var data Perangkat
		var jumlahdata int
		fmt.Print("Masukkan jumlah data yang ingin ditambahkan : ")
		fmt.Scan(&jumlahdata)
		if jumlahdata <= 0 {
			fmt.Println("Jumlah data tidak valid!")
			return *banyakdata
		}

		if *banyakdata+jumlahdata > Max-1 {
			fmt.Println("Data Penuh!")
			return *banyakdata
		}
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		for i := 0; i < jumlahdata; i++ {
			fmt.Print("Masukkan nama perangkat : ")
			data.perangkat, _ = reader.ReadString('\n')
			data.perangkat = strings.TrimSpace(data.perangkat)

			fmt.Print("Durasi pemakaian perangkat dalam sehari (jam) : ")
			fmt.Scan(&data.durasi)
			fmt.Print("Daya listrik perangkat (Watt) : ")
			fmt.Scan(&data.watt)

			reader.ReadString('\n')

			fmt.Print("Nama ruangan : ")
			data.ruangan, _ = reader.ReadString('\n')
			data.ruangan = strings.TrimSpace(data.ruangan)

			fmt.Println("Data berhasil disimpan.")
			fmt.Println()
			arr[*banyakdata] = data
			*banyakdata++
		}

		return *banyakdata
	}

	func showData(arr arrPerangkat, banyakdata int) {
		if banyakdata == 0 {
			fmt.Println("Data Kosong! Silakan tambah data terlebih dahulu.")
			return
		}
		fmt.Println()
		fmt.Println("Berikut data perangkat yang sudah ditambahkan :")
		for i := 0; i < banyakdata; i++ {
			fmt.Printf("%d. %s ", (i + 1), arr[i].perangkat)
			fmt.Printf("%d jam ", arr[i].durasi)
			fmt.Printf("%d Watt ", arr[i].watt)
			fmt.Printf("%s\n", arr[i].ruangan)
		}

	}

	func changeData(arr *arrPerangkat, banyakdata int) {
		var idx int

		if banyakdata == 0 {
			fmt.Println("Data Kosong! Silakan tambah data terlebih dahulu.")
			return
		}
		showData(*arr, banyakdata)
		fmt.Print("Masukkan index yang ingin di ubah : ")
		fmt.Scan(&idx)

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		if idx >= 1 && idx <= banyakdata {
			fmt.Print("Masukkan perangkat baru : ")
			arr[idx-1].perangkat, _ = reader.ReadString('\n')
			arr[idx-1].perangkat = strings.TrimSpace(arr[idx-1].perangkat)

			fmt.Print("Masukan durasi pemakaian baru : ")
			fmt.Scan(&arr[idx-1].durasi)

			fmt.Print("Daya listrik perangkat baru : ")
			fmt.Scan(&arr[idx-1].watt)

			reader.ReadString('\n')

			fmt.Print("Ruangan baru : ")
			arr[idx-1].ruangan, _ = reader.ReadString('\n')
			arr[idx-1].ruangan = strings.TrimSpace(arr[idx-1].ruangan)

			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Index tidak valid!")
		}

	}

	func deleteData(arr *arrPerangkat, banyakdata *int) int {
		var idx int
		if *banyakdata == 0 {
			fmt.Println("Data Kosong! Silakan tambah data terlebih dahulu.")
			return *banyakdata
		}
		showData(*arr, *banyakdata)
		fmt.Print("Masukkan index yang ingin dihapus : ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= *banyakdata {
			for i := idx - 1; i < *banyakdata-1; i++ {
				(*arr)[i] = (*arr)[i+1]
			}
			fmt.Println("Data berhasil dihapus.")
			*banyakdata--
		} else {
			fmt.Print("Index tidak valid!")
			fmt.Println()
		}

		return *banyakdata
	}

	func searchData(arr arrPerangkat, banyakdata int) int {
		var found int
		var k int
		var X string
		if banyakdata == 0 {
			return -1
		}
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		fmt.Print("Masukkan Nama Perangkat yang akan dicari : ")
		X, _ = reader.ReadString('\n')
		X = strings.TrimSpace(X)

		found = -1
		k = 0
		for found == -1 && k < banyakdata {
			if arr[k].perangkat == X {
				found = k
			}
			k++
		}

		return found

	}

	func sortData(arr *arrPerangkat, banyakdata int) {

		var idxMax int
		if banyakdata == 0 {
			fmt.Println("Data Kosong! Silakan tambah data terlebih dahulu.")
			return
		}
		for i := 0; i < banyakdata-1; i++ {
			idxMax = i
			for j := i + 1; j < banyakdata; j++ {
				if (*arr)[j].watt > (*arr)[idxMax].watt {
					idxMax = j
				}
			}
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[idxMax]
			(*arr)[idxMax] = temp
		}

		showData(*arr, banyakdata)

	}

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
		var pilihan int
		fmt.Println(" ————— Aplikasi Pencatatan Konsumsi Listrik Perangkat (PowerLog) ————— ")
		for {
			fmt.Println()
			fmt.Println("Silakan pilih menu di bawah ini : ")
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

			switch pilihan {
			case 1:
				addData(&arr, &banyakdata)
			case 2:
				showData(arr, banyakdata)
			case 3:
				changeData(&arr, banyakdata)
			case 4:
				deleteData(&arr, &banyakdata)
			case 5:
				idx := searchData(arr, banyakdata)
				if banyakdata == 0 {
					fmt.Println("Data kosong! Silakan tambah data terlebih dahulu.")
				} else if idx != -1 {
					fmt.Println("Data ditemukan!")
					fmt.Println("Nama Perangkat  :", arr[idx].perangkat)
					fmt.Println("Durasi :", arr[idx].durasi)
					fmt.Println("Watt :", arr[idx].watt)
					fmt.Println("Nama Ruangan :", arr[idx].ruangan)
				} else {
					fmt.Println("Data tidak ditemukan!")
				}

			case 6:
				sortData(&arr, banyakdata)
			case 7:
				analysisData(arr, banyakdata)
			case 0:
				fmt.Println("Terima kasih telah menggunakan PowerLog!")
				return
			default:
				fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
			}
		}
	}
