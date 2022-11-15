package main

import "github.com/go-worknotes/code/design-patterns/abstractfactory"

// "github.com/go-worknotes/code/design-patterns/factory"


func main() {
	// env1 := "production"
	// env2 := "development"

	// db1 := factory.DatabaseFactory(env1)
	// db2 := factory.DatabaseFactory(env2)

	// db1.PutData("test", "this is mongoDB")
	// fmt.Println(db1.GetData("test"))

	// db2.PutData("test2", "this is sqlite")
	// fmt.Println(db2.GetData("test2"))


	// fmt.Println(reflect.TypeOf(db1).Name())
	// fmt.Println(reflect.TypeOf(&db1).Name())
	// fmt.Println(reflect.TypeOf(db2).Name())
	// fmt.Println(reflect.TypeOf(&db1).Name())


}


func SetupConstructors(env string) (abstractfactory.Database, abstractfactory.FileSystem) {
	fs := abstractfactory.AbstractFactory("filesystem")
	db := abstractfactory.AbstractFactory("database")
	return db(env).(abstractfactory.Database), fs(env).(abstractfactory.FileSystem)
}
