// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/mohuishou/new-project/config/initializer"
	"github.com/mohuishou/new-project/internal/server/repo"
	"github.com/mohuishou/new-project/internal/server/service"
	"github.com/mohuishou/new-project/internal/server/usecase"
)

// Injectors from wire.go:

// NewServices NewServices
func NewServices() (*services, error) {
	db := initializer.NewGorm()
	iArticleRepo := repo.NewArticleRepo(db)
	iArticleUsecase := usecase.NewArticleUsecase(iArticleRepo)
	artcile := service.NewArticleService(iArticleUsecase)
	mainServices := &services{
		article: artcile,
	}
	return mainServices, nil
}