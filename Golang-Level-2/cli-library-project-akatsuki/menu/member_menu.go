package menu

import (
	"bufio"
	"cli-library-project-akatsuki/menu/crud"
	"fmt"
	"strings"
)

func MemberMenuAkatsuki(reader *bufio.Reader) {

	for {
		/// ===== Tampilan Menu Member Akatsuki=====
		fmt.Println("\n===== Menu Member Akatsuki =====")
		fmt.Println("1. Add Member")
		fmt.Println("2. View All Member")
		fmt.Println("3. Update Member")
		fmt.Println("4. Deactivate Member")
		fmt.Println("5. Hapus Data Member")
		fmt.Println("0. Back to Main Menu")
		fmt.Print("Pilih menu member :")
		choiceMember, _ := reader.ReadString('\n')
		choiceMember = strings.TrimSpace(choiceMember)

		// ===== Routing Menu Member=====
		switch choiceMember {
		case "1":
			//Add Member
			crud.AddMembers(reader)
		case "2":
			//View All Member
			crud.ViewAllMembers()
		case "3":
			//Update Member
			crud.UpdateMembers(reader)
		case "4":
			//Deactivate Member
		case "5":
			//Hapus Data Member
		case "0":
			//Back to Main Menu
			return
		}
	}
}
