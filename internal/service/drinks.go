// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================
package service

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
)

type (
	IDRINKS interface {
		Add(ctx context.Context, in *entity.Drinks) (err error)
		Get(ctx context.Context, in *do.Drinks) (out *entity.Drinks, err error)
	}
)

var (
	localIDRINKS IDRINKS
)

func Idrinks() IDRINKS {
	if localIDRINKS == nil {
		panic("implement not found for interface IDRINKS, forgot register?")
	}
	return localIDRINKS
}

func RegisterIdrinks(i IDRINKS) {
	localIDRINKS = i
}
