package huck

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var counterStore *CounterStorage = InitCounterStorage()

type CounterStorage struct {
	counterMap map[string]uint64

	lock *sync.RWMutex
}

func InitCounterStorage() *CounterStorage {
	return &CounterStorage{
		counterMap: make(map[string]uint64),
		lock:       new(sync.RWMutex),
	}
}

func (store *CounterStorage) Save(key string, value uint64) {
	store.counterMap[key] = value
	go store.persistenceSync()
}

func (store *CounterStorage) persistenceSync() {

	store.lock.Lock()

	f, err := os.OpenFile(DEFAULT_COUNTER_FILE_NAME, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = f.Write(store.serialize())

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		f.Close()
		store.lock.Unlock()
	}()
}

func (store *CounterStorage) serialize() []byte {
	content := ""

	for k, v := range store.counterMap {
		content += k
		content += ": "
		content += strconv.FormatUint(v, 10)
		content += "\n"
	}

	return StringToBytes(content)
}
