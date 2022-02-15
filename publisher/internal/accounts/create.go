package accounts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/andygeiss/meridian59-build/pkg/maintenance"
)

// Create ...
func Create() http.HandlerFunc {
	host := os.Getenv("M59_HOST")
	port := os.Getenv("M59_PORT")
	return func(rw http.ResponseWriter, r *http.Request) {
		// init connection to maintenance port
		m := maintenance.NewHandler()
		addr := fmt.Sprintf("%s:%s", host, port)
		m.Connect(addr)
		defer m.Close()
		// define request and response
		var request struct {
			User  string `json:"user"`
			Pass  string `json:"pass"`
			Pass2 string `json:"pass2"`
			EMail string `json:"email"`
		}
		var response struct {
			Error string `json:"error"`
		}
		// decode the JSON body
		json.NewDecoder(r.Body).Decode(&request)
		// create an account
		cmd := fmt.Sprintf("create account user %s %s %s", request.User, request.Pass, request.EMail)
		m.Send(cmd)
		out := m.Receive()
		if err := m.Error(); err != nil {
			log.Printf("error: %v", err)
		}
		lines := strings.Split(out, "\n")
		if len(lines) > 0 {
			// handle errors
			if strings.Contains(lines[1], "already exists") {
				lines[1] = strings.ReplaceAll(lines[1], "\u0000", "")
				response.Error = lines[1]
			}
			// handle success
			if strings.Contains(out, "Created") {
				parts := strings.Split(lines[1], " ")
				numPoint := strings.Split(parts[2], ".")
				m.Send(fmt.Sprintf("create user %s", numPoint[0]))
				m.Receive()
				m.Send(fmt.Sprintf("create user %s", numPoint[0]))
				m.Receive()
			}
			// send response
		}
		json.NewEncoder(rw).Encode(response)
	}
}
