package service

import (
	"context"
	"wharehouse-control/internal/dto"
)

func (s *Service) UpdateItem(ctx context.Context, updateItem *dto.UpdateItem) error {
	return s.storage.UpdateItem(ctx, updateItem)
}
