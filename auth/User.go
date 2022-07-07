package auth

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Name             string `json:"name" bson:"name"`
	Password         string `json:"-" bson:"password"` // todo: hide password for json
}

func NewUser(email, name, password string) *User {
	return &User{Email: email, Name: name, Password: password}
}

func CreateUser(email, name, password string) (*User, error) {
	usr := NewUser(email, name, password)

	/* hash password */
	usr.Password, _ = HashPassword(usr.Password)

	/* save to db */
	return usr, mgm.Coll(usr).Create(usr)
}

func FindUser(id string) *User {
	usr := &User{}
	if err := mgm.Coll(usr).FindByID(id, usr); err != nil {
		return nil
	}
	return usr
}

func FindUserByEmail(email string) *User {
	usr := &User{}
	if err := mgm.Coll(usr).First(bson.M{"email": email}, usr); err != nil {
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

func AuthenticateUser(email, password string) (*User, string) {
	/* Login the user */
	usr := FindUserByEmail(email)
	if usr == nil {
		log.Println("Login attempt failed. Email", email, "not found")
		return nil, "Email not found"
	}
	if !CheckPasswordHash(password, usr.Password) {
		log.Println("Login attempt failed. Wrong password for", email)
		return nil, "Wrong password"
	}
	log.Println("Successful login attempt.", email, "is authenticated")
	return usr, "Successfully logged in"
}
