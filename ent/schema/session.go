package schema

import "github.com/facebookincubator/ent"

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
		Unique(),
		field.[]byte("data"),
		field.Time("expiry"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
