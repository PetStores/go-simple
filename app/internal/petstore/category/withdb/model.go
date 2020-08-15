package withdb

import "time"

//go:generate reform

//reform:category
type category struct {
	ID        int64     `reform:"id,pk"`
	Name      string    `reform:"name"`
	IsVisible bool      `reform:"is_visible"`
	CreatedAt time.Time `reform:"created_at"`
	UpdatedAt time.Time `reform:"updated_at"`
}

// BeforeInsert set CreatedAt and UpdatedAt.
func (c *category) BeforeInsert() error {
	c.CreatedAt = time.Now().UTC().Truncate(time.Second)
	c.UpdatedAt = c.CreatedAt
	c.IsVisible = true
	return nil
}

// BeforeUpdate set UpdatedAt.
func (c *category) BeforeUpdate() error {
	c.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return nil
}
