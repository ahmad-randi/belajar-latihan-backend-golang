package cmd

import (
	"bufio"
	"cli-library-project-akatsuki/menu"
	"cli-library-project-akatsuki/util"
	"fmt"
	"os"
	"strings"
)

func RunMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		util.PrintBanner()
		fmt.Println("1. Member Management")
		fmt.Println("2. Mission Management")
		fmt.Println("3. Organization Stats")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		//Routing Menu
		switch choice {
		case "1":
			//Member Management
			menu.MemberMenuAkatsuki(reader)
		case "2":
			//Mission Management
		case "3":
			//Organization Stats
		case "0":
			//Exit Program
			fmt.Println("Perdamaian hanyalah ilusi...")
			fmt.Println("Program CLI Ditutup, trimakasih by Bedul...👋")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
