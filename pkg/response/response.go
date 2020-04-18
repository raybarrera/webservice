package response

type DataReader interface{}

// ApiError contains client-readable information on a given API error.
// ReferenceCode refers to the internal code the implementer can check against.
// Message is the human-readable description of what caused the error.
type ApiError struct {
	ReferenceCode string `json:"reference_code"`
	Message       string `json:"message"`
}

// accountCreateResponse is used to package the response back to the client in a uniform way.
type AccountCreateResponse struct {
	Status   string     `json:"status"`
	Code     int        `json:"code"`
	Messages []string   `json:"messages"`
	Data     DataReader `json:"data"`
	Errors   []apiError `json:"errors"`
}
