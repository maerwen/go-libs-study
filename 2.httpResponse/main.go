package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", handler6)
	http.ListenAndServe(":8080", nil)
}

// 利用curl辅助测试
// 不需要返回任何的数据，只是返回一个header
func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "a go web server!")
	w.WriteHeader(200)
}

// 返回文本
func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am a Gopher!\n"))
}

// 返回json
type Person struct {
	Name string `json:"name"`
	Age  int    `json:",omitempty"`
}

func handler3(w http.ResponseWriter, r *http.Request) {
	p := new(Person)
	p = &Person{
		Name: "韩立", Age: 15,
	}
	result, err := json.Marshal(p)
	if err != nil {
		log.Println("json marshal error:", err)
		return
	}
	w.Write(result)
	w.Write([]byte("\n"))
}

// 返回xml
type Toy struct {
	Id      int `xml:",attr"`
	Name    string
	Comment string `xml:",cdata"`
}

func handler4(w http.ResponseWriter, r *http.Request) {
	var t Toy
	t = Toy{
		Id:      1,
		Name:    "rabbit",
		Comment: "她很爱这只兔子",
	}
	result, err := xml.MarshalIndent(t, "", "   ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
	w.Write([]byte("\n"))
}

// 返回文件到页面
func handler5(w http.ResponseWriter, r *http.Request) {
	file := path.Join("2.httpResponse/images", "20170502141600329.gif")
	// file := path.Join("2.httpResponse/images", "下载.png")
	http.ServeFile(w, r, file)
}

// 返回html页面
type Profile struct {
	Name    string
	Hobbies []string
}

func handler6(w http.ResponseWriter, r *http.Request) {
	p := Profile{Name: "韩立", Hobbies: []string{"杀人夺宝", "跑路"}}
	fp := path.Join("2.httpResponse/templates", "index.html")
	t, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
