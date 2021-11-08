package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router()*mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	return router
}

func Test_getPosts (t *testing.T){
	req, _:= http.NewRequest("GET","/posts",nil)
	resp := httptest.NewRecorder()
	Router().ServeHTTP(resp,req)
	assert.Equal(t, http.StatusOK,resp.Code,"OK response is expected")
	fmt.Println(resp)
}