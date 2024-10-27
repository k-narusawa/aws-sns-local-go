package aws

type CreateTopicInput struct {
	Name                 string `json:"name"`
	Attributes           string `json:"attributes"`
	Tags                 string `json:"tags"`
	DataProtectionPolicy string `json:"data_protection_policy"`
}
