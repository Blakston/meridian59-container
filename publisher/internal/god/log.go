package god

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Log ...
func Log() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// define request and response
		var lines []string

		content, _ := ioutil.ReadFile("server/channel/god.txt")

		lines = strings.Split(string(content), "\n")
		reverse(lines)
		lines = filter(lines)

		// send response
		json.NewEncoder(rw).Encode(lines)
	}
}

func filter(lines []string) (filtered []string) {
	for _, line := range lines {
		// filter who commands
		if strings.Contains(line, "Session: M Command:") {
			continue
		}
		if strings.Contains(line, "create user") {
			continue
		}
		filtered = append(filtered, line)
	}
	return
}

func reverse(slice []string) {
	last := len(slice) - 1
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[last-i] = slice[last-i], slice[i]
	}
}
