package service

import (
	"context"

	"tax_service/internal/repository"

	"github.com/Kitten-King/tax_sdk/pb"
)

type TaxServer struct {
	pb.UnimplementedTaxServiceServer
	repo *repository.TaxRepository
}

func NewTaxServer(repo *repository.TaxRepository) *TaxServer {
	return &TaxServer{repo: repo}
}

func (s *TaxServer) CalculateTripTax(ctx context.Context, req *pb.TaxRequest) (*pb.TaxResponse, error) {
	rate, err := s.repo.GetRateByCityID(ctx, int(req.CityId))
	if err != nil {
		return nil, err
	}

	basePrice := float64(req.TripDurationSeconds) * 5.0
	taxAmount := basePrice * rate

	return &pb.TaxResponse{
		TaxAmount: taxAmount,
		TaxRate:   rate,
	}, nil
}
