package main

import "fmt"

const NMAX int = 100

type Barang struct {
	nama, kategori               string
	modal, harga, idBarang, stok int
}

type BarangTransaksi struct {
	idBarang     int
	nama         string
	jumlahBarang int
}

type arrBarang [NMAX]Barang
type arrTransaksi [NMAX]BarangTransaksi

func main() {
	var nBarang, nTransaksi, pilih int
	var A arrBarang
	var B arrTransaksi
	nBarang = 0
	pilih = 0
	for pilih != 14 {
		menu()
		fmt.Println("Silahkan pilih menu yang ingin di akses (1/2/3/4/5/6/7/8/9/10/11/12/13/14)")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputBarang(&A, &nBarang)
		case 2:
			editBarang(&A, nBarang)
		case 3:
			hapusBarang(&A, &nBarang)
		case 4:
			fmt.Println("Modal barang keseluruhan", modalBarang(A, nBarang))
		case 5:
			defaultShow(A, nBarang)
		case 6:
			showByID(&A, nBarang)
		case 7:
			showByPrice(&A, nBarang)
		case 8:
			searchByID(A, nBarang)
		case 9:
			inputTransaksi(&A, nBarang, &B, &nTransaksi)
		case 10:
			editTransaksi(&A, nBarang, &B, nTransaksi)
		case 11:
			deleteTransaksi(&A, nBarang, &B, &nTransaksi)
		case 12:
			income(A, nBarang, B, nTransaksi)
		case 13:
			top5(&B, nTransaksi)
		case 14:
			fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Ini")
		}
	}
}

func menu() {
	fmt.Println("-----------------------------")
	fmt.Println("<<Aplikasi Jual Beli Barang>>")
	fmt.Println("-----------------------------")
	fmt.Println("<<Menu Database Barang>>")
	fmt.Println("(1)  Tambah Data Barang")
	fmt.Println("(2)  Edit Data Barang")
	fmt.Println("(3)  Hapus Data Barang")
	fmt.Println("(4)  Total Modal Barang")
	fmt.Println("(5)  Menampilkan Data Barang Seperti Input")
	fmt.Println("(6)  Menampilkan Data Barang Terurut berdasarkan ID Barang")
	fmt.Println("(7)  Menampilkan Data Barang Terurut berdasarkan Harga Barang termahal ke termurah")
	fmt.Println("(8)  Cari barang Berdasarkan ID Barang")
	fmt.Println("-----------------------------")
	fmt.Println("<<Menu Database Transaksi>>")
	fmt.Println("(9)  Tambah Data Transaksi")
	fmt.Println("(10) Edit Data Transaksi")
	fmt.Println("(11) Hapus Data Transaksi")
	fmt.Println("(12) Total Pendapatan Transaksi")
	fmt.Println("(13) Menampilkan 5 Barang Terbanyak Terjual")
	fmt.Println("(14) Keluar Aplikasi")
}

func inputBarang(A *arrBarang, n *int) {
	//menambahkan data barang
	var nama string = "A"
	for *n < NMAX && nama != "none" && nama != "None" && nama != "NONE" {
		fmt.Println("Barang", *n+1)
		fmt.Print("Nama: ")
		fmt.Scan(&nama)
		if nama != "none" && nama != "None" && nama != "NONE" {
			A[*n].nama = nama
			fmt.Print("ID Barang: ")
			fmt.Scan(&A[*n].idBarang)
			fmt.Print("Kategori: ")
			fmt.Scan(&A[*n].kategori)
			fmt.Print("Modal: ")
			fmt.Scan(&A[*n].modal)
			fmt.Print("Harga: ")
			fmt.Scan(&A[*n].harga)
			fmt.Print("Total Barang: ")
			fmt.Scan(&A[*n].stok)
			fmt.Println("-----------------------------")
			*n++
		}
	}
}

func editBarang(A *arrBarang, n int) {
	//melakukan edit data barang
	var index, Pilihan int
	var Dicari string
	fmt.Println("Masukkan Nama Barang yang ingin di edit:")
	fmt.Scan(&Dicari)
	index = seqSearchNamaA(*A, n, Dicari)
	if index == -1 {
		fmt.Println("Barang yang dicari tidak ada")
	}
	fmt.Println("(1) ID Barang")
	fmt.Println("(2) Nama")
	fmt.Println("(3) Kategori")
	fmt.Println("(4) Modal")
	fmt.Println("(5) Harga")
	fmt.Println("(6) Banyak Barang")
	fmt.Println("Pilih data yang ingin di ubah (1/2/3/4/5/6)")
	fmt.Scan(&Pilihan)
	switch Pilihan {
	case 1:
		fmt.Println("Masukkan ID Barang:")
		fmt.Scan(&A[index].idBarang)
	case 2:
		fmt.Println("Masukkan nama Barang:")
		fmt.Scan(&A[index].nama)
	case 3:
		fmt.Println("Masukkan Kategori:")
		fmt.Scan(&A[index].kategori)
	case 4:
		fmt.Println("Masukkan Modal:")
		fmt.Scan(&A[index].modal)
	case 5:
		fmt.Println("Masukkan Harga:")
		fmt.Scan(&A[index].harga)
	case 6:
		fmt.Println("Masukkan Banyak Barang:")
		fmt.Scan(&A[index].stok)
	}
}

func seqSearchNamaA(A arrBarang, n int, x string) int {
	//Mencari index berdasarkan nama barang
	var i, index int
	i = -1
	for i = 0; i < n; i++ {
		if x == A[i].nama {
			index = i
		}
	}
	return index
}

func hapusBarang(A *arrBarang, n *int) {
	var dicari string
	fmt.Print("Silahkan masukkan nama yang ingin dihapus: ")
	fmt.Scan(&dicari)
	index := seqSearchNamaA(*A, *n, dicari)
	if index == -1 {
		fmt.Println("Tempat Wisata tidak ditemukan.")
	}
	for i := index; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func modalBarang(A arrBarang, n int) int {
	var i, total int
	total = 0
	for i = 0; i < n; i++ {
		total = total + (A[i].modal * A[i].stok)
	}
	return total
}

func defaultShow(A arrBarang, n int) {
	fmt.Printf("%15s %15s %15s %15s %15s %15s\n", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%15d %15s %15s %15d %15d %15d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
	}
}

func showByID(A *arrBarang, n int) {
	fmt.Println("ID Barang")
	selSortID(A, n)
	fmt.Printf("%15s %15s %15s %15s %15s %15s\n", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%15d %15s %15s %15d %15d %15d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
	}
}

func selSortID(A *arrBarang, n int) {
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

func showByPrice(A *arrBarang, n int) {
	var i int
	var j int
	var temp Barang
	//insertion sort ascending
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].harga < temp.harga {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = temp
	}
	fmt.Printf("%15s %15s %15s %15s %15s %15s\n", "ID Barang", "Nama Barang", "Kategori", "Modal", "Harga", "Stok")
	for i = 0; i < n; i++ {
		fmt.Printf("%15d %15s %15s %15d %15d %15d\n", A[i].idBarang, A[i].nama, A[i].kategori, A[i].modal, A[i].harga, A[i].stok)
	}
}

func searchByID(A arrBarang, n int) {
	var i, j, minIdx, idx, kiri, kanan, tengah, index int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].idBarang < A[minIdx].idBarang {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
	fmt.Println("Masukkan ID Barang yang ingin dicari")
	fmt.Scan(&idx)
	index = -1
	kiri, kanan = 0, n-1
	for i = 0; i < n; i++ {
		tengah = (kiri + kanan) / 2
		if A[tengah].idBarang < idx {
			kiri = tengah + 1
		} else if A[tengah].idBarang > idx {
			kanan = tengah - 1
		} else {
			index = tengah
		}
	}
	if index == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Berikut data barang yang anda cari")
		fmt.Println("Nama Barang: ", A[index].nama)
		fmt.Println("Kategori: ", A[index].kategori)
		fmt.Println("Modal: ", A[index].modal)
		fmt.Println("Harga: ", A[index].harga)
		fmt.Println("Stok: ", A[index].stok)
	}
}

func inputTransaksi(A *arrBarang, n int, B *arrTransaksi, k *int) {
	var nama string = "A"
	var nBarang, index, IDBarang, temp int
	for nama != "none" && nama != "None" && nama != "NONE" {
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&nama)
		if nama != "none" && nama != "NONE" && nama != "None" {
			fmt.Print("Masukkan ID Barang: ")
			fmt.Scan(&IDBarang)
			fmt.Print("Masukkan Banyak Barang: ")
			fmt.Scan(&nBarang)
			fmt.Println("-----------------------------")
			index = seqSearchNamaA(*A, n, nama)
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
}

func editTransaksi(A *arrBarang, n int, B *arrTransaksi, k int) {
	var nama string = "A"
	var indexB, Pilihan int
	fmt.Println("Masukkan nama barang yang ingin di edit")
	fmt.Scan(&nama)
	indexB = seqSearchNamaB(*B, k, nama)
	if indexB == -1 {
		fmt.Println("Barang tidak ditemukan")
	}
	fmt.Println("(1) ID Barang")
	fmt.Println("(2) Nama Barang")
	fmt.Println("(3) Jumlah Barang")
	fmt.Println("Silahkan pilih data yang ingin di edit (1/2/3/4)")
	fmt.Scan(&Pilihan)
	switch Pilihan {
	case 1:
		fmt.Print("Masukkan IDBarang: ")
		fmt.Scan(&B[indexB].idBarang)
	case 2:
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&B[indexB].nama)
	case 3:
		var temp, indexA int
		var nama string
		nama = B[indexB].nama
		indexA = seqSearchNamaA(*A, n, nama)
		fmt.Print("Masukkan Jumlah Barang: ")
		temp = B[indexB].jumlahBarang
		fmt.Scan(&B[indexB].jumlahBarang)
		nama = B[indexB].nama
		indexA = seqSearchNamaA(*A, n, nama)
		if temp > B[indexB].jumlahBarang {
			A[indexA].stok = A[indexA].stok + temp - B[indexB].jumlahBarang
		} else if temp < B[indexB].jumlahBarang {
			A[indexA].stok = A[indexA].stok + temp - B[indexB].jumlahBarang
		} else {
			A[indexA].stok = B[indexB].jumlahBarang
		}
	}
}

func seqSearchNamaB(A arrTransaksi, n int, x string) int {
	//Mencari index berdasarkan nama barang
	var i, index int
	i = -1
	for i = 0; i < n; i++ {
		if x == A[i].nama {
			index = i
		}
	}
	return index
}

func deleteTransaksi(A *arrBarang, n int, B *arrTransaksi, k *int) {
	var i, indexA, indexB, Qty int
	var nama string
	fmt.Println("Masukkan ID Barang yang ingin dihapus data barangnya:")
	indexA = binSearchID(A, n)
	nama = A[indexA].nama
	indexB = seqSearchNamaB(*B, *k, nama)
	Qty = B[indexB].jumlahBarang
	A[indexA].stok += Qty
	for i = indexB; i < n; i++ {
		B[i] = B[i+1]
	}
}

func binSearchID(A *arrBarang, n int) int {
	var i, j, minIdx, idx, kiri, kanan, tengah, index int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].idBarang < A[minIdx].idBarang {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
	fmt.Scan(&idx)
	kiri, kanan = 0, n-1
	for i = 0; i < n; i++ {
		tengah = (kiri + kanan) / 2
		if A[tengah].idBarang < idx {
			kiri = tengah + 1
		} else if A[tengah].idBarang > idx {
			kanan = tengah - 1
		} else {
			index = tengah
		}
	}
	return index
}

func income(A arrBarang, n int, B arrTransaksi, k int) {
	var total, i, indexA int
	var nama string
	for i = 0; i < k; i++ {
		nama = B[i].nama
		indexA = seqSearchNamaA(A, n, nama)
		total = total + (A[indexA].harga * B[i].jumlahBarang)
	}
	fmt.Println("Total Pendapatan :", total)
}

func top5(B *arrTransaksi, n int) {
	var i, j int
	var temp BarangTransaksi
	for i = 1; i < n; i++ {
		temp = B[i]
		j = i - 1
		for j >= 0 && B[j].jumlahBarang < temp.jumlahBarang {
			B[j+1] = B[j]
			j = j - 1
		}
		B[j+1] = temp
	}
	fmt.Println("5 Barang dengan penjualan terbanyak")
	fmt.Printf("%15s %15s %15s\n", "ID Barang", "Nama Barang", "Jumlah")
	for i = 0; i < 5; i++ {
		fmt.Printf("%15d %15s %15d\n", B[i].idBarang, B[i].nama, B[i].jumlahBarang)
	}
}
