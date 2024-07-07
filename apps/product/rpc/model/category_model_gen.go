// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	categoryFieldNames          = builder.RawFieldNames(&Category{})
	categoryRows                = strings.Join(categoryFieldNames, ",")
	categoryRowsExpectAutoSet   = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	categoryRowsWithPlaceHolder = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Category, error)
		Update(ctx context.Context, data *Category) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Category struct {
		Id         uint64    `db:"id"`          // 分类id
		Parentid   int64     `db:"parentid"`    // 父类别id当id=0时说明是根节点,一级类别
		Name       string    `db:"name"`        // 类别名称
		Status     int64     `db:"status"`      // 类别状态1-正常,2-已废弃
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newCategoryModel(conn sqlx.SqlConn) *defaultCategoryModel {
	return &defaultCategoryModel{
		conn:  conn,
		table: "`category`",
	}
}

func (m *defaultCategoryModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id uint64) (*Category, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
	var resp Category
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, categoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status)
	return ret, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status, data.Id)
	return err
}

func (m *defaultCategoryModel) tableName() string {
	return m.table
}
