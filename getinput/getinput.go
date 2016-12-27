package getinput

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func get(year int, day int, sessionKey string) ([]byte, error) {
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
	return contents, nil
}

func Get(year int, day int, sessionKey string) ([]string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("could not get current user: %v", err)
	}
	cacheFilename := filepath.Join(os.TempDir(), "advent", u.Username, strconv.Itoa(year), strconv.Itoa(day))

	var b []byte
	if _, err := os.Stat(cacheFilename); err != nil {
		// ensure read and write for current user
		if sessionKey == "" {
			return nil, errors.New("no session key provided")
		}
		err := os.MkdirAll(filepath.Dir(cacheFilename), os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("could not create advent cache dir: %v", err)
		}

		b, err = get(year, day, sessionKey)
		if err != nil {
			return nil, fmt.Errorf("error fetching puzzle: %v", err)
		}
		// give read-write to current user, read to others
		if err := ioutil.WriteFile(cacheFilename, b, 0644); err != nil {
			return nil, fmt.Errorf("error saving puzzle: %v", err)
		}

	} else {
		var err error
		b, err = ioutil.ReadFile(cacheFilename)
		if err != nil {
			return nil, fmt.Errorf("error reading cache: %v", err)
		}
	}

	ret := strings.Split(string(b), "\n")

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
