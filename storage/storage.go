package storage

type Storage struct {
	data map[string]string
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New() *Storage {
	return &Storage{make(map[string]string)}
}

func (store *Storage) Put(key string, value string) {
	store.data[key] = value
}

func (store Storage) Get(key string) (d Data, exists bool) {
	v, exists := store.data[key]
	return Data{key, v}, exists
}

// Delete only sets the value to null, but leaves the entry
func (store *Storage) Delete(key string) {
	delete(store.data, key)
}

func (store Storage) List() []Data {
	var d []Data
	for key, value := range store.data {
		d = append(d, Data{key, value})
	}
	return d
}
