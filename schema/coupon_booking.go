package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Uint64("booking_id"),
		field.Int32("status").GoType(coupon.CouponUsedStatus(1)),
	}
}

func (CouponBooking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coupon", Coupon.Type).Ref("coupon_bookings").Unique().Required().Field("coupon_id"),
	}
}
