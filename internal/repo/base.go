package repo

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/coin50etf/coin-market/internal/pkg/database"
)

// BaseRepository 提供通用的 CRUD 操作
type BaseRepository[T any] struct {
	PostgresDB *database.PostgresDB
}

// Create 新增单条数据（支持事务）
func (r *BaseRepository[T]) Create(data *T, tx ...*gorm.DB) error {
	return r.getDB(tx...).Create(data).Error
}

// CreateInBatches 批量新增数据
func (r *BaseRepository[T]) CreateInBatches(data []*T, batchSize int, tx ...*gorm.DB) error {
	return r.getDB(tx...).CreateInBatches(data, batchSize).Error
}

// UpdateByCondition 通过条件更新指定字段
func (r *BaseRepository[T]) UpdateByCondition(condition map[string]interface{}, updates map[string]interface{}, tx ...*gorm.DB) error {
	return r.getDB(tx...).Model(new(T)).Where(condition).Updates(updates).Error
}

// DeleteByCondition 根据条件删除
func (r *BaseRepository[T]) DeleteByCondition(condition map[string]interface{}, tx ...*gorm.DB) error {
	result := r.getDB(tx...).Where(condition).Delete(new(T))
	if result.RowsAffected == 0 {
		return errors.New("no records found to delete")
	}

	return result.Error
}

// FindByID 通过 ID 查询单条数据
func (r *BaseRepository[T]) FindByID(id uint64, tx ...*gorm.DB) (*T, error) {
	var entity T
	if err := r.getDB(tx...).First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAll 获取所有数据（支持查询修改器，如分页、排序）
func (r *BaseRepository[T]) FindAll(queryModifier ...func(*gorm.DB) *gorm.DB) ([]*T, error) {
	var entities []*T
	db := r.getDB()

	for _, modifier := range queryModifier {
		db = modifier(db)
	}

	err := db.Find(&entities).Error

	return entities, err
}

// FindByFieldIn 通过字段值列表查询数据
func (r *BaseRepository[T]) FindByFieldIn(field string, values []interface{}, tx ...*gorm.DB) ([]*T, error) {
	var results []*T
	if err := r.getDB(tx...).Where(fmt.Sprintf("`%s` IN (?)", field), values).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// FindByCondition 通过条件查询单条数据（支持查询修改器）
func (r *BaseRepository[T]) FindByCondition(condition map[string]interface{}, queryModifier ...func(*gorm.DB) *gorm.DB) (*T, error) {
	var entity T
	db := r.getDB().Where(condition)

	for _, modifier := range queryModifier {
		db = modifier(db)
	}

	if err := db.First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAllByCondition 通过条件查询数据列表（支持查询修改器）
func (r *BaseRepository[T]) FindAllByCondition(condition map[string]interface{}, queryModifier ...func(*gorm.DB) *gorm.DB) ([]*T, error) {
	var entities []*T
	db := r.getDB().Where(condition)

	for _, modifier := range queryModifier {
		db = modifier(db)
	}

	err := db.Find(&entities).Error

	return entities, err
}

// Count 计算符合条件的数据数量
func (r *BaseRepository[T]) Count(condition map[string]interface{}, tx ...*gorm.DB) (int64, error) {
	var total int64
	err := r.getDB(tx...).Model(new(T)).Where(condition).Count(&total).Error

	return total, err
}

// Exists 判断是否存在符合条件的数据
func (r *BaseRepository[T]) Exists(condition map[string]interface{}, tx ...*gorm.DB) (bool, error) {
	count, err := r.Count(condition, tx...)

	return count > 0, err
}

// FindByPage 通过条件分页获取数据
func (r *BaseRepository[T]) FindByPage(condition map[string]interface{}, page int, pageSize int, orderBy string, queryModifier ...func(*gorm.DB) *gorm.DB) ([]T, int64, error) {
	var entities []T
	var total int64
	db := r.getDB().Model(new(T)).Where(condition)

	for _, modifier := range queryModifier {
		db = modifier(db)
	}

	// 先统计总条数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	query := db.Limit(pageSize).Offset(offset)

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	result := query.Find(&entities)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return entities, total, nil
}

// RunTransaction 执行事务
func (r *BaseRepository[T]) RunTransaction(txFunc func(tx *gorm.DB) error) error {
	return r.getDB().Transaction(txFunc)
}

// getDB 获取数据库实例（支持事务）
func (r *BaseRepository[T]) getDB(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 && tx[0] != nil {
		return tx[0]
	}

	return r.PostgresDB.DB
}
