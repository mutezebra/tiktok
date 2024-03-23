package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Mutezebra/tiktok/app/domain/repository"
)

func userModel(t *testing.T, id int64, name string) *repository.User {
	t.Helper()
	return &repository.User{
		ID:             id,
		UserName:       name,
		Email:          "asd",
		PasswordDigest: "asd",
		Gender:         1,
		Avatar:         "avatar",
		Fans:           4,
		Follows:        6,
		TotpEnable:     false,
		TotpSecret:     "secret",
		CreateAt:       time.Now().Unix(),
		UpdateAt:       time.Now().Unix(),
		DeleteAt:       0,
	}
}

func isEqual(t *testing.T, u1, u2 *repository.User) bool {
	t.Helper()
	return u1.ID == u2.ID && u1.UserName == u2.UserName
}

func TestCreateUser(t *testing.T) {
	if err := mysqlInit(t); err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepository()
	ctx := context.Background()
	id := int64(67)
	name := "mu"
	model := userModel(t, id, name)
	if err := repo.CreateUser(ctx, model); err != nil {
		t.Fatal(err)
	}

	u1, err := repo.UserInfoByID(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println()
	u2, err := repo.UserInfoByName(ctx, name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(isEqual(t, u1, u2))
}
