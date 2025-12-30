package crud

import (
	"bufio"
	"cli-library-project-akatsuki/dto/member_dto"
	"cli-library-project-akatsuki/service"
	"cli-library-project-akatsuki/util"
	"fmt"
	"strings"
)

/*
AddMember meneangani input user
untuk menembah member baru
*/
func AddMember(reader *bufio.Reader) {
	fmt.Print("Nama : ")
	nama, _ := reader.ReadString('\n')
	fmt.Print("Patner : ")
	patner, _ := reader.ReadString('\n')
	fmt.Print("Rank : ")
	rank, _ := reader.ReadString('\n')

	req := member_dto.CreateMemberRequestUser{
		Nama:   strings.TrimSpace(nama),
		Patner: strings.TrimSpace(patner),
		Rank:   strings.TrimSpace(rank),
	}

	if err := service.AddMember(req); err != nil {
		fmt.Println("Gagal: ", err.Error())
	}

	fmt.Println("add member succesfully ✅")
}

/*
menuListMember menampilkan seluruh data member
menangani input user view all member
*/
func ViewAllMember() {
	members := service.GetAllMember()

	if len(members) == 0 {
		fmt.Println("Belum ada member")
		return
	}

	util.HeaderViewAllMmeber()

	for _, value := range members {
		status := "Aktif"
		if value.Status {
			status = "Nonaktif"
		}

		fmt.Printf(
			" %-4d | %-15s | %-15s | %-10s\n",
			value.ID,
			value.Name,
			value.Partner,
			status,
		)
	}

	util.FooterViewAllMmeber(len(members))
}
