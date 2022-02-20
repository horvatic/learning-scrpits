package interpreter

type DataStore struct {
	data map[string]interface{}
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]interface{}),
	}
}

func (d *DataStore) AddData(label string, data interface{}) {
	d.data[label] = data
}

func (d *DataStore) GetData(label string) interface{} {
	return d.data[label]
}
