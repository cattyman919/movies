package response

// NewGuessSessionResponse creates a new guest session id
type NewGuestSessionResponse struct {
	generalTMDBResponse
	Guest_Session_Id string `json:"guest_session_id"`
	Expires_At       string `json:"expires_at"`
}
