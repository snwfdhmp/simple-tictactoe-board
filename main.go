package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	var field [3][3]int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			formToInt := func(key string) int {
				return int(r.FormValue(key)[0] - '0')
			}
			x, y := formToInt("x"), formToInt("y")
			field[x][y] = (field[x][y] + 1) % 3
		}
		json.NewEncoder(w).Encode(field)
	})
	http.ListenAndServe(":8081", nil)
}
