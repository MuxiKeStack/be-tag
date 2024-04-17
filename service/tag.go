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
	GetAssessmentTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64) ([]tagv1.AssessmentTag, error)
	GetFeatureTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64) ([]tagv1.FeatureTag, error)
	CountAssessmentTagsByCourseTagger(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountAssessmentItem, error)
	CountFeatureTagsByCourseTaggerRequest(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountFeatureItem, error)
}

type tagService struct {
	repo repository.TagRepository
}

func (s *tagService) CountAssessmentTagsByCourseTagger(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountAssessmentItem, error) {
	return s.repo.GetCountAssessmentTagsByCourseTagger(ctx, courseId, taggerIds)
}

func (s *tagService) CountFeatureTagsByCourseTaggerRequest(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountFeatureItem, error) {
	return s.repo.GetCountFeatureTagsByCourseTaggerRequest(ctx, courseId, taggerIds)
}

func (s *tagService) GetAssessmentTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64) ([]tagv1.AssessmentTag, error) {
	// 这个地方可以将Assessment切换为参数domain.TagTypeAssessment，向下传，
	// 但是上面的Count无法这样做因为，他的返回参数比较特别，我觉得这里应该对不同类型的tag进行实现不同的handler处理
	tags, err := s.repo.GetTagsByTaggerBizType(ctx, taggerId, biz, bizId, domain.TagTypeAssessment)
	return slice.Map(tags, func(idx int, src int32) tagv1.AssessmentTag {
		return tagv1.AssessmentTag(src)
	}), err
}

func (s *tagService) GetFeatureTagsByTaggerBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64) ([]tagv1.FeatureTag, error) {
	tags, err := s.repo.GetTagsByTaggerBizType(ctx, taggerId, biz, bizId, domain.TagTypeFeature)
	return slice.Map(tags, func(idx int, src int32) tagv1.FeatureTag {
		return tagv1.FeatureTag(src)
	}), err
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
