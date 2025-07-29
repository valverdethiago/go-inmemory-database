// Problem Statement - Level 1
// Implement a basic in-memory database that supports basic operations to manipulate records, fields, and values.

// Interface for Level 1
package level1

type InMemoryDBLevel1 interface {
	Set(key string, field string, value string)
	Get(key string, field string) (string, bool)
	Delete(key string, field string) bool
}