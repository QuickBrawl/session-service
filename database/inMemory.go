package database

import (
	"errors"
	"slices"
)

type inMemoryDB struct {
	userToSessionMap map[string][]string
}

func (db *inMemoryDB) initialize() {
	db.userToSessionMap = make(map[string][]string)
}

func (db *inMemoryDB) StoreSessionID(id string) error {
	if _, ok := db.userToSessionMap[id]; ok {
		return errors.New("Session does already exists!")
	}
	db.userToSessionMap[id] = make([]string, 0)
	return nil
}

func (db *inMemoryDB) SessionExists(id string) (bool, error) {
	if _, ok := db.userToSessionMap[id]; ok {
		return true, nil
	}
	return false, nil
}

func (db *inMemoryDB) RemoveSessionID(id string) error {
	delete(db.userToSessionMap, id)
	return nil
}

func (db *inMemoryDB) AddUserToSession(userId string, sessionId string) error {
	if ok, _ := db.SessionExists(sessionId); !ok {
		return errors.New("Session does not exists!")
	}

	if _, ok := db.userToSessionMap[sessionId]; !ok {
		db.userToSessionMap[sessionId] = make([]string, 0)
	}
	db.userToSessionMap[sessionId] = append(db.userToSessionMap[sessionId], userId)

	return nil
}

func (db *inMemoryDB) RemoveUserFromSession(userId string, sessionId string) error {
	if ok, _ := db.SessionExists(sessionId); !ok {
		return errors.New("Session does not exists!")
	}

	if _, ok := db.userToSessionMap[sessionId]; !ok {
		return nil
	}

	index := slices.Index(db.userToSessionMap[sessionId], userId)
	if index >= 0 {
		db.userToSessionMap[sessionId][index] = db.userToSessionMap[sessionId][len(db.userToSessionMap[sessionId])-1]
		db.userToSessionMap[sessionId] = db.userToSessionMap[sessionId][:len(db.userToSessionMap[sessionId])-1]
	}

	return nil
}

func (db *inMemoryDB) Close() {}
