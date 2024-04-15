package repository

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/domain"
	"github.com/MuxiKeStack/be-tag/repository/dao"
	"github.com/ecodeclub/ekit/slice"
)

type TagRepository interface {
	BindTagsToBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64, tags []uint32, tagType domain.TagType) error
}

type tagRepository struct {
	dao dao.TagDAO
}

func NewTagRepository(dao dao.TagDAO) TagRepository {
	return &tagRepository{dao: dao}
}

func (repo *tagRepository) BindTagsToBiz(ctx context.Context, taggerId int64, biz tagv1.Biz, bizId int64,
	tags []uint32, tagType domain.TagType) error {
	return repo.dao.BatchCreate(ctx, slice.Map(tags, func(idx int, src uint32) dao.Tag {
		return dao.Tag{
			TaggerId: taggerId,
			Biz:      uint32(biz),
			BizId:    bizId,
			Tag:      src,
			TagType:  tagType.Uint8(),
		}
	}))
}
