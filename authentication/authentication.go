package authentication

import (
	"io/ioutil"
	"strings"
)

//read from token file
//store the token in some sort of static variable so that i do not have to re-read it
// or read the file at the start of running and store token in memory

func readFile() (string, error){
	data, err := ioutil.ReadFile("./authentication/token.txt")
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func GetToken() (string, error){
	token, err := readFile()
	if err != nil{
		return "", err
	}
	return token, nil
}
