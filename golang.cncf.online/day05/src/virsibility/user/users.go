package user

type address struct {
	Region string
	number string
	street string
}

type User struct {
	ID        int
	Name      string
	birthDate string
	addr      address
}
