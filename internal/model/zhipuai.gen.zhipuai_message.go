package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ZhipuaiMessageMgr struct {
	*_BaseMgr
}

// ZhipuaiMessageMgr open func
func ZhipuaiMessageMgr(db *gorm.DB) *_ZhipuaiMessageMgr {
	if db == nil {
		panic(fmt.Errorf("ZhipuaiMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZhipuaiMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zhipuai_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZhipuaiMessageMgr) GetTableName() string {
	return "zhipuai_message"
}

// Reset 重置gorm会话
func (obj *_ZhipuaiMessageMgr) Reset() *_ZhipuaiMessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZhipuaiMessageMgr) Get() (result ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZhipuaiMessageMgr) Gets() (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZhipuaiMessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZhipuaiMessageMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleName role_name获取
func (obj *_ZhipuaiMessageMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithContent content获取
func (obj *_ZhipuaiMessageMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithUUID uuid获取
func (obj *_ZhipuaiMessageMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithDeleted deleted获取
func (obj *_ZhipuaiMessageMgr) WithDeleted(deleted int64) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithCreatedAt created_at获取
func (obj *_ZhipuaiMessageMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZhipuaiMessageMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZhipuaiMessageMgr) GetByOption(opts ...Option) (result ZhipuaiMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZhipuaiMessageMgr) GetByOptions(opts ...Option) (results []*ZhipuaiMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZhipuaiMessageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZhipuaiMessage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where(options.query)
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
func (obj *_ZhipuaiMessageMgr) GetFromID(id int64) (result ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromID(ids []int64) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容
func (obj *_ZhipuaiMessageMgr) GetFromRoleName(roleName string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`role_name` = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromRoleName(roleNames []string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`role_name` IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_ZhipuaiMessageMgr) GetFromContent(content string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromContent(contents []string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_ZhipuaiMessageMgr) GetFromUUID(uuid string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromUUID(uuids []string) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_ZhipuaiMessageMgr) GetFromDeleted(deleted int64) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromDeleted(deleteds []int64) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZhipuaiMessageMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZhipuaiMessageMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZhipuaiMessageMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZhipuaiMessageMgr) FetchByPrimaryKey(id int64) (result ZhipuaiMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

func (obj *_ZhipuaiMessageMgr) Insert(model *ZhipuaiMessage) (err error) {
	err = obj.DB.WithContext(obj.ctx).Create(model).Error
	return
}

func (obj *_ZhipuaiMessageMgr) GetBySize(limit int64) ([]*ZhipuaiMessage, error) {
	var result []*ZhipuaiMessage
	err := obj.DB.WithContext(obj.ctx).Model(ZhipuaiMessage{}).Limit(int(limit)).Find(&result).Error
	return result, err
}
