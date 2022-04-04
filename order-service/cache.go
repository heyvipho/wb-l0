package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type cacheLifetime struct {
	Key string `json:"key"`
}

type cacheData struct {
	Items        map[string]interface{} `json:"items"`
	LifetimeList []cacheLifetime        `json:"lifetimeList"`
}

type Cache struct {
	mutex    sync.RWMutex
	fileName string
	max      int
	data     cacheData
}

func CreateCache(fileName string, max int) *Cache {
	data := cacheData{
		Items:        make(map[string]interface{}),
		LifetimeList: make([]cacheLifetime, 0),
	}

	cache := Cache{
		mutex:    sync.RWMutex{},
		fileName: fileName,
		max:      max,
		data:     data,
	}

	return &cache
}

func (v *Cache) Get(key string) interface{} {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return v.data.Items[key]
}

func (v *Cache) Set(key string, value interface{}) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if len(v.data.LifetimeList) >= v.max {
		delCandidate := v.data.LifetimeList[0]
		v.data.LifetimeList = v.data.LifetimeList[1:]
		delete(v.data.Items, delCandidate.Key)
	}

	v.data.Items[key] = value
	v.data.LifetimeList = append(v.data.LifetimeList, cacheLifetime{
		Key: key,
	})

	v.Save()
}

func (v *Cache) Save() {
	jsonObj, err := json.Marshal(v.data)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(v.fileName, jsonObj, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (v *Cache) Restore() {
	fileBody, err := ioutil.ReadFile(v.fileName)
	if err != nil {
		return
	}

	var fileData cacheData
	err = json.Unmarshal(fileBody, &fileData)
	if err != nil {
		return
	}

	if fileData.Items != nil {
		v.data.Items = fileData.Items
	}

	if fileData.LifetimeList != nil {
		v.data.LifetimeList = fileData.LifetimeList
	}
}
