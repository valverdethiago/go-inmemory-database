package level4_test

import (
	"testing"

	"github.com/stretchr/testify/assert" 
	"github.com/valverdethiago/go-inmemory-database/db/level4"

)

func TestLevel4_BackupAndRestore(t *testing.T) {
	db := level4.NewDB()
	db.SetWithTTL("config", "featureX", "enabled", 100, 50)
	db.SetWithTTL("config", "featureY", "disabled", 100, 10)
	db.SetWithTTL("config", "featureZ", "experimental", 100, 5)

	count := db.Backup(104)
	assert.Equal(t, 3, count)

	db.SetWithTTL("config", "featureW", "deprecated", 110, 5)
	db.Backup(115)
	db.SetWithTTL("config", "featureW", "removed", 120, 1)

	db.Restore(105, 130)

	val, ok := db.GetWithTimestamp("config", "featureX", 130)
	assert.True(t, ok)
	assert.Equal(t, "enabled", val)

	val, ok = db.GetWithTimestamp("config", "featureY", 130)
	assert.False(t, ok) // expired after restore

	val, ok = db.GetWithTimestamp("config", "featureZ", 130)
	assert.False(t, ok)

	val, ok = db.GetWithTimestamp("config", "featureW", 130)
	assert.False(t, ok) // not part of 105 backup
}