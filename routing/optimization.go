package routing

type OptimizationParameters struct {
	OptimizationProblemID string          `http:"optimization_problem_id,omitempty" json:"-"`
	Reoptimize            bool            `http:"reoptimize,omitempty" json:"-"`
	ShowDirections        bool            `http:"show_directions,omitempty" json:"-"`
	Parameters            RouteParameters `json:"parameters,omitempty"`
	Addresses             []Address       `json:"addresses,omtempty"`
}
