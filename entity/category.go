package entities

import "time"

type Category struct {
	ID            string
	CompanyId     string
	IntegrationId string
	ExternalId    string
	Key           string
	Name          string
	UpdateTime    time.Time
}
