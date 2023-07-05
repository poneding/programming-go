package database

import (
	"fmt"
)

type Orm struct {
	Database *database
}

const (
	DefaultListLimit = 100
	DefaultPageNo    = 1
	DefaultPageSize  = 10
	OrderInDesc      = "DESC"
)

func (r *Orm) IsExist(id interface{}, entry interface{}) bool {
	if err := r.Database.First(entry, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

func (r *Orm) Get(id, entry interface{}) error {
	if err := r.Database.First(entry, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *Orm) GetAndInclude(id, entry interface{}, includes ...string) error {
	query := r.Database.DB
	if len(includes) > 0 {
		for _, in := range includes {
			query = query.Preload(in)
		}
		query = query.Preload("")
	}
	if err := query.First(entry, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

type FirstOption struct {
	Condition *condition
}

func (r *Orm) First(entry interface{}, opt ...FirstOption) error {
	var firstOption FirstOption
	if len(opt) > 0 {
		firstOption = opt[0]
	}
	query := r.Database.Model(entry)
	if firstOption.Condition != nil {
		query = query.Where(firstOption.Condition.Expression, firstOption.Condition.Arguments...)
	}
	return query.First(entry).Error
}

func (r *Orm) Create(entry interface{}) error {
	return r.Database.Model(entry).Create(entry).Error
}

type UpdateOption struct {
	Condition *condition
	Columns   []string
}

func (r *Orm) Update(entry interface{}, opt ...UpdateOption) error {
	var updateOption UpdateOption
	if len(opt) > 0 {
		updateOption = opt[0]
	}
	query := r.Database.Model(entry)
	if updateOption.Condition != nil {
		query = query.Where(updateOption.Condition.Expression, updateOption.Condition.Arguments...)
	}
	return query.Select(updateOption.Columns).Updates(entry).Error
}

func (r *Orm) Delete(id interface{}, entry interface{}) error {
	return r.Database.Where("id = ?", id).Delete(entry).Error
}

type SweepOption struct {
	Condition *condition
}

func (r *Orm) Sweep(entry interface{}, opt SweepOption) error {
	if opt.Condition == nil {
		return nil
	}
	return r.Database.Where(opt.Condition.Expression, opt.Condition.Arguments...).Delete(entry).Error
}

type condition struct {
	Expression string
	Arguments  []interface{}
}

func Condition(exp string, args ...interface{}) *condition {
	return &condition{
		Expression: exp,
		Arguments:  args,
	}
}

type ListOption struct {
	Condition *condition
	OrderBy   string
	Desc      bool
}

func (r *Orm) List(entries interface{}, opt ...ListOption) error {
	var listOption ListOption
	if len(opt) > 0 {
		listOption = opt[0]
	}
	query := r.Database.DB

	if listOption.Condition != nil {
		query = query.Where(listOption.Condition.Expression, listOption.Condition.Arguments...)
	}

	if len(listOption.OrderBy) > 0 {
		if listOption.Desc {
			listOption.OrderBy = fmt.Sprintf("%s %s", listOption.OrderBy, OrderInDesc)
		}
		query = query.Order(listOption.OrderBy)
	}
	return query.Find(entries).Error
}

type PageOption struct {
	No        int
	Size      int
	OrderBy   string
	Desc      *bool
	Condition *condition
}

func (r *Orm) Page(entries interface{}, opt ...PageOption) (int64, error) {
	var pageOption PageOption
	if len(opt) > 0 {
		pageOption = opt[0]
		if pageOption.Size == 0 {
			pageOption.Size = 10
		}
	} else {
		pageOption = PageOption{
			No:   DefaultPageNo,
			Size: DefaultPageSize,
		}
	}
	query := r.Database.DB

	if pageOption.Condition != nil {
		query = query.Where(pageOption.Condition.Expression, pageOption.Condition.Arguments...)
	}
	if len(pageOption.OrderBy) == 0 {
		pageOption.OrderBy = "updated_at"
	}
	desc := pageOption.Desc
	if pageOption.Desc == nil || *desc {
		pageOption.OrderBy = fmt.Sprintf("%s %s", pageOption.OrderBy, OrderInDesc)
	}
	query = query.Order(pageOption.OrderBy)
	var total int64
	query.Model(entries).Count(&total)
	err := query.Offset((pageOption.No - 1) * pageOption.Size).Limit(pageOption.Size).Find(entries).Error
	return total, err
}

func (r *Orm) All(entries interface{}) error {
	return r.Database.Find(entries).Error
}
