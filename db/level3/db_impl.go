package level3

import (
	"fmt"
	"sort"
	"strings"
)

type Field struct {
	Value     string
	Timestamp int
	TTL       int
}

type memoryDB struct {
	data map[string]map[string]Field
}

func NewDB() InMemoryDBLevel3 {
	return &memoryDB{
		data: make(map[string]map[string]Field),
	}
}

func (m *memoryDB) SetWithTimestamp(key, field, value string, timestamp int) {
	if _, ok := m.data[key]; !ok {
		m.data[key] = make(map[string]Field)
	}
	m.data[key][field] = Field{Value: value, Timestamp: timestamp}
}

func (m *memoryDB) SetWithTTL(key, field, value string, timestamp, ttl int) {
	if _, ok := m.data[key]; !ok {
		m.data[key] = make(map[string]Field)
	}
	m.data[key][field] = Field{Value: value, Timestamp: timestamp, TTL: ttl}
}

func (m *memoryDB) GetWithTimestamp(key, field string, timestamp int) (string, bool) {
	if fields, ok := m.data[key]; ok {
		if f, exists := fields[field]; exists {
			if isExpired(f, timestamp) {
				delete(fields, field)
				if len(fields) == 0 {
					delete(m.data, key)
				}
				return "", false
			}
			return f.Value, true
		}
	}
	return "", false
}

func (m *memoryDB) DeleteWithTimestamp(key, field string, timestamp int) bool {
	if fields, ok := m.data[key]; ok {
		if f, exists := fields[field]; exists {
			if isExpired(f, timestamp) {
				delete(fields, field)
				if len(fields) == 0 {
					delete(m.data, key)
				}
				return false
			}
			delete(fields, field)
			if len(fields) == 0 {
				delete(m.data, key)
			}
			return true
		}
	}
	return false
}

func (m *memoryDB) ScanWithTimestamp(key string, timestamp int) []string {
	result := []string{}
	if fields, ok := m.data[key]; ok {
		keys := []string{}
		for f, val := range fields {
			if isExpired(val, timestamp) {
				delete(fields, f)
				continue
			}
			keys = append(keys, f)
		}
		if len(fields) == 0 {
			delete(m.data, key)
		}
		sort.Strings(keys)
		for _, f := range keys {
			result = append(result, fmt.Sprintf("%s=%s", f, fields[f].Value))
		}
	}
	return result
}

func (m *memoryDB) ScanByPrefixWithTimestamp(key, prefix string, timestamp int) []string {
	result := []string{}
	if fields, ok := m.data[key]; ok {
		keys := []string{}
		for f, val := range fields {
			if isExpired(val, timestamp) {
				delete(fields, f)
				continue
			}
			if strings.HasPrefix(f, prefix) {
				keys = append(keys, f)
			}
		}
		if len(fields) == 0 {
			delete(m.data, key)
		}
		sort.Strings(keys)
		for _, f := range keys {
			result = append(result, fmt.Sprintf("%s=%s", f, fields[f].Value))
		}
	}
	return result
}

func isExpired(f Field, now int) bool {
	return f.TTL > 0 && now >= f.Timestamp+f.TTL
}
