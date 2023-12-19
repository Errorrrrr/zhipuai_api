package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ZhipuaiSendLogMgr struct {
	*_BaseMgr
}

// ZhipuaiSendLogMgr open func
func ZhipuaiSendLogMgr(db *gorm.DB) *_ZhipuaiSendLogMgr {
	if db == nil {
		panic(fmt.Errorf("ZhipuaiSendLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZhipuaiSendLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zhipuai_send_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZhipuaiSendLogMgr) GetTableName() string {
	return "zhipuai_send_log"
}

// Reset 重置gorm会话
func (obj *_ZhipuaiSendLogMgr) Reset() *_ZhipuaiSendLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZhipuaiSendLogMgr) Get() (result ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZhipuaiSendLogMgr) Gets() (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZhipuaiSendLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZhipuaiSendLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithReq req获取
func (obj *_ZhipuaiSendLogMgr) WithReq(req string) Option {
	return optionFunc(func(o *options) { o.query["req"] = req })
}

// WithResp resp获取
func (obj *_ZhipuaiSendLogMgr) WithResp(resp string) Option {
	return optionFunc(func(o *options) { o.query["resp"] = resp })
}

// WithCreatedAt created_at获取
func (obj *_ZhipuaiSendLogMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_ZhipuaiSendLogMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZhipuaiSendLogMgr) GetByOption(opts ...Option) (result ZhipuaiSendLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZhipuaiSendLogMgr) GetByOptions(opts ...Option) (results []*ZhipuaiSendLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZhipuaiSendLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZhipuaiSendLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where(options.query)
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
func (obj *_ZhipuaiSendLogMgr) GetFromID(id int64) (result ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZhipuaiSendLogMgr) GetBatchFromID(ids []int64) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromReq 通过req获取内容
func (obj *_ZhipuaiSendLogMgr) GetFromReq(req string) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`req` = ?", req).Find(&results).Error

	return
}

// GetBatchFromReq 批量查找
func (obj *_ZhipuaiSendLogMgr) GetBatchFromReq(reqs []string) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`req` IN (?)", reqs).Find(&results).Error

	return
}

// GetFromResp 通过resp获取内容
func (obj *_ZhipuaiSendLogMgr) GetFromResp(resp string) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`resp` = ?", resp).Find(&results).Error

	return
}

// GetBatchFromResp 批量查找
func (obj *_ZhipuaiSendLogMgr) GetBatchFromResp(resps []string) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`resp` IN (?)", resps).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZhipuaiSendLogMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ZhipuaiSendLogMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ZhipuaiSendLogMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ZhipuaiSendLogMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZhipuaiSendLogMgr) FetchByPrimaryKey(id int64) (result ZhipuaiSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZhipuaiSendLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

func (obj *_ZhipuaiSendLogMgr) Insert(model *ZhipuaiSendLog) error {
	err := obj.DB.WithContext(obj.ctx).Create(model).Error
	return err
}
