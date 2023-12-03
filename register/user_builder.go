package register

import (
	"pkg-hub/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName         string `json:"user_name,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	HashPassword     string `json:"hash_password,omitempty"`
	HashPasswordSalt string `json:"hash_password_salt,omitempty"`
	PhoneNumber      int    `json:"phone_number,omitempty"`
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{}}
}

func (ub *UserBuilder) SetUserName(userName string) *UserBuilder {
	ub.user.UserName = userName
	return ub
}

func (ub *UserBuilder) SetFirstName(firstName string) *UserBuilder {
	ub.user.FirstName = firstName
	return ub
}

func (ub *UserBuilder) SetLastName(lastName string) *UserBuilder {
	ub.user.LastName = lastName
	return ub
}

func (ub *UserBuilder) SetEmail(email string) *UserBuilder {
	ub.user.Email = email
	return ub
}

func (ub *UserBuilder) SetPassword(password string) {
	salt, _ := utils.GenerateRandomSalt(10)
	ub.user.HashPasswordSalt = string(salt)
	bPass, _ := bcrypt.GenerateFromPassword([]byte(password+ub.user.HashPasswordSalt), bcrypt.DefaultCost)
	ub.user.HashPassword = string(bPass)
}

func (ub *UserBuilder) SetPhoneNumber(phoneNumber int) {
	ub.user.PhoneNumber = phoneNumber
}

func (ub *UserBuilder) Build() User {
	return ub.user
}
