package viewModels

// ApproxResponse is response to approxRect/approxCircle queries.
type ApproxResponse struct {
	CellIDs []string `json:"cellIDs"`
}
