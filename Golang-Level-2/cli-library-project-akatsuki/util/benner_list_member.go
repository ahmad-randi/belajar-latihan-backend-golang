package util

import "fmt"

func HeaderViewAllMmeber() {
	// Header tabel
	fmt.Println("================================================================")
	fmt.Printf(" %-4s | %-15s | %-15s | %-10s\n", "ID", "NAME", "PARTNER", "STATUS")
	fmt.Println("================================================================")
}

func FooterViewAllMmeber(total int) {
	// Header tabel
	fmt.Println("================================================================")
	fmt.Printf("Total Member: %d\n", total)
}
