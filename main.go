package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	indexId  = MakeIndexId()
	indexSex = MakeIndexSex()
)

const (
	BaseDir   = "data"
	Delimeter = "/"
	httpPort  = ":80"
)

type (
	AccountId uint32

	Like struct {
		Id AccountId `json:"id"`
		Ts uint32    `json:"ts"`
	}

	Account struct {
		Birth     uint32    `json:"birth"`
		City      string    `json:"city"`
		Country   string    `json:"country"`
		Email     string    `json:"email"`
		Fname     string    `json:"fname"`
		Id        AccountId `json:"id"`
		Interests []string  `json:"interests"`
		Joined    uint32    `json:"joined"`
		Likes     []Like    `json:"likes"`
		Phone     string    `json:"phone"`
		Sex       string    `json:"sex"`
		Sname     string    `json:"sname"`
		Status    string    `json:"status"`
	}
)

func parseJson(filename string) {

	// ToDo: add bufio.Reader to read from json file
	jsonFile, err := os.Open(filename)
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
			jsonReady := jsonStorage.String()[jsonStart:searchStart]

			account := Account{}
			err = json.Unmarshal([]byte(jsonReady), &account)

			if err != nil {
				log.Panic(err)
			}

			indexId.Add(&account)
			indexSex.Add(&account)

			tail = jsonStorage.String()[searchStart:]
			jsonStorage.Reset()
			jsonStorage.WriteString(tail)

			jsonStart = 0
			searchStart = 0
		}
	}

}

func main() {

	files, err := ioutil.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "accounts") {
			parseJson(strings.Join([]string{BaseDir, file.Name()}, Delimeter))
		}
	}

	log.Printf("Accounts total: %d\n", indexId.Len())
	log.Printf("Starting http server\n")

	http.HandleFunc("/", httpMultiplexer)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
