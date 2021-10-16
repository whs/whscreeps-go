package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type codeBody struct {
	Branch  string                 `json:"branch"`
	Modules map[string]interface{} `json:"modules"`
}

type binaryModule []byte

func (m binaryModule) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{\"binary\":\"")
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write(m)
	encoder.Close()
	buf.WriteString("\"}")

	return buf.Bytes(), nil
}

func main() {
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")
	token := flag.String("token", "", "")
	branch := flag.String("branch", "sim", "Branch to upload to")
	wasmAsJs := flag.Bool("wasm-as-js", false, "Encode WASM as JS file")
	flag.Parse()

	body := codeBody{
		Branch:  *branch,
		Modules: map[string]interface{}{},
	}

	for _, file := range flag.Args() {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		moduleName := path.Base(file)
		if path.Ext(file) == ".js" {
			moduleName = strings.Replace(moduleName, path.Ext(moduleName), "", 1)
			body.Modules[moduleName] = string(content)
		} else {
			if *wasmAsJs {
				body.Modules[moduleName] = encodeBinaryAsJs(content)
			} else {
				body.Modules[moduleName] = binaryModule(content)
			}
		}
	}

	rawBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Payload size: %d\n", len(rawBody))

	req, _ := http.NewRequest(http.MethodPost, "https://screeps.com/api/user/code", bytes.NewBuffer(rawBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "whscreeps/2.0 cmd/deploy")
	if *username != "" && *password != "" {
		req.SetBasicAuth(*username, *password)
	}
	if *token != "" {
		req.Header.Set("X-Token", *token)
	}
	out, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(out.Status)
	respBody, _ := ioutil.ReadAll(out.Body)
	fmt.Println(string(respBody))
}

func encodeBinaryAsJs(content []byte) string {
	var buf bytes.Buffer

	buf.WriteString("module.exports=Uint8Array.from([")

	first := true
	for _, b := range content {
		if first {
			first = false
		} else {
			buf.WriteRune(',')
		}
		buf.WriteString(strconv.FormatUint(uint64(b), 10))
	}

	buf.WriteString("])")

	return buf.String()
}
