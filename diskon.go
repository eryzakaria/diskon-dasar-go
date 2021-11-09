package main

import "fmt"

func main() {
	var diskon, jb float32

	fmt.Printf("Inputkan Jumlah Barang: ")
	var j float32
	fmt.Scanln(&j)

	fmt.Printf("Inputkan Harga Barang: ")
	var h float32
	fmt.Scanln(&h)

	hrbg := j * h

	fmt.Printf("Jumlah Harga Barang : %.2f\n", hrbg)

	if hrbg >= 50000 {
		diskon = 0.2 * hrbg
		jb = hrbg - diskon
	} else if hrbg < 50000 {
		diskon = 0.1 * hrbg
		jb = hrbg - diskon
	}

	fmt.Printf("Jumlah Diskon : %2.f\n", diskon)
	fmt.Printf("Jumlah Bayar : %2.f\n", jb)
}
