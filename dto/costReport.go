package dto

type CostResponse struct {
	Id           int     `json:"id"`
	Project      string  `json:"project"`
	ProjectName  string  `json:"project_env"`
	Cluster      string  `json:"cluster"`
	Installation int     `json:"installation"`
	Tenant       string  `json:"tenant"`
	Budget       float32 `json:"budget"`
	ReportedOn   string  `json:"reported_on"`
	CPU          int     `json:"cpu"`
	EBSGB        float32 `json:"ebs_gb"`
	EFSGB        float32 `json:"efs_gb"`
	ELBGB        float32 `json:"elb_gb"`
	MemoryGB     float32 `json:"mem_gb"`
	Pods         int     `json:"pods"`
	S3GB         float32 `json:"s3_gb"`
}

type CostResponseList struct {
	Items      []CostResponse `json:"items"`
	NextPageID int            `json:"next_page_id,omitempty" example:"10"`
	PrevPageID int            `json:"prev_page_id,omitempty" example:"10"`
}
