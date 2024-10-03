package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/Users/daffa/OneDrive/Documents/INTERN/TLKM/golang-workouts/golang-workouts/Pemrograman Go Dasar/A50 File/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// membuat file baru
func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> file berhasil dibuat", path)
}

// mengedit file
func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil diisi")
}

// membaca file
func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				return
			}
		}
		if n == 0 {
			break
		}

		fmt.Println("==> file berhasil dibaca")
		fmt.Println(string(text[:n])) // hanya mencetak bagian yang terbaca
	}
}

// Menghapus File
func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("==> file berhasil dihapus")
}

// Menu
func menu() {
	var choose int
	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Buat file")
		fmt.Println("2. Tulis ke file")
		fmt.Println("3. Baca file")
		fmt.Println("4. Hapus file")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choose)

		switch choose {
		case 1:
			createFile()
		case 2:
			writeFile()
		case 3:
			readFile()
		case 4:
			deleteFile()
		case 5:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func main() {
	menu()
}
