package main

import (
	"Nodeprj/controllers/stdhttp"
	"Nodeprj/gate/psg"
	"fmt"
	"net/http"
)

func main() {
	s, err := psg.NewPsg("postgres://st10:123@localhost:5432/db1", "st10", "123")
	if err != nil {
		fmt.Println("Error")
	}
	stdhttp.NewController("localhost:8080", s)
	http.ListenAndServe(":8080", nil)
}
