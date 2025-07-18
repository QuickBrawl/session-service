package database

type Database interface {
	StoreSessionID(sessionId string) error
	SessionExists(sessionId string) (bool, error)
	RemoveSessionID(sessionId string) error
	AddUserToSession(userId string, sessionId string) error
	RemoveUserFromSession(userId string, sessionIs string) error
	Close()
}

func New() Database {
	db := new(inMemoryDB)
	db.initialize()
	return db
}
