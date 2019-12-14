package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

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
		edge.To("planets", Planet.Type),
	}
}
