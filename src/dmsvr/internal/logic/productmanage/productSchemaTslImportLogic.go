package productmanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/events"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/spf13/cast"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSchemaTslImportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB *relationDB.ProductInfoRepo
}

func NewProductSchemaTslImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSchemaTslImportLogic {
	return &ProductSchemaTslImportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
	}
}

// 删除产品
func (l *ProductSchemaTslImportLogic) ProductSchemaTslImport(in *dm.ProductSchemaTslImportReq) (*dm.Response, error) {
	l.Infof("%s req:%v", utils.FuncName(), in)
	_, err := l.PiDB.FindOneByFilter(l.ctx, relationDB.ProductFilter{ProductIDs: []string{in.ProductID}})
	if err != nil {
		if errors.Cmp(err, errors.NotFind) {
			return nil, errors.Parameter.AddDetail("not find ProductID id:" + cast.ToString(in.ProductID))
		}
		return nil, err
	}
	t, err := schema.ValidateWithFmt([]byte(in.Tsl))
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.HubLogRepo.InitProduct(
		l.ctx, in.ProductID); err != nil {
		l.Errorf("%s.DeviceLogRepo.InitProduct failure,err:%v", utils.FuncName(), err)
		return nil, err
	}
	{ //更新td物模型表
		oldT, err := l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, in.ProductID)
		if err != nil {
			l.Errorf("%s.SchemaManaRepo.GetSchemaModel failure,err:%v", utils.FuncName(), err)
			return nil, err
		}
		if err := l.svcCtx.SchemaManaRepo.DeleteProduct(l.ctx, oldT, in.ProductID); err != nil {
			l.Errorf("%s.SchemaManaRepo.InitProduct failure,err:%v", utils.FuncName(), err)
			return nil, err
		}
		if err := l.svcCtx.SchemaManaRepo.InitProduct(l.ctx, t, in.ProductID); err != nil {
			l.Errorf("%s.SchemaManaRepo.InitProduct failure,err:%v", utils.FuncName(), err)
			return nil, err
		}
		defer l.svcCtx.SchemaRepo.ClearCache(l.ctx, in.ProductID)
	}
	err = l.svcCtx.SchemaRepo.TslImport(l.ctx, in.ProductID, t)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DataUpdate.ProductSchemaUpdate(l.ctx, &events.DeviceUpdateInfo{ProductID: in.ProductID})
	if err != nil {
		return nil, err
	}
	return &dm.Response{}, nil
}
