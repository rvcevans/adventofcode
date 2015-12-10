package getinput

import (
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

func Get(day int, sessionKey string) ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://adventofcode.com/day/%d/input", day), nil)

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

func MustGet(day int, sessionKey string) []string {
	in, err := Get(day, sessionKey)
	if err != nil {
		panic(err)
	}

	return in
}
