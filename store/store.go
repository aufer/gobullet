package main

import (
	"log"
	"os"
)

type Store struct {
	entries map[string]interface{}
}

func New() *Store {
	storeEntries := make(map[string]interface{})

	f, err := os.ReadFile("store.db")
	if err != nil {
		log.Fatal(err)
	}

	// read/parse and create entries

	return &Store{
		entries: storeEntries,
	}
}

func (s *Store) Put(key string, value interface{}) {
	s.entries[key] = value
}

func (s *Store) Get(key string) interface{} {
	return s.entries[key]
}
