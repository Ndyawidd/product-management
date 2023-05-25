package main

import (
	"fmt"
	"os"
)

const NMAX = 1000

type Product struct {
	id, name   string
	price, qty int
}

type Products struct {
	data      [NMAX]Product
	totalData int
}

// Query
func findIdxProduct(T Products, name string) int {
	for i := 0; i < T.totalData; i++ {
		if T.data[i].name == name {
			return i
		}
	}
	return -1
}

// Commands
func insertProduct(T *Products) {
	var product Product
	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&product.name)
	var idx = findIdxProduct(*T, product.name)
	if idx == -1 {
		product.id = string(rune(T.totalData))
		fmt.Print("Masukkan jumlah barang: ")
		fmt.Scan(&product.qty)
		fmt.Print("Masukkan harga barang: ")
		fmt.Scan(&product.price)
		T.data[T.totalData] = product
		T.totalData++
	} else {
		product = T.data[idx]
		fmt.Print("Masukkan jumlah barang: ")
		fmt.Scan(&product.qty)
	}
}

func updateProduct(T *Products, idx int, product Product) {
	T.data[idx] = product
}

func deleteProduct(T *Products, idx int) {
	T.totalData--
	if idx == T.totalData {
		T.data[idx] = Product{}
	} else {
		for i := idx; i < T.totalData; i++ {
			T.data[i] = T.data[i+1]
		}
	}
}

func main() {
	var T Products
	var idx int
	var selectedMenu string
	var product Product
	for true {
		fmt.Printf("-----\tMenu Inventarisasi Barang\t-----\n1. Tambahkan Barang\n2. Tampilkan Barang\n3. Ubah Barang\n4. Hapus Barang\n5. Cari Barang\n0. Keluar\n")
		for true {
			fmt.Print("Masukkan pilihan menu: ")
			fmt.Scan(&selectedMenu)
			if selectedMenu == "0" || selectedMenu == "1" || selectedMenu == "2" || selectedMenu == "3" || selectedMenu == "4" || selectedMenu == "5" {
				break
			}
		}
		switch selectedMenu {
		case "1":
			insertProduct(&T)
		case "2":
			// Show All data Product
		case "3":
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&product.name)
			idx = findIdxProduct(T, product.name)
			if idx >= 0 {
				product = T.data[idx]
				fmt.Print("Masukkan harga barang: ")
				fmt.Scan(&product.price)
				updateProduct(&T, idx, product)
				fmt.Print("------\tSuccess Mengubah Barang\t------\n")
			} else {
				fmt.Print("------\tData Tidak Ditemukan\t------\n")
			}
		case "4":
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&product.name)
			idx = findIdxProduct(T, product.name)
			if idx >= 0 {
				deleteProduct(&T, idx)
				fmt.Print("------\tSuccess Menghapus Barang\t------\n")
			} else {
				fmt.Print("------\tData Tidak Ditemukan\t------\n")
			}
		case "5":
			//Find Product
		}
		if selectedMenu == "0" {
			fmt.Println("Terimakasih")
			os.Exit(0)
		}
	}
}
