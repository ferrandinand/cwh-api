package domain

import (
	"time"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/dto"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
	"github.com/jmoiron/sqlx"
)

type CostList struct {
	Items      []Cost `json:"items"`
	NextPageID int    `json:"next_page_id,omitempty" example:"10"`
	PrevPageID int    `json:"prev_page_id,omitempty" example:"10"`
}

type Cost struct {
	Id           int     `db:"cost_id"`
	Project      string  `db:"project"`
	Cluster      string  `db:"cluster"`
	ProjectName  string  `db:"project_name"`
	Installation int     `db:"installation"`
	ReportedOn   string  `db:"reported_on"`
	Tenant       string  `db:"tenant"`
	Budget       float32 `db:"budget"`
	CPU          int     `db:"cpu"`
	EBSCost      float32 `db:"ebs_cost"`
	EBSGB        float32 `db:"ebs_gb"`
	EFSCost      float32 `db:"efs_cost"`
	EFSGB        float32 `db:"efs_gb"`
	ELBCost      float32 `db:"elb_cost"`
	ELBGB        float32 `db:"elb_gb"`
	MemoryGB     float32 `db:"mem_gb"`
	Pods         int     `db:"pods"`
	PodsCost     float32 `db:"pods_cost"`
	RDSCost      float32 `db:"rds_cost"`
	S3Cost       float32 `db:"s3_cost"`
	S3GB         float32 `db:"s3_gb"`
	SecretsCost  float32 `db:"secrets_cost"`
	TransferCost float32 `db:"transfer_cost"`
	OthersCost   float32 `db:"others_cost"`
	Total        float32 `db:"total"`
}

type CostRepository interface {
	FindAll(pageId int) (CostList, *errs.AppError)
	ById(projectId string) (CostList, *errs.AppError)
}

func (c Cost) ToDto() dto.CostResponse {
	return dto.CostResponse{
		Project:      c.Project,
		Cluster:      c.Cluster,
		ProjectName:  c.ProjectName,
		Installation: c.Installation,
		ReportedOn:   c.ReportedOn,
		Tenant:       c.Tenant,
		Budget:       c.Budget,
		CPU:          c.CPU,
		EBSGB:        c.EBSGB,
		EFSGB:        c.EFSGB,
		ELBGB:        c.ELBGB,
		MemoryGB:     c.MemoryGB,
		Pods:         c.Pods,
		S3GB:         c.S3GB,
	}
}

func NewCost(
	projectName string,
	installation int,
	tenant string,
	cluster string,
	project string,
	budget float32,
	cpu int,
	ebsCost float32,
	ebsGB float32,
	efsCost float32,
	efsGB float32,
	elbCost float32,
	elbGB float32,
	memoryGB float32,
	pods int,
	podsCost float32,
	rdsCost float32,
	s3Cost float32,
	s3GB float32,
	secretsCost float32,
	transferCost float32,
	othersCost float32,
	total float32) Cost {

	return Cost{
		ProjectName:  projectName,
		Installation: installation,
		ReportedOn:   time.Now().Format("2006-01-02 15:04:05"),
		Tenant:       tenant,
		Cluster:      cluster,
		Project:      project,
		Budget:       budget,
		CPU:          cpu,
		EBSCost:      ebsCost,
		EBSGB:        ebsGB,
		EFSCost:      efsCost,
		EFSGB:        efsGB,
		ELBCost:      elbCost,
		ELBGB:        elbGB,
		MemoryGB:     memoryGB,
		Pods:         pods,
		PodsCost:     podsCost,
		RDSCost:      rdsCost,
		S3Cost:       s3Cost,
		S3GB:         s3GB,
		SecretsCost:  secretsCost,
		TransferCost: transferCost,
		OthersCost:   othersCost,
		Total:        total,
	}
}

func NewCostRepositoryDb(dbClient *sqlx.DB) CostRepositoryDb {
	return CostRepositoryDb{dbClient}
}
