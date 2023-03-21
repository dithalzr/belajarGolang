package main

import (
	"fmt"
	"os"
)

type friend struct {
	nama    string
	alamat  string
	pekerjaan string
	alasan  string
}

func main() {
	friendList := []friend{
		friend{"Aldi", "Bandung", "Mahasiswa", "Ingin meningkatkan kemampuan programming"},
		friend{"Dena", "Surabaya", "Fresh Graduate", "Mau belajar bahasa pemrograman baru"},
		friend{"Tomy", "Yogyakarta", "Karyawan", "Ingin mengembangkan aplikasi website"},
		friend{"Johan", "Jakarta", "Mahasiswa", "Ingin menambah ilmu programming"},
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go <no>")
		return
	}

	no := args[1]
	for i, f := range friendList {
		if no == fmt.Sprintf("%d", i+1) {
			fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", f.nama, f.alamat, f.pekerjaan, f.alasan)
			return
		}
	}

	fmt.Println("Data teman tidak ditemukan")
}