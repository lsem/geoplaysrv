package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/geo/s2"
)

func isFloat(v string) bool {
	_, err := strconv.ParseFloat(v, 64)
	return err == nil
}

func isInt(v string) bool {
	_, err := strconv.ParseInt(v, 10, 32)
	return err == nil
}

func asFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func asInt(v string) int {
	i, _ := strconv.ParseInt(v, 10, 32)
	return int(i)
}

// ApproxResponse is response to approxRect/approxCircle queries.
type ApproxResponse struct {
	CellIDs []string `json:"cellIDs"`
}

func approxRect(w http.ResponseWriter, r *http.Request) {
	south := r.URL.Query().Get("south")
	west := r.URL.Query().Get("west")
	north := r.URL.Query().Get("north")
	east := r.URL.Query().Get("east")
	minLvl := r.URL.Query().Get("minLvl")
	maxLvl := r.URL.Query().Get("maxLvl")
	maxCells := r.URL.Query().Get("maxCells")

	if len(minLvl) == 0 || !isInt(minLvl) {
		http.Error(w, "minLvl missing or bad", 400)
		return
	}
	if len(maxLvl) == 0 || !isInt(maxLvl) {
		http.Error(w, "maxLvl missing or bad", 400)
		return
	}
	if len(maxCells) == 0 || !isInt(maxCells) {
		http.Error(w, "maxCells missing or bad", 400)
		return
	}
	fmt.Println("maxCells: ", maxCells)
	if len(south) == 0 || !isFloat(south) {
		http.Error(w, "south missing", 400)
		return
	}
	if len(west) == 0 || !isFloat(west) {
		http.Error(w, "west missing", 400)
		return
	}
	if len(north) == 0 || !isFloat(north) {
		http.Error(w, "north missing", 400)
		return
	}
	if len(east) == 0 || !isFloat(east) {
		http.Error(w, "east missing", 400)
		return
	}

	rect := s2.RectFromLatLng(s2.LatLngFromDegrees(asFloat(south), asFloat(west)))
	rect = rect.AddPoint(s2.LatLngFromDegrees(asFloat(north), asFloat(east)))
	fmt.Println("rect: ", rect)

	region := s2.Region(rect)
	fmt.Println("region: ", region)

	rc := s2.RegionCoverer{MinLevel: asInt(minLvl), MaxLevel: asInt(maxLvl),
		LevelMod: 0, MaxCells: asInt(maxCells)}
	covering := rc.Covering(region)

	fmt.Println("covering: ", covering)

	response := ApproxResponse{CellIDs: nil}

	for _, c := range covering {
		response.CellIDs = append(response.CellIDs, c.String())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("server starting ...")
	http.HandleFunc("/approxRect", approxRect)
	http.ListenAndServe(":8000", nil)
}
