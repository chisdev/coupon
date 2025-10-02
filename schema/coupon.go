package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	coupon "github.com/chisdev/coupon/api"
)

type Coupon struct {
	ent.Schema
}

func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
		field.Float("value").Default(0.0),
		field.String("store_id").NotEmpty(),
		field.Time("expire_at").Optional().Nillable(),
		field.String("customer_id").Optional().Nillable(),
		field.Strings("service_ids").Default([]string{}),
		field.Int32("type").GoType(coupon.CouponType(1)),
		field.Uint64("currency_id").Optional(),
		field.Int32("usage_limit").Default(1),
		field.Int32("used_count").Default(0),
		field.Int32("status").GoType(coupon.CouponStatus(0)),
		field.Int32("reserved_count").Default(0),
	}
}

func (Coupon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("currency", Currency.Type).Ref("coupons").Unique().Field("currency_id"),
	}
}
