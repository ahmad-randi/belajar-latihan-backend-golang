package user

// users berperan sebagai DATABASE SEMENTARA
var users []User

func Save(u User) {
	users = append(users, u)
}

func GetAllUsers() []User {
	return users
}

func Update(index int, u User) {
	users[index] = u
}

func Delete(index int) {
	users = append(users[:index], users[index+1:]...)
}
