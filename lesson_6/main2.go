package main

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

type User interface {
	ChangeFIO(newFio string)
	ChangeAdress(newAddress string)
}

func (u *user) ChangeFIO(newFio string) {
	u.FIO = newFio
}

func (u *user) ChangeAdress(newAddress string) {
	u.Address = newAddress
}

func NewUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u
}


