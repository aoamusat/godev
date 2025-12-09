// Server2 is a minimal "echo" and counter server.
package webserver3

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func requestHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", request.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
