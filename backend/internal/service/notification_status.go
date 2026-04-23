package service

import "context"

func (s *serviceImpl) UpdateStatus(ctx context.Context, id, status string, retry int) error {
	return s.database.UpdateStatus(ctx, id, status, retry)
}
