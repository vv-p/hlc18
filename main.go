package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Like struct {
	Id uint32 `json:"id"`
	Ts uint32 `json:"ts"`
}

type Account struct {
	Birth     uint32   `json:"birth"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Email     string   `json:"email"`
	Fname     string   `json:"fname"`
	Id        uint32   `json:"id"`
	Interests []string `json:"interests"`
	Joined    uint32   `json:"joined"`
	Likes     []Like   `json:"likes"`
	Phone     string   `json:"phone"`
	Sex       string   `json:"sex"`
	Sname     string   `json:"sname"`
	Status    string   `json:"status"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

func main() {
	var accounts Accounts

	jsonFile, err := os.Open("data/accounts_1.json")
	if err != nil {
		log.Fatal(err) // os.Exit(1) will be here too
	}
	defer jsonFile.Close()

	rawData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(rawData), &accounts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Accounts[0] = %s\n", accounts.Accounts[1])
}
