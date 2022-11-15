package abstractfactory

import "fmt"

type mongoDB struct {
	database map[string]string
}

type sqlite struct {
	database map[string]string
}


type file struct {
	name string
	content string
}

type ntfs struct {
	files map[string]file
}

type ext4 struct {
	files map[string]file
}

// Database interface
type Database interface {
	PutData(string, string)
	GetData(string) string
}

// FileSystem interface
type FileSystem interface {
	CreateFile(string)
	FindFile(string) file
}


type Factory func(string) interface{}

// Database methods
func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}
	return mdb.database[query]
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}
	return sql.database[query]
}


// FileSystem methods
func (ntfs ntfs) CreateFile(path string) {
	file := file{content: "NTFS File", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext4 ext4) CreateFile(path string) {
	file := file{name: path, content: "EXT4 File"}
	ext4.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}
	return ntfs.files[path]
}

func (ext4 ext4) FindFile(path string) file {
	if _, ok := ext4.files[path]; !ok {
		return file{}
	}
	return ext4.files[path]
}


func DatabaseFactory(env string) interface{} {
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


func FileSystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}


func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FileSystemFactory
	default:
		return nil
	}
}
