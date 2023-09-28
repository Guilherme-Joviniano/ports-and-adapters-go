package connections

import (
	"database/sql"
	"sync"
)

var once = &sync.Once{}

type Singleton struct {
	Connection *sql.DB
}

var singletonInstance *Singleton

func GetSqliteInstance() *Singleton {
	if singletonInstance == nil {
		once.Do(
			func() {
				db, err := sql.Open("sqlite3", "db.sqlite")

				if err != nil {
					panic(err)
				}
				singletonInstance = &Singleton{
					Connection: db,
				}
			})
	}
	return singletonInstance
}
