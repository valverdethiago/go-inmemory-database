// Problem Statement - Level 2
// Extend the in-memory database to support querying a specific record based on a given filter.
// Specifically, you must be able to retrieve all fields for a key and filter fields by prefix.

// Interface for Level 2
package db

type InMemoryDBLevel2 interface {
	// Level 1
	Set(key string, field string, value string)
	Get(key string, field string) (string, bool)
	Delete(key string, field string) bool

	// Level 2
	Scan(key string) []string
	ScanByPrefix(key string, prefix string) []string
}
