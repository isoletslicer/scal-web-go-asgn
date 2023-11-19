package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func findFriend(queriedName string, friendList []Friend) (int, Friend) {
	// cari lalu return index & data teman tsbt
	for index, friend := range friendList {
		if friend.Nama == queriedName {
			return index, friend
		}
	}
	// jaga kalo ga nemu, returning obj kosong
	fmt.Println("Nama teman", queriedName, "tidak ditemukan")
	return -1, Friend{}
}

func displayDetail(index int, friend Friend) {
	fmt.Println("ID:", index)
	fmt.Println("Nama:", friend.Nama)
	fmt.Println("Alamat:", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan:", friend.Alasan)
}

func main() {
	// pengecekan validasi argument inputan
	if len(os.Args) < 2 {
		fmt.Println("Masukan perintah: go run biodata.go <nama>")
		fmt.Println("Note: nama tanpa <>")
		return
	}

	friendList := []Friend{
		{"Thomas", "Shinjiku", "Sekuriti", "Belajar dikala gabut"},
		{"Edward", "Liverpool", "Researcher", "Belajar coba mendalami kodingan"},
		{"Budi", "Jakarta", "Pengangguran", "Belajar buat cari kerja golang mulu info loker"},
		{"Ariel", "Chiapas", "OB", "Belajar buat banting setir untuk hidup lebih baik"},
		{"Mikel", "London", "Mahasiswa", "Belajar buat cari magang saat masih muda"},
	}

	parsedFriendName := os.Args[1]
	index, friend := findFriend(parsedFriendName, friendList)
	if friend.Nama != "" {
		displayDetail(index, friend)
	}
}
