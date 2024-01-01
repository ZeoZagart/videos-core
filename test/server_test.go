package server_test

import (
	"context"
	"testing"

	DB "go.videos.core/db"
	authproto "go.videos.core/protos"
	"google.golang.org/grpc"
)

func setup(t *testing.T) (authproto.AuthServiceClient, *grpc.ClientConn) {
	cc, err := grpc.Dial("0.0.0.0:1001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authClient := authproto.NewAuthServiceClient(cc)
	return authClient, cc
}

func tearDown(cc *grpc.ClientConn, t *testing.T) {
	cc.Close()
}

func TestCreateAccountWorksWithNewUserAndFailsWithOld(t *testing.T) {
	client, cc := setup(t)
	defer tearDown(cc, t)
	user := getUser()

	resp, err := createAccount(client, user, t)
	if err != nil || resp.Email != user.Email {
		t.Errorf("Error creating signup %v", err)
	}

	resp, err = createAccount(client, user, t)
	if err == nil || resp != nil {
		t.Errorf("Created account twice %v", resp)
	}
}

func TestGetAccountWorks(t *testing.T) {
	client, cc := setup(t)
	defer tearDown(cc, t)

	user := getUser()
	resp, err := getAccount(client, user, t)
	if err != nil || resp.Email != user.Email {
		t.Errorf("Error fetching account err: %+v acc: %+v", err, resp)
	}
}

func createAccount(client authproto.AuthServiceClient, user DB.User, t *testing.T) (*authproto.LoginResponse, error) {
	request := &authproto.SignUpRequest{Email: user.Email, Username: user.Username, Password: user.PasswordHash}
	resp, err := client.CreateAccount(context.Background(), request)
	t.Logf("\nresp: %+v\nerror: %+v\n", resp, err)
	return resp, err
}

func getAccount(client authproto.AuthServiceClient, user DB.User, t *testing.T) (*authproto.LoginResponse, error) {
	request := &authproto.LoginRequest{Email: user.Email, Password: user.PasswordHash}
	resp, err := client.Login(context.Background(), request)
	t.Logf("\nresp: %+v\nerror: %+v\n", resp, err)
	return resp, err
}

func getUser() DB.User {
	return DB.User{Email: "abhashasd5@gmail.com", Username: "abhash", PasswordHash: "random"}
}
