package presenter

type UserResponse struct {
	ID    string `json:"user_id"`
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
}
