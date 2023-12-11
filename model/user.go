package model

// Product struct
type (
	User struct {
		ID       int    `json:"id"`
		FullName string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Users struct {
		Users []User `json:"users"`
	}
)
