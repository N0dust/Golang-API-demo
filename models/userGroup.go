package models

type UserGroup struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`

	// CreateTime time.Time `json:"create_time"`
	// UpdateTime time.Time `json:"update_time"`
	// DeleteTime time.Time `json:"delete_time"`
}
