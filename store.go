package huck

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var counterStore = InitCounterStorage()

// CounterStorage is Counter persistence device.
type CounterStorage struct {
	counterMap map[string]uint64

	fieldLocks map[string]*sync.Mutex
	fileLock   *sync.Mutex
}

// InitCounterStorage to initializer a CounterStorage.
func InitCounterStorage() *CounterStorage {
	store := &CounterStorage{
		counterMap: make(map[string]uint64),
		fieldLocks: make(map[string]*sync.Mutex),
		fileLock:   new(sync.Mutex),
	}

	store.LoadFromDisk(DefaultCounterFileName)

	return store
}

// InitCounterFieldLocks to init fieldLocks
func (store *CounterStorage) InitCounterFieldLocks(key string) {
	store.fieldLocks[key] = new(sync.Mutex)
}

// LoadFromDisk is Locate the persistent counter from the hard disk and load it into Storage.
func (store *CounterStorage) LoadFromDisk(filename string) {
	if !IsFileExist(filename) {
		return
	}

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(err.Error())
	}

	content := BytesToString(b)
	store.LoadFromString(content)
}

// LoadFromString is Locate the persistent counter from the String and load it into Storage.
func (store *CounterStorage) LoadFromString(content string) {

	counters := strings.Split(content, "\n")
	for _, counter := range counters {
		counterArr := strings.Split(counter, ":")

		if len(counterArr) == 1 {
			continue
		} else if len(counterArr) != 2 {
			log.Fatalln(StatisticsFileFormatError)
		}

		count, err := strconv.ParseUint(strings.TrimSpace(counterArr[1]), 10, 64)
		if err != nil {
			log.Println(StatisticsFileFormatError)
			log.Fatalln(err.Error())
		}

		store.counterMap[strings.TrimSpace(counterArr[0])] = count
	}
}

// Save the new counter value and persist it. persist is Asynchronous.
func (store *CounterStorage) Save(key string, value uint64) {
	store.fieldLocks[key].Lock()
	store.counterMap[key] = value
	store.fieldLocks[key].Unlock()

	store.persistenceToDisk()
}

func (store *CounterStorage) persistenceToDisk() {

	store.fileLock.Lock()

	f, err := os.OpenFile(DefaultCounterFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = f.Write(store.serialize())

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		f.Close()
		store.fileLock.Unlock()
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
