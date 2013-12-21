package main

import (
	"fmt"
	"github.com/metasansana/pelau"
)

func handler(req pelau.Request, res pelau.Response) {

	fmt.Printf("[%s] request detected on path %s.\n", req.Method(), req.Url().Path)
	fmt.Printf("Printing Submatches: %v\n", req.Params())

}

func main() {

	pelau.Init(pelau.DefaultRouter()).
		Get(pelau.Static("/static/get", handler)).
		Post(pelau.Static("/static/post", handler)).
		Put(pelau.Static("/static/put", handler)).
		Head(pelau.Static("/static/head", handler)).
		Delete(pelau.Static("/static/delete", handler)).
		Get(pelau.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Post(pelau.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Put(pelau.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Head(pelau.Regex("/regex/[a-zA-Z0-9]+", handler)).
		Delete(pelau.Regex("/regex/([a-zA-Z0-9]+)", handler)).
		Bind(":8080")

}
