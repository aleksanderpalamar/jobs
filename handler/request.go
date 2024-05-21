package handler

import "fmt"

func errParamIsRequest(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role        string `json:"role"`
	Level       string `json:"level"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
}

func (request *CreateOpeningRequest) Validate() error {
	if request.Role == "" && request.Level == "" && request.Company == "" && request.Location == "" && request.Remote == nil && request.Salary <= 0 && request.Title == "" && request.Description == "" && request.Keywords == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if request.Role == "" {
		return errParamIsRequest("role", "string")
	}
	if request.Level == "" {
		return errParamIsRequest("level", "string")
	}
	if request.Company == "" {
		return errParamIsRequest("company", "string")
	}
	if request.Location == "" {
		return errParamIsRequest("location", "string")
	}
	if request.Remote == nil {
		return errParamIsRequest("remote", "bool")
	}
	if request.Salary <= 0 {
		return errParamIsRequest("salary", "string")
	}
	if request.Title == "" {
		return errParamIsRequest("title", "string")
	}
	if request.Description == "" {
		return errParamIsRequest("description", "string")
	}
	if request.Keywords == "" {
		return errParamIsRequest("keywords", "string")
	}
	return nil
}
