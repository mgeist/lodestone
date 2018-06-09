package lodestone

import (
	"io/ioutil"
	"log"
	"net/http"
)

type callback func([]byte)

var hooks = make(map[string]callback)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	hooks[r.URL.Path](body)
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
