package level2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valverdethiago/go-inmemory-database/db/level2"
)

func TestLevel2_SetGetDelete(t *testing.T) {
	db := level2.NewDB()

	db.Set("user1", "name", "Alice")
	db.Set("user1", "email", "alice@example.com")

	val, ok := db.Get("user1", "name")
	assert.True(t, ok)
	assert.Equal(t, "Alice", val)

	ok = db.Delete("user1", "email")
	assert.True(t, ok)
	_, ok = db.Get("user1", "email")
	assert.False(t, ok)
}

func TestLevel2_Scan(t *testing.T) {
	db := level2.NewDB()
	db.Set("user2", "z", "zzz")
	db.Set("user2", "a", "aaa")
	db.Set("user2", "m", "mmm")

	scan := db.Scan("user2")
	assert.Equal(t, []string{"a=aaa", "m=mmm", "z=zzz"}, scan)
}

func TestLevel2_ScanByPrefix(t *testing.T) {
	db := level2.NewDB()
	db.Set("user3", "addr_street", "Main St")
	db.Set("user3", "addr_zip", "12345")
	db.Set("user3", "email", "bob@example.com")

	filtered := db.ScanByPrefix("user3", "addr_")
	assert.Equal(t, []string{"addr_street=Main St", "addr_zip=12345"}, filtered)
}
