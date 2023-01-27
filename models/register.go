package models

type Register struct {
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"passowrd"`
}

func (u Register) GetName() string {
	return u.Name
}

func (u *Register) SetName(name string) {
	u.Name = name
}

func (u Register) GetEmail() string {
	return u.Email
}

func (u *Register) SetEmail(email string) {
	u.Email = email
}

func (u Register) GetPassword() string {
	return u.Password
}

func (u *Register) SetPassword(password string) {
	u.Password = password
}
