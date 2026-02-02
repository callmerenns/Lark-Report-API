package response

import "time"

type RecordCreatedData struct {
	ID                  string    `json:"id" example:"1"`
	IncidentTitle       string    `json:"incident_title" example:"Machine Error"`
	IncidentDescription string    `json:"incident_description" example:"Description of the machine error"`
	MachineModel        string    `json:"machine_model" example:"Model X"`
	LightboxType        string    `json:"lightbox_type" example:"Type A"`
	SetType             string    `json:"set_type" example:"Set 1"`
	TypesOfSpareparts   string    `json:"types_of_spareparts" example:"Sparepart A, Sparepart B"`
	HowToHandle         string    `json:"how_to_handle" example:"Handle with care"`
	CollectDataset      string    `json:"collect_dataset" example:"Yes"`
	ProcessingModel     string    `json:"processing_model" example:"Model Y"`
	DateOfIncident      string    `json:"date_of_incident" example:"2026-01-29"`
	TimeOfIncident      string    `json:"time_of_incident" example:"14:30:00"`
	DayOfIncident       string    `json:"day_of_incident" example:"Monday"`
	ProcessReason       string    `json:"process_reason" example:"Routine check"`
	CreatedBy           string    `json:"created_by" example:"admin"`
	ModifiedBy          string    `json:"modified_by" example:"admin"`
	Status              string    `json:"status" example:"OPEN"`
	CreatedAt           time.Time `json:"created_at" example:"2026-01-29T16:30:00Z"`
}
