package mongo

import (
	"context"
	"reflect"
	"testing"

	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/util"
)

func createUser(t *testing.T) *model.User {
	user := util.RandomUser()
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

func TestUpdate(t *testing.T) {
	t.Run("should find existing user", func(t *testing.T) {
		user := createUser(t)
		newName := "new name"
		update := &model.User{
			ID:   user.ID,
			Name: &newName,
		}
		err := testStore.Update(context.Background(), user.ID, update)
		if err != nil {
			t.Errorf("%v", err)
		}

		found, err := testStore.FindOne(context.Background(), user.ID)
		if err != nil {
			t.Errorf("%v", err)
		}

		if found.ID != user.ID {
			t.Errorf("%v", err)
		}

		if found.Name == user.Name {
			t.Errorf("%v", err)
		}
	})
}

func TestFind(t *testing.T) {
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
