package logic

import (
	"context"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogDeleteLogic) SysLogDelete(in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	err := l.svcCtx.SysLogModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.SysLogDeleteResp{}, nil
}
