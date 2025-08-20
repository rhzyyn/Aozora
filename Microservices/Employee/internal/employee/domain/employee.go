package domain

import "time"

type Employee struct {
	UUID         string    `json:"uuid"`
	Credentials  string    `json:"credentials,omitempty"`
	EmployeeID   int       `json:"employee_id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name,omitempty"`
	LastName     string    `json:"last_name"`
	Education    string    `json:"education,omitempty"`
	Email        string    `json:"email"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Ages         int       `json:"ages"`
	Gender       string    `json:"gender"`
	Religion     string    `json:"religion,omitempty"`
	Marital      string    `json:"marital,omitempty"`
	SocialNumber string    `json:"social_number,omitempty"`
	NPWP         string    `json:"npwp,omitempty"`
	FamilyNumber string    `json:"family_number,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by,omitempty"`
}
