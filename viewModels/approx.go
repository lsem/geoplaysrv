package viewModels

// LatLng represent project-local equialent of Latitude/Longitude pair
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Cell represent one S2 cell
type Cell struct {
	Level    int      `json:"level"`
	CellID   string   `json:"cellID"`
	Vertices []LatLng `json:"vertices"`
}

// ApproxResponse is response to approxRect/approxCircle queries.
type ApproxResponse struct {
	Cells []Cell `json:"cells"`
}
