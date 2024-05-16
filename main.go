package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SibaPrasadTech/go-http-standard/middleware"
)

func main() {
	fmt.Println("Hello World!")
	router := http.NewServeMux()
	middlewareRouter := middleware.Logging(router)
	router.HandleFunc("/ping",func (rs http.ResponseWriter, req *http.Request){
		rs.Write([]byte("PING PONG"))
	})
	router.HandleFunc("GET /intpath/{int}",func (rs http.ResponseWriter, req *http.Request){
		pathValue := req.PathValue("int")
		intPath, err := strconv.Atoi(pathValue);
		if (err != nil){
			rs.WriteHeader(http.StatusBadRequest)
			rs.Write([]byte(fmt.Sprintf("The Path Value for intpath : %v is an invalid integer",pathValue)))
			return
		}
		rs.WriteHeader(http.StatusOK)
		rs.Write([]byte(fmt.Sprintf("The Path Value for intpath is : %v",intPath)))
	})
	router.HandleFunc("/strpath/{str}",func (rs http.ResponseWriter, req *http.Request){
		pathValue := req.PathValue("str")
		rs.WriteHeader(http.StatusOK)
		rs.Write([]byte(fmt.Sprintf("The Path Value for strpath is : %v",pathValue)))
	})
	router.HandleFunc("POST /strpath/{str}",func (rs http.ResponseWriter, req *http.Request){
		pathValue := req.PathValue("str")
		rs.WriteHeader(http.StatusOK)
		rs.Write([]byte(fmt.Sprintf("{POST strpath : %v",pathValue)))
	})
	server := http.Server{
		Addr: ":8000",
		Handler: middlewareRouter,
	}
	fmt.Println("Server listening on Port :8000")
	server.ListenAndServe()
}