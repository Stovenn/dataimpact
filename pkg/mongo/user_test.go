package mongo

import (
	"context"
	"testing"

	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func createUser(t *testing.T) *model.User {
	user := util.RandomUser()
	err := testStore.Create(context.Background(), user)
	require.NoError(t, err)

	return user
}

func TestCreate(t *testing.T) {
	createUser(t)
}

func TestFindOne(t *testing.T) {
	t.Run("should find existing user", func(t *testing.T) {
		user := createUser(t)
		found, err := testStore.FindOne(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, found)

		assert.Equal(t, user, found)
	})
	t.Run("should not find user", func(t *testing.T) {
		user, err := testStore.FindOne(context.Background(), "unknown")
		assert.Error(t, err)
		assert.Empty(t, user)
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
		assert.NoError(t, err)

		found, err := testStore.FindOne(context.Background(), user.ID)
		assert.NoError(t, err)

		assert.Equal(t, found.ID, user.ID)
		assert.NotEqual(t, found.Name, user.Name)
	})
}

func TestFind(t *testing.T) {
	for i := 0; i < 5; i++ {
		createUser(t)
	}

	found, err := testStore.Find(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, found)
}

func TestDelete(t *testing.T) {
	user := createUser(t)

	err := testStore.DeleteOne(context.Background(), user.ID)
	assert.NoError(t, err)

	foundUser, err := testStore.FindOne(context.Background(), user.ID)
	assert.Empty(t, foundUser)
	assert.Error(t, err)
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
}
