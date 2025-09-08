package dto

type CreateUserResponse struct {
	User                 *UserResponse `json:"user"`
	VerificationToken    string        `json:"verification_token,omitempty"`
	VerificationRequired bool          `json:"verification_required"`
}
