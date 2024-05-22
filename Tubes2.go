package main

import "fmt"

const Start int = 10000

type Barang struct {
	nama, kategori                                string
	modal, harga, idBarang, banyak, banyakTerjual int
}

type arrBarang [Start]Barang

func main() {
	menu()
}

func menu() {
	var pilihan int
	var n int
	var data arrBarang
	fmt.Println("Aplikasi Jual Beli Barang")
	fmt.Println(" ")
	fmt.Println("Silahkan Pilh Opsi yang Ingin di Akses")
	fmt.Println("(1) Akses Database Barang")
	fmt.Println("(2) Akses Database Transaksi")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menuDatabaseBarang(&data, &n)
	case 2:
		menuDatabaseTransaksi(&data, &n)
	default:
		fmt.Println("Pilihan tidak tersedia")
		menu()
	}
}

func menuDatabaseBarang(data *arrBarang, n *int) {
	var pilihan int
	fmt.Println("Menu Database Barang")
	fmt.Println(" ")
	fmt.Println("Silahkan Pilh Opsi yang Ingin di Akses")
	fmt.Println("(1) Tambah Data Barang")
	fmt.Println("(2) Edit Data Barang")
	fmt.Println("(3) Hapus Data Barang")
	fmt.Println("(4) Total Modal Barang")
	fmt.Println("(5) Menampilkan Data Barang")
	fmt.Println("(6) Cari barang")
	fmt.Println("(7) Kembali Ke Menu Awal")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahDataBarang(data, n)
	case 2:
		editDataBarang(data, *n)
	case 3:
		hapusDataBarang(*data, *n)
	case 4:
		fmt.Println("Modal seluruh barang :")
		totalModalBarang(*data, *n)
	case 5:
		tampilDataBarang(data, n)
	case 6:
		mencariBarang(data, n)
	}
}

func menuDatabaseTransaksi(data *arrBarang, n *int) {
	var data arrBarang
	var pilihan, n int
	fmt.Println("Menu Database Transaksi")
	fmt.Println(" ")
	fmt.Println("Silahkan Pilh Opsi yang Ingin di Akses")
	fmt.Println("(1) Tambah Data Transaksi")
	fmt.Println("(2) Edit Data Transaksi")
	fmt.Println("(3) Hapus Data Transaksi")
	fmt.Println("(4) Total Pendapatan Transaksi")
	fmt.Println("(5) Menampilkan 5 Barang Terbanyak Terjual")
	fmt.Println("(6) Kembali Ke Menu Awal")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahDataTransaksi(&data, &n)
	case 2:
		editDataTransaksi(&data, n)
	case 3:
		hapusDataTransaksi(data, n)
	case 4:
		totalPendapatanTransaksi(data, n)
	case 5:
		barangTerbanyakTerjual(data, n)
	case 6:
		menu()
	default:
		fmt.Println("Pilihan tidak tersedia")
		menuDatabaseTransaksi()
	}
}

func tambahDataBarang(A *arrBarang, n *int) {
	var nama string
	var i int
	for i = 0; i < Start && (nama != "None" || nama != "none"); i++ {
		fmt.Println("ID Barang: ")
		fmt.Scan(&A[i].idBarang)
		fmt.Println("Nama: ")
		fmt.Scan(&A[i].nama)
		fmt.Println("Kategori: ")
		fmt.Scan(&A[i].kategori)
		fmt.Println("Modal: ")
		fmt.Scan(&A[i].modal)
		fmt.Println("Harga: ")
		fmt.Scan(&A[i].harga)
		fmt.Println("Total Barang: ")
		fmt.Scan(&A[i].banyak)
		nama = A[i].nama
		*n++
	}
	if nama == "None" || nama == "none" {
		menuDatabaseBarang(A, n)
	}
}

func editDataBarang(A *arrBarang, n int) {
	var indexBaru, Pilihan int
	var Dicari string
	fmt.Println("Masukkan Nama Barang yang ingin di edit:")
	fmt.Scan(&Dicari)
	indexBaru = searchEditIndex(A, n, Dicari)
	fmt.Println("Apa yang ingin di edit?")
	fmt.Println("(1) ID Barang")
	fmt.Println("(2) Nama")
	fmt.Println("(3) Kategori")
	fmt.Println("(4) Modal")
	fmt.Println("(5) Harga")
	fmt.Println("(6) Banyak Barang")
	fmt.Println("(7) Kembali ke Menu Awal")
	fmt.Scan(&Pilihan)
	switch Pilihan {
	case 1:
		fmt.Println("Masukkan ID Barang:")
		fmt.Scan(&A[indexBaru].idBarang)
	case 2:
		fmt.Println("Masukkan nama Barang:")
		fmt.Scan(&A[indexBaru].nama)
	case 3:
		fmt.Println("Masukkan Kategori:")
		fmt.Scan(&A[indexBaru].kategori)
	case 4:
		fmt.Println("Masukkan Modal:")
		fmt.Scan(&A[indexBaru].modal)
	case 5:
		fmt.Println("Masukkan Harga:")
		fmt.Scan(&A[indexBaru].harga)
	case 6:
		fmt.Println("Masukkan Banyak Barang:")
		fmt.Scan(&A[indexBaru].banyak)
	}
}

func hapusDataBarang(A arrBarang, n int) {
	var i, indexbaru int
	var dicari int
	fmt.Println("Masukkan ID Barang yang ingin dihapus data barangnya:")
	fmt.Scanln(&dicari)
	indexbaru = searchEditIndex2(&A, n, dicari)
	for i = indexbaru; i < n; i++ {
		A[i] = A[i+1]
	}
}

func totalModalBarang(A arrBarang, n int) int {
	var i, total int
	total = 0
	for i = 0; i < n; i++ {
		total = total + (A[i].modal * A[i].banyak)
	}
	return total
}

func tambahDataTransaksi(A *arrBarang, n *int) {
	var i int
	var nama string = "A"
	for i = 0; i < Start && (nama != "None" || nama != "none"); i++ {
		fmt.Print("Nama Barang Transaksi :")
		fmt.Scan(&nama)
		if nama == A[i].nama {
			fmt.Println("Banyak Barang:")
			fmt.Scan(&A[i].banyakTerjual)
		}
	}
}

func editDataTransaksi(A *arrBarang, n int) {
	var indexBaru, Pilihan int
	var Dicari string
	fmt.Println("Masukkan Nama Barang yang ingin di edit:")
	fmt.Scan(&Dicari)
	indexBaru = searchEditIndex(A, n, Dicari)
	fmt.Println("Apa yang ingin di edit?")
	fmt.Println("(1) Banyak Barang")
	fmt.Println("(2) Kembali ke menu awal")
	fmt.Scan(&Pilihan)
	switch Pilihan {
	case 1:
		fmt.Println("Masukkan Banyak Barang:")
		fmt.Scan(&A[indexBaru].banyakTerjual)
	case 2:
		menu()
	}
}

func hapusDataTransaksi(A arrBarang, n int) {

}

func totalPendapatanTransaksi(A arrBarang, n int) int {
	var i, total int
	total = 0
	for i = 0; i < n; i++ {
		total = total + (A[i].harga * A[i].banyakTerjual)
	}
	return total
}

func searchEditIndex(A *arrBarang, n int, x string) int {
	var i, index int
	for i = 0; i < n; i++ {
		if x == A[i].nama {
			index = i
		}
	}
	return index
}

func searchEditIndex2(A *arrBarang, n int, x int) int {
	var kiri, kanan, tengah, index int
	var i int
	for i = 0; i < n; i++ {
		tengah = (kiri + kanan) / 2
		if A[tengah].idBarang < x {
			kiri = tengah + 1
		} else if A[tengah].idBarang > x {
			kanan = tengah - 1
		} else {
			index = tengah
		}
	}
	return index
}
