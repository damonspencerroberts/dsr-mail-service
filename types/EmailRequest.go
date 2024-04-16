package types

type EmailRequest struct {
	ToEmail   string `json:"to_email" validate:"required,email"`
	Subject   string `json:"subject" validate:"required"`
	Message   string `json:"message" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}