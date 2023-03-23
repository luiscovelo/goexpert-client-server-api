package repository

import (
	"context"
	"time"

	"github.com/luiscovelo/goexpert-client-server-api/server/model"
	"gorm.io/gorm"
)

type quotationRepository struct {
	db *gorm.DB
}

func NewQuotationRepository(db *gorm.DB) model.QuotationRepository {
	return &quotationRepository{
		db: db,
	}
}

func (r *quotationRepository) Save(quotation *model.Quotation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	return r.db.WithContext(ctx).Create(quotation).Error
}
