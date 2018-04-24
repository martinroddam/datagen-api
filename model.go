package main

// RecordSet is the high level object that contains the generated records and associated statistics
type RecordSet struct {
	Statistics Stats    `json:"statistics"`
	Records    []Record `json:"records"`
}

// Record contains the various fields that make up a personal data record
type Record struct {
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	Email       string `json:"email"`
	HomePhone   string `json:"homePhone"`
	MobilePhone string `json:"mobilePhone"`
}

// Name is made up of Prefix, First Name and Last Name
type Name struct {
	Prefix string
	First  string
	Last   string
}

// Stats contains aggregate statistics for the RecordSet
type Stats struct {
	CountryCounts map[string]int `json:"countryCounts"`
	RecordCount   int            `json:"recordCount"`
}

// EndpointList contains a list of endpoints available for this app
type EndpointList struct {
	Endpoints []string `json:"endpoints"`
}
