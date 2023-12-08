package models

type Practice struct {
	ObjectID    string `json:"objectID,omitempty"`
	Practice    string `json:"practice,omitempty"`
	Code        string `json:"code,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	Address3    string `json:"address3,omitempty"`
	Address4    string `json:"address4,omitempty"`
	Address5    string `json:"address5,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
	Tel         string `json:"tel,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Hcps        []Hcps `json:"hcps,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsGPSurgery string `json:"is_gp_surgery,omitempty"`
}

type Hcps struct {
	GMC       string `json:"gmc"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
