package main

import (
	"fmt"
	"io/ioutil"
)
//อ่านไฟล์
func main(){
	bs, err := ioutil.ReadFile("output.txt")
	if err != nil {
		return
	}
	
	fmt.Println(bs)
	str := string(bs)
	fmt.Println(str)
	
}