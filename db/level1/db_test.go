package level1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
    "github.com/valverdethiago/go-inmemory-database/db/level1"
)

func TestLevel1_SetGetDelete(t *testing.T) {
	db := level1.NewDB()

	db.Set("user1", "name", "Alice")
	db.Set("user1", "email", "alice@example.com")

	val, ok := db.Get("user1", "name")
	assert.True(t, ok)
	assert.Equal(t, "Alice", val)

	val, ok = db.Get("user1", "email")
	assert.True(t, ok)
	assert.Equal(t, "alice@example.com", val)

	ok = db.Delete("user1", "name")
	assert.True(t, ok)
	_, ok = db.Get("user1", "name")
	assert.False(t, ok)

	ok = db.Delete("user1", "nonexistent")
	assert.False(t, ok)
} 
