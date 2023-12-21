package main

import (
	"net/http"
	"github.com/Didjacome/varacode-api/pkg/load"
	"github.com/Didjacome/varacode-api/pkg/apiRequest"
)

func main() {
	load.Load()
	cr := load.GetApiCredetial()
	r :=  apirequest.ApiRequest(cr.ApiKeyID, cr.ApiKeySecret, cr.ApiUrl, http.MethodGet)
	print(r)
}