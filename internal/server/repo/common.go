package repo

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	_db "github.com/utlai/utl/internal/pkg/db"
	"github.com/utlai/utl/internal/server/domain"
	"gorm.io/gorm"
	"strings"
)

type CommonRepo struct {
}

func NewCommonRepo() *CommonRepo {
	return &CommonRepo{}
}
func (r *CommonRepo) Defer(tx *gorm.DB, code *int) {
	if *code == 1 {
		//提交事务
		tx.Commit()
	} else {
		//回滚
		tx.Rollback()
	}
}

// GetAll 批量查询
func (r *CommonRepo) GetAll(model interface{}, s *serverDomain.Search) *gorm.DB {
	db := _db.GetInst().DB().Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))

	db.Scopes(r.FoundByWhere(s.Fields), r.Relation(s.Relations))

	return db
}

// Found 查询条件
func (r *CommonRepo) Found(s *serverDomain.Search) *gorm.DB {
	return _db.GetInst().DB().Scopes(r.Relation(s.Relations), r.FoundByWhere(s.Fields))
}

// IsNotFound 判断是否是查询不存在错误
func (r *CommonRepo) IsNotFound(err error) bool {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
		color.Yellow("查询数据不存在")
		return true
	}
	return false
}

// UpdateObj 更新
func (r *CommonRepo) UpdateObj(v, d interface{}, id uint) error {
	if err := _db.GetInst().DB().Model(v).Where("id = ?", id).Updates(d).Error; err != nil {
		color.Red(fmt.Sprintf("UpdateObj %+v to %+v\n", v, d))
		return err
	}
	return nil
}

// Relation 加载关联关系
func (r *CommonRepo) Relation(relates []*serverDomain.Relate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(relates) > 0 {
			for _, re := range relates {
				if len(re.Value) > 0 {
					if re.Func != nil {
						db = db.Preload(re.Value, re.Func)
					} else {
						db = db.Preload(re.Value)
					}
				}
				color.Yellow(fmt.Sprintf("Preoad %s", re))
			}
		}
		return db
	}
}

// FoundByWhere 查询条件
func (r *CommonRepo) FoundByWhere(fields []*serverDomain.Filed) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						color.Red(fmt.Sprintf("未知数据类型：%+v", field.Value))
					}
				}
			}
		}
		return db
	}
}

// GetRelations 转换前端获取关联关系为 []*Relate
func (r *CommonRepo) GetRelations(relation string, fs map[string]interface{}) []*serverDomain.Relate {
	var relates []*serverDomain.Relate
	if len(relation) > 0 {
		arr := strings.Split(relation, ";")
		for _, item := range arr {
			relate := &serverDomain.Relate{
				Value: item,
			}
			// 增加关联过滤
			for key, f := range fs {
				if key == item {
					relate.Func = f
				}
			}
			relates = append(relates, relate)
		}

	}
	color.Yellow(fmt.Sprintf("relation :%s , relates:%+v", relation, relates))
	return relates
}

// GetSearch 转换前端查询关系为 *Filed
func (r *CommonRepo) GetSearch(key, search string) *serverDomain.Filed {
	if len(search) > 0 {
		if strings.Contains(search, ":") {
			searches := strings.Split(search, ":")
			if len(searches) == 2 {
				value := searches[0]
				if strings.ToLower(searches[1]) == "like" {
					value = fmt.Sprintf("%%%s%%", searches[0])
				}

				return &serverDomain.Filed{
					Condition: searches[1],
					Key:       key,
					Value:     value,
				}

			} else if len(searches) == 1 {
				return &serverDomain.Filed{
					Condition: "=",
					Key:       key,
					Value:     searches[0],
				}
			}
		} else {
			return &serverDomain.Filed{
				Condition: "=",
				Key:       key,
				Value:     search,
			}
		}
	}
	return nil
}

// Paginate 分页
func (r *CommonRepo) Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

type SumRes struct {
	Total int64 `json:"total"`
}

// Filed 查询字段结构体
type Filed struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type Relate struct {
	Value string
	Func  interface{}
}

// Search 查询参数结构体
type Search struct {
	Fields    []*Filed  `json:"fields"`
	Relations []*Relate `json:"relations"`
	OrderBy   string    `json:"order_by"`
	Sort      string    `json:"sort"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}
