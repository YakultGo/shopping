package server

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	smsPb "shopping/api/sms"
	"shopping/pkg/util"
)

//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/verify_code.lua
var luaVerifyCode string

type SmsServer struct {
	smsPb.UnimplementedSmsServer
	client redis.Cmdable
}

func NewSmsServer(client redis.Cmdable) *SmsServer {
	return &SmsServer{
		client: client,
	}
}

func (s *SmsServer) SendCode(ctx context.Context, req *smsPb.SendCodeRequest) (*smsPb.SendCodeResponse, error) {
	code := util.GetRandomCode()
	res, err := s.client.Eval(ctx, luaSetCode, []string{s.key(req.Biz, req.Phone)}, code).Int()
	if err != nil {
		return nil, err
	}
	switch res {
	case 0:
		return &smsPb.SendCodeResponse{
			Code: code,
		}, nil
	case -1:
		return nil, errors.New("发送验证码太频繁")
	default:
		return nil, errors.New("系统错误")
	}
}

func (s *SmsServer) VerifyCode(ctx context.Context, req *smsPb.VerifyCodeRequest) (*smsPb.VerifyCodeResponse, error) {
	res, err := s.client.Eval(ctx, luaVerifyCode, []string{s.key(req.Biz, req.Phone)}, req.Code).Int()
	if err != nil {
		return nil, err
	}
	switch res {
	case 0:
		return &smsPb.VerifyCodeResponse{Success: true}, nil
	case -1:
		// 频繁出现可能是被恶意攻击
		return &smsPb.VerifyCodeResponse{Success: false}, errors.New("验证码错误次数太多")
	case -2:
		return &smsPb.VerifyCodeResponse{Success: false}, errors.New("验证码错误")
	default:
		return &smsPb.VerifyCodeResponse{Success: false}, errors.New("系统错误")
	}
}

func (s *SmsServer) key(biz string, phone string) string {
	return fmt.Sprintf("%s:%s", biz, phone)
}
