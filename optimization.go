package route4me

import (
        "net/url"
        "encoding/json"
)

type OptimizationStateEnum uint

const (
        _ = iota
        OPTIMIZATIONSTATE_INITIAL OptimizationStateEnum = iota
        OPTIMIZATIONSTATE_MATRIXPROCESSING
        OPTIMIZATIONSTATE_OPTIMIZING
        OPTIMIZATIONSTATE_OPTIMIZED
        OPTIMIZATIONSTATE_ERROR
        OPTIMIZATIONSTATE_COMPUTINGDIRECTIONS
)

type Optimization struct {
        Optimization_problem_id string `json:"optimization_problem_id,omitempty"`
        State OptimizationStateEnum `json:"state,omitempty"`
        Parameters RouteParameters `json:"parameters,omitempty"`
        Send_to_background bool `json:"send_to_background,omitempty"`
        Addresses []Address `json:"addresses"`
        Routes []Route `json:"routes,omitempty"`
        Links struct {
                View string `json:"view,omitempty"`
        } `json:"links"`
}


func (r4m *Route4Me) GetOptimization(optimization_problem_id string) (Optimization, *Exception, error) {

        var container Optimization

        requestParams := url.Values{}
        requestParams.Set("optimization_problem_id", optimization_problem_id)

        res, err := r4m.Get("/api.v4/optimization_problem.php", requestParams)

        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*Optimization); ok {
                return *response, exception, err
        } else {
                return container, exception, err
        }
}

func (r4m *Route4Me) Reoptimize(optimization_problem_id string) (Optimization, *Exception, error) {
        request := NewOptimizationUpdateParams()
        request.Reoptimize = true

        return r4m.UpdateOptimization(optimization_problem_id, request)
}

type OptimizationNewParams struct {
        Addresses []Address
        Parameters *RouteParameters
        Directions bool
        Route_path_output string
        Optimized_callback_url string
        Reoptimize bool
}


func NewOptimizationNewParams() (OptimizationNewParams) {
        return OptimizationNewParams{
                Directions: true,
                Route_path_output: "None",
        }
}

func (r4m *Route4Me) NewOptimization(request OptimizationNewParams) (Optimization, *Exception, error) {
        var container Optimization

        requestParams := url.Values{}

        if request.Directions != true {
                requestParams.Set("directions", "0")
        }

        if request.Route_path_output != "None" {
                requestParams.Set("route_path_output", request.Route_path_output)
        }


        if request.Optimized_callback_url != "" {
                requestParams.Set("optimized_callback_url", request.Optimized_callback_url)
        }


        requestBody := map[string]interface{}{
                "addresses": request.Addresses,
                "parameters": request.Parameters,
        }

        requestBodyBytes, err := json.Marshal(requestBody)
        if err != nil {
                return container, nil, err
        }

        res, err := r4m.Post("/api.v4/optimization_problem.php", requestParams, requestBodyBytes)

        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*Optimization); ok {
                return *response, exception, err
        } else {
                return container, exception, err
        }
}

func NewOptimizationUpdateParams() (OptimizationNewParams) {
        return OptimizationNewParams{
                Directions: true,
                Route_path_output: "None",
                Reoptimize: false,
        }
}


func (r4m *Route4Me) UpdateOptimization(optimization_problem_id string, request OptimizationNewParams) (Optimization, *Exception, error) {
        var container Optimization

        requestParams := url.Values{}

        requestParams.Set("optimization_problem_id", optimization_problem_id)

        if request.Reoptimize == true {
                requestParams.Set("reoptimize", "1")
        }

        if request.Directions != true {
                requestParams.Set("directions", "0")
        }

        if request.Route_path_output != "None" {
                requestParams.Set("route_path_output", request.Route_path_output)
        }


        if request.Optimized_callback_url != "" {
                requestParams.Set("optimized_callback_url", request.Optimized_callback_url)
        }


        requestBody := map[string]interface{}{}

        if request.Addresses != nil {
                requestBody["addresses"] = request.Addresses
        }

        if request.Parameters != nil {
                requestBody["parameters"] = request.Parameters
        }

        requestBodyBytes, err := json.Marshal(requestBody)
        if err != nil {
                return container, nil, err
        }

        res, err := r4m.Put("/api.v4/optimization_problem.php", requestParams, requestBodyBytes)
        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*Optimization); ok {
                return *response, exception, err
        } else {
                return container, exception, err
        }
}
