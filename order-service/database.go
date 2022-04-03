package main

type Database struct {
	c     *Config
	cache *Cache
}

func CreateDatabase(c *Config) (*Database, error) {
	cache := CreateCache()

	db := Database{
		c:     c,
		cache: cache,
	}

	return &db, nil
}
