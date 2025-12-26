package crud

import (
	"bufio"
	"cli-library-project-akatsuki/dto/member_dto"
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

	fmt.Println("add member succesfully ✅", req)
}
