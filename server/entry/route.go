package entry

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//EnableRouteAdd enables the route add with given url extension
func (e *Entry) EnableRouteAdd(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		w.Write([]byte("entry added"))
	})
}
