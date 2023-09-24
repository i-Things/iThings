package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/stores"
	"gorm.io/gorm"
)

type GroupInfoRepo struct {
	db *gorm.DB
}

type GroupInfoFilter struct {
	GroupID     int64
	GroupNames  []string
	ParentID    int64
	GroupName   string
	Tags        map[string]string
	WithProduct bool
}

func NewGroupInfoRepo(in any) *GroupInfoRepo {
	return &GroupInfoRepo{db: stores.GetCommonConn(in)}
}

func (p GroupInfoRepo) fmtFilter(ctx context.Context, f GroupInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if f.WithProduct {
		db = db.Preload("ProductInfo")
	}
	if f.GroupID != 0 {
		db = db.Where("group_id=?", f.GroupID)
	}
	if len(f.GroupNames) != 0 {
		db = db.Where("group_name in ?", f.GroupNames)
	}
	if f.ParentID != 0 {
		db = db.Where("parent_id=?", f.ParentID)
	}
	if f.GroupName != "" {
		db = db.Where("group_name like ?", "%"+f.GroupName+"%")
	}
	if f.Tags != nil {
		for k, v := range f.Tags {
			db = db.Where("JSON_CONTAINS(`tags`, JSON_OBJECT(?,?))",
				k, v)
		}
	}
	return db
}

func (g GroupInfoRepo) Insert(ctx context.Context, data *DmGroupInfo) error {
	result := g.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (g GroupInfoRepo) FindOneByFilter(ctx context.Context, f GroupInfoFilter) (*DmGroupInfo, error) {
	var result DmGroupInfo
	db := g.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}
func (p GroupInfoRepo) FindByFilter(ctx context.Context, f GroupInfoFilter, page *def.PageInfo) ([]*DmGroupInfo, error) {
	var results []*DmGroupInfo
	db := p.fmtFilter(ctx, f).Model(&DmGroupInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p GroupInfoRepo) CountByFilter(ctx context.Context, f GroupInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&DmGroupInfo{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (g GroupInfoRepo) Update(ctx context.Context, data *DmGroupInfo) error {
	err := g.db.WithContext(ctx).Where("group_id = ?", data.GroupID).Save(data).Error
	return stores.ErrFmt(err)
}

func (g GroupInfoRepo) DeleteByFilter(ctx context.Context, f GroupInfoFilter) error {
	db := g.fmtFilter(ctx, f)
	err := db.Delete(&DmGroupInfo{}).Error
	return stores.ErrFmt(err)
}
