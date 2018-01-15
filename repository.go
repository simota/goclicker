package goclicker

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository() LinkRepository {
	return LinkRepository{db: GetDatabase()}
}

func (r *LinkRepository) Find(id uint64) *Link {
	var item Link
	if r.db.First(&item, id).RecordNotFound() {
		return nil
	}
	return &item
}

func (r *LinkRepository) FindByKey(s string) *Link {
	var item Link
	if r.db.Where("`key` = ?", s).First(&item).RecordNotFound() {
		return nil
	}
	return &item
}

func (r *LinkRepository) FindByURL(s string) *Link {
	var item Link
	if r.db.Where("url = ?", s).First(&item).RecordNotFound() {
		return nil
	}
	return &item
}

func (r *LinkRepository) Save(item Link) {
	if item.ID == 0 {
		item.Key = CreateKey(item.URL)
		item.CreatedAt = time.Now()
	}
	item.UpdatedAt = time.Now()
	r.db.Save(&item)
}

func (r *LinkRepository) All() []Link {
	var items []Link
	r.db.Order("id desc").Find(&items)
	return items
}

type AccessLogRepository struct {
	db *gorm.DB
}

func NewAccessLogRepository() AccessLogRepository {
	return AccessLogRepository{db: GetDatabase()}
}

func (r *AccessLogRepository) Save(item AccessLog) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	r.db.Save(&item)
}

func (r *AccessLogRepository) FindByLinkID(id uint64) []AccessLog {
	var items []AccessLog
	r.db.Where("link_id = ?", id).Order("id desc").Find(&items)
	return items
}
