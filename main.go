package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Hash struct {
	Name string
	John string
	Hashcat int32
	Extended bool
}

type HashTest struct {
	Regex string `json:"regex"`
	Modes []Hash
}

func (hashTest *HashTest) test(input string) bool {
	match, _ := regexp.MatchString(hashTest.Regex, input)

	return match
}

func createListOfHashes() ([]HashTest, error) {
	jsonFile, err := os.Open("prototypes.json")

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	var hashTests []HashTest
	err = json.Unmarshal(byteValue, &hashTests)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return hashTests, nil
}

func identify(input string) ([]Hash, error) {
	if input == "" {
		return nil, errors.New("no input passed")
	}

	hashTests, err := createListOfHashes()

	if err != nil {
		return nil, err
	}

	var list []Hash

	for _, hashTest := range hashTests {
		if hashTest.test(input) {
			list = append(list, hashTest.Modes...)
		}
	}

	return list, nil
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func main() {
	hash := "$mongodb$0$sa$75692b1d11c072c6c79332e248c4f699"
	listOfMatchingHashes, err := identify(hash)

	if err != nil {
		log.Fatal(err)
	}

	_ = PrettyPrint(listOfMatchingHashes)
}
