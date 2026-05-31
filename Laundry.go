package main

import "fmt"

const NMAX int = 1000

type Pelanggan struct {
	ID   string
	Nama string
	noHp string
}
type tabPelanggan [NMAX]Pelanggan

var dataPelanggan tabPelanggan

type Transaksi struct {
	IDPelanggan string
	Berat       float64
	Tanggal     string
	Layanan     string
	TotalHarga  int
}
type tabTransaksi [NMAX]Transaksi

var dataTransaksi tabTransaksi

func main() {
	var selesai bool
	var jumlahPelanggan, jumlahTransaksi int

	fmt.Println("========================================")
	fmt.Println("          LAUNDRY KICAU MANIA           ")
	fmt.Println("   Tugas Besar Algoritma Pemrograman    ")
	fmt.Println("      103032500021 & 103032500038       ")
	fmt.Println("========================================")
	fmt.Println()

	for !selesai {
		selesai = Menu(&jumlahPelanggan, &jumlahTransaksi)
	}
}

func Menu(nP *int, nT *int) bool {
	var pilihan int
	fmt.Println("========================================")
	fmt.Println("               Menu Utama               ")
	fmt.Println("========================================")
	fmt.Println("1. Tampilkan Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Keluar")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menuTampil(nP, nT)
		fmt.Println()
	case 2:
		menuEdit(nP, nT)
		fmt.Println()
	case 3:
		fmt.Println("========================================")
		fmt.Println("   Terima kasih telah menggunakan       ")
		fmt.Println("       Laundry Kicau Mania!             ")
		fmt.Println("========================================")
		return true
	default:
		fmt.Println("Mohon Masukan pilihan yang valid.")
	}
	return false
}

// Menu lanjutan untuk Menampilkan Transaksi
func menuTampil(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========================================")
	fmt.Println("          Menu Tampilkan Data           ")
	fmt.Println("========================================")
	fmt.Println("1. Tampilkan Semua Transaksi")
	fmt.Println("2. Tampilkan Sesuai Kata Kunci")
	fmt.Println("3. Tampilkan Berdasarkan Urutan Transaksi")
	fmt.Println("4. Cari Berdasarkan Total Harga")
	fmt.Println("5. Pendapatan Total")
	fmt.Println("6. Kembali")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tampilSemua(nP, nT)
		fmt.Println()
	case 2:
		tampilKey(nP, nT)
		fmt.Println()
	case 3:
		tampilUrutan(nP, nT)
		fmt.Println()
	case 4:
		binarySearchHarga(nT)
		fmt.Println()
	case 5:
		pendapatanTotal(nT)
		fmt.Println()
	case 6:
		fmt.Println()
		return
	default:
		fmt.Println("Mohon Masukan pilihan yang valid.")
	}
}

// Menu lanjutan untuk Mengedit Transaksi
func menuEdit(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========================================")
	fmt.Println("            Menu Edit Data              ")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Pelanggan")
	fmt.Println("2. Edit Pelanggan")
	fmt.Println("3. Hapus Pelanggan")
	fmt.Println("4. Tambah Transaksi")
	fmt.Println("5. Edit Transaksi")
	fmt.Println("6. Hapus Transaksi")
	fmt.Println("7. Kembali")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println()
		tambahPelanggan(nP)
	case 2:
		fmt.Println()
		editPelanggan(nP)
	case 3:
		fmt.Println()
		hapusPelanggan(nP, nT)
	case 4:
		fmt.Println()
		tambahTransaksi(nP, nT)
	case 5:
		fmt.Println()
		editTransaksi(nP, nT)
	case 6:
		fmt.Println()
		hapusTransaksi(nP, nT)
	case 7:
		fmt.Println()
		return
	default:
		fmt.Println("Mohon masukan pilihan yang valid.")
	}
}

func tambahPelanggan(nP *int) {
	var p Pelanggan

	fmt.Println("=========== Tambah Pelanggan ===========")
	fmt.Print("Nama  : ")
	fmt.Scan(&p.Nama)

	fmt.Print("No Hp : ")
	fmt.Scan(&p.noHp)

	p.ID = p.Nama + p.noHp
	dataPelanggan[*nP] = p
	*nP++
	fmt.Println("Pelanggan berhasil ditambah.")
	fmt.Println("========================================")
	fmt.Println()
}

func editPelanggan(nP *int) {
	var noHp, namaBaru, noBaru string
	var pilihan int
	var ketemu bool

	fmt.Println("=========== Edit Pelanggan ===========")
	fmt.Print("Masukkan no HP: ")
	fmt.Scan(&noHp)
	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].noHp == noHp {
			ketemu = true

			fmt.Println("---------------Data Lama---------------")
			fmt.Println("Nama: ", dataPelanggan[i].Nama)
			fmt.Println("No HP: ", dataPelanggan[i].noHp)
			fmt.Println("---------------------------------------")

			fmt.Println("Data Pelanggan yang ingin diubah?")
			fmt.Println("1. Nama")
			fmt.Println("2. No HP")
			fmt.Println("3. Kembali")
			fmt.Print("Masukan pilihan Anda: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Print("Input nama baru: ")
				fmt.Scan(&namaBaru)
				dataPelanggan[i].Nama = namaBaru
			case 2:
				fmt.Print("Input nomor baru: ")
				fmt.Scan(&noBaru)
				dataPelanggan[i].noHp = noBaru
			case 3:
				return
			default:
				fmt.Println("Mohon masukan pilihan yang valid.")
			}

			fmt.Println()
			fmt.Println("---------------Data Baru---------------")
			fmt.Println("Data Baru")
			fmt.Println("ID    :", dataPelanggan[i].ID)
			fmt.Println("Nama  :", dataPelanggan[i].Nama)
			fmt.Println("No HP :", dataPelanggan[i].noHp)
			fmt.Println("---------------------------------------")
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
	}
	fmt.Println("========================================")
}

func hapusPelanggan(nP *int, nT *int) {
	var noHp string
	var idx int = -1
	var idTarget string

	fmt.Println("=========== Hapus Pelanggan ===========")
	fmt.Print("Masukkan no HP pelanggan yang akan dihapus: ")
	fmt.Scan(&noHp)

	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].noHp == noHp {
			idx = i
			idTarget = dataPelanggan[i].ID
		}
	}

	if idx == -1 {
		fmt.Println("Pelanggan tidak ditemukan.")
		fmt.Println("========================================")
		return
	}
	for i := 0; i < *nT; i++ {
		if dataTransaksi[i].IDPelanggan == idTarget {
			dataTransaksi[i].IDPelanggan = "ANON"
		}
	}

	for i := idx; i < *nP-1; i++ {
		dataPelanggan[i] = dataPelanggan[i+1]
	}

	dataPelanggan[*nP-1] = Pelanggan{}
	*nP--

	fmt.Println("Pelanggan berhasil dihapus.")
	fmt.Println("========================================")

}

func tambahTransaksi(nP *int, nT *int) {
	var t Transaksi
	var noHp string
	var ketemu bool

	fmt.Println("=========== Tambah Transaksi ===========")
	fmt.Print("Masukkan no HP            : ")
	fmt.Scan(&noHp)

	for i := 0; i < *nP && !ketemu; i++ {
		if dataPelanggan[i].noHp == noHp {
			t.IDPelanggan = dataPelanggan[i].ID
			ketemu = true
			fmt.Print("Berat                     : ")
			fmt.Scan(&t.Berat)
			fmt.Print("Layanan (Express/Reguler) : ")
			fmt.Scan(&t.Layanan)
			for t.Layanan != "Express" && t.Layanan != "Reguler" {
				fmt.Print("Mohon masukkan Express atau Reguler: ")
				fmt.Scan(&t.Layanan)
			}
			fmt.Print("Tanggal (dd/mm/yy)        : ")
			fmt.Scan(&t.Tanggal)

			t.TotalHarga = hitungHarga(t.Berat, t.Layanan)

			dataTransaksi[*nT] = t
			*nT++
			fmt.Println("Transaksi berhasil ditambah")
			fmt.Println("========================================")
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
		tambahPelanggan(nP)
		tambahTransaksi(nP, nT)
	}
}

func hitungHarga(berat float64, layanan string) int {
	if layanan == "Express" {
		return int(berat * 9000)
	}
	return int(berat * 6000)
}

func editTransaksi(nP *int, nT *int) {
	var namaCari, tanggalCari string
	var idPelanggan string
	var idxTrans int = -1
	var pilihan int
	var ketemuPelanggan bool

	fmt.Println("============ Edit Transaksi ============")
	fmt.Print("Masukkan Nama Pelanggan    : ")
	fmt.Scan(&namaCari)
	fmt.Print("Masukkan Tanggal Transaksi : ")
	fmt.Scan(&tanggalCari)

	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].Nama == namaCari {
			idPelanggan = dataPelanggan[i].ID
			ketemuPelanggan = true
		}
	}

	if !ketemuPelanggan {
		fmt.Println("Pelanggan dengan nama tersebut tidak ditemukan.")
		fmt.Println("========================================")
		return
	}

	for i := 0; i < *nT; i++ {
		if dataTransaksi[i].IDPelanggan == idPelanggan && dataTransaksi[i].Tanggal == tanggalCari {
			idxTrans = i
		}
	}

	if idxTrans == -1 {
		fmt.Println("Transaksi tidak ditemukan untuk pelanggan dan tanggal tersebut.")
		fmt.Println("========================================")
		return
	}

	fmt.Println("--------Data Transaksi Ditemukan--------")
	fmt.Println("Nama Pelanggan :", namaCari)
	fmt.Println("Berat Lama     :", dataTransaksi[idxTrans].Berat, "kg")
	fmt.Println("Layanan Lama   :", dataTransaksi[idxTrans].Layanan)
	fmt.Println("Tanggal Lama   :", dataTransaksi[idxTrans].Tanggal)
	fmt.Println("Total Harga    :", dataTransaksi[idxTrans].TotalHarga)
	fmt.Println("----------------------------------------")

	fmt.Println("Data Transaksi yang ingin diubah: ")
	fmt.Println("1. Berat")
	fmt.Println("2. Layanan")
	fmt.Println("3. Kembali")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Berat Baru: ")
		fmt.Scan(&dataTransaksi[idxTrans].Berat)
		dataTransaksi[idxTrans].TotalHarga = hitungHarga(dataTransaksi[idxTrans].Berat, dataTransaksi[idxTrans].Layanan)
		fmt.Println("Berat berhasil diubah.")
	case 2:
		fmt.Print("Masukkan Layanan Baru (Express/Reguler): ")
		fmt.Scan(&dataTransaksi[idxTrans].Layanan)
		for dataTransaksi[idxTrans].Layanan != "Express" && dataTransaksi[idxTrans].Layanan != "Reguler" {
			fmt.Print("Mohon masukkan Express atau Reguler: ")
			fmt.Scan(&dataTransaksi[idxTrans].Layanan)
		}
		dataTransaksi[idxTrans].TotalHarga = hitungHarga(dataTransaksi[idxTrans].Berat, dataTransaksi[idxTrans].Layanan)
		fmt.Println("Layanan berhasil diubah.")
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("---------Data Telah Diperbarui---------")
	fmt.Println("Berat Baru     :", dataTransaksi[idxTrans].Berat, "kg")
	fmt.Println("Layanan Baru   :", dataTransaksi[idxTrans].Layanan)
	fmt.Println("Tanggal Baru   :", dataTransaksi[idxTrans].Tanggal)
	fmt.Println("Total Harga    :", dataTransaksi[idxTrans].TotalHarga)
	fmt.Println("----------------------------------------")
	fmt.Println("   Transaksi berhasil diperbarui! ")
	fmt.Println("========================================")
}

func hapusTransaksi(nP *int, nT *int) {
	var noCari, tanggalCari, idPelanggan string
	var ketemuPelanggan bool
	var idx int = -1

	fmt.Println("=========== Hapus Transaksi ===========")
	fmt.Print("Masukan No Pelanggan      : ")
	fmt.Scan(&noCari)
	fmt.Print("Masukan Tanggal Transaksi : ")
	fmt.Scan(&tanggalCari)

	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].noHp == noCari {
			idPelanggan = dataPelanggan[i].ID
			ketemuPelanggan = true
		}
	}

	if !ketemuPelanggan {
		fmt.Println("Pelanggan tidak ditemukan.")
		fmt.Println("========================================")
		return
	}

	for i := 0; i < *nT; i++ {
		if dataTransaksi[i].IDPelanggan == idPelanggan && dataTransaksi[i].Tanggal == tanggalCari {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		fmt.Println("========================================")
		return
	}

	for i := idx; i < *nT-1; i++ {
		dataTransaksi[i] = dataTransaksi[i+1]
	}
	dataTransaksi[*nT-1] = Transaksi{}
	*nT--
	fmt.Println("Transaksi berhasil dihapus.")
	fmt.Println("========================================")
}

func tampilSemua(nP *int, nT *int) {
	if *nT == 0 {
		fmt.Println("Tidak ada transaksi, silahkan tambahkan transaksi terlebih dahulu")
		fmt.Println()
		return
	}
	fmt.Println("=========== Semua Transaksi ===========")
	for i := 0; i < *nT; i++ {
		if dataTransaksi[i].IDPelanggan == "ANON" {
			fmt.Println("----------------------------------------")
			fmt.Println("ID Pelanggan : ANON")
			fmt.Println("Nama         : Anonymous")
			fmt.Println("No HP        : -")
			fmt.Println("Berat        :", dataTransaksi[i].Berat)
			fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
			fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
			fmt.Println("Total Harga  :", dataTransaksi[i].TotalHarga)
			fmt.Println("----------------------------------------")

		} else {
			for j := 0; j < *nP; j++ {
				if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
					fmt.Println("----------------------------------------")
					fmt.Println("ID Pelanggan :", dataTransaksi[i].IDPelanggan)
					fmt.Println("Nama         :", dataPelanggan[j].Nama)
					fmt.Println("No HP        :", dataPelanggan[j].noHp)
					fmt.Println("Berat        :", dataTransaksi[i].Berat)
					fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
					fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
					fmt.Println("Total Harga  :", dataTransaksi[i].TotalHarga)
					fmt.Println("----------------------------------------")
				}
			}
		}
		fmt.Println("========================================")
	}
}

func tampilKey(nP *int, nT *int) {
	var key string
	var found bool
	fmt.Println("===== Cari Berdasarkan Kata Kunci =====")
	fmt.Print("Masukkan Kata Kunci (No HP/Tanggal): ")
	fmt.Scan(&key)

	fmt.Println("------------Hasil Pencarian------------")
	for i := 0; i < *nT; i++ {
		idx := -1
		for j := 0; j < *nP; j++ {
			if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
				idx = j
			}
		}
		if dataTransaksi[i].Tanggal == key || (idx != -1 && dataPelanggan[idx].noHp == key) {
			found = true
			fmt.Println("Tanggal     :", dataTransaksi[i].Tanggal)
			fmt.Println("ID          :", dataTransaksi[i].IDPelanggan)
			fmt.Println("Nama        :", dataPelanggan[idx].Nama)
			fmt.Println("Total Harga :", dataTransaksi[i].TotalHarga)
			fmt.Println("-----------------------------------------")
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
	fmt.Println("========================================")
}

func pendapatanTotal(nT *int) {
	total := hitungTotalRekursif(0, *nT)
	fmt.Println("=========== Total Pendapatan ===========")
	fmt.Println("Jumlah transaksi         :", *nT)
	fmt.Println("Total seluruh pendapatan :", total)
	fmt.Println("========================================")
}

func hitungTotalRekursif(i int, n int) int {
	if i >= n {
		return 0
	}
	return dataTransaksi[i].TotalHarga + hitungTotalRekursif(i+1, n)
}

func tampilUrutan(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========== Urutkan Transaksi ==========")
	fmt.Println("1. Urutkan dari yang terkecil ke terbesar")
	fmt.Println("2. Urutkan dari yang terbesar ke terkecil")
	fmt.Println("=======================================")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		selectionSortAscending(nT)
		fmt.Println("---------Data Ascending---------")
		tampilSemua(nP, nT)
	case 2:
		insertionSortDescending(nT)
		fmt.Println("---------Data Descending---------")
		tampilSemua(nP, nT)
	default:
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("========================================")
	}
}

func selectionSortAscending(nT *int) {
	var pass, idx, i int
	var temp Transaksi
	pass = 0
	for pass < *nT-1 {
		idx = pass
		i = pass + 1
		for i < *nT {
			if dataTransaksi[idx].TotalHarga > dataTransaksi[i].TotalHarga {
				idx = i
			}
			i++
		}
		temp = dataTransaksi[pass]
		dataTransaksi[pass] = dataTransaksi[idx]
		dataTransaksi[idx] = temp
		pass++
	}
}

func insertionSortDescending(nT *int) {
	var pass, i int
	var temp Transaksi
	pass = 1
	for pass < *nT {
		temp = dataTransaksi[pass]
		i = pass - 1
		for i >= 0 && temp.TotalHarga > dataTransaksi[i].TotalHarga {
			dataTransaksi[i+1] = dataTransaksi[i]
			i--
		}
		dataTransaksi[i+1] = temp
		pass++
	}
}

func binarySearchHarga(nT *int) {
	selectionSortAscending(nT)
	var cari int
	var left, right, mid int
	var ketemu bool

	fmt.Println("===== Cari Berdasarkan Total Harga =====")
	fmt.Print("Masukkan total harga yang dicari: ")
	fmt.Scan(&cari)

	left = 0
	right = *nT - 1

	for left <= right && !ketemu {
		mid = (left + right) / 2

		if dataTransaksi[mid].TotalHarga == cari {
			ketemu = true
			fmt.Println("-------------Data Ditemukan-------------")
			fmt.Println("ID Pelanggan :", dataTransaksi[mid].IDPelanggan)
			fmt.Println("Tanggal      :", dataTransaksi[mid].Tanggal)
			fmt.Println("Layanan      :", dataTransaksi[mid].Layanan)
			fmt.Println("Total Harga  :", dataTransaksi[mid].TotalHarga)
			fmt.Println("----------------------------------------")
		} else if cari < dataTransaksi[mid].TotalHarga {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
	fmt.Println("========================================")
}
