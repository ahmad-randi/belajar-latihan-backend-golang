package crud

import (
	"bufio"
	"cli-library-project-akatsuki/dto/member_dto"
	"cli-library-project-akatsuki/service"
	"cli-library-project-akatsuki/util"
	"fmt"
	"strconv"
	"strings"
)

/*
AddMember meneangani input user
untuk menembah member baru
*/
func AddMembers(reader *bufio.Reader) {
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

	if err := service.AddMemberService(req); err != nil {
		fmt.Println("Gagal: ", err.Error())
	}

	fmt.Println("add member succesfully ✅")
}

/*
menuListMember menampilkan seluruh data member
menangani input user view all member
*/
func ViewAllMembers() {
	members := service.GetAllMemberService()

	if len(members) == 0 {
		fmt.Println("Belum ada member")
		return
	}

	util.HeaderViewAllMember()

	for _, value := range members {
		status := "Aktif"
		if value.Status {
			status = "Nonaktif"
		}

		fmt.Printf(
			" %-4d | %-15s | %-15s | %-10s | %-12d | %-12d\n",
			value.ID,
			value.Name,
			value.Partner,
			status,
			value.TotalMissions,
			value.TotalReward,
		)
	}

	util.FooterViewAllMmeber(len(members))
}

/*
UpdateMember mengubah data member
menangani input user untuk update data member
*/
func UpdateMembers(reader *bufio.Reader) {
	fmt.Println("ID Mmeber : ")
	idMmeber, _ := reader.ReadString('\n')

	id, err := strconv.Atoi(strings.TrimSpace(idMmeber))
	if err != nil {
		fmt.Println("ID harus angka")
		return
	}

	fmt.Print("Nama : ")
	nama, _ := reader.ReadString('\n')
	fmt.Print("Patner : ")
	patner, _ := reader.ReadString('\n')
	fmt.Print("Rank : ")
	rank, _ := reader.ReadString('\n')

	req := member_dto.UpdateMemberRequestUser{
		ID:     id,
		Nama:   strings.TrimSpace(nama),
		Patner: strings.TrimSpace(patner),
		Rank:   strings.TrimSpace(rank),
	}

	if err := service.UpdateMmeberService(req); err != nil {
		fmt.Println("Gagal: ", err.Error())
	}

	fmt.Println("Member berhasil diubah ✏️", req)
}
