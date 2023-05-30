package main

import (
	"fmt"
	"os"
	"strconv"
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

type Transaction struct {
	id         string
	products   Products
	status     string
	totalPrice int
}

type Transactions struct {
	data      [NMAX]Transaction
	totalData int
}

// Query
func findIdxProduct(P Products, name string) int {
	for i := 0; i < P.totalData; i++ {
		if P.data[i].name == name {
			return i
		}
	}
	return -1
}

func findAllTransactions(T Transactions) {
	for i := 0; i < T.totalData; i++ {
		fmt.Printf("============ ID : %v ============\nTotal Produk: %v\n", T.data[i].id, T.data[i].products.totalData)
		fmt.Print("Produk: [\n")
		// findAllProduct(T.data[i].products)
		if T.data[i].status == "OUT" {
			fmt.Printf("]\nTotal Harga: %d\n==========\n", T.data[i].totalPrice)
		} else {
			fmt.Print("]\n==========\n")
		}
	}
}

// Commands
func insertProduct(P *Products) {
	var product Product
	product.id = "PRD" + "0"
	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&product.name)
	var idx = findIdxProduct(*P, product.name)
	if idx == -1 {
		if P.totalData > 0 {
			id, err := strconv.Atoi(P.data[P.totalData-1].id[3:])
			if err != nil {
				return
			}
			product.id = "PRD" + strconv.Itoa(id+1)
		}
		fmt.Print("Masukkan jumlah barang: ")
		fmt.Scan(&product.qty)
		fmt.Print("Masukkan harga barang: ")
		fmt.Scan(&product.price)
		P.data[P.totalData] = product
		P.totalData++
	} else {
		product = P.data[idx]
		fmt.Print("Masukkan jumlah barang: ")
		fmt.Scan(&product.qty)
	}
}

func updateProduct(P *Products, idx int, product Product) {
	P.data[idx] = product
}

func deleteProduct(P *Products, idx int) {
	P.totalData--
	if idx == P.totalData {
		P.data[idx] = Product{}
	} else {
		for i := idx; i < P.totalData; i++ {
			P.data[i] = P.data[i+1]
		}
		P.data[P.totalData] = Product{}
	}
}

func updateQtyManyProducts(P *Products, selectedProducts Products) {
	for i := 0; i < selectedProducts.totalData; i++ {
		name := selectedProducts.data[i].name
		idx := findIdxProduct(*P, name)
		isExistData := P.data[idx]
		isExistData.qty -= selectedProducts.data[i].qty
		updateProduct(P, idx, isExistData)
	}
}

func insertProductTransaction(P Products, selectedProducts *Products) {
	var product Product
	fmt.Printf("# Selesai\n* Batalkan\n")
	for true {
		fmt.Print("Masukkan nama barang: ")
		fmt.Scan(&product.name)
		if product.name == "#" || product.name == "*" {
			if product.name == "*" {
				selectedProducts.totalData = 0
			}
			break
		}
		var idx = findIdxProduct(P, product.name)
		if idx == -1 {
			fmt.Println("Masukkan Data yang benar")
		} else {
			product = P.data[idx]
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&product.qty)
			if product.qty <= P.data[idx].qty {
				P.data[idx].qty -= product.qty
				var idxSelected = findIdxProduct(*selectedProducts, product.name)
				if idxSelected >= 0 {
					selectedProducts.data[idxSelected].qty += product.qty
				} else {
					selectedProducts.data[selectedProducts.totalData] = product
					selectedProducts.totalData++
				}
			} else {
				fmt.Println("Stok kurang")
			}
		}
	}
}

func insertTransaction(T *Transactions, P Products, status string) {
	ctx := "TRS"
	var transaction Transaction
	var totalPrice int
	var costumedId string = ctx + "0"
	if T.totalData > 0 {
		id, err := strconv.Atoi(T.data[T.totalData-1].id[3:])
		if err != nil {
			return
		}
		costumedId = ctx + strconv.Itoa(id+1)
	}

	if status == "OUT" {
		totalPrice = calculatePrice(P)
	}

	transaction = Transaction{id: costumedId, products: P, status: status, totalPrice: totalPrice}

	T.data[T.totalData] = transaction
	T.totalData++
}

// Helper
func calculatePrice(P Products) int {
	var totalPrice int
	for i := 0; i < P.totalData; i++ {
		totalPrice += P.data[i].price * P.data[i].qty
	}
	return totalPrice
}

// Menu
func managementMenu(P *Products, T *Transactions) {
	var idx int
	var selectedMenu string
	var product Product
	for true {
		fmt.Printf("-----\tMenu Inventarisasi Barang\t-----\n1. Tambahkan Barang\n2. Tampilkan Barang\n3. Ubah Barang\n4. Hapus Barang\n5. Cari Barang\n0. Kembali\n")
		for true {
			fmt.Print("Masukkan pilihan menu: ")
			fmt.Scan(&selectedMenu)
			if selectedMenu == "0" || selectedMenu == "1" || selectedMenu == "2" || selectedMenu == "3" || selectedMenu == "4" || selectedMenu == "5" {
				break
			}
		}
		switch selectedMenu {
		case "1":
			insertProduct(P)
			insertTransaction(T, *P, "IN")
		case "2":
			// Show All data Product
		case "3":
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&product.name)
			idx = findIdxProduct(*P, product.name)
			if idx >= 0 {
				product = P.data[idx]
				fmt.Print("Masukkan harga barang: ")
				fmt.Scan(&product.price)
				updateProduct(P, idx, product)
				fmt.Print("------\tSuccess Mengubah Barang\t------\n")
			} else {
				fmt.Print("------\tData Tidak Ditemukan\t------\n")
			}
		case "4":
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&product.name)
			idx = findIdxProduct(*P, product.name)
			if idx >= 0 {
				deleteProduct(P, idx)
				fmt.Print("------\tSuccess Menghapus Barang\t------\n")
			} else {
				fmt.Print("------\tData Tidak Ditemukan\t------\n")
			}
		case "5":
			//Find Product
		case "0":
			return
		}
	}
}

func transactionMenu(T *Transactions, P *Products) {
	var selectedMenu string
	var selectedProducts Products
	for true {
		fmt.Printf("===== Menu Transaksi =====\n1. Transaksi\n2. Lihat History Transaksi\n0. Kembali\n========================================\n")
		for true {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&selectedMenu)
			if selectedMenu == "0" || selectedMenu == "1" || selectedMenu == "2" {
				break
			}
		}
		switch selectedMenu {
		case "1":
			nPrev := T.totalData
			insertProductTransaction(*P, &selectedProducts)
			if selectedProducts.totalData > 0 {
				insertTransaction(T, selectedProducts, "OUT")
				updateQtyManyProducts(P, selectedProducts)
				if nPrev == T.totalData {
					fmt.Println("Pembelian Dibatalkan")
				} else {
					fmt.Println("Pembelian Berhasil")
				}
			}
		case "2":
			// function findTransaction
		case "0":
			return
		}
	}
}

func main() {
	var P Products
	var T Transactions
	var selectedMenu string
	for true {
		fmt.Printf("===== Main Menu =====\n1. Management Produk\n2. Transaksi Produk\n0. Keluar\n=====================\n")
		for true {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&selectedMenu)
			if selectedMenu == "0" || selectedMenu == "1" || selectedMenu == "2" {
				break
			}
		}
		switch selectedMenu {
		case "1":
			managementMenu(&P, &T)
		case "2":
			transactionMenu(&T, &P)
		case "0":
			os.Exit(1)
		}
	}
}
