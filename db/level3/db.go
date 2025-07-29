// Problem Statement - Level 3
// Enhance the in-memory database to support TTL (Time-To-Live) functionality for some records.
// Records with TTL should expire and be inaccessible after their time is up.
// All read/write operations should take a timestamp parameter to determine expiration.

// Interface for Level 3
package level3


type InMemoryDBLevel3 interface {
	SetWithTimestamp(key string, field string, value string, timestamp int)
	SetWithTTL(key string, field string, value string, timestamp int, ttl int)
	GetWithTimestamp(key string, field string, timestamp int) (string, bool)
	DeleteWithTimestamp(key string, field string, timestamp int) bool
	ScanWithTimestamp(key string, timestamp int) []string
	ScanByPrefixWithTimestamp(key string, prefix string, timestamp int) []string
}
