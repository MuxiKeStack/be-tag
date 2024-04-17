package service

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/domain"
	"github.com/MuxiKeStack/be-tag/repository"
)

type GeneralTagService interface {
	AttachTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []int32, tagType domain.TagType) error
	GetTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tagType domain.TagType) ([]int32, error)
	CountTagsByBizTagger(ctx context.Context, biz tagv1.Biz, bizId int64, taggerIds []int64, tagType domain.TagType) ([]domain.CountTagItem, error)
}

// generalTagService 使用中的实现
type generalTagService struct {
	repo repository.TagRepository
}

func (s *generalTagService) CountTagsByBizTagger(ctx context.Context, biz tagv1.Biz, bizId int64, taggerIds []int64,
	tagType domain.TagType) ([]domain.CountTagItem, error) {
	return s.repo.GetCountTagsByBizTagger(ctx, biz, bizId, taggerIds, tagType)
}

func NewGeneralTagService(repo repository.TagRepository) GeneralTagService {
	return &generalTagService{repo: repo}
}

func (s *generalTagService) AttachTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []int32,
	tagType domain.TagType) error {
	return s.repo.BindTagsToBiz(ctx, taggerId, biz, bizId, tags, tagType)
}

func (s *generalTagService) GetTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64,
	tagType domain.TagType) ([]int32, error) {
	return s.repo.GetTagsByTaggerBizType(ctx, taggerId, biz, bizId, tagType)
}
