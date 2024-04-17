package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type TagDAO interface {
	BatchCreate(ctx context.Context, tags []Tag) error
	GetTagsByTaggerBizType(ctx context.Context, taggerId int64, biz int32, bizId int64, tagType uint8) ([]int32, error)
	CountTagsByBizTagger(ctx context.Context, biz int32, bizId int64, taggerIds []int64, tagType uint8) ([]CountTagItem, error)
	GetCountTagsByCourseTaggerType(ctx context.Context, courseId int64, taggerIds []int64, tagType uint8) ([]CountTagItem, error)
}

type GORMTagDAO struct {
	db *gorm.DB
}

func (dao *GORMTagDAO) CountTagsByBizTagger(ctx context.Context, biz int32, bizId int64, taggerIds []int64,
	tagType uint8) ([]CountTagItem, error) {
	var countItems []CountTagItem
	err := dao.db.WithContext(ctx).
		Select("tag, count(*) as count").
		Model(&Tag{}).
		Where("biz = ? and biz_id = ? and tag_type = ? and tagger_id in ?", biz, bizId, tagType, taggerIds).
		Group("tag").
		Find(&countItems).Error
	return countItems, err
}

func (dao *GORMTagDAO) GetCountTagsByCourseTaggerType(ctx context.Context, courseId int64, taggerIds []int64,
	tagType uint8) ([]CountTagItem, error) {
	var countItems []CountTagItem
	err := dao.db.WithContext(ctx).
		Select("tag, count(*) as count").
		Model(&Tag{}).
		Where("course_id = ? and tag_type = ? and tagger_id in ?", courseId, tagType, taggerIds).
		Group("tag").
		Find(&countItems).Error
	return countItems, err
}

func (dao *GORMTagDAO) GetTagsByTaggerBizType(ctx context.Context, taggerId int64, biz int32, bizId int64, tagType uint8) ([]int32, error) {
	var tags []int32
	err := dao.db.WithContext(ctx).
		Model(&Tag{}).
		Select("tag").
		Where("tagger_id = ? and biz = ? and biz_id = ? and tag_type = ?", taggerId, biz, bizId, tagType).
		Find(&tags).Error
	return tags, err
}

func NewGORMTagDAO(db *gorm.DB) TagDAO {
	return &GORMTagDAO{db: db}
}

func (dao *GORMTagDAO) BatchCreate(ctx context.Context, tags []Tag) error {
	if len(tags) == 0 {
		return nil
	}
	now := time.Now().UnixMilli()
	for i := range tags {
		tags[i].Ctime = now
		tags[i].Utime = now
	}
	firstTag := tags[0]
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("tagger_id = ? and biz = ? and biz_id = ? and tag_type = ?",
			firstTag.TaggerId, firstTag.Biz, firstTag.BizId, firstTag.TagType).
			Delete(&Tag{}).Error
		if err != nil {
			return err
		}
		return tx.Create(tags).Error
	})
}

type Tag struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	TaggerId int64 `gorm:"uniqueIndex:tagger_biz_bizId_tagType_tag"`
	Biz      int32 `gorm:"uniqueIndex:tagger_biz_bizId_tagType_tag"`
	BizId    int64 `gorm:"uniqueIndex:tagger_biz_bizId_tagType_tag"`
	TagType  uint8 `gorm:"uniqueIndex:tagger_biz_bizId_tagType_tag"`
	Tag      int32 `gorm:"uniqueIndex:tagger_biz_bizId_tagType_tag"`
	//EvaluationId int64 // 查询某个课评的标签，1. 冗余这个字段，但是会导致该服务的通用性降低
	// 2. 不冗余，每次先从evaluation查到tagger和biz,bizId
	// 我选择了 2
	Utime int64
	Ctime int64
}

type CountTagItem struct {
	Tag   int32
	Count int64
}
