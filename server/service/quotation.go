package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/luiscovelo/goexpert-client-server-api/server/dto"
	"github.com/luiscovelo/goexpert-client-server-api/server/model"
)

type QuotationService struct {
	repo model.QuotationRepository
}

func NewQuotationService(repo model.QuotationRepository) *QuotationService {
	return &QuotationService{
		repo: repo,
	}
}

func (svc *QuotationService) GetQuotationAndSave() (*dto.QuotationResponseDTO, error) {
	quotationDTO, err := svc.getQuotationFromAPI()
	if err != nil {
		return nil, err
	}

	quotation := &model.Quotation{
		Code:       quotationDTO.USDBRL.Code,
		Codein:     quotationDTO.USDBRL.Codein,
		Name:       quotationDTO.USDBRL.Name,
		High:       quotationDTO.USDBRL.High,
		Low:        quotationDTO.USDBRL.Low,
		VarBid:     quotationDTO.USDBRL.VarBid,
		PctChange:  quotationDTO.USDBRL.PctChange,
		Bid:        quotationDTO.USDBRL.Bid,
		Ask:        quotationDTO.USDBRL.Ask,
		Timestamp:  quotationDTO.USDBRL.Timestamp,
		CreateDate: quotationDTO.USDBRL.CreateDate,
	}

	err = svc.save(quotation)
	if err != nil {
		return nil, err
	}

	resp := &dto.QuotationResponseDTO{
		Bid: quotation.Bid,
	}

	return resp, nil
}

func (svc *QuotationService) getQuotationFromAPI() (*dto.QuotationDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	endpoint := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quotation dto.QuotationDTO
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return nil, err
	}

	return &quotation, nil
}

func (svc *QuotationService) save(quotation *model.Quotation) error {
	return svc.repo.Save(quotation)
}
