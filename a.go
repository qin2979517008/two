package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Student struct {
	Xh string `json:"xh"`
	Xm string   `json:"xm"`
	XmEn string  `json:"xmEn"`
	Xb string  `json:"xb"`
	Bj string  `json:"bj"`
	Zyh string   `json:"zyh"`
	Zym string  `json:"zym"`
	Yxh string   `json:"yxh"`
	Yxm string	`json:"yxm"`
	Nj string	`json:"nj"`
	Csrq string	`json:"csrq"`
	Xjzt string	`json:"xjzt"`
	Rxrq string	`json:"rxrq"`
	Yxmen string	`json:"yxmen"`
	ZymEn string	`json:"zymEn"`
	Xz int	`json:"xz"`
	Mz string	`json:"mz"`
}
func main() {
	var name Student
	 file, err := os.OpenFile("Students.txt", os.O_RDONLY, 0644)
	 if err != nil {
 	fmt.Println("Open file error!")
	 }
	 str,err:= ioutil.ReadAll(file)
	 if err !=nil{
 	fmt.Println(err)
	 }
	err =json.Unmarshal([]byte(str), &name)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(name.Xm)
}
