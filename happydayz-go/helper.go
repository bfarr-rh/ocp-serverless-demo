package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func prettyPrint(req *http.Request) string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%v %v %v %v\n", req.Method, req.URL, req.Proto, req.Host)
	for k, vv := range req.Header {
		for _, v := range vv {
			fmt.Fprintf(b, "  %v: %v\n", k, v)
		}
	}

	if req.Method == "POST" {
		req.ParseForm()
		fmt.Fprintln(b, "Body:")
		for k, v := range req.Form {
			fmt.Fprintf(b, "  %v: %v\n", k, v)
		}
	}

	return b.String()
}

func parsPost(req *http.Request) string {
	var data Simple
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Can not umarshal thingy")
	}
	switch data.Oper {
	case "+":
		return fmt.Sprintf("{'response': '%v'}", strconv.Itoa(data.VarA-data.VarB))
	case "-":
		return fmt.Sprintf("{'response': '%v'}", strconv.Itoa(data.VarA-data.VarB))
	}
	return "UNKNOWN"
}
