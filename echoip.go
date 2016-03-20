package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "",
		"Address on which to serve. [default: :{{port}}]")
	port := flag.Int("port", 8888,
		"Port on which to serve. (Ignored if -addr is present) [default: 8888]")
	flag.Parse()

	if len(*addr) == 0 {
		*addr = fmt.Sprintf(":%d", *port)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.RemoteAddr)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
