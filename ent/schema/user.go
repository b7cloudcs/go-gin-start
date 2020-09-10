package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Entity
type User struct {
	ent.Schema
}

// Fields
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Optional(),
		field.String("password").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
