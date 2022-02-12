package accounts

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/andygeiss/meridian59-build/pkg/maintenance"
)

// Show ...
func Show() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// init connection to maintenance port
		m := maintenance.NewHandler()
		m.Connect("127.0.0.1:59595")
		defer m.Close()
		// get players online
		m.Send("who")
		out := m.Receive()
		// extract toon names from the who list
		var toons []string
		lines := strings.Split(out, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Game -") {
				parts := strings.Split(line, "Game -")
				toonWithPort := strings.Split(parts[1], "(")
				toons = append(toons, toonWithPort[0])
			}
		}
		// write json encoded result
		json.NewEncoder(rw).Encode(toons)
	}
}
