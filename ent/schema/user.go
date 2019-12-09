package schema

import "github.com/facebookincubator/ent"

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}

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
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("planets", Planet.Type)
	}
}
