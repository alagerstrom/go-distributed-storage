package storage

type Storage struct {
	data map[string]Data
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New() *Storage {
	return &Storage{make(map[string]Data)}
}

func (store *Storage) Put(key string, value string) {
	store.data[key] = Data{Key: key, Value: value}
}

func (store Storage) Get(key string) (d Data, exists bool) {
	d, exists = store.data[key]
	return d, exists
}

func (store Storage) List() []Data {
	var d []Data
	for _, data := range store.data {
		d = append(d, data)
	}
	return d
}
