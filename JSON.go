package pelau

import (
	"encoding/json"
	"log"
	"net/http"
)

type goson struct {
	v interface{}
}

func (js *goson) Output(w http.ResponseWriter) {

	bits, err := json.Marshal(js.v)

	if err != nil {

		log.Fatal(err)
	}
	w.Write(bits)

}

func JSON(v interface{}) Data {

	return &goson{v}

}
