package apierrors

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Meta   struct {
		Attribute string `json:"attribute"`
		Message   string `json:"message"`
	} `json:"meta"`
	Status string `json:"status"`
	Title  string `json:"title"`
}
