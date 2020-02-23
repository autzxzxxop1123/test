package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	
	
	
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

		for i := 0; i < len(G.Name()); i++ { 
			if G.Name()[i:i+1] == "." { 
				

				if G.Name()[i+1:] == typeFile { 
					tempFile := "path : " + dirRoot + " , name file : " + G.Name() + " ,  file size : " + strconv.FormatInt(int64(G.Size()), 10)
					arr = append(arr, tempFile) 
				}
			}
		}
	}
	return arr 
}
func createFile(pathFile []string) {
	f, err := os.Create("output.txt") 
	if err != nil {                   
		fmt.Println(err)
		return
	}
	for _, v := range pathFile { 
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("*****************************************")
	fmt.Println("*              completed                *")
	fmt.Println("*****************************************")

	err = f.Close() 
	if err != nil {
		fmt.Println(err)
		return
	}
}