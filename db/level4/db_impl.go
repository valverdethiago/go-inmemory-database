package level4

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

type Backup struct {
	Timestamp int
	Data      map[string]map[string]Field
}

type memoryDB struct {
	data    map[string]map[string]Field
	backups []Backup
}

func NewDB() InMemoryDBLevel4 {
	return &memoryDB{
		data:    make(map[string]map[string]Field),
		backups: []Backup{},
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
		if f, exists := fields[field]; exists && !isExpired(f, timestamp) {
			return f.Value, true
		}
	}
	return "", false
}

func (m *memoryDB) DeleteWithTimestamp(key, field string, timestamp int) bool {
	if fields, ok := m.data[key]; ok {
		if f, exists := fields[field]; exists && !isExpired(f, timestamp) {
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
			if !isExpired(val, timestamp) {
				keys = append(keys, f)
			}
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
			if strings.HasPrefix(f, prefix) && !isExpired(val, timestamp) {
				keys = append(keys, f)
			}
		}
		sort.Strings(keys)
		for _, f := range keys {
			result = append(result, fmt.Sprintf("%s=%s", f, fields[f].Value))
		}
	}
	return result
}

func (m *memoryDB) Backup(timestamp int) int {
	copyData := make(map[string]map[string]Field)
	total := 0
	for k, fields := range m.data {
		copyFields := make(map[string]Field)
		for f, v := range fields {
			if !isExpired(v, timestamp) {
				copyFields[f] = v
				total++
			}
		}
		if len(copyFields) > 0 {
			copyData[k] = copyFields
		}
	}
	m.backups = append(m.backups, Backup{Timestamp: timestamp, Data: copyData})
	return total
}

func (m *memoryDB) Restore(requestedTs int, now int) {
	var best Backup
	found := false
	for _, b := range m.backups {
		if b.Timestamp <= requestedTs && (!found || b.Timestamp > best.Timestamp) {
			best = b
			found = true
		}
	}
	if !found {
		return
	}

	restored := make(map[string]map[string]Field)
	for k, fields := range best.Data {
		restored[k] = make(map[string]Field)
		for f, v := range fields {
			adjustedTTL := v.TTL - (now - v.Timestamp)
			if adjustedTTL <= 0 && v.TTL > 0 {
				continue // would be expired by now
			}
			restored[k][f] = Field{Value: v.Value, Timestamp: now, TTL: adjustedTTL}
		}
	}
	m.data = restored
}

func isExpired(f Field, now int) bool {
	return f.TTL > 0 && now >= f.Timestamp+f.TTL
}
