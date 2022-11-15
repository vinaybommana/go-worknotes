package factory

import "fmt"

type mongoDB struct {
	database map[string]string
}

type sqlite struct {
	database map[string]string
}

// Database Interface
type Database interface {
	GetData(string) string
	PutData(string, string)
}


// mongoDB methods
func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

// sqlite methods
func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("SQLite")
	return sql.database[query]
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}


func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}
