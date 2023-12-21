package main

import (
	"net/http"
	"github.com/Didjacome/varacode-api/pkg/load"
	"github.com/Didjacome/varacode-api/pkg/apiRequest"
)

func main() {
	err := load.Load()
	if err != nil {
		panic(err)
	}

	cr := load.GetApiCredetial()
	r :=  apirequest.ApiRequest(cr.ApiKeyID, cr.ApiKeySecret, cr.ApiUrl, http.MethodGet)
	print(r)
}