package main

import bhu_go_sdk "github.com/bhu/bhugo"

func main() {
	client := bhu_go_sdk.NewSession()
	err := neugo.Use(client).WithAuth("student_id", "password").Login(CAS)
	token := neugo.About(client).Token(CAS)
}
