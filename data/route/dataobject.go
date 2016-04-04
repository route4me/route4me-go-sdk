package route

//OptimizationState describes an optimization problem can be at one state at any given time
//every state change invokes a socket notification to the associated member id
//every state change invokes a callback webhook event invocation if it was provided during the initial optimization
type OptimizationState uint

const (
	Initial OptimizationState = iota + 1
	MatrixProcessing
	Optimizing
	Optimized
	Error
	ComputingDirections
)

type DataObject struct {
	OptimizationProblemID string            `json:"optimization_problem_id"`
	State                 OptimizationState `json:"state"`
}
