package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", stats())
	http.ListenAndServe(":8888", nil)
}

// Запрос: ?type=a ?type=b ?type=c
// Файл со статистикой:
// type; count
func stats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		typ := r.URL.Query().Get("type")
		if !isValid(typ) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Type is not valid, should be a, b or c."))
			return
		}

		for {
			fdlock, err := os.OpenFile("./lockfile", os.O_CREATE, 0600)
			if err == nil {
				defer func() {
					fdlock.Close()
					os.Remove("./lockfile")
				}()
				break
			}
			time.Sleep(3 * time.Millisecond)
		}

		fd, err := os.OpenFile("./dump.json", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Panicf("can't open file: %v", err)
		}
		defer fd.Close()

		dump := new(map[string]uint64)
		json.NewDecoder(fd).Decode(dump)

		switch r.Method {
		case http.MethodPost:
			saveSt(w, typ, dump)
			return
		case http.MethodGet:
			readSt(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func saveSt(w http.ResponseWriter, typ string, dump map[string]uint64) {
	if _, ok := dump[typ]; ok {
		dump[typ]++
	} else {
		dump[typ] = 1
	}

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
