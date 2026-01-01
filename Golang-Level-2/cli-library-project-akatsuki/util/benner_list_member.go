package util

import "fmt"

func HeaderViewAllMember() {
	fmt.Println("==============================================================================================")
	fmt.Printf(
		" %-4s | %-15s | %-15s | %-10s | %-12s | %-12s\n",
		"ID",
		"NAME",
		"PARTNER",
		"STATUS",
		"TOTAL MISI",
		"TOTAL REWARD",
	)
	fmt.Println("==============================================================================================")
}

func FooterViewAllMmeber(total int) {
	// Header tabel
	fmt.Println("==============================================================================================")
	fmt.Printf("Total Member: %d\n", total)
}
