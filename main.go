package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/:link", testHandle)

	http.ListenAndServe(":3000", router)
}

func testHandle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	link := ps.ByName("link")
	http.Redirect(w, req, link, http.StatusSeeOther)
}
