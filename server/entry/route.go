package entry

import "net/http"

//EnableRouteAdd enables the route add with given url extension
func (e *Entry) EnableRouteAdd(urlExt string) {
	e.router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("entry added"))
	})
}
