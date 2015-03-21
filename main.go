package main

import (
	"github.com/Townsita/data-adapter-dummy"
	"github.com/Townsita/townsita"
	"log"
	"net/http"
	"os"
)

func main() {
	da := dummy.New()
	handler := townsita.New(da).GetHTTPHandler(os.Args)
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
