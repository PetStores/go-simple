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

// BeforeInsert set CreatedAt and UpdatedAt.
func (p *pet) BeforeInsert() error {
	p.CreatedAt = time.Now().UTC().Truncate(time.Second)
	p.UpdatedAt = p.CreatedAt
	return nil
}

// BeforeUpdate set UpdatedAt.
func (p *pet) BeforeUpdate() error {
	p.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return nil
}
