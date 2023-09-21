package service

import (
	"app/repository"
	"context"
)

type FeedServiceImpl struct {
	repo repository.Repo
}

func NewFeedService(
	repo repository.Repo) FeedService {
	return &FeedServiceImpl{repo}
}

func (r *FeedServiceImpl) GetNewsfeed(ctx context.Context, ids_arr []int32, limit, offset int32) ([]int32, error) {
	arr, err := r.repo.GetNewsfeed(ctx, repository.GetNewsfeedParams{
		Column1: ids_arr,
		Limit:   limit,
		Offset:  offset,
	})

	res := []int32{}
	for _, e := range arr.([]interface{}) {
		res = append(res, e.(int32))
	}

	return res, err
}
