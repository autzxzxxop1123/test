package main

import (
	"fmt"
	"io/ioutil"
	
	
	
)

func main(){
	var searchPath string
	var typeFile string
	fmt.Println("Please Enter Your SearchPath")
	fmt.Scanln(&searchPath)
	fmt.Println("Please Enter Your typeFile")
	fmt.Scanln(&typeFile)
}

func searchFile(dirRoot string, typeFile string) []string {

	arr := []string{}
	files, err := ioutil.ReadDir(dirRoot) 
	if err != nil {                       
		return []string{"Search file not found !"} 
	}
	for _, G := range files { 

		}
	
	return arr 
}

