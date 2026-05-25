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
func analysisData(arr arrPerangkat, banyakdata int)

func main() {
	var arr arrPerangkat
	var banyakdata int
	var pilihan int

	fmt.Println("Aplikasi PowerLog")

	for {
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Data")
		fmt.Println("6. Urutkan Data")
		fmt.Println("7. Analisis Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
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
			searchData(arr, banyakdata)
		case 6:
			sortData(&arr, banyakdata)
		case 7:
			analysisData(arr, banyakdata)
		case 0:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}