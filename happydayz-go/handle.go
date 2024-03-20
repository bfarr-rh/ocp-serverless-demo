package function

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	App     string `json:"app"`
	Message string `json:"message"`
}

type Simple struct {
	VarA int    `json:"vara"`
	VarB int    `json:"varb"`
	Oper string `json:"oper"`
}

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	var payload Response
	fmt.Println("Received request")
	fmt.Println(prettyPrint(req))
	// fmt.Println(req.URL.Path) // echo to local output
	// fmt.Fprint(res, prettyPrint(req)) // echo to caller
	if msg := os.Getenv("MESSAGE"); msg != "" {
		payload = Response{
			Message: msg,
			App:     "GoLang Serverless",
		}
	} else {
		payload = Response{
			Message: "I 💕 GoLang Serverless",
			App:     "GoLang Serverless",
		}
	}
	jsonResponse, _ := json.Marshal(payload)

	if req.Method == "Post" {
		res.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(parsPost(req))
		res.Write(jsonResponse)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}
