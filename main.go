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

func findAllTransactions(T Transactions, flag string, isSort, isAsc bool) {
	if isSort {
		insertionSort(&T, flag, isAsc)
	}
	for i := 0; i < T.totalData; i++ {
		fmt.Printf("\n=======\nID : %v\nTotal Produk: %v\n", T.data[i].id, T.data[i].products.totalData)
		fmt.Print("Produk: [\n")
		printProduct(T.data[i].products)
		if T.data[i].status == "OUT" {
			fmt.Printf("]\nTotal Harga: %d\n", T.data[i].totalPrice)
		} else {
			fmt.Print("]\n")
		}
		fmt.Printf("Status: %v\n=======\n", T.data[i].status)
	}
}

func maxMin(T Transactions, isMax bool) {
	var result Transaction
	result = T.data[0]
	for i := 1; i < T.totalData; i++ {
		if isMax && result.totalPrice < T.data[i].totalPrice && T.data[i].status == "OUT" {
			result = T.data[i]
		} else if !isMax && result.totalPrice > T.data[i].totalPrice && T.data[i].status == "OUT" {
			result = T.data[i]
		} else if result.status == "IN" {
			result = T.data[i]
		}
	}
	fmt.Println(result.status)
	if result.status == "OUT" {
		fmt.Printf("============ ID : %v ============\nTotal Produk: %v\n", result.id, result.products.totalData)
		fmt.Print("Produk: [\n")
		printProduct(result.products)
		fmt.Printf("]\nTotal Harga: %d\n==========\n", result.totalPrice)
	}
}

func sort(T *Products, sorting string, flag string) {
	for pass := 1; pass <= T.totalData-1; pass++ {
		idx := pass - 1
		for i := pass; i < T.totalData; i++ {
			switch flag {
			case "1":
				if sorting == "1" && T.data[idx].id > T.data[i].id {
					idx = i
				} else if sorting == "2" && T.data[idx].id < T.data[i].id {
					idx = i
				}
			case "2":
				if sorting == "1" && T.data[idx].name > T.data[i].name {
					idx = i
				} else if sorting == "2" && T.data[idx].name < T.data[i].name {
					idx = i
				}
			case "3":
				if sorting == "1" && T.data[idx].price > T.data[i].price {
					idx = i
				} else if sorting == "2" && T.data[idx].price < T.data[i].price {
					idx = i
				}
			case "4":
				if sorting == "1" && T.data[idx].qty > T.data[i].qty {
					idx = i
				} else if sorting == "2" && T.data[idx].qty < T.data[i].qty {
					idx = i
				}

			}
		}
		T.data[pass-1], T.data[idx] = T.data[idx], T.data[pass-1]
	}
}

func search(T Products, flag string) Products {
	var price, qty int
	var id, name string
	var arrProduct Products
	switch flag {
	case "1":
		fmt.Print("Masukkan ID yang dicari: ")
		fmt.Scan(&id)
		for i := 0; i < T.totalData; i++ {
			if id == T.data[i].id {
				arrProduct.data[arrProduct.totalData] = T.data[i]
				arrProduct.totalData++
			}
		}
	case "2":
		fmt.Print("Masukkan nama yang dicari: ")
		fmt.Scan(&name)
		for i := 0; i < T.totalData; i++ {
			if name == T.data[i].name {
				arrProduct.data[arrProduct.totalData] = T.data[i]
				arrProduct.totalData++
			}
		}
	case "3":
		fmt.Print("Masukkan harga yang dicari: ")
		fmt.Scan(&price)
		for i := 0; i < T.totalData; i++ {
			if price == T.data[i].price {
				arrProduct.data[arrProduct.totalData] = T.data[i]
				arrProduct.totalData++
			}
		}
	case "4":
		fmt.Print("Masukkan jumlah yang dicari: ")
		fmt.Scan(&qty)
		for i := 0; i < T.totalData; i++ {
			if qty == T.data[i].qty {
				arrProduct.data[arrProduct.totalData] = T.data[i]
				arrProduct.totalData++
			}
		}
	}
	return arrProduct
}

func printProduct(P Products) {
	for i := 0; i < P.totalData; i++ {
		fmt.Printf("\nID: %v \nNama: %s \nHarga: %v \nJumlah: %v\n\n", P.data[i].id, P.data[i].name, P.data[i].price, P.data[i].qty)
	}
}

func insertionSort(T *Transactions, flag string, isAsc bool) {
	for pass := 1; pass < T.totalData; pass++ {
		var i int = pass
		var transaction Transaction = T.data[i]
		switch flag {
		case "1":
			if isAsc {
				for i > 0 && T.data[i-1].id > transaction.id {
					T.data[i] = T.data[i-1]
					i--
				}
			} else {
				for i > 0 && T.data[i-1].id < transaction.id {
					T.data[i] = T.data[i-1]
					i--
				}
			}
		case "2":
			if isAsc {
				for i > 0 && (T.data[i-1].totalPrice > transaction.totalPrice || T.data[i-1].status == "IN") {
					T.data[i] = T.data[i-1]
					i--
				}
			} else {
				for i > 0 && T.data[i-1].totalPrice < transaction.totalPrice {
					T.data[i] = T.data[i-1]
					i--
				}
			}
		}
		T.data[i] = transaction
	}
}

func findTransactionByID(T Transactions, id string) Transaction {
	var min, mid, max int
	max = T.totalData - 1
	for min <= max {
		mid = (min + max) / 2
		if T.data[mid].id == id {
			return T.data[mid]
		} else if T.data[mid].id > id {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return Transaction{}
}

func printTransaction(T Transaction) {
	fmt.Printf("\n=======\nID : %v\nTotal Produk: %v\n", T.id, T.products.totalData)
	fmt.Print("Produk: [\n")
	printProduct(T.products)
	if T.status == "OUT" {
		fmt.Printf("]\nTotal Harga: %d\n", T.totalPrice)
	} else {
		fmt.Print("]\n")
	}
	fmt.Printf("Status: %v\n=======\n", T.status)
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
		_, err := fmt.Scanf("%d", &product.qty)
		if err != nil {
			fmt.Println("Data yang anda masukan tidak valid")
			return
		}
		fmt.Print("Masukkan harga barang: ")
		_, err = fmt.Scanf("%d", &product.price)
		if err != nil {
			fmt.Println("Data yang anda masukan tidak valid")
			return
		}
		P.data[P.totalData] = product
		P.totalData++
	} else {
		product = P.data[idx]
		fmt.Print("Masukkan jumlah barang: ")
		_, err := fmt.Scanf("%d", &product.qty)
		if err != nil {
			fmt.Println("Data yang anda masukan tidak valid")
			return
		}
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
			_, err := fmt.Scanf("%d", &product.qty)
			if err != nil {
				break
			}
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
			fmt.Printf("-----\tData Diurutkan Berdasarkan:\t-----\n1. ID\n2. Nama\n3. Harga\n4. Jumlah\n5. Tidak diurutkan\n")
			var s string
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scan(&s)
			if s != "5" {
				fmt.Printf("-----\tCara Pengurutan:\t-----\n1. Ascending\n2. Descending\n")
				var sorting string
				fmt.Print("Masukkan Pilihan: ")
				fmt.Scan(&sorting)
				sort(P, sorting, s)
			}
			printProduct(*P)
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
			fmt.Printf("-----\tCari Berdasarkan:\t-----\n1.ID \n2. Nama \n3. Harga\n4. Jumlah")
			var s string
			fmt.Printf("\nMasukkan pilihan: ")
			fmt.Scan(&s)
			var data Products = search(*P, s)
			printProduct(data)
		case "0":
			return
		}
	}
}

func transactionMenu(T *Transactions, P *Products) {
	var selectedMenu string
	var selectedProducts Products
	for true {
		fmt.Printf("-----\tMenu Transaksi\t-----\n1. Transaksi\n2. Lihat History Transaksi\n3. Data Penjualan\n4. Pencarian Dengan ID\n0. Kembali\n---------------------------	\n")
		for true {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&selectedMenu)
			if selectedMenu == "0" || selectedMenu == "1" || selectedMenu == "2" || selectedMenu == "3" || selectedMenu == "4" {
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
				selectedProducts = Products{}
			}
		case "2":
			var flag, isAsc string
			var isSort bool
			fmt.Printf("Urutkan Data Sesuai: \n1. ID\n2. Total Harga\n3. Tidak diurutkan\nMasukkan pilihan: ")
			for true {
				fmt.Scan(&flag)
				if flag == "1" || flag == "2" {
					fmt.Printf("Mengurutkan dengan\n1. Ascending\n2. Descending\nMasukkan pilihan: ")
					fmt.Scan(&isAsc)
					isSort = true
					break
				} else if flag == "3" {
					isSort = false
					break
				}
			}
			findAllTransactions(*T, flag, isSort, isAsc == "1")
		case "3":
			var s int
			fmt.Printf("----- Pencarian Transaksi -----\n1. Terbesar\n2. Terkecil\n")
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scan(&s)
			maxMin(*T, s == 1)
		case "4":
			var id string
			fmt.Printf("------ Cari Transaksi berdasatkan ID ------\nMasukkan ID: ")
			fmt.Scan(&id)
			isExistData := findTransactionByID(*T, id)
			if isExistData.id != "" {
				printTransaction(isExistData)
			} else {
				fmt.Print("------\tData Tidak Ditemukan\t------\n")
			}
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
