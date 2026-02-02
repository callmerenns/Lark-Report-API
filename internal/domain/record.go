package domain

type Record struct {
	IncidentDescription string `json:"incident_description"`
	MachineModel        string `json:"machine_model"`
	LightboxType        string `json:"lightbox_type"`
	SetType             string `json:"set_type"`
	TypesOfSpareparts   string `json:"types_of_spareparts"`
	HowToHandle         string `json:"how_to_handle"`
	CollectDataset      string `json:"collect_dataset"`
	ProcessingModel     string `json:"processing_model"`
	DateOfIncident      string `json:"date_of_incident"`
	TimeOfIncident      string `json:"time_of_incident"`
	DayOfIncident       string `json:"day_of_incident"`
	ProcessReason       string `json:"process_reason"`
	CreatedBy           string `json:"created_by"`
	ModifiedBy          string `json:"modified_by"`
	Status              string `json:"status"`
}
