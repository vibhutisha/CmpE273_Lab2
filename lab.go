package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Greet string `json:"greeting"`
}

func (res *Response) UnmarshalJSON(data []byte) error {
	var req Request

	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}

	res.Greet = "Hello, " + req.Name + "!!!"

	return nil
}

func post_h(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	v := Request{}

	json.NewDecoder(r.Body).Decode(&v)

	uj, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	s := uj

	var m Response
	if err := json.Unmarshal(s, &m); err != nil {
		fmt.Println(err)
		return
	}

	un, _ := json.Marshal(m)
	fmt.Fprintf(w, "%s", un)

}

func getuser(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func main() {

	r := httprouter.New()

	r.GET("/hello/:name", getuser)
	r.POST("/hello", post_h)

	http.ListenAndServe("localhost:8080", r)
}
