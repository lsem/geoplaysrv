package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	hlp "github.com/lsem/geoplaysrv/helpers"
	"github.com/lsem/geoplaysrv/viewModels"

	"github.com/golang/geo/s2"
)

// ApproxRect is a handler for /approxRect HTTP requests.
// Note, this handler assumes it is already decorated to set proper
// HTTP response headers.
func ApproxRect(w http.ResponseWriter, r *http.Request) {
	south := r.URL.Query().Get("south")
	west := r.URL.Query().Get("west")
	north := r.URL.Query().Get("north")
	east := r.URL.Query().Get("east")
	minLvl := r.URL.Query().Get("minLvl")
	maxLvl := r.URL.Query().Get("maxLvl")
	maxCells := r.URL.Query().Get("maxCells")

	fmt.Println("south: ", south)
	fmt.Println("west: ", west)
	fmt.Println("north: ", north)
	fmt.Println("east: ", east)
	fmt.Println("minLvl: ", minLvl)
	fmt.Println("maxLvl: ", maxLvl)
	fmt.Println("maxCells: ", maxCells)

	if len(minLvl) == 0 || !hlp.IsInt(minLvl) {
		http.Error(w, "minLvl missing or bad", 400)
		return
	}
	if len(maxLvl) == 0 || !hlp.IsInt(maxLvl) {
		http.Error(w, "maxLvl missing or bad", 400)
		return
	}
	if len(maxCells) == 0 || !hlp.IsInt(maxCells) {
		http.Error(w, "maxCells missing or bad", 400)
		return
	}
	if len(south) == 0 || !hlp.IsFloat(south) {
		fmt.Println("bad request: south: ", south)
		http.Error(w, "south missing", 400)
		return
	}
	if len(west) == 0 || !hlp.IsFloat(west) {
		http.Error(w, "west missing", 400)
		return
	}
	if len(north) == 0 || !hlp.IsFloat(north) {
		http.Error(w, "north missing", 400)
		return
	}
	if len(east) == 0 || !hlp.IsFloat(east) {
		http.Error(w, "east missing", 400)
		return
	}

	rect := s2.RectFromLatLng(s2.LatLngFromDegrees(hlp.AsFloat(south), hlp.AsFloat(west)))
	rect = rect.AddPoint(s2.LatLngFromDegrees(hlp.AsFloat(north), hlp.AsFloat(east)))

	region := s2.Region(rect)

	rc := s2.RegionCoverer{MinLevel: hlp.AsInt(minLvl), MaxLevel: hlp.AsInt(maxLvl),
		LevelMod: 0, MaxCells: hlp.AsInt(maxCells)}
	fmt.Println("Calculating covering ")
	covering := rc.Covering(region)
	fmt.Println("Calculating covering DONE")
	fmt.Println()

	response := viewModels.ApproxResponse{}

	for _, cid := range covering {
		cell := viewModels.Cell{Level: 10, CellID: cid.String(), Vertices: nil}
		s2Cell := s2.CellFromCellID(cid)
		for vidx := 0; vidx < 4; vidx++ {
			latLng := s2.LatLngFromPoint(s2Cell.Vertex(vidx))
			cell.Vertices = append(cell.Vertices,
				viewModels.LatLng{Lat: hlp.RadToDegrees(float64(latLng.Lat)),
					Lng: hlp.RadToDegrees(float64(latLng.Lng))})
		}
		response.Cells = append(response.Cells, cell)
	}
	json.NewEncoder(w).Encode(response)
}
