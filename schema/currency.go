package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Currency struct {
	ent.Schema
}

func (Currency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Currency) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

func (Currency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coupons", Coupon.Type),
		edge.To("reward", Reward.Type),
	}
}
