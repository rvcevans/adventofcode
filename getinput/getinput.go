package getinput

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
)

func Get(year int, day int, sessionKey string) ([]string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://adventofcode.com/%d/day/%d/input", year, day)
	log.Printf("Fetching input: %s", url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionKey})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := strings.Split(string(contents), "\n")

	if len(ret[len(ret)-1]) == 0 {
		// Last one is empty, discard
		ret = ret[:len(ret)-1]
	}

	return ret, nil
}

func MustGet(year int, day int, sessionKey string) []string {
	in, err := Get(year, day, sessionKey)
	if err != nil {
		panic(err)
	}

	return in
}
