// Problem Statement - Level 4
// Extend the in-memory database to support backup and restore functionality.
// You must be able to save the state of the database and restore it to a given timestamp.
// TTL values should be recalculated during restore to reflect the time elapsed.

// Interface for Level 4
package db

type InMemoryDBLevel4 interface {
	// Level 1
	Set(key string, field string, value string)
	Get(key string, field string) (string, bool)
	Delete(key string, field string) bool

	// Level 2
	Scan(key string) []string
	ScanByPrefix(key string, prefix string) []string

	// Level 3
	SetWithTimestamp(key string, field string, value string, timestamp int)
	SetWithTTL(key string, field string, value string, timestamp int, ttl int)
	GetWithTimestamp(key string, field string, timestamp int) (string, bool)
	DeleteWithTimestamp(key string, field string, timestamp int) bool
	ScanWithTimestamp(key string, timestamp int) []string
	ScanByPrefixWithTimestamp(key string, prefix string, timestamp int) []string

	// Level 4
	Backup(timestamp int) int
	Restore(requestedTs int, now int)
}
