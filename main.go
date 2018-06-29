package main

import (
	"os"
	"errors"
)

func fileFromCommandLine() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("no file in")
	} else {
		return os.Args[1], nil
	}
}

type Product struct {
	name string
	url  string
}

func Test(a []string) []string {
	a = append(a, "d", "e")
	return a
}
func main() {
	//inFileName,erro := fileFromCommandLine()
	//if erro != nil{
	//	fmt.Println(erro)
	//	os.Exit(-1)
	//}
	//inFile, erro := os.Open(inFileName)
	//defer inFile.Close()
	//bufInFile := bufio.NewReader(inFile)
	//line, _,_ := bufInFile.ReadLine()
	//fmt.Println(string(line))
	var a Product = Product{}
	//a.name = "aaa"
	//a.url = "bbb"
	//fmt.Println(a)
	//fmt.Printf("%v", a)
	//a := [...]string{"a","b","c"}
	//fmt.Println(a)
	//b := a[:]
	//fmt.Println(b)
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(b))
	//c:=Test(b)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
}
