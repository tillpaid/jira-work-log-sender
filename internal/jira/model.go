package jira

type tempoCreateWorklogEngineeringActivitiesAttribute struct {
	Name            string `json:"name"`
	WorkAttributeId int    `json:"workAttributeId"`
	Value           string `json:"value"`
}

type tempoCreateWorklogAttributes struct {
	EngineeringActivities tempoCreateWorklogEngineeringActivitiesAttribute `json:"_EngineeringActivities_"`
}

type TempoCreateWorklogRequest struct {
	Attributes            tempoCreateWorklogAttributes `json:"attributes"`
	BillableSeconds       interface{}                  `json:"billableSeconds"`
	OriginId              int                          `json:"originId"`
	Worker                string                       `json:"worker"`
	Comment               string                       `json:"comment"`
	Started               string                       `json:"started"`
	TimeSpentSeconds      int                          `json:"timeSpentSeconds"`
	OriginTaskId          string                       `json:"originTaskId"`
	RemainingEstimate     string                       `json:"remainingEstimate"`
	EndDate               interface{}                  `json:"endDate"`
	IncludeNonWorkingDays bool                         `json:"includeNonWorkingDays"`
}
