package schema

import "github.com/facebookincubator/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").
		Unique(),
		field.String("username").
		Unique(),
		field.String("email").
		Unique(),
		field.String("password").
		Sensitive(),
		field.Time("created_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("planets", Planet.Type)
	}
}
