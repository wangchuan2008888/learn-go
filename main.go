package main

import (
	"os"
	"errors"
	"fmt"
	"bufio"
)

func fileFromCommandLine()(string, error){
	if len(os.Args) <2{
		return "", errors.New("no file in")
	} else {
		return os.Args[1],nil
	}
}
func main()  {
	inFileName,erro := fileFromCommandLine()
	if erro != nil{
		fmt.Println(erro)
		os.Exit(-1)
	}
	inFile, erro := os.Open(inFileName)
	defer inFile.Close()
	bufInFile := bufio.NewReader(inFile)
	line, _,_ := bufInFile.ReadLine()
	fmt.Println(string(line))
}
