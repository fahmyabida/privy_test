package main

import "fmt"

func main() {
	a := fungsiA([]string{"aku", "adalah", "seorang", "kapiten"})
	fmt.Println(a)
}

// fungsi A memindahkan slice pada param ke slice baru, dengan isi yang sama namun bisa jadi urutannya berbeda
func fungsiA(slice []string) []string {
	// inisialisasi variable dengan nama fungsi map bertipe map[string]struct{}
	fungsiMap := make(map[string]struct{})
	// mekakukan looping sebanyak length yang ada pada varible 'slice'(bertipe []string)
	for _, v := range slice{
		// memasukkan struct kosong pada varible fungsi map dengan index[v],
		// dimana v merupakan string yang ada didalam slice string
		fungsiMap[v] = struct{}{}
	}

	// inisialisasi variable fungsiSlice dengan panjang slice 0 dan kapsitas sepanjang slice
	fungsiSlice := make([]string, 0, len(slice))
	// me looping fungsiMap
	for v := range fungsiMap {
		// menambahkan v pada variable fungsiSlice
		fungsiSlice = append(fungsiSlice, v)
	}
	return fungsiSlice
}