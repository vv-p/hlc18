package main

import (
	"log"
	"net/http"
	"strconv"
)

func httpMultiplexer(w http.ResponseWriter, r *http.Request) {
	var (
		limit uint64
	)

	q := r.URL.Query()
	limit, _ = strconv.ParseUint(q.Get("limit"), 10, 64) // default is zero
	sex_eq := q.Get("sex_eq")

	as := indexSex.Get(sex_eq, limit)
	log.Printf("as = %q\n", as)
	w.Write(buildResponse(as))
}

func buildResponse(as []*Account) []byte {
	response := make([]byte, 0, 128)
	response = append(response, "{\"accounts\": ["...)

	for _, a := range as {
		response = strconv.AppendUint(response, uint64(a.Id), 10)
		response = append(response, ", "...)
	}

	response = append(response, "]}"...)
	return response
}
