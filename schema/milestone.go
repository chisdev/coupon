package schema

import (
	"entgo.io/ent"
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
		field.Time("expire_at").Optional().Nillable(),
		field.Strings("service_ids").Default([]string{}),
		field.Int32("coupon_type").GoType(coupon.CouponType(0)),
		field.Int32("milestone_type").GoType(coupon.MilestoneType(0)),
		field.Uint64("currency_id").Optional(),
		field.Int32("usage_limit").Default(1),
		field.Int32("threshold").Default(0),
		field.Int32("step").Default(0),
		field.Float("coupon_value").Default(0.0),
	}
}

func (Milestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("currency", Currency.Type).Ref("milestones").Unique().Field("currency_id")}
}
