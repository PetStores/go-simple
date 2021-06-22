package main

import (
	"fmt"
	"net/http"
	"sync"
)

type statistics map[string]uint64

var mx *sync.Mutex
var st statistics = map[string]uint64{
	"a": 0,
	"b": 0,
	"c": 0,
}

func main() {
	mx = new(sync.Mutex)
	http.HandleFunc("/", stats())
	http.ListenAndServe(":8888", nil)
}

// Запрос: ?type=a ?type=b ?type=c
// Файл со статистикой:
// type; count
func stats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			saveSt(w, r)
			return
		case http.MethodGet:
			readSt(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func saveSt(w http.ResponseWriter, r *http.Request) {
	typ := r.URL.Query().Get("type")
	if !isValid(typ) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Type is not valid, should be a, b or c."))
		return
	}

	mx.Lock()
	defer mx.Unlock()
	st[typ]++
	w.WriteHeader(http.StatusCreated)
}

func readSt(w http.ResponseWriter, r *http.Request) {
	typ := r.URL.Query().Get("type")
	if !isValid(typ) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Type is not valid, should be a, b or c."))
		return
	}

	mx.Lock()
	defer mx.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Statistics for %s is %d", typ, st[typ])
}

func isValid(typ string) bool {
	switch typ {
	case "a", "b", "c":
		return true
	}

	return false
}
