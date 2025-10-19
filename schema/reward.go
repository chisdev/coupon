package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	coupon "github.com/chisdev/coupon/api"
)

type Reward struct {
	ent.Schema
}

func (Reward) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Reward) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("milestone_id"),
		field.Float("expired_duration").Optional().Nillable(),
		field.JSON("service_ids", []uint64{}).Default([]uint64{}),
		field.Int32("coupon_type").GoType(coupon.CouponType(0)),
		field.Uint64("currency_id").Optional().Nillable(),
		field.Int32("usage_limit").Default(1),
		field.Float("coupon_value").Default(0.0),
	}
}

func (Reward) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("currency", Currency.Type).Ref("reward").Unique().Field("currency_id"),
		edge.From("milestone", Milestone.Type).Ref("reward").Required().Unique().Field("milestone_id")}

}
