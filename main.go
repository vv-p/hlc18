package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func readChunk(jsonFile *os.File) {
	for {
		var b []byte

		n, err := jsonFile.Read(b)

		if err == io.EOF {
			fmt.Printf("EOF: n=%d\n", n)
			break
		}
	}
}

func toJson(rawJson string) {
	fmt.Printf("Wow: %s\n", rawJson)
}

func main() {
	// var accounts Accounts

	jsonFile, err := os.Open("data/accounts_1.json")
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Seek(14, 0)
	if err != nil {
		log.Panic(err)
	}

	var jsonStorage strings.Builder
	var leftIndex, rightIndex, searchStart, jsonStart, depth int
	var tail string
	b := make([]byte, 64)

	searchStart = 0
	depth = 0
	for {

		// read next chunk only we haven't any "{" or "}" in current data
		if s := jsonStorage.String(); !strings.ContainsAny(s[searchStart:], "{}") {
			_, err = jsonFile.Read(b)
			if err == io.EOF {
				break
			}
			jsonStorage.Write(b)
		}

		leftIndex = strings.Index(jsonStorage.String()[searchStart:], "{")
		rightIndex = strings.Index(jsonStorage.String()[searchStart:], "}")

		if depth == 0 && leftIndex > -1 {
			jsonStart = leftIndex
		}

		if leftIndex > -1 && rightIndex > -1 { // both "{" and "}" were found
			if leftIndex < rightIndex { // found "{" before "}", it means next level of json
				searchStart = searchStart + leftIndex + 1
				depth++
				continue
			}
			if leftIndex > rightIndex { // found "}" before "{", it means return to previous level
				searchStart = searchStart + rightIndex + 1
				depth--
			}
		}

		if leftIndex > -1 && rightIndex == -1 { // only "{" was found
			searchStart = searchStart + leftIndex + 1
			depth++
			continue
		}

		if leftIndex == -1 && rightIndex > -1 { // only "}" was found
			searchStart = searchStart + rightIndex + 1
			depth--
		}

		if depth == 0 { // we've complete json here in the jsonStorage[jsonStart:searchStart]
			fmt.Printf("JSON: %s\n", jsonStorage.String()[jsonStart:searchStart])

			tail = jsonStorage.String()[searchStart:]
			jsonStorage.Reset()
			jsonStorage.WriteString(tail)

			jsonStart = 0
			searchStart = 0
		}

	}
	//err = json.Unmarshal([]byte(rawData), &accounts)
	//if err != nil {
	//	log.Panic(err)
	//}

	//fmt.Printf("Accounts[0] = %s\n", accounts.Accounts[1])
}
