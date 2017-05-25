package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/lsem/geosrv/helpers"
	"github.com/lsem/geosrv/viewModels"

	"github.com/golang/geo/s2"
)

// ApproxRect is a handler for /approxRect HTTP requests.
func ApproxRect(w http.ResponseWriter, r *http.Request) {
	south := r.URL.Query().Get("south")
	west := r.URL.Query().Get("west")
	north := r.URL.Query().Get("north")
	east := r.URL.Query().Get("east")
	minLvl := r.URL.Query().Get("minLvl")
	maxLvl := r.URL.Query().Get("maxLvl")
	maxCells := r.URL.Query().Get("maxCells")

	if len(minLvl) == 0 || !IsInt(minLvl) {
		http.Error(w, "minLvl missing or bad", 400)
		return
	}
	if len(maxLvl) == 0 || !IsInt(maxLvl) {
		http.Error(w, "maxLvl missing or bad", 400)
		return
	}
	if len(maxCells) == 0 || !IsInt(maxCells) {
		http.Error(w, "maxCells missing or bad", 400)
		return
	}
	if len(south) == 0 || !IsFloat(south) {
		http.Error(w, "south missing", 400)
		return
	}
	if len(west) == 0 || !IsFloat(west) {
		http.Error(w, "west missing", 400)
		return
	}
	if len(north) == 0 || !IsFloat(north) {
		http.Error(w, "north missing", 400)
		return
	}
	if len(east) == 0 || !IsFloat(east) {
		http.Error(w, "east missing", 400)
		return
	}

	rect := s2.RectFromLatLng(s2.LatLngFromDegrees(AsFloat(south), AsFloat(west)))
	rect = rect.AddPoint(s2.LatLngFromDegrees(AsFloat(north), AsFloat(east)))
	fmt.Println("rect: ", rect)

	region := s2.Region(rect)
	fmt.Println("region: ", region)

	rc := s2.RegionCoverer{MinLevel: AsInt(minLvl), MaxLevel: AsInt(maxLvl),
		LevelMod: 0, MaxCells: AsInt(maxCells)}
	covering := rc.Covering(region)

	fmt.Println("covering: ", covering)

	response := viewModels.ApproxResponse{CellIDs: nil}

	for _, c := range covering {
		response.CellIDs = append(response.CellIDs, c.String())
	}
	json.NewEncoder(w).Encode(response)
}
