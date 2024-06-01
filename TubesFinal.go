package main

import (
	"fmt"
)

const Start int = 10000

type Barang struct {
	nama, kategori               string
	modal, harga, idBarang, stok int
}

type BarangTransaksi struct {
	idBarang     int
	nama         string
	jumlahBarang int
}

type arrBarang [Start]Barang
type arrTransaksi [Start]BarangTransaksi

func main() {
	menu()
}

func menu() {
	var pilihan int
	var n, k int = 0, 0
	var data arrBarang
	var data2 arrTransaksi
	fmt.Println("Aplikasi Jual Beli Barang")
	fmt.Println("Silahkan Pilh Opsi yang Ingin di Akses")
	fmt.Println("(1) Akses Database Barang")
	fmt.Println("(2) Akses Database Transaksi")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menuDatabaseBarang(&data, &n)
	case 2:
		menuDatabaseTransaksi(&data, &data2, &n, &k)
	default:
		fmt.Println("Pilihan tidak tersedia")
		menu()
	}
}

func menuDatabaseBarang(data *arrBarang, n *int) {
	var pilihan int
	fmt.Println("Menu Database Barang")
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
		hapusDataBarang(data, *n)
	case 4:
		fmt.Println("Modal seluruh barang :")
		totalModalBarang(*data, *n)
	case 5:
		menutampilDataBarang(data, *n)
	case 6:
		menucaribarang(*data, *n)
	}
}

func menutampilDataBarang(A *arrBarang, n int) {
	var pilihan int
	fmt.Println("Tampilkan Data terurut berdasarkan:")
	fmt.Println("(1) ID Barang")
	fmt.Println("(2) Harga Tertinggi ke rendah")
	fmt.Println("(3) Harga Terendah ke tertinggi")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("ID Barang")
		selsortID(A, n)
		fmt.Printf("%40s %40s %40s %40s %40s %40s\n", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
		for i := 0; i < n; i++ {
			fmt.Printf("%40d %40s %40s %40d %40d %40d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
		}
		menu()
	case 2:
		fmt.Println("Harga Tertinggi ke rendah")
		insortHNaik(A, n)
		fmt.Printf("%40s %40s %40s %40s %40s %40s\n", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
		for i := 0; i < n; i++ {
			fmt.Printf("%40d %40s %40s %40d %40d %40d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
		}
		menu()
	case 3:
		fmt.Println("Harga Terendah ke tertinggi")
		insortHTurun(A, n)
		fmt.Printf("%40s %40s %40s %40s %40s %40s", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
		for i := 0; i < n; i++ {
			fmt.Printf("%40d %40s %40s %40d %40d %40d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
		}
		fmt.Println("Kembali Kemenu awal?")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			menu()
		} else {
			menu()
		}
	}
}

func menuDatabaseTransaksi(data *arrBarang, data2 *arrTransaksi, n, k *int) {
	var pilihan int
	fmt.Println("Menu Database Transaksi")
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
		tambahDataTransaksi(data, data2, *n, k)
	case 2:
		editDataTransaksi(data, data2, *n, *k)
	case 3:
		hapusDataTransaksi(data, data2, *n, k)
	case 4:
		totalPendapatanTransaksi(*data, *data2, *n, *k)
	case 5:
		barangTerbanyakTerjual(data2, *k)
	case 6:
		menu()
	}
}

func menucaribarang(A arrBarang, n int) {
	var pilihan int
	fmt.Println("Mau cari berdasarkan:")
	fmt.Println("(1) ID Barang")
	fmt.Println("(2) Nama Barang")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		var id, index2 int
		fmt.Println("Masukkan ID Barang:")
		fmt.Scanln(&id)
		index2 = searchIndexbyidBarangA(A, n, id)
		fmt.Println("Data Barang yang dicari:")
		fmt.Println("ID Barang:", A[index2].idBarang)
		fmt.Println("Nama Barang:", A[index2].nama)
		fmt.Println("Kategori:", A[index2].kategori)
		fmt.Println("Modal:", A[index2].modal)
		fmt.Println("Harga:", A[index2].harga)
		fmt.Println("Total Barang:", A[index2].stok)
		fmt.Println("Kembali Kemenu awal?")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			menu()
		} else {
			menu()
		}

	case 2:
		var nama string
		var index int
		fmt.Println("Masukkan Nama Barang:")
		fmt.Scanln(&nama)
		index = searchIndexbynamaA(A, n, nama)
		fmt.Println("Data Barang yang dicari:")
		fmt.Println("ID Barang:", A[index].idBarang)
		fmt.Println("Nama Barang:", A[index].nama)
		fmt.Println("Kategori:", A[index].kategori)
		fmt.Println("Modal:", A[index].modal)
		fmt.Println("Harga:", A[index].harga)
		fmt.Println("Total Barang:", A[index].stok)
		fmt.Println("Kembali Kemenu awal?")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			menu()
		} else {
			menu()
		}
	default:
		fmt.Println("Pilihan tidak tersedia")
		menucaribarang(A, n)
	}
}

func tambahDataBarang(A *arrBarang, n *int) {
	var nama string = "A"
	var i int
	for i = 0; i < Start && nama != "none" && nama != "None"; i++ {
		fmt.Println("Barang", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&nama)
		if nama != "none" && nama != "None" {
			A[i].nama = nama
			fmt.Print("ID Barang: ")
			fmt.Scan(&A[i].idBarang)
			fmt.Print("Kategori: ")
			fmt.Scan(&A[i].kategori)
			fmt.Print("Modal: ")
			fmt.Scan(&A[i].modal)
			fmt.Print("Harga: ")
			fmt.Scan(&A[i].harga)
			fmt.Print("Total Barang: ")
			fmt.Scan(&A[i].stok)
			*n++
		}
	}
	if nama == "None" || nama == "none" {
		menuDatabaseBarang(A, n)
	}
}

func editDataBarang(A *arrBarang, n int) {
	var index, Pilihan int
	var Dicari string
	fmt.Println("Masukkan Nama Barang yang ingin di edit:")
	fmt.Scan(&Dicari)
	index = searchIndexbynamaA(*A, n, Dicari)
	if index != -1 {
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
			fmt.Scan(&A[index].idBarang)
			menu()
		case 2:
			fmt.Println("Masukkan nama Barang:")
			fmt.Scan(&A[index].nama)
			menu()
		case 3:
			fmt.Println("Masukkan Kategori:")
			fmt.Scan(&A[index].kategori)
			menu()
		case 4:
			fmt.Println("Masukkan Modal:")
			fmt.Scan(&A[index].modal)
			menu()
		case 5:
			fmt.Println("Masukkan Harga:")
			fmt.Scan(&A[index].harga)
			menu()
		case 6:
			fmt.Println("Masukkan Banyak Barang:")
			fmt.Scan(&A[index].stok)
			menu()
		}
	} else {
		fmt.Println("Barang yang dicari tidak ada")
	}
}

func hapusDataBarang(A *arrBarang, n int) {
	var i, indexbaru int
	var dicari int
	fmt.Println("Masukkan ID Barang yang ingin dihapus data barangnya:")
	fmt.Scanln(&dicari)
	indexbaru = searchIndexbyidBarangA(*A, n, dicari)
	for i = indexbaru; i < n; i++ {
		A[i] = A[i+1]
	}
}

func totalModalBarang(A arrBarang, n int) int {
	var i, total int
	total = 0
	for i = 0; i < n; i++ {
		total = total + (A[i].modal * A[i].stok)
	}
	return total
}

func tambahDataTransaksi(A *arrBarang, B *arrTransaksi, n int, k *int) {
	var nama string = "A"
	var nBarang, index, IDBarang, temp int
	for nama != "none" && nama != "None" {
		fmt.Print("Masukkan ID Barang: ")
		fmt.Scan(&IDBarang)
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan Banyak Barang: ")
		fmt.Scan(&nBarang)
		index = searchIndexbynamaA(*A, n, nama)
		if index != -1 {
			temp = A[index].stok
			if A[index].stok != 0 {
				A[index].stok -= nBarang
				if A[index].stok >= 0 {
					B[*k].idBarang = IDBarang
					B[*k].nama = nama
					B[*k].jumlahBarang = nBarang
					*k++
				} else {
					A[index].stok = temp
					fmt.Println("Stok Barang Tidak mencukupi, silahkan input ulang")

				}
			} else {
				fmt.Println("Stok Barang Habis, silahkan input ulang")
			}
		} else {
			fmt.Println("Barang tidak ditemukan")
		}
	}
}

func editDataTransaksi(A *arrBarang, B *arrTransaksi, n, k int) {
	var nama string = "A"
	var indexB, Pilihan int
	fmt.Println("Masukkan nama barang yang ingin di edit")
	indexB = searchIndexbynamaB(B, k, nama)
	if indexB != -1 {
		fmt.Println("Apa yang mau di edit?")
		fmt.Println("(1) ID Barang")
		fmt.Println("(2) Nama Barang")
		fmt.Println("(3) Jumlah Barang")
		fmt.Scanln(&Pilihan)
		switch Pilihan {
		case 1:
			fmt.Print("Masukkan IDBarang: ")
			fmt.Scan(&B[indexB].idBarang)
		case 2:
			fmt.Print("Masukkan Nama Barang: ")
			fmt.Scan(&B[indexB].nama)
		case 3:
			var temp, Qty, indexA int
			var nama string
			nama = B[indexB].nama
			indexA = searchIndexbynamaA(*A, n, nama)
			fmt.Print("Masukkan Jumlah Barang: ")
			temp = B[indexB].jumlahBarang
			fmt.Scan(&B[indexB].jumlahBarang)
			nama = B[indexB].nama
			indexA = searchIndexbynamaA(*A, n, nama)
			B[indexB].jumlahBarang = Qty
			if temp > Qty {
				A[indexA].stok -= Qty
			} else if temp < Qty {
				A[indexA].stok += Qty
			} else {
				A[indexA].stok = Qty
			}
		}
	}
}

func hapusDataTransaksi(A *arrBarang, B *arrTransaksi, n int, k *int) {
	var i, indexbaru, dicari, indexB, Qty int
	var nama string
	fmt.Println("Masukkan ID Barang yang ingin dihapus data barangnya:")
	fmt.Scanln(&dicari)
	indexbaru = searchIndexbyidBarangA(*A, n, dicari)
	nama = A[indexbaru].nama
	indexB = searchIndexbynamaB(B, *k, nama)
	Qty = B[indexB].jumlahBarang
	A[indexbaru].stok += Qty
	for i = indexbaru; i < n; i++ {
		B[i] = B[i+1]
	}
}

func totalPendapatanTransaksi(A arrBarang, B arrTransaksi, n, k int) int {
	var total, i, indexA int
	var nama string
	for i = 0; i < k; i++ {
		nama = B[i].nama
		indexA = searchIndexbynamaA(A, n, nama)
		total = total + (A[indexA].harga * B[i].jumlahBarang)
	}
	return total
}

func barangTerbanyakTerjual(B *arrTransaksi, n int) {
	var i, j int
	var temp BarangTransaksi
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if B[i].jumlahBarang < B[j].jumlahBarang {
				temp = B[i]
				B[i] = B[j]
				B[j] = temp
			}
		}
	}
	fmt.Println("Barang yang paling banyak terjual")
	for i = 0; i < 5; i++ {
		fmt.Println(B[i].nama, B[i].jumlahBarang)
	}
}

func searchIndexbynamaA(A arrBarang, n int, x string) int {
	var i, index int
	i = -1
	for i = 0; i < n; i++ {
		if x == A[i].nama {
			index = i
		}
	}
	return index
}

func searchIndexbynamaB(A *arrTransaksi, n int, x string) int {
	var i, index int
	i = -1
	for i = 0; i < n; i++ {
		if x == A[i].nama {
			index = i
		}
	}
	return index
}

func searchIndexbyidBarangA(A arrBarang, n int, x int) int {
	var kiri, kanan, tengah, index int
	var i int
	selsortID(&A, n)
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

func selsortID(A *arrBarang, n int) {
	var minIdx, i, j int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].idBarang < A[minIdx].idBarang {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

func selsortIDB(A *arrTransaksi, n int) {
	var minIdx, i, j int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].idBarang < A[minIdx].idBarang {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

func insortHNaik(A *arrBarang, n int) {
	var i, pass int
	var temp Barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp.harga = A[1].harga
		for i > 0 && temp.harga < A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass = pass + 1
	}
}

func insortHTurun(A *arrBarang, n int) {
	var i, pass int
	var temp Barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp.harga = A[1].harga
		for i > 0 && temp.harga > A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass = pass + 1
	}

}
