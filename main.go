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
	var leftIndex, rightIndex, currentIndex, startIndex, depth int 
	var tail string
	b := make([]byte, 4)

	currentIndex = -1
	depth = 0
	for {
		_, err = jsonFile.Read(b)
		if err == io.EOF {
			break
		}

		jsonStorage.Write(b)
		fmt.Printf("%d %s\n", len(jsonStorage.String()), jsonStorage.String())

		leftIndex  = strings.Index(jsonStorage.String()[currentIndex+1:], "{")
		rightIndex = strings.Index(jsonStorage.String()[currentIndex+1:], "}")

		if depth == 0 && leftIndex > -1 {
			startIndex = leftIndex
		}

		fmt.Printf("left = %d, right = %d, currentIndex = %d\n", leftIndex, rightIndex, currentIndex)

		if leftIndex > -1 && rightIndex > -1 {  // both "{" and "}" were found
			if leftIndex < rightIndex {
				currentIndex = leftIndex
				depth++
				continue
			}
			if leftIndex > rightIndex {
				currentIndex = rightIndex
				depth--
			}
		}

		if leftIndex > -1 && rightIndex == -1 {  // only "{" was found
			// ToDo: add here new chunk of file to b
			currentIndex = leftIndex
			depth++
			fmt.Printf("depth++ = %d\n", depth)
			continue
		}

		if leftIndex == -1 && rightIndex > -1 {  // only "}" was found
			currentIndex = rightIndex
			depth--
		}

		if depth == 0 {  //we've found json here
			fmt.Printf("JSON: %s\n", jsonStorage.String()[startIndex:rightIndex+startIndex+2])
			tail = jsonStorage.String()[rightIndex+startIndex+2:]
			jsonStorage.Reset()
			jsonStorage.WriteString(tail)
			currentIndex = -1
			leftIndex = -1
			rightIndex = -1

			fmt.Printf("tail=\"%s\" jsonStorage=\"%s\"\n", tail, jsonStorage.String())			
		}			

	}
	//err = json.Unmarshal([]byte(rawData), &accounts)
	//if err != nil {
	//	log.Panic(err)
	//}

	//fmt.Printf("Accounts[0] = %s\n", accounts.Accounts[1])
}
