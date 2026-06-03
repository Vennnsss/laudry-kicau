package main

import "fmt"

const NMAX int = 1000

type Pelanggan struct {
	ID, Nama, noHp string
}
type tabPelanggan [NMAX]Pelanggan

var dataPelanggan tabPelanggan

type Transaksi struct {
	IDPelanggan, Tanggal, Layanan string
	Berat                         float64
	TotalHarga                    int
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
	fmt.Println("2. Tambah Data")
	fmt.Println("3. Edit Data")
	fmt.Println("4. Keluar")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		menuTampil(nP, nT)
		fmt.Println()
	case 2:
		menuTambah(nP, nT)
		fmt.Println()
	case 3:
		menuEdit(nP, nT)
		fmt.Println()
	case 4:
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

func menuTambah(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========================================")
	fmt.Println("           Menu Tambah Data             ")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Pelanggan")
	fmt.Println("2. Tambah Transaksi")
	fmt.Println("3. Kembali")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println()
		tambahPelanggan(nP)
	case 2:
		fmt.Println()
		tambahTransaksi(nP, nT)
	default:
		fmt.Println("Mohon masukan pilihan yang valid.")
	}
}

func menuTampil(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========================================")
	fmt.Println("          Menu Tampilkan Data           ")
	fmt.Println("========================================")
	fmt.Println("1. Tampilkan Semua Transaksi")
	fmt.Println("2. Tampilkan Sesuai Kata Kunci")
	fmt.Println("3. Urutkan Berdasarkan Harga")
	fmt.Println("4. Urutkan Berdasarkan Tanggal")
	fmt.Println("5. Cari Berdasarkan Total Harga")
	fmt.Println("6. Pendapatan Periode")
	fmt.Println("7. Kembali")
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
		urutHarga(nP, nT)
		fmt.Println()
	case 4:
		urutTanggal(nP, nT)
		fmt.Println()
	case 5:
		binarySearchHarga(nT)
		fmt.Println()
	case 6:
		pendapatanPeriode(nT)
	case 7:
		return
	default:
		fmt.Println("Mohon Masukan pilihan yang valid.")
	}
}

func menuEdit(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========================================")
	fmt.Println("            Menu Edit Data              ")
	fmt.Println("========================================")
	fmt.Println("1. Edit Pelanggan")
	fmt.Println("2. Hapus Pelanggan")
	fmt.Println("3. Edit Transaksi")
	fmt.Println("4. Hapus Transaksi")
	fmt.Println("5. Kembali")
	fmt.Println("========================================")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println()
		editPelanggan(nP)
	case 2:
		fmt.Println()
		hapusPelanggan(nP, nT)
	case 3:
		fmt.Println()
		editTransaksi(nP, nT)
	case 4:
		fmt.Println()
		hapusTransaksi(nP, nT)
	case 5:
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
			fmt.Println("Data pelanggan berhasil diubah.")
			fmt.Println("========================================")
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
		fmt.Println("========================================")
	}
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

	*nP--

	fmt.Println("Pelanggan berhasil dihapus.")
	fmt.Println("========================================")
}

func tambahTransaksi(nP *int, nT *int) {
	var t Transaksi
	var noHp string
	var ketemu bool
	var hari, bulan, tahun string

	fmt.Println("=========== Tambah Transaksi ===========")
	fmt.Print("Masukkan no HP            : ")
	fmt.Scan(&noHp)

	for i := 0; i < *nP && !ketemu; i++ {
		if dataPelanggan[i].noHp == noHp {
			ketemu = true
			t.IDPelanggan = dataPelanggan[i].ID

			fmt.Print("Berat                     : ")
			fmt.Scan(&t.Berat)

			fmt.Print("Layanan (Express/Reguler) : ")
			fmt.Scan(&t.Layanan)
			for t.Layanan != "Express" && t.Layanan != "Reguler" {
				fmt.Print("Mohon masukkan Express atau Reguler: ")
				fmt.Scan(&t.Layanan)
			}

			fmt.Print("Tanggal (dd mm yyyy)      : ")
			fmt.Scan(&hari, &bulan, &tahun)

			t.Tanggal = tahun + bulan + hari
			t.TotalHarga = hitungHarga(t.Berat, t.Layanan)

			dataTransaksi[*nT] = t
			*nT++

			fmt.Println("Transaksi berhasil ditambah")
			fmt.Println("========================================")
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
		fmt.Println("========================================")
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
	var idx int = -1
	var pilihan int
	var hari, bulan, tahun string

	fmt.Println("============ Edit Transaksi ============")
	fmt.Print("Masukkan Nama Pelanggan                 : ")
	fmt.Scan(&namaCari)
	fmt.Print("Masukkan Tanggal Transaksi (dd mm yyyy) : ")
	fmt.Scan(&hari, &bulan, &tahun)
	tanggalCari = tahun + bulan + hari

	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].Nama == namaCari {
			idPelanggan = dataPelanggan[i].ID
		}
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
	fmt.Println("Data Transaksi yang ingin diubah: ")
	fmt.Println("1. Berat")
	fmt.Println("2. Layanan")
	fmt.Println("3. Kembali")
	fmt.Print("Masukan pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Berat Baru: ")
		fmt.Scan(&dataTransaksi[idx].Berat)
	case 2:
		fmt.Print("Masukkan Layanan Baru (Express/Reguler): ")
		fmt.Scan(&dataTransaksi[idx].Layanan)
		for dataTransaksi[idx].Layanan != "Express" && dataTransaksi[idx].Layanan != "Reguler" {
			fmt.Print("Mohon masukkan Express atau Reguler: ")
			fmt.Scan(&dataTransaksi[idx].Layanan)
		}
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}
	dataTransaksi[idx].TotalHarga = hitungHarga(dataTransaksi[idx].Berat, dataTransaksi[idx].Layanan)

	fmt.Println("Transaksi berhasil diperbarui!")
	fmt.Println("========================================")
}

func hapusTransaksi(nP *int, nT *int) {
	var noCari, tanggalCari, idPelanggan string
	var idx int = -1
	var hari, bulan, tahun string

	fmt.Println("=========== Hapus Transaksi ===========")
	fmt.Print("Masukan No Pelanggan                    : ")
	fmt.Scan(&noCari)
	fmt.Print("Masukkan Tanggal Transaksi (dd mm yyyy) : ")
	fmt.Scan(&hari, &bulan, &tahun)
	tanggalCari = tahun + bulan + hari

	for i := 0; i < *nP; i++ {
		if dataPelanggan[i].noHp == noCari {
			idPelanggan = dataPelanggan[i].ID
		}
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
		nama := "Anonymous"
		nohp := "-"

		for j := 0; j < *nP; j++ {
			if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
				nama = dataPelanggan[j].Nama
				nohp = dataPelanggan[j].noHp
			}
		}
		fmt.Println("----------------------------------------")
		fmt.Println("ID Pelanggan :", dataTransaksi[i].IDPelanggan)
		fmt.Println("Nama         :", nama)
		fmt.Println("No HP        :", nohp)
		fmt.Println("Berat        :", dataTransaksi[i].Berat)
		fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
		fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
		fmt.Println("Total Harga  :", dataTransaksi[i].TotalHarga)
	}
	fmt.Println("========================================")
}

func tampilKey(nP *int, nT *int) {
	var key string
	var found bool
	var pilihan int

	fmt.Println("===== Cari Berdasarkan Kata Kunci =====")
	fmt.Println("1. Cari Nama")
	fmt.Println("2. Cari Tanggal")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama / No HP: ")
		fmt.Scan(&key)
	case 2:
		var hari, bulan, tahun string
		fmt.Print("Masukkan tanggal (dd mm yyyy): ")
		fmt.Scan(&hari, &bulan, &tahun)
		key = tahun + bulan + hari
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	for i := 0; i < *nT; i++ {

		nama := ""
		for j := 0; j < *nP; j++ {
			if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
				nama = dataPelanggan[j].Nama
			}
		}

		if nama == key || dataTransaksi[i].Tanggal == key {
			found = true
			fmt.Println("----------------------------------------")
			fmt.Println("Nama         :", nama)
			fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
			fmt.Println("Total Harga  :", dataTransaksi[i].TotalHarga)
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}

	fmt.Println("========================================")
}

func urutHarga(nP *int, nT *int) {
	var pilihan int

	fmt.Println("========== Urutkan Transaksi ==========")
	fmt.Println("1. Urutkan dari yang terkecil ke terbesar")
	fmt.Println("2. Urutkan dari yang terbesar ke terkecil")
	fmt.Println("=======================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	for i := 1; i < *nT; i++ {
		temp := dataTransaksi[i]
		j := i - 1
		for j >= 0 && dataTransaksi[j].TotalHarga > temp.TotalHarga {
			dataTransaksi[j+1] = dataTransaksi[j]
			j--
		}
		dataTransaksi[j+1] = temp
	}
	switch pilihan {
	case 1:
		tampilSemua(nP, nT)
	case 2:
		for i := 0; i < *nT/2; i++ {
			temp := dataTransaksi[i]
			dataTransaksi[i] = dataTransaksi[*nT-1-i]
			dataTransaksi[*nT-1-i] = temp
		}
		tampilSemua(nP, nT)
	default:
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("========================================")
	}
}

func urutTanggal(nP *int, nT *int) {
	var pilihan int
	fmt.Println("========== Urutkan Transaksi ==========")
	fmt.Println("1. Urutkan dari yang terkecil ke terbesar")
	fmt.Println("2. Urutkan dari yang terbesar ke terkecil")
	fmt.Println("=======================================")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	for i := 0; i < *nT-1; i++ {
		idx := i
		for j := i + 1; j < *nT; j++ {
			if dataTransaksi[j].Tanggal < dataTransaksi[idx].Tanggal {
				idx = j
			}
		}
		temp := dataTransaksi[i]
		dataTransaksi[i] = dataTransaksi[idx]
		dataTransaksi[idx] = temp
	}
	switch pilihan {
	case 1:
		tampilSemua(nP, nT)
	case 2:
		for i := 0; i < *nT/2; i++ {
			temp := dataTransaksi[i]
			dataTransaksi[i] = dataTransaksi[*nT-1-i]
			dataTransaksi[*nT-1-i] = temp
		}
		tampilSemua(nP, nT)
	default:
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("========================================")
	}
}

func binarySearchHarga(nT *int) {
	var cari int
	var left, right, mid int
	var ketemu bool

	for i := 1; i < *nT; i++ {
		temp := dataTransaksi[i]
		j := i - 1

		for j >= 0 && dataTransaksi[j].TotalHarga > temp.TotalHarga {
			dataTransaksi[j+1] = dataTransaksi[j]
			j--
		}

		dataTransaksi[j+1] = temp
	}

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

func pendapatanPeriode(nT *int) {
	var awal, akhir string
	var h1, b1, t1 string
	var h2, b2, t2 string

	fmt.Println("=========== Total Pendapatan ===========")

	fmt.Print("Tanggal awal (dd mm yyyy) : ")
	fmt.Scan(&h1, &b1, &t1)

	fmt.Print("Tanggal akhir (dd mm yyyy): ")
	fmt.Scan(&h2, &b2, &t2)

	awal = t1 + b1 + h1
	akhir = t2 + b2 + h2

	total := hitungTotalRekursif(0, *nT, awal, akhir)

	fmt.Println("Total pendapatan :", total)
	fmt.Println("========================================")
}

func hitungTotalRekursif(i int, n int, awal string, akhir string) int {

	if i >= n {
		return 0
	}

	total := 0

	if dataTransaksi[i].Tanggal >= awal &&
		dataTransaksi[i].Tanggal <= akhir {

		total = dataTransaksi[i].TotalHarga
	}

	return total + hitungTotalRekursif(i+1, n, awal, akhir)
}

// tes buat dicommit
