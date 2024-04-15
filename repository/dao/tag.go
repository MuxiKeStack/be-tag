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
	TaggerId int64
	Biz      uint32
	BizId    int64
	Tag      uint32
	TagType  uint8
	Utime    int64
	Ctime    int64
}
