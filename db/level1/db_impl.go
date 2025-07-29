package level1

type memoryDB struct {
	data map[string]map[string]string
}

func NewDB() InMemoryDBLevel1 {
	return &memoryDB{
		data: make(map[string]map[string]string),
	}
}

func (m *memoryDB) Set(key string, field string, value string) {
	if _, ok := m.data[key]; !ok {
		m.data[key] = make(map[string]string)
	}
	m.data[key][field] = value
}

func (m *memoryDB) Get(key string, field string) (string, bool) {
	if fields, ok := m.data[key]; ok {
		val, exists := fields[field]
		return val, exists
	}
	return "", false
}

func (m *memoryDB) Delete(key string, field string) bool {
	if fields, ok := m.data[key]; ok {
		if _, exists := fields[field]; exists {
			delete(fields, field)
			if len(fields) == 0 {
				delete(m.data, key)
			}
			return true
		}
	}
	return false
}
