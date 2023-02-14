package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestServerInit(t *testing.T) {

	indx := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	}

	hll := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	}

	router := httprouter.New()
	router.GET("/", indx)
	router.GET("/hello/:name", hll)
	//router.METHOD

	//t.Log(http.ListenAndServe(":8080", router))
}
