package database

type Database interface {
	StoreSessionID(id string) error
	Close() error
}

func New()
