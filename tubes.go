package main

import (
	"fmt"
	"os"

)

type port struct {
	tanggalCatat string
	namaSite     string
	nomorRak     int
	nomorSlot    int
	namaBoard    string
	versiBoard   float64
	inv_id       string
}

type arrPort [2048]port

func main() {
	var menu int
	var tabPort arrPort
	var n int
	for {
		fmt.Println("=========================")
		fmt.Println("Aplikasi Manajemen Board ")
		fmt.Println("1. Tambah Data Port      ")
		fmt.Println("2. Delete Data Port      ")
		fmt.Println("3. Update Data Port      ")
		fmt.Println("4. Cari Data Port Berdasarkan Id        ")
		fmt.Println("5. Cari Data Port Berdasarkan Nama Site")
		fmt.Println("6. Show Data Port        ")
		fmt.Println("7. Exit                  ")
		fmt.Println("======================== ")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			inputData(&tabPort, &n)
			showData(&tabPort, n)
		case 2:
			deleteData(&tabPort, &n)
		case 3:
			mengubahData(&tabPort, n)
		case 4:
			mencariData(&tabPort, n)
		case 5:
			cariNama(&tabPort,n)
		case 6:
			selectionSort(&tabPort, n)
			showData(&tabPort, n)
		case 7:
			fmt.Println("Have A Nice Day!")
			os.Exit(0)
		}
	}
}

func inputData(T *arrPort, n *int) {
	fmt.Scan(&T[*n].inv_id)
	*n = 0
	for T[*n].inv_id != "STOP" {
		fmt.Scan(&T[*n].tanggalCatat, &T[*n].namaSite, &T[*n].nomorRak, &T[*n].nomorSlot, &T[*n].namaBoard, &T[*n].versiBoard)
		*n++
		fmt.Scan(&T[*n].inv_id)
	}
}

func deleteData(board *arrPort, n *int) {
	var edit string
	fmt.Println("Data yang ingin di hapus :")
	fmt.Scan(&edit)
	
		if searchSeq(board, *n, edit) != -1{
			idx := searchSeq(board, *n, edit)
			for j := idx; j < *n-1; j++ {
				board[j] = board[j+1]
			}

			board[*n-1] = port{}
			*n--

			fmt.Println("Data Port Berhasil Dihapus!")
			return
		}
	
}

func mengubahData(board *arrPort, n int) {
	var i int
	var edit string
	var nama_site, nama_board string
	var nomor_rak, nomor_slot int
	var tanggal_catat string
	var versi_board float64

	fmt.Println("Data yang ingin di rubah :")
	fmt.Scan(&edit)
	if searchSeq(board, n, edit) != -1 {
		i = searchSeq(board, n, edit)
		fmt.Println("Update Tanggal Catat :")
		fmt.Scan(&tanggal_catat)
		fmt.Println("Update Nama Site :")
		fmt.Scan(&nama_site)
		fmt.Println("Update Nomor Rak :")
		fmt.Scan(&nomor_rak)
		fmt.Println("Update Nomor Slot :")
		fmt.Scan(&nomor_slot)
		fmt.Println("Update Nama Board: ")
		fmt.Scan(&nama_board)
		fmt.Println("Update Versi Board: ")
		fmt.Scan(&versi_board)

		board[i].tanggalCatat = tanggal_catat
		board[i].namaSite = nama_site
		board[i].nomorRak = nomor_rak
		board[i].nomorSlot = nomor_slot
		board[i].namaBoard = nama_board
		board[i].versiBoard = versi_board

		fmt.Println("Data Port berhasil di update!")
		return
	} else {
		fmt.Println("Data Port Tidak Terdaftar")
	}
}

func mencariData(board *arrPort, n int) {
	var search string
	var idx int
	fmt.Println("Masukkan Kode Invoice :")
	fmt.Scan(&search)
	for i := 0; i < n; i++ {
		if search == board[i].inv_id {
			idx = i
			fmt.Println("Invoice ID : ", board[idx].inv_id)
			fmt.Println("Tanggal Catat :", board[idx].tanggalCatat)
			fmt.Println("Nama Site :", board[idx].namaSite)
			fmt.Println("Nomor Rak Board :", board[idx].nomorRak)
			fmt.Println("Nomor Slot Board :", board[idx].nomorSlot)
			fmt.Println("Nama Board :", board[idx].namaBoard)
			fmt.Println("Versi Board :", board[idx].versiBoard)
		}
	}
}

func cariNama(board *arrPort, n int){
	var search string
	var idx int
	fmt.Println("Masukkan Nama Site :")
	fmt.Scan(&search)
	insertionSort(board,n)
	idx = binarySearch(board,n,search)
	if idx != -1 {
		fmt.Println("Invoice ID : ", board[idx].inv_id)
		fmt.Println("Tanggal Catat :", board[idx].tanggalCatat)
		fmt.Println("Nama Site :", board[idx].namaSite)
		fmt.Println("Nomor Rak Board :", board[idx].nomorRak)
		fmt.Println("Nomor Slot Board :", board[idx].nomorSlot)
		fmt.Println("Nama Board :", board[idx].namaBoard)
		fmt.Println("Versi Board :", board[idx].versiBoard)

	}else {
		fmt.Println("Data Tidak Di temukan!!")
	}
}

func showData(board *arrPort, n int) {
	for i := 0; i < n; i++ {
		println("")
		fmt.Println("Invoice ID : ", board[i].inv_id)
		fmt.Println("Tanggal Catat :", board[i].tanggalCatat)
		fmt.Println("Nama Site :", board[i].namaSite)
		fmt.Println("Nomor Rak Board :", board[i].nomorRak)
		fmt.Println("Nomor Slot Board :", board[i].nomorSlot)
		fmt.Println("Nama Board :", board[i].namaBoard)
		fmt.Println("Versi Board :", board[i].versiBoard)
	}
}

func selectionSort(T *arrPort, n int) {
	var idx_min, i, j int
	var t port
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = 1
		for j < n {
			if T[idx_min].inv_id > T[j].inv_id {
				idx_min = j
			}
			j += 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i += 1

	}
}

func insertionSort(T *arrPort, n int) {
	var i,j int
	var temp port
	i = 1
	for i <= n-1{
		j = i
		temp= T[j]
		for j > 0 && temp.namaSite < T[j-1].namaSite{
			T[j] = T[j-1]
			j = j-1
		}
		T[j] = temp
		i = i + 1
	}
}

func searchSeq(T *arrPort, n int, cariInv string) int {
	for i := 0; i < n; i++ {
		if T[i].inv_id == cariInv {
			return i
		}
	}
	return -1
}

func binarySearch(T *arrPort,n int, X string)int{
	var found int = -1
	var med int 
	var kr int = 0
	var kn int = n - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X < T[med].namaSite {
			kn = med - 1
		}else if X > T[med].namaSite{
			kr = med + 1
		}else {
			found = med
		}		
	}
	return found
}