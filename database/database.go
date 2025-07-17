package database

type Database interface {
	StoreSessionID(id string) error
	RemoveSessionID(id string) error
	Close()
}

func New() Database {
	db := new(inMemoryDB)
	db.initialize()
	return db
}
