package huck

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var counterStore *CounterStorage = InitCounterStorage()

type CounterStorage struct {
	counterMap map[string]uint64

	lock *sync.RWMutex
}

func InitCounterStorage() *CounterStorage {
	store := &CounterStorage{
		counterMap: make(map[string]uint64),
		lock:       new(sync.RWMutex),
	}

	store.LoadFromDisk(DEFAULT_COUNTER_FILE_NAME)

	return store
}

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

func (store *CounterStorage) LoadFromString(content string) {

	counters := strings.Split(content, "\n")
	for _, counter := range counters {
		counterArr := strings.Split(counter, ":")

		if len(counterArr) == 1 {
			continue
		} else if len(counterArr) != 2 {
			log.Fatalln(STATISTICS_FILE_FORMAT_ERROR)
		}

		count, err := strconv.ParseUint(strings.TrimSpace(counterArr[1]), 10, 64)
		if err != nil {
			log.Println(STATISTICS_FILE_FORMAT_ERROR)
			log.Fatalln(err.Error())
		}

		store.counterMap[strings.TrimSpace(counterArr[0])] = count
	}
}

func (store *CounterStorage) Save(key string, value uint64) {
	store.counterMap[key] = value
	go store.persistenceToDisk()
}

func (store *CounterStorage) persistenceToDisk() {

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
