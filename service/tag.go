package service

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/domain"
	"github.com/MuxiKeStack/be-tag/repository"
	"github.com/ecodeclub/ekit/slice"
)

type TagService interface {
	AttachAssessmentTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []tagv1.AssessmentTag) error
	AttachFeatureTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []tagv1.FeatureTag) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) AttachAssessmentTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []tagv1.AssessmentTag) error {
	return s.repo.BindTagsToBiz(ctx, taggerId, biz, bizId, slice.Map(tags, func(idx int, src tagv1.AssessmentTag) int32 {
		return int32(src)
	}), domain.TagTypeAssessment)
}

func (s *tagService) AttachFeatureTags(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []tagv1.FeatureTag) error {
	return s.repo.BindTagsToBiz(ctx, taggerId, biz, bizId, slice.Map(tags, func(idx int, src tagv1.FeatureTag) int32 {
		return int32(src)
	}), domain.TagTypeFeature)
}
