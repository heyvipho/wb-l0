package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v4"
)

var ERR_DB_NOT_FOUND = errors.New("DB: Not found")

type Database struct {
	config *Config
	cache  *Cache
	conn   *pgx.Conn
}

func CreateDatabase(config *Config) *Database {
	cache := CreateCache()

	db := Database{
		config: config,
		cache:  cache,
	}

	return &db
}

func (v *Database) GetOrder(id string) (Order, error) {
	var order Order

	var jsonObj []byte
	err := v.conn.QueryRow(context.Background(), "SELECT json FROM orders WHERE id=$1", id).Scan(&jsonObj)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(jsonObj, &order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (v *Database) AddOrder(obj Order) error {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	_, err = v.conn.Exec(context.Background(), "INSERT INTO orders VALUES ($1, $2)", obj.OrderUid, jsonObj)

	return err
}

func (v *Database) Connect() error {
	conn, err := pgx.Connect(context.Background(), v.config.DatabaseURL)
	v.conn = conn

	return err
}

func (v *Database) Close() error {
	err := v.conn.Close(context.Background())

	return err
}
