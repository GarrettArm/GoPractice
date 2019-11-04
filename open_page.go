// opens a web browser to a specific page

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type userinfo struct {
	Expiration string `json:"expiration"`
}

func main() {
	url := "http://localhost:8000/?id="
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	user := userinfo{}
	jsonErr := json.Unmarshal(b, &user)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1)
	}
	exp, err := time.Parse("2006-01-02", user.Expiration)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if exp.After(time.Now()) {
		fmt.Println("active")
	} else {
		fmt.Println("inactive")
	}
}
