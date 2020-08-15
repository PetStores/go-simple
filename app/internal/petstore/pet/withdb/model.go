package withdb

import "time"

//go:generate reform

//reform:pet
type pet struct {
	ID         int64     `reform:"id,pk"`
	Name       string    `reform:"name"`
	CategoryID int64     `reform:"category_id"`
	CreatedAt  time.Time `reform:"created_at"`
	UpdatedAt  time.Time `reform:"updated_at"`
}
