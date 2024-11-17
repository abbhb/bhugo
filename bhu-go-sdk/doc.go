// Package beihuago
//
// This package encapsulates some operations on the campus SSO service, in particular CAS, of beihua (cn).
//
// Examples:
//
// E1. Log in to the CAS using account and get the token.
//
//	// you can use your own *http.Client instead of creating a new session.
//	client := beihuago.NewSession()
//	err := beihuago.Use(client).WithAuth("student_id", "password").Login(CAS)
//	token := beihuago.About(client).Token(CAS)
//
// E2. Request a service via Web VPN using token.
//
//	  client := beihuago.NewSession()
//	  err := beihuago.Use(client).WithToken("your_webvpn_token").Login(WebVPN)
//	  serviceURL := "https://ipgw.beihua.edu.cn"
//		 reqURL := beihuago.EncryptURLToWebVPN(serviceURL)
//	  req := http.NewRequest("GET", reqURL, nil)
//	  resp, err := client.Do(req)
package bhu_go_sdk
