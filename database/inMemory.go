package database

import "slices"

type inMemoryDB struct {
	sessions []string
}

func (db *inMemoryDB) initialize() {
	db.sessions = make([]string, 0)
}

func (db *inMemoryDB) StoreSessionID(id string) error {
	db.sessions = append(db.sessions, id)
	return nil
}

func (db *inMemoryDB) RemoveSessionID(id string) error {
	index := slices.Index(db.sessions, id)
	if index >= 0 {
		db.sessions[index] = db.sessions[len(db.sessions)-1]
		db.sessions = db.sessions[:len(db.sessions)-1]
	}

	return nil
}

func (db *inMemoryDB) Close() {}
