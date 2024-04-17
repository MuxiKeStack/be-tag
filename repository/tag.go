package repository

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/domain"
	"github.com/MuxiKeStack/be-tag/repository/dao"
	"github.com/ecodeclub/ekit/slice"
)

type TagRepository interface {
	BindTagsToBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []int32, tagType domain.TagType) error
	GetTagsByTaggerBizType(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tagType domain.TagType) ([]int32, error)
	GetCountAssessmentTagsByCourseTagger(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountAssessmentItem, error)
	GetCountFeatureTagsByCourseTaggerRequest(ctx context.Context, courseId int64, taggerIds []int64) ([]*tagv1.CountFeatureItem, error)
	GetCountTagsByBizTagger(ctx context.Context, biz tagv1.Biz, bizId int64, taggerIds []int64, tagType domain.TagType) ([]domain.CountTagItem, error)
}

type tagRepository struct {
	dao dao.TagDAO
}

func (repo *tagRepository) GetCountTagsByBizTagger(ctx context.Context, biz tagv1.Biz, bizId int64,
	taggerIds []int64, tagType domain.TagType) ([]domain.CountTagItem, error) {
	items, err := repo.dao.CountTagsByBizTagger(ctx, int32(biz), bizId, taggerIds, tagType.Uint8())
	return slice.Map(items, func(idx int, src dao.CountTagItem) domain.CountTagItem {
		return repo.toDomainCountTagItem(src)
	}), err
}

func (repo *tagRepository) BindTagsToBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64,
	tags []int32, tagType domain.TagType) error {
	return repo.dao.BatchCreate(ctx, slice.Map(tags, func(idx int, src int32) dao.Tag {
		return dao.Tag{
			TaggerId: taggerId,
			Biz:      int32(biz),
			BizId:    bizId,
			TagType:  tagType.Uint8(),
			Tag:      src,
		}
	}))
}

func (repo *tagRepository) GetTagsByTaggerBizType(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tagType domain.TagType) ([]int32, error) {
	return repo.dao.GetTagsByTaggerBizType(ctx, taggerId, int32(biz), bizId, tagType.Uint8())
}

func (repo *tagRepository) GetCountAssessmentTagsByCourseTagger(ctx context.Context, courseId int64,
	taggerIds []int64) ([]*tagv1.CountAssessmentItem, error) {
	items, err := repo.dao.GetCountTagsByCourseTaggerType(ctx, courseId, taggerIds, domain.TagTypeAssessment.Uint8())
	return slice.Map(items, func(idx int, src dao.CountTagItem) *tagv1.CountAssessmentItem {
		return &tagv1.CountAssessmentItem{
			Tag:   tagv1.AssessmentTag(src.Tag),
			Count: src.Count,
		}
	}), err
}

func (repo *tagRepository) GetCountFeatureTagsByCourseTaggerRequest(ctx context.Context, courseId int64,
	taggerIds []int64) ([]*tagv1.CountFeatureItem, error) {
	items, err := repo.dao.GetCountTagsByCourseTaggerType(ctx, courseId, taggerIds, domain.TagTypeFeature.Uint8())
	return slice.Map(items, func(idx int, src dao.CountTagItem) *tagv1.CountFeatureItem {
		return &tagv1.CountFeatureItem{
			Tag:   tagv1.FeatureTag(src.Tag),
			Count: src.Count,
		}
	}), err
}

func NewTagRepository(dao dao.TagDAO) TagRepository {
	return &tagRepository{dao: dao}
}

func (repo *tagRepository) toDomainCountTagItem(item dao.CountTagItem) domain.CountTagItem {
	return domain.CountTagItem{
		Tag:   item.Tag,
		Count: item.Count,
	}
}
