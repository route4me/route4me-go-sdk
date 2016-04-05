package query

import "github.com/route4me/route4me-go-sdk/data/route"

type Optimization struct {
	OptimizationProblemID string           `http:"optimization_problem_id,omitempty" json:"-"`
	Reoptimize            bool             `http:"reoptimize,omitempty" json:"-"`
	ShowDirections        bool             `http:"show_directions,omitempty" json:"-"`
	Parameters            route.Parameters `json:"parameters,omitempty"`
	Addresses             []Address        `json:"addresses,omtempty"`
}
