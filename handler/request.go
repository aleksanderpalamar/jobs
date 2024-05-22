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
}

func (request *CreateOpeningRequest) Validate() error {
	if request.Role == "" && request.Level == "" && request.Company == "" && request.Location == "" && request.Remote == nil && request.Salary <= 0 && request.Title == "" && request.Description == "" {
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
	return nil
}

// UpdateOpening

type UpdateOpeningRequest struct {
	Role        string `json:"role"`
	Level       string `json:"level"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (request *UpdateOpeningRequest) Validate() error {
	// If any field is promided, validation is truthy
	if request.Role != "" || request.Level != "" || request.Company != "" || request.Location != "" || request.Remote != nil || request.Salary > 0 || request.Title != "" || request.Description != "" {
		return nil
	}
	// if none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")

}
