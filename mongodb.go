package database

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)

	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(
		&Person{"sam", "12345733"},
		&Person{"peu", "94043785"},
	)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "sam"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Phone for %s is %s", result.Name, result.Phone)
}
