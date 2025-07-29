package level3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valverdethiago/go-inmemory-database/db/level3"
)

func TestLevel3_TimestampAndTTL(t *testing.T) {
	db := level3.NewDB()
	db.SetWithTimestamp("item1", "name", "ProductA", 100)
	db.SetWithTTL("item1", "promo", "10%OFF", 100, 5)
	db.SetWithTTL("item1", "expire-soon", "SoonGone", 100, 2)

	val, ok := db.GetWithTimestamp("item1", "name", 102)
	assert.True(t, ok)
	assert.Equal(t, "ProductA", val)

	val, ok = db.GetWithTimestamp("item1", "promo", 104)
	assert.True(t, ok)
	assert.Equal(t, "10%OFF", val)

	val, ok = db.GetWithTimestamp("item1", "promo", 106)
	assert.False(t, ok)

	val, ok = db.GetWithTimestamp("item1", "expire-soon", 103)
	assert.False(t, ok)
}

func TestLevel3_ScanAndDelete(t *testing.T) {
	db := level3.NewDB()
	db.SetWithTTL("session1", "token", "abc", 200, 5)
	db.SetWithTimestamp("session1", "user", "bob", 200)

	scan := db.ScanWithTimestamp("session1", 202)
	assert.ElementsMatch(t, []string{"token=abc", "user=bob"}, scan)

	ok := db.DeleteWithTimestamp("session1", "token", 204)
	assert.True(t, ok)

	ok = db.DeleteWithTimestamp("session1", "user", 204)
	assert.True(t, ok)

	scan = db.ScanWithTimestamp("session1", 205)
	assert.Empty(t, scan)
}
