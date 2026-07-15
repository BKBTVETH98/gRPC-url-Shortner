package auth

import (
	"context"
	"fmt"
	"testing"

	ssov1 "github.com/BKBTVETH98/protos/gen/go/sso"
)

func TestLogin(t *testing.T) {
	server := &serverAPI{}

	req := &ssov1.LoginRequest{
		Email:    "52",
		Password: "52",
	}

	resp, err := server.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
		fmt.Println()
	}
	if resp == nil {
		t.Fatal("Response is nil")
		fmt.Println()
	}

	if resp.Token != "" {
		fmt.Println("Response token:", resp.Token)
		fmt.Println()
	}

	if resp.Token == "" {
		t.Fatal("Token is empty")
		fmt.Println()
	}
}

func TestRegister(t *testing.T) {
	server := serverAPI{}

	req := &ssov1.RegisterRequest{
		Email:    "52@mail.ru",
		Password: "52ALBLAK52",
	}
	resp, err := server.Register(context.Background(), req)

	if err != nil {
		fmt.Println("Register failed:", err)
		fmt.Println(resp)
	}
	if resp == nil {
		fmt.Println("Response is nil", resp)
		fmt.Println()
	}

	if resp != nil {
		fmt.Println("Response User id", resp.UserId)
		fmt.Println()
	}
}
