package db

import "sync"

var (
	idMu sync.Mutex
	id   int64

	dbMu sync.Mutex
	db   = make(map[int64]*Book)
)

// Save - method save book to storage
func Save(b *Book) int64 {
	id := NextID()
	dbMu.Lock()
	db[id] = b
	dbMu.Unlock()
	b.ID = id
	return b.ID
}

// FindAll - return all books
func FindAll() []*Book {
	bs := make([]*Book, 0)
	dbMu.Lock()
	for _, b := range db {
		bs = append(bs, b)
	}
	dbMu.Unlock()
	return bs
}

// NextID - returns next books ID
func NextID() int64 {
	incID()
	return getID()
}

func incID() {
	idMu.Lock()
	id++
	idMu.Unlock()
}

func getID() int64 {
	var v int64
	idMu.Lock()
	v = id
	idMu.Unlock()
	return v
}
