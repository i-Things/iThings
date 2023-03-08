// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dmProductSchemaFieldNames          = builder.RawFieldNames(&DmProductSchema{})
	dmProductSchemaRows                = strings.Join(dmProductSchemaFieldNames, ",")
	dmProductSchemaRowsExpectAutoSet   = strings.Join(stringx.Remove(dmProductSchemaFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	dmProductSchemaRowsWithPlaceHolder = strings.Join(stringx.Remove(dmProductSchemaFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	dmProductSchemaModel interface {
		Insert(ctx context.Context, data *DmProductSchema) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DmProductSchema, error)
		FindOneByProductIDIdentifier(ctx context.Context, productID string, identifier string) (*DmProductSchema, error)
		Update(ctx context.Context, data *DmProductSchema) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDmProductSchemaModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DmProductSchema struct {
		Id          int64        `db:"id"`
		ProductID   string       `db:"productID"`  // 产品id
		Tag         int64        `db:"tag"`        // 物模型标签 1:自定义 2:可选 3:必选  必选不可删除
		Type        int64        `db:"type"`       // 物模型类型 1:property属性 2:event事件 3:action行为
		Identifier  string       `db:"identifier"` // 标识符
		Name        string       `db:"name"`       // 功能名称
		Desc        string       `db:"desc"`       // 描述
		Required    int64        `db:"required"`   // 是否必须,1是 2否
		Affordance  string       `db:"affordance"` // 各类型的自定义功能定义
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime time.Time    `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
	}
)

func newDmProductSchemaModel(conn sqlx.SqlConn) *defaultDmProductSchemaModel {
	return &defaultDmProductSchemaModel{
		conn:  conn,
		table: "`dm_product_schema`",
	}
}

func (m *defaultDmProductSchemaModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDmProductSchemaModel) FindOne(ctx context.Context, id int64) (*DmProductSchema, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dmProductSchemaRows, m.table)
	var resp DmProductSchema
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmProductSchemaModel) FindOneByProductIDIdentifier(ctx context.Context, productID string, identifier string) (*DmProductSchema, error) {
	var resp DmProductSchema
	query := fmt.Sprintf("select %s from %s where `productID` = ? and `identifier` = ? limit 1", dmProductSchemaRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, productID, identifier)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmProductSchemaModel) Insert(ctx context.Context, data *DmProductSchema) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, dmProductSchemaRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.Tag, data.Type, data.Identifier, data.Name, data.Desc, data.Required, data.Affordance)
	return ret, err
}

func (m *defaultDmProductSchemaModel) Update(ctx context.Context, newData *DmProductSchema) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dmProductSchemaRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ProductID, newData.Tag, newData.Type, newData.Identifier, newData.Name, newData.Desc, newData.Required, newData.Affordance, newData.Id)
	return err
}

func (m *defaultDmProductSchemaModel) tableName() string {
	return m.table
}
