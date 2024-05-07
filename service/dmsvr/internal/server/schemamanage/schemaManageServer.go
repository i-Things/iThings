// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/logic/schemamanage"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

type SchemaManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedSchemaManageServer
}

func NewSchemaManageServer(svcCtx *svc.ServiceContext) *SchemaManageServer {
	return &SchemaManageServer{
		svcCtx: svcCtx,
	}
}

func (s *SchemaManageServer) CommonSchemaInit(ctx context.Context, in *dm.Empty) (*dm.Empty, error) {
	l := schemamanagelogic.NewCommonSchemaInitLogic(ctx, s.svcCtx)
	return l.CommonSchemaInit(in)
}

// 更新产品物模型
func (s *SchemaManageServer) CommonSchemaUpdate(ctx context.Context, in *dm.CommonSchemaUpdateReq) (*dm.Empty, error) {
	l := schemamanagelogic.NewCommonSchemaUpdateLogic(ctx, s.svcCtx)
	return l.CommonSchemaUpdate(in)
}

// 新增产品
func (s *SchemaManageServer) CommonSchemaCreate(ctx context.Context, in *dm.CommonSchemaCreateReq) (*dm.Empty, error) {
	l := schemamanagelogic.NewCommonSchemaCreateLogic(ctx, s.svcCtx)
	return l.CommonSchemaCreate(in)
}

// 删除产品
func (s *SchemaManageServer) CommonSchemaDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := schemamanagelogic.NewCommonSchemaDeleteLogic(ctx, s.svcCtx)
	return l.CommonSchemaDelete(in)
}

// 获取产品信息列表
func (s *SchemaManageServer) CommonSchemaIndex(ctx context.Context, in *dm.CommonSchemaIndexReq) (*dm.CommonSchemaIndexResp, error) {
	l := schemamanagelogic.NewCommonSchemaIndexLogic(ctx, s.svcCtx)
	return l.CommonSchemaIndex(in)
}
