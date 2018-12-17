package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

const (
	BaseDir = "/tmp/out"
	Delimeter = "/"
	httpPort = ":80"
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
	accounts map[uint32]Account
}

func getNewAccounts() Accounts {
	return Accounts{accounts: map[uint32]Account{}}
}

func (a *Accounts) Add(account Account) {
	a.accounts[account.Id] = account
}

func (a *Accounts) Get(id uint32) Account {
	return a.accounts[id]
}

func (a *Accounts) Len() int {
	return len(a.accounts)
}

func parseJson(filename string, accounts *Accounts) {
	var account Account

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
			err = json.Unmarshal([]byte(jsonReady), &account)

			if err != nil {
				log.Panic(err)
			}

			accounts.Add(account)

			tail = jsonStorage.String()[searchStart:]
			jsonStorage.Reset()
			jsonStorage.WriteString(tail)

			jsonStart = 0
			searchStart = 0
		}
	}

}

func main() {
	accounts := getNewAccounts()

	files, err := ioutil.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "accounts") {
			parseJson(strings.Join([]string{BaseDir, file.Name()}, Delimeter), &accounts)
		}
	}

	log.Printf("Accounts total: %d\n", accounts.Len())
	log.Printf("Starting http server\n")

	router := httprouter.New() // https://github.com/julienschmidt/httprouter

	// Default 404 page
	router.GET("/", defaultNotFound)

	// Readers
	// It's not possible to register wildcard route here without some tricks
	// See https://github.com/julienschmidt/httprouter/issues/175 for details
	router.GET("/accounts/:id/", getMultiplexer)
	router.GET("/accounts/:id/recommend/", accountsRecommend)
	router.GET("/accounts/:id/suggest/", accountsSuggest)

	// Writers
	router.POST("/accounts/:id/", postMultiplexer)
	
	log.Fatal(http.ListenAndServe(httpPort, router))
}



func getMultiplexer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// filter and group

	switch ps.ByName("id") {
	case "filter":
		accountsFilter(w, r, ps)
	case "group":
		accountsGroup(w, r, ps)
	default:
		defaultNotFound(w, r, ps)
	}
}

func postMultiplexer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// new, likes and id

	switch ps.ByName("id") {
	case "new":
		accountsNew(w, r, ps)
	case "likes":
		accountsLikes(w, r, ps)
	default:
		accountsId(w, r, ps)
	}
}

// Readers

func accountsFilter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"accounts\": [ ]}")
}

func accountsGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"groups\": [ ]}")
}

func accountsRecommend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"accounts\": [ ]}")
}

func accountsSuggest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{\"accounts\": [ ]}")
}

// Writers

func accountsNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{}")
	w.WriteHeader(http.StatusCreated)
}

func accountsId(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{}")
	w.WriteHeader(http.StatusAccepted)
}

func accountsLikes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "{}")
	w.WriteHeader(http.StatusAccepted)
}

// Default 404 page
func defaultNotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("request: %s\n", r.URL.Path)
	http.NotFound(w, r)
}
