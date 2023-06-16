package prototest

import (
	"fmt"
	"io/ioutil"
	"vbi-cosmos-basic/example"

	"github.com/golang/protobuf/proto"
)

func Execution() {
	person := &example.Person{
		Name: "Thang",
		Age:  22,
	}

	// Encode Person object to byte array
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshal:", err)
		return
	}

	// Write byte array to file
	err = ioutil.WriteFile("person.dat", data, 0644)
	if err != nil {
		fmt.Println("Error write file:", err)
		return
	}

	// Read byte array from file
	data, err = ioutil.ReadFile("person.dat")
	if err != nil {
		fmt.Println("Errro read file:", err)
		return
	}

	// Decode byte array to Person object
	var restoredPerson example.Person
	err = proto.Unmarshal(data, &restoredPerson)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	// Show info Person ojbect
	fmt.Println("Name:", restoredPerson.Name)
	fmt.Println("Age:", restoredPerson.Age)
}
