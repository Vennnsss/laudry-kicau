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
var jumlahPelanggan int

type Transaksi struct {
	IDPelanggan string
	Berat       float64
	Tanggal     string
	Layanan     string
	TotalHarga  int
}
type tabTransaksi [NMAX]Transaksi

var dataTransaksi tabTransaksi
var jumlahTransaksi int

func main() {
	var selesai bool

	for !selesai {
		selesai = Menu()
	}
}

func Menu() bool {
	var pilihan int
	fmt.Println("======= Menu Utama =======")
	fmt.Println("1. Tampilkan Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Keluar")
	fmt.Println("==========================")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menuTampil()
	case 2:
		menuEdit()
	case 3:
		return true
	default:
		fmt.Println("Mohon Masukan pillihan yang valid.")
	}
	return false
}

// Menu lanjutan untuk Menampilkan Transaksi
func menuTampil() {
	var pilihan int

	fmt.Println("=========== Menu Tampilkan Data ===========")
	fmt.Println("1. Tampilkan Semua Transaksi")
	fmt.Println("2. Tampilkan Sesuai Kata Kunci")
	fmt.Println("3. Tampilkan Berdasarkan Urutan Transaksi")
	fmt.Println("4. Pendapatan Total")
	fmt.Println("5. Kembali")
	fmt.Println("===========================================")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tampilSemua()
	case 2:
		tampilKey()
	case 3:
		tampilUrutan()
	case 4:
		pendapatanTotal()
	case 5:
		return
	default:
		fmt.Println("Mohon Masukan pillihan yang valid.")
	}

}

// Menu lanjutan untuk Mengedit Transaksi
func menuEdit() {
	var pilihan int

	fmt.Println("=========== Menu Edit Data ===========")
	fmt.Println("1. Tambah Pelanggan")
	fmt.Println("2. Edit Pelanggan")
	fmt.Println("3. Hapus Pelanggan")
	fmt.Println("4. Tambah Transaksi")
	fmt.Println("5. Edit Transaksi")
	fmt.Println("6. Hapus Transaksi")
	fmt.Println("7. Kembali")
	fmt.Println("======================================")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahPelanggan()
	case 2:
		editPelanggan()
	case 3:
		hapusPelanggan()
	case 4:
		tambahTransaksi()
	case 5:
		editTransaksi()
	case 6:
		hapusTransaksi()
	case 7:
		return
	default:
		fmt.Println("Mohon masukan pillihan yang valid.")
	}
}

func tambahPelanggan() {
	var p Pelanggan

	fmt.Print("Nama: ")
	fmt.Scan(&p.Nama)

	fmt.Print("No Hp: ")
	fmt.Scan(&p.noHp)

	p.ID = p.Nama + p.noHp
	dataPelanggan[jumlahPelanggan] = p
	jumlahPelanggan++
	fmt.Println("Pelanggan Berhasil Ditambah")
}

// ini ga perlu ditampilin tapi keep dulu aja
func tampilPelanggan() {
	for i := 0; i < jumlahPelanggan; i++ {
		fmt.Println(dataPelanggan[i].ID)
		fmt.Println(dataPelanggan[i].Nama)
		fmt.Println(dataPelanggan[i].noHp)
	}
}

func editPelanggan() {
	var noHp, namaBaru, noBaru string
	var pilihan int
	var ketemu bool

	fmt.Print("Masukkan no HP: ")
	fmt.Scan(&noHp)
	for i := 0; i < jumlahPelanggan; i++ {
		if dataPelanggan[i].noHp == noHp {
			ketemu = true

			fmt.Println("Data Lama")
			fmt.Println("Nama: ", dataPelanggan[i].Nama)
			fmt.Println("No HP: ", dataPelanggan[i].noHp)
			fmt.Println()

			fmt.Println("Apa yang mau diganti?")
			fmt.Println("1. Nama")
			fmt.Println("2. No HP")
			fmt.Println("3. Kembali")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println("Input nama baru: ")
				fmt.Scan(&namaBaru)
				dataPelanggan[i].Nama = namaBaru
			case 2:
				fmt.Println("Input nomor baru: ")
				fmt.Scan(&noBaru)
				dataPelanggan[i].noHp = noBaru
			case 3:
				return
			default:
				fmt.Println("Mohon masukan pillihan yang valid.")
			}

			fmt.Println()
			fmt.Println("Data Baru")
			fmt.Println("Nama:", dataPelanggan[i].Nama)
			fmt.Println("No HP:", dataPelanggan[i].noHp)
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
	}
}

// HAPUS PELANGGAN

func tambahTransaksi() {
	var t Transaksi
	var noHp string
	var ketemu bool

	fmt.Print("Masukkan no HP: ")
	fmt.Scan(&noHp)

	for i := 0; i < jumlahPelanggan && !ketemu; i++ {
		if dataPelanggan[i].noHp == noHp {
			t.IDPelanggan = dataPelanggan[i].ID
			ketemu = true

			fmt.Print("Berat: ")
			fmt.Scan(&t.Berat)
			fmt.Print("Layanan: ")
			fmt.Scan(&t.Layanan)
			fmt.Print("Tanggal: ")
			fmt.Scan(&t.Tanggal)

			t.TotalHarga = hitungHarga(t.Berat, t.Layanan)

			dataTransaksi[jumlahTransaksi] = t
			jumlahTransaksi++
			fmt.Println("Transaksi berhasil ditambah")
		}
	}
	if !ketemu {
		fmt.Println("Pelanggan tidak ditemukan, masukan pelanggan terlebih dahulu.")
		tambahPelanggan()
		tambahTransaksi()
	}
}

func hitungHarga(berat float64, layanan string) int {
	if layanan == "Express" {
		return int(berat * 9000)
	}
	return int(berat * 6000)
}

//EDIT TRANSAKSI

//HAPUS TRANSAKSI

func tampilSemua() {
	for i := 0; i < jumlahTransaksi; i++ {
		for j := 0; j < jumlahPelanggan; j++ {
			if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
				fmt.Println("ID Pelanggan :", dataTransaksi[i].IDPelanggan)
				fmt.Println("Nama		  :", dataPelanggan[j].Nama)
				fmt.Println("No HP		  :", dataPelanggan[j].noHp)
				fmt.Println("Berat        :", dataTransaksi[i].Berat)
				fmt.Println("Layanan      :", dataTransaksi[i].Layanan)
				fmt.Println("Tanggal      :", dataTransaksi[i].Tanggal)
				fmt.Println("Total Harga  :", dataTransaksi[i].TotalHarga)
				fmt.Println("==============================")
			}
		}
	}
}

func tampilKey() {
	var key string
	var found bool
	fmt.Print("Masukkan Kata Kunci (No HP atau Tanggal): ")
	fmt.Scan(&key)

	fmt.Println("======= Hasil Pencarian =======")
	for i := 0; i < jumlahTransaksi; i++ {
		idx := -1
		for j := 0; j < jumlahPelanggan; j++ {
			if dataPelanggan[j].ID == dataTransaksi[i].IDPelanggan {
				idx = j
			}
		}

		if dataTransaksi[i].Tanggal == key || (idx != -1 && dataPelanggan[idx].noHp == key) {
			found = true
			fmt.Printf("Tgl: %s | ID: %s | Total: %d\n",
				dataTransaksi[i].Tanggal, dataTransaksi[i].IDPelanggan, dataTransaksi[i].TotalHarga)
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

//PENDAPATAN TOTAL

//MENAMPILKAN PAKE SORTING
