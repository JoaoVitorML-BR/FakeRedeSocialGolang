package models

type Register struct {
	email    string
	name     string
	password string
}

type login struct {
}

type Post struct {
}

func (u Register) Name() string {
	return u.name
}

func (u *Register) SetName(name string) {
	u.name = name
}

func (u Register) Email() string {
	return u.email
}

func (u *Register) SetEmail(email string) {
	u.email = email
}

func (u Register) Password() string {
	return u.password
}

func (u *Register) SetPassword(password string){
	u.password = password
}