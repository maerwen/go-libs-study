package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	test3()
}

func test1() {
	/*
		go run test.go -name kkk
		go run test.go
	*/
	name := flag.String("name", "maerwen", "Input your name")
	flag.Parse()
	fmt.Println("Welcome to here," + *name)
}
func test2() {
	/*
		go run test.go -a 2121 -bbb=2
		a=2121,地址：842350584024(0xc42001e0d8)
		参数个数：      4
		参数列表：
		0:      /tmp/go-build849455815/command-line-arguments/_obj/exe/test
		1:      -a
		2:      2121
		3:      -bbb=2
	*/
	a := flag.Int("a", 0, "a")
	var b int
	flag.IntVar(&b, "bbb", 0, "b")
	flag.Parse()
	fmt.Printf("a=%d,地址：%d(%v)\n", *a, a, a)
	fmt.Printf("参数个数：\t%d\n", len(os.Args))
	fmt.Printf("参数列表：\n")
	for i, j := range os.Args {
		fmt.Printf("%d:\t%s\n", i, j)
	}
}

type StringArray []string

func (s *StringArray) String() string {
	return fmt.Sprint([]string(*s))
}
func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func test3() {
	/*
		go run test.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0 -time 2m0s
		go run test.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0 -time 2m0s33
	*/
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	verFlag := flagSet.String("ver", "", "version")
	xtimeFlag := flagSet.Duration("time", 10*time.Minute, "time duration")
	addrFlag := StringArray{}
	flagSet.Var(&addrFlag, "a", "b")
	fmt.Println("os.Arg[0]:", os.Args[0])
	flagSet.Parse(os.Args[1:])
	fmt.Println("当前命令行参数类型个数：", flagSet.NFlag())
	// for i,j:=range flagSet.N
	for i := 0; i != flagSet.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("\n参数值:")
	fmt.Println("ver:", *verFlag)
	fmt.Println("xtimeFlag:", *xtimeFlag)
	fmt.Println("addrFlag:", addrFlag.String())
	for i, param := range flag.Args() {
		fmt.Printf("---#%d :%s\n", i, param)
	}
}
