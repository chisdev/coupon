package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Progress struct {
	ent.Schema
}

func (Progress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Progress) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_id").NotEmpty(),
		field.Uint64("milestone_id"),
		field.Int32("progress").Default(0),
		field.Int32("pass_count").Default(0),
	}
}

func (Progress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("milestone", Milestone.Type).Ref("progress").Required().Unique().Field("milestone_id")}

}
