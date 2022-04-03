package main

type Database struct {
	config *Config
	cache  *Cache
}

func CreateDatabase(config *Config) *Database {
	cache := CreateCache()

	db := Database{
		config: config,
		cache:  cache,
	}

	return &db
}
