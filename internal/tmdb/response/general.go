package response

type generalTMDBResponse struct {
	Success bool `json:"success"`
}

type paginationResponse struct {
	Page          int `json:"page"`
	Total_pages   int `json:"total_pages"`
	Total_results int `json:"total_results"`
}

type ErrorTMDBResponse struct {
	HTTP_Status    int    `json:"-"`
	Status_Code    int    `json:"status_code"`
	Status_Message string `json:"status_message"`
}

func (e *ErrorTMDBResponse) Error() string {
	return e.Status_Message
}
