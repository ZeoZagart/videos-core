package DB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	authproto "go.videos.core/protos"
	"golang.org/x/crypto/bcrypt"
)

type IAuthTable interface {
	CreateAccount(*authproto.SignUpRequest) (User, error)
	VerifyUser(context.Context, *authproto.LoginRequest) (User, error)
}

func getAuthTable(authCollection *mongo.Collection) IAuthTable {
	return &authTable{authCollection}
}

type authTable struct {
	table *mongo.Collection
}

func (a *authTable) CreateAccount(r *authproto.SignUpRequest) (User, error) {
	passHash, err := getPassHash(r.Password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Email:        r.Email,
		Username:     r.Username,
		PasswordHash: string(passHash),
	}
	id, err := a.table.InsertOne(context.Background(), user)
	fmt.Printf("id: %+v\n", id)
	return User{Email: r.Email, PasswordHash: user.PasswordHash}, err
}

func (a *authTable) VerifyUser(ctxt context.Context, r *authproto.LoginRequest) (User, error) {
	user := new(User)
	err := a.table.FindOne(ctxt, bson.M{}).Decode(user)
	fmt.Printf("user: %+v, error: %+v\n", user, err)
	return *user, err
}

func getPassHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
