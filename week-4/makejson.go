package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main () {
	var data = make(map[string]string)

	var inputName, inputAddress string
	fmt.Println("Enter your first name: ")
	fmt.Scan(&inputName)
	data["name"] = inputName
	fmt.Println("Enter your address: ")

	reader := bufio.NewReader(os.Stdin)
	inputAddress, err := reader.ReadString('\n')
	data["address"] = inputAddress

	datajson, err := json.Marshal(data)

	if err !=nil {
		fmt.Println("There was an error: ", err)
	}
	datajsonString := string(datajson)
	fmt.Println("The JSON object is: ", datajsonString)
	
}



