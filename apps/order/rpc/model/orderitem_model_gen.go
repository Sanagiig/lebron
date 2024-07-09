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
	orderitemFieldNames          = builder.RawFieldNames(&Orderitem{})
	orderitemRows                = strings.Join(orderitemFieldNames, ",")
	orderitemRowsExpectAutoSet   = strings.Join(stringx.Remove(orderitemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderitemRowsWithPlaceHolder = strings.Join(stringx.Remove(orderitemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	orderitemModel interface {
		Insert(ctx context.Context, data *Orderitem) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Orderitem, error)
		Update(ctx context.Context, data *Orderitem) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultOrderitemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Orderitem struct {
		Id           uint64    `db:"id"`            // 订单子表id
		OrderId      string    `db:"order_id"`      // 订单id
		UserId       uint64    `db:"user_id"`       // 用户id
		ProductId    uint64    `db:"product_id"`    // 商品id
		ProductName  string    `db:"product_name"`  // 商品名称
		ProductImage string    `db:"product_image"` // 商品图片地址
		CurrentPrice float64   `db:"current_price"` // 生成订单时的商品单价，单位是元,保留两位小数
		Quantity     int64     `db:"quantity"`      // 商品数量
		TotalPrice   float64   `db:"total_price"`   // 商品总价,单位是元,保留两位小数
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newOrderitemModel(conn sqlx.SqlConn) *defaultOrderitemModel {
	return &defaultOrderitemModel{
		conn:  conn,
		table: "`orderitem`",
	}
}

func (m *defaultOrderitemModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderitemModel) FindOne(ctx context.Context, id uint64) (*Orderitem, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderitemRows, m.table)
	var resp Orderitem
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

func (m *defaultOrderitemModel) Insert(ctx context.Context, data *Orderitem) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderitemRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.UserId, data.ProductId, data.ProductName, data.ProductImage, data.CurrentPrice, data.Quantity, data.TotalPrice)
	return ret, err
}

func (m *defaultOrderitemModel) Update(ctx context.Context, data *Orderitem) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderitemRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.UserId, data.ProductId, data.ProductName, data.ProductImage, data.CurrentPrice, data.Quantity, data.TotalPrice, data.Id)
	return err
}

func (m *defaultOrderitemModel) tableName() string {
	return m.table
}
