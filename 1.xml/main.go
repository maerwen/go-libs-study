package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

var ch = make(chan []byte)

func main() {
	go encode()
	decode()
}

type Address struct {
	Province string
	City     string
}
type Person struct {
	// XMLName     xml.Name `xml:"person"`
	Id          int    `xml:",attr"`
	FirstName   string `xml:"name>first"`
	LastName    string `xml:"name>last"`
	Age         int    `xml:",omitempty"`
	Sex         bool
	Address     Address `xml:",omitempty"`
	Comment     string  `xml:",comment"`
	Description string  `xml:",cdata"`
}

func encode() {
	var p = Person{
		Id:          1,
		FirstName:   "韩",
		LastName:    "立",
		Age:         18,
		Sex:         true,
		Comment:     "天机不可泄露",
		Description: "天择",
	}
	p.Address = Address{
		Province: "四川",
		City:     "成都",
	}
	// result, err := xml.Marshal(p)
	result, err := xml.MarshalIndent(p, " ", "   ")
	if err != nil {
		log.Println("xml marshal err:", err)
	}
	// fmt.Printf("%s\n", result)
	ch <- result
}
func decode() {
	result := <-ch
	var p Person
	// xml.Unmarshal第二个参数必须是指针
	err := xml.Unmarshal(result, &p)
	if err != nil {
		log.Println("xml unmarshal err:", err)
	}
	fmt.Println("successful!")
	fmt.Printf("%s\n", result)
}
