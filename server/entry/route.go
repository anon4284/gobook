package entry

import (
	"encoding/json"
	"net/http"
	"projects/resellbay/server/util"

	"github.com/gorilla/mux"
)

//EnableRouteAdd enables the route add with given url extension
func (e *Entry) EnableRouteAdd(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var input EntryAddInput
		err := decoder.Decode(&input)
		util.CheckErr(err)

		output := validate(input)

		if output.Valid {
			_, err := e.Add(input)
			if err != nil {
				b, _ := json.Marshal(EntryAddOutput{false, "db", "Email already exists"})
				w.Write(b)
			} else {
				b, _ := json.Marshal(output)
				w.Write(b)
			}

		} else {
			b, _ := json.Marshal(output)
			w.Write(b)
		}

	})
}

//EnableRouteGet to get all the guestbook entries
func (e *Entry) EnableRouteGet(router *mux.Router, urlExt string) {
	router.HandleFunc(urlExt, func(w http.ResponseWriter, r *http.Request) {
		items, err := e.Scan()
		if err != nil {
			w.Write([]byte(`{"Valid":"false"}`))
		} else {

			b, _ := json.Marshal(items)
			w.Write(b)
		}

	})
}
