package response

import "time"

/*
|--------------------------------------------------------------------------
| SUCCESS EXAMPLES
|--------------------------------------------------------------------------
*/

// WelcomeSuccessExample swagger example
type WelcomeSuccessExample struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Welcome to API Lark"`
	Data    struct {
		Name    string `json:"name" example:"Lark Webhook API"`
		Version string `json:"version" example:"v1"`
		Status  string `json:"status" example:"running"`
	} `json:"data"`
}

// RecordCreatedSuccessExample digunakan untuk @Success 201
type RecordCreatedSuccessExample struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Record created successfully"`
	Data    struct {
		ID                  string    `json:"id" example:"1"`
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
	} `json:"data"`
}

// SuccessOKWithDataExample digunakan untuk @Success 200 (GET / detail)
type SuccessOKWithDataExample struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Request successful"`
	Data    struct {
		ID                  string `json:"id" example:"1"`
		IncidentDescription string `json:"incident_description" example:"Description of the machine error"`
		MachineModel        string `json:"machine_model" example:"Model X"`
		LightboxType        string `json:"lightbox_type" example:"Type A"`
		SetType             string `json:"set_type" example:"Set 1"`
		TypesOfSpareparts   string `json:"types_of_spareparts" example:"Sparepart A, Sparepart B"`
		HowToHandle         string `json:"how_to_handle" example:"Handle with care"`
		CollectDataset      string `json:"collect_dataset" example:"Yes"`
		ProcessingModel     string `json:"processing_model" example:"Model Y"`
		DateOfIncident      string `json:"date_of_incident" example:"2026-01-29"`
		TimeOfIncident      string `json:"time_of_incident" example:"14:30:00"`
		DayOfIncident       string `json:"day_of_incident" example:"Monday"`
		ProcessReason       string `json:"process_reason" example:"Routine check"`
		CreatedBy           string `json:"created_by" example:"admin"`
		ModifiedBy          string `json:"modified_by" example:"admin"`
		Status              string `json:"status" example:"OPEN"`
	} `json:"data"`
}

/*
|--------------------------------------------------------------------------
| AUTH / SECURITY ERRORS
|--------------------------------------------------------------------------
*/

// UnauthorizedErrorExample digunakan untuk 401
type UnauthorizedErrorExample struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Unauthorized"`
	Error   struct {
		Code    string `json:"code" example:"MISSING_AUTHORIZATION"`
		Details string `json:"details" example:"Missing Authorization header"`
	} `json:"error"`
}

// ForbiddenErrorExample digunakan untuk 403
type ForbiddenErrorExample struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Forbidden"`
	Error   struct {
		Code    string `json:"code" example:"FORBIDDEN"`
		Details string `json:"details" example:"Access is not allowed"`
	} `json:"error"`
}

/*
|--------------------------------------------------------------------------
| VALIDATION ERRORS
|--------------------------------------------------------------------------
*/

// BadRequestErrorExample digunakan untuk 400
type BadRequestErrorExample struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid payload"`
	Error   struct {
		Code    string `json:"code" example:"INVALID_PAYLOAD"`
		Details string `json:"details" example:"Payload validation failed"`
	} `json:"error"`
}

/*
|--------------------------------------------------------------------------
| RATE LIMIT ERRORS
|--------------------------------------------------------------------------
*/

// RateLimitErrorExample digunakan untuk 429
type RateLimitErrorExample struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Too many requests"`
	Error   struct {
		Code    string `json:"code" example:"RATE_LIMITED"`
		Details string `json:"details" example:"Too many requests"`
		Limit   int    `json:"limit" example:"5"`
		Window  string `json:"window" example:"1m0s"`
	} `json:"error"`
}

/*
|--------------------------------------------------------------------------
| SERVER ERRORS
|--------------------------------------------------------------------------
*/

// InternalServerErrorExample digunakan untuk 500
type InternalServerErrorExample struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Internal server error"`
	Error   struct {
		Code    string `json:"code" example:"INTERNAL_SERVER_ERROR"`
		Details string `json:"details" example:"Unexpected error occurred"`
	} `json:"error"`
}
