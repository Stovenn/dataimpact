package mongo

import (
	"context"
	"reflect"
	"testing"

	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/util"
)

func createUser(t *testing.T) *model.User {
	name := util.RandomString(10)
	user := &model.User{ID: util.RandomString(25), Name: &name}
	err := testStore.Create(context.Background(), user)
	if err != nil {
		t.Errorf("%v", err)
	}
	return user
}

func TestCreate(t *testing.T) {
	createUser(t)
}

func TestFindOne(t *testing.T) {
	t.Run("should find existing user", func(t *testing.T) {
		user := createUser(t)
		found, err := testStore.FindOne(context.Background(), user.ID)
		if err != nil {
			t.Errorf("%v", err)
		}

		if user.ID != found.ID {
			t.Errorf("%v", err)
		}
	})
	t.Run("should not find user", func(t *testing.T) {
		_, err := testStore.FindOne(context.Background(), "unknown")
		if err == nil {
			t.Errorf("%v", err)
		}
	})

}

func TestFindAll(t *testing.T) {
	var users []*model.User
	for i := 0; i < 5; i++ {
		users = append(users, createUser(t))
	}

	found, err := testStore.Find(context.Background())
	if err != nil {
		t.Errorf("%v", err)
	}

	if reflect.DeepEqual(users, found) {
		t.Errorf("%v", err)
	}
}

func TestDelete(t *testing.T) {
	user := createUser(t)

	err := testStore.DeleteOne(context.Background(), user.ID)
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = testStore.FindOne(context.Background(), user.ID)
	if err == nil {
		t.Errorf("%v", err)
	}
}
