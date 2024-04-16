package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type TagDAO interface {
	BatchCreate(ctx context.Context, tags []Tag) error
}

type GORMTagDAO struct {
	db *gorm.DB
}

func NewGORMTagDAO(db *gorm.DB) TagDAO {
	return &GORMTagDAO{db: db}
}

func (dao *GORMTagDAO) BatchCreate(ctx context.Context, tags []Tag) error {
	if len(tags) == 0 {
		return nil
	}
	now := time.Now().UnixMilli()
	for _, t := range tags {
		t.Ctime = now
		t.Utime = now
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
	Utime    int64
	Ctime    int64
}
