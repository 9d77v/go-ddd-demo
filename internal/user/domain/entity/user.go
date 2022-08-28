package entity

import (
	"log"

	"github.com/9d77v/go-ddd-demo/internal/user/domain/enum"

	"github.com/9d77v/go-ddd-demo/pkg/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db.Entity
	Phone    string
	Password string
	Nickname string
	Gender   enum.GenderEnum
}

func NewUser(phone, password, nickname string, gender enum.GenderEnum) *User {
	var u = new(User)
	u.ID = uuid.NewString()
	u.Phone = phone
	u.Password = u.encryptPassword(password)
	u.Nickname = nickname
	u.Gender = gender
	return u
}

func (u *User) encryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panicf("generate password failed:%v/n", err)
	}
	return string(bytes)
}
