package main

type cacheLifetime struct {
	key    string
	expiry int
}

type Cache struct {
	items        map[string]interface{}
	lifetimeList []cacheLifetime
}

func CreateCache() *Cache {
	cache := Cache{
		items:        make(map[string]interface{}),
		lifetimeList: make([]cacheLifetime, 0),
	}

	return &cache
}
