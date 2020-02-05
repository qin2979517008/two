package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Student struct {
	Xh string `json:"xh"`
	Xm string `json:"xm"`
}

func main() {
	var name Student
	file, err := os.OpenFile("Students.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Open file error!")
	}
	str, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	regall := `{"xh":"\d{7}573[\s\S]{21}`
	rega := regexp.MustCompile(regall)
	a := rega.FindAllStringSubmatch(string(str), 1)
	js := append(a[0], "}")
	justString := strings.Join(js, " ")

	 
	err = json.Unmarshal([]byte(justString), &name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("需要找的人姓名为", name.Xm)

}
