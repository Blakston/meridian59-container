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

// Online ...
func Online() http.HandlerFunc {
	host := os.Getenv("M59_HOST")
	port := os.Getenv("M59_PORT")
	return func(rw http.ResponseWriter, r *http.Request) {
		// init connection to maintenance port
		m := maintenance.NewHandler()
		addr := fmt.Sprintf("%s:%s", host, port)
		m.Connect(addr)
		defer m.Close()
		// get players online
		m.Send("who")
		out := m.Receive()
		if err := m.Error(); err != nil {
			log.Printf("error: %v", err)
		}
		// extract toon names from the who list
		toons := []string{}
		lines := strings.Split(out, "\n")
		for _, line := range lines {
			line = strings.ReplaceAll(line, "\u0000", "")
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
