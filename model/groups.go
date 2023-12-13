package model

// Group struct
type (
	Group struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Permissions string `json:"permissions"` // dunno, permissions ?
		UserId      int    `json:"user_id"`
	}

	Groups struct {
		Groups []Group `json:"groups"`
	}
)
