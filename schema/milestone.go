package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	coupon "github.com/chisdev/coupon/api"
)

type Milestone struct {
	ent.Schema
}

func (Milestone) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Milestone) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Nillable(),
		field.String("store_id").NotEmpty(),
		field.Int32("milestone_type").GoType(coupon.MilestoneType(0)),
		field.Int32("threshold").Default(0),
		field.Int32("step").Default(0),
	}
}

func (Milestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reward", Reward.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
		edge.To("progress", Progress.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
	}
}
