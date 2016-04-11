package routing

type OptimizationParameters struct {
	ProblemID      string          `http:"optimization_problem_id" json:"-"`
	Reoptimize     bool            `http:"reoptimize" json:"-"`
	ShowDirections bool            `http:"show_directions" json:"-"`
	Parameters     RouteParameters `json:"parameters,omitempty"`
	Addresses      []Address       `json:"addresses,omtempty"`
}
