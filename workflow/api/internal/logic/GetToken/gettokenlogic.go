package GetToken

import (
	"context"
	"time"
	"zero-workflow/common"
	"zero-workflow/common/jwtx"
	"zero-workflow/common/msg"

	"zero-workflow/workflow/api/internal/svc"
	"zero-workflow/workflow/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenLogic) GetToken(req *types.GetTokenRequest) (resp *types.Response, err error) {
	// 生成Token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, req.Uid, common.UserTokenType, req.NickName, req.TenantId)

	var resData LoginResponse
	resData.AccessToken = accessToken
	resData.AccessExpire = now + accessExpire

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: resData,
	}, nil

}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}
