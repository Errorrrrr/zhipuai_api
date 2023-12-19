package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ZhipuaiSessionMgr struct {
	*_BaseMgr
}

// ZhipuaiSessionMgr open func
func ZhipuaiSessionMgr(db *gorm.DB) *_ZhipuaiSessionMgr {
	if db == nil {
		panic(fmt.Errorf("ZhipuaiSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZhipuaiSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zhipuai_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZhipuaiSessionMgr) GetTableName() string {
	return "zhipuai_session"
}

// Reset 重置gorm会话
func (obj *_ZhipuaiSessionMgr) Reset() *_ZhipuaiSessionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZhipuaiSessionMgr) Get() (result ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZhipuaiSessionMgr) Gets() (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZhipuaiSessionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZhipuaiSessionMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_ZhipuaiSessionMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithUUID uuid获取
func (obj *_ZhipuaiSessionMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithDeleted deleted获取
func (obj *_ZhipuaiSessionMgr) WithDeleted(deleted int64) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithCreatedAt created_at获取
func (obj *_ZhipuaiSessionMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZhipuaiSessionMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZhipuaiSessionMgr) GetByOption(opts ...Option) (result ZhipuaiSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZhipuaiSessionMgr) GetByOptions(opts ...Option) (results []*ZhipuaiSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZhipuaiSessionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZhipuaiSession, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ZhipuaiSessionMgr) GetFromID(id int64) (result ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromID(ids []int64) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_ZhipuaiSessionMgr) GetFromTitle(title string) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromTitle(titles []string) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_ZhipuaiSessionMgr) GetFromUUID(uuid string) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromUUID(uuids []string) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_ZhipuaiSessionMgr) GetFromDeleted(deleted int64) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromDeleted(deleteds []int64) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZhipuaiSessionMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZhipuaiSessionMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZhipuaiSessionMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZhipuaiSessionMgr) FetchByPrimaryKey(id int64) (result ZhipuaiSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Where("`id` = ?", id).First(&result).Error

	return
}

func (obj *_ZhipuaiSessionMgr) Insert(model *ZhipuaiSession) (err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSession{}).Create(model).Error

	return
}
