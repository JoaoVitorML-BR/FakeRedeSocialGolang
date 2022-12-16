package models

type Register struct {
	email    string
	name     string
	password string
}

//type Login struct {
//	emailogin string
//	passwordlogin string
//}

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

func (u *Register) SetPassword(password string) {
	u.password = password
}

//func (l Login) Email() string {
//	return  l.emailogin
//}
//
//func (l *Login) SetLogin(emailogin string) {
//	l.emailogin = emailogin
//}

//func (l Login) Password() string {
//	return l.passwordlogin
//}
//
//func (l *Login) SetPawordLogin(passwordlogin string){
//	l.passwordlogin = passwordlogin
//}
