package repository

import "github.com/chisdev/coupon/pkg/ent"

type Repository struct {
	ent *ent.Client
}

func New(ent *ent.Client) *Repository {
	return &Repository{ent: ent}
}
