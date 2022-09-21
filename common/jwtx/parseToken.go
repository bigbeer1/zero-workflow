package jwtx

import (
	"context"
	"fmt"
)

func ParseToken(ctx context.Context) *tokenData {
	// 用户登录信息
	var uidInterface = ctx.Value("uid")
	var nickNameInterface = ctx.Value("nickName")
	var tenantIdInterface = ctx.Value("tenantId")
	uid := fmt.Sprintf("%v", uidInterface)
	tokenType := fmt.Sprintf("%v", nickNameInterface)
	tenantId := fmt.Sprintf("%v", tenantIdInterface)
	return &tokenData{
		Uid:      uid,
		NickName: tokenType,
		TenantId: tenantId,
	}
}

type tokenData struct {
	Uid      string `json:"uid"`
	NickName string `json:"nick_name"` // 昵称
	TenantId string `json:"tenantId"`  // 租户Id
}
