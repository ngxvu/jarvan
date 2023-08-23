package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	"gitlab.com/merakilab9/j4/pkg/model"

	"gorm.io/gorm"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

func (h *MigrationHandler) Migrate(ctx *gin.Context) {
	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230721154905",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Exec(`
					CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
				`).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&model.CateCrawl{}, &model.CateUrl{}, &model.ShopIdUrl{}, &model.ShopDetail{}, &model.OfficialShop{}); err != nil {
					return err
				}
				return nil
			},
		},
	})
	err := migrate.Migrate()
	if err != nil {
		panic(err)
	}
}
