package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	done := make(chan bool)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
		fmt.Println(string(dump))
	}))
	fmt.Println(ts.URL)
	<-done
	defer ts.Close()
}
