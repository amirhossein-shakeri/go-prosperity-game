package auth

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Name             string `json:"name" bson:"name"`
	Password         string `json:"password" bson:"password"`
}

func New(email, name, password string) *User {
	return &User{Email: email, Name: name, Password: password}
}

func Create(email, name, password string) (*User, error) {
	usr := New(email, name, password)

	/* hash password */
	usr.Password, _ = HashPassword(usr.Password)

	/* save to db */
	return usr, mgm.Coll(usr).Create(usr)
}

func Find(id string) *User {
	usr := &User{}
	if err := mgm.Coll(usr).FindByID(id, usr); err != nil {
		return nil
	}
	return usr
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
