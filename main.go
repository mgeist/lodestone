package lodestone

import (
	"log"
	"net/http"
)

type callback func(*http.Request)

var hooks = make(map[string]callback)

func handler(w http.ResponseWriter, r *http.Request) {
	hooks[r.URL.Path](r)
}

func RegisterEndpoint(endpoint string, fn callback) {
	hooks[endpoint] = fn
}

func StartServer(addr string) {
	for hook, _ := range hooks {
		http.HandleFunc(hook, handler)
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
