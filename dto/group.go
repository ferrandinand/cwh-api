package dto

type GroupRequest struct {
	Name      string
	CreatedOn string
}

type GroupResponse struct {
	Id         int                    `json:"id"`
	Name       string                 `json:"name"`
	CreatedOn  string                 `json:"created_on"`
	Attributes map[string]interface{} `json:"attributes"`
	Status     string                 `json:"status"`
}

type GroupResponseList struct {
	Items      []GroupResponse `json:"items"`
	NextPageID int             `json:"next_page_id,omitempty" example:"10"`
}
