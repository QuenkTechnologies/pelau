package main

import (
	"fmt"
	"github.com/quenktech/pelau"
	"github.com/quenktech/pelau/mid"
)

func handler(req pelau.Request, res pelau.Response) {

	fmt.Printf("[%s] request detected on path %s.\n", req.Raw().Method, req.Raw().URL.Path)
	fmt.Printf("Printing Submatches: %v\n", req.Params())
	res.Head("X-Men", "Was a good show")
	//	res.Header("Content-Type", "application/json")
	res.Send(struct{ FIELD string }{"NAME"})

}

func main() {

	pelau.DefaultRouter().
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
		Use(mid.JSONEncoding).
		Use(func(req pelau.Request, res pelau.Response, c *pelau.Context) {

		fmt.Println("This message will print regardless of what request method was used.")
		c.Next(req, res, c)
	}).
		Bind("127.0.0.1:8080", func() {

		fmt.Println("Sever started.")

	})

}
