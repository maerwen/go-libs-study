package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

var ch = make(chan []byte)

func main() {
	encode()
	// go encode()
	// decode()
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
	Address1    Address   `xml:",omitempty"`
	Address2    Address   `xml:",chardata"`
	Address3    Address   `xml:",innerxml"`
	Comment     string    `xml:",comment"`
	Description string    `xml:",cdata"`
	Addresses   []Address `xml:"hobbies"`
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
	p.Address1 = Address{
		Province: "四川",
		City:     "成都",
	}
	p.Address2 = Address{
		Province: "上海",
		City:     "上海",
	}
	p.Address3 = Address{
		Province: "广东",
		City:     "深圳",
	}
	// <hobbies>足球</hobbies>
	// <hobbies>唱歌</hobbies>
	// <hobbies>跳舞</hobbies>
	p.Addresses = []Address{p.Address1, p.Address2, p.Address3}
	// result, err := xml.Marshal(p)
	result, err := xml.MarshalIndent(p, " ", "   ")
	if err != nil {
		log.Println("xml marshal err:", err)
	}
	fmt.Printf("%s\n", result)
	// ch <- result
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
