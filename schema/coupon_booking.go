package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	coupon "github.com/chisdev/coupon/api"
)

type CouponBooking struct {
	ent.Schema
}

func (CouponBooking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (CouponBooking) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("coupon_id"),
		field.String("booking_id"),
		field.Int32("status").GoType(coupon.CouponUsedStatus(1)),
		field.Strings("service_ids").Optional(),
		field.String("customer_id").Optional().Nillable(),
	}
}

func (CouponBooking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coupon", Coupon.Type).Ref("coupon_bookings").Unique().Required().Field("coupon_id"),
	}
}

func (CouponBooking) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coupon_id", "booking_id").Unique(),
	}
}
