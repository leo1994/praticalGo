package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/leo1994")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	var r struct {
		Name         string `json:"name,omitempty"`
		Public_Repos int    `json:"public_repos,omitempty"`
	}

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("%#v\n", r)
	fmt.Println(githubInfo("leo1994"))
}

func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	var r struct {
		Name         string `json:"name,omitempty"`
		Public_Repos int    `json:"public_repos,omitempty"`
	}

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.Public_Repos, nil

}
