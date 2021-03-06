package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateDeleteLogic {
	return &FeightTemplateDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(in *pms.FeightTemplateDeleteReq) (*pms.FeightTemplateDeleteResp, error) {
	err := l.svcCtx.PmsFeightTemplateModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.FeightTemplateDeleteResp{}, nil
}
