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
	c := townsita.NewConfig()
	if err := c.Load(os.Args); err != nil {
		log.Fatal(err)
	}
	handler := townsita.New(c, da).GetHTTPHandler()
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
