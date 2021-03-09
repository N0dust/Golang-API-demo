package models


type User struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	GroupID string `json:"group_id"`
	// CreateTime time.Time `json:"create_time"`
	// UpdateTime time.Time `json:"update_time"`
	// DeleteTime time.Time `json:"delete_time"`
}
