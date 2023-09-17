package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"github.com/mssola/user_agent"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {

	ua := user_agent.New(ctxs.GetUserCtx(l.ctx).Os)
	browser, _ := ua.Browser()
	os := ua.OS()

	l.Infof("%s req=%+v", utils.FuncName(), req)
	if req.LoginType == "pwd" && req.Code != "" && req.CodeID != "" {
		if l.svcCtx.Captcha.Verify(req.CodeID, req.Code) == false {
			return nil, errors.Captcha
		}
	}
	uResp, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &sys.UserLoginReq{
		Account:   req.Account,
		PwdType:   req.PwdType,
		Password:  req.Password,
		LoginType: req.LoginType,
		Code:      req.Code,
		CodeID:    req.CodeID,
		Ip:        ctxs.GetUserCtx(l.ctx).IP,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.Login req=%v err=%+v", utils.FuncName(), req, er)
		//登录失败记录
		l.svcCtx.LogRpc.LoginLogCreate(l.ctx, &sys.LoginLogCreateReq{
			UserID:        0,
			UserName:      req.Account, //todo 这里也要调整
			IpAddr:        ctxs.GetUserCtx(l.ctx).IP,
			LoginLocation: GetCityByIp(ctxs.GetUserCtx(l.ctx).IP),
			Browser:       browser,
			Os:            os,
			Msg:           er.Error(),
			Code:          400,
		})
		return nil, er
	}
	if uResp == nil {
		l.Errorf("%s.rpc.Register return nil req=%v", utils.FuncName(), req)
		//登录失败记录
		l.svcCtx.LogRpc.LoginLogCreate(l.ctx, &sys.LoginLogCreateReq{
			UserID:        0,
			UserName:      req.Account,
			IpAddr:        ctxs.GetUserCtx(l.ctx).IP,
			LoginLocation: GetCityByIp(ctxs.GetUserCtx(l.ctx).IP),
			Browser:       browser,
			Os:            os,
			Msg:           "register core rpc return nil",
			Code:          400,
		})
		return nil, errors.System.AddDetail("register core rpc return nil")
	}
	//登录成功记录
	l.svcCtx.LogRpc.LoginLogCreate(l.ctx, &sys.LoginLogCreateReq{
		UserID:        uResp.Info.UserID,
		UserName:      uResp.Info.UserName,
		IpAddr:        ctxs.GetUserCtx(l.ctx).IP,
		LoginLocation: GetCityByIp(ctxs.GetUserCtx(l.ctx).IP),
		Browser:       browser,
		Os:            os,
		Msg:           "登录成功",
		Code:          200,
	})

	return &types.UserLoginResp{
		Info: types.UserInfo{
			UserID:      uResp.Info.UserID,
			UserName:    uResp.Info.UserName,
			Password:    "",
			Email:       uResp.Info.Email,
			Phone:       uResp.Info.Phone,
			Wechat:      uResp.Info.Wechat,
			LastIP:      uResp.Info.LastIP,
			RegIP:       uResp.Info.RegIP,
			NickName:    uResp.Info.NickName,
			City:        uResp.Info.City,
			Country:     uResp.Info.Country,
			Province:    uResp.Info.Province,
			Language:    uResp.Info.Language,
			HeadImgUrl:  uResp.Info.HeadImgUrl,
			CreatedTime: uResp.Info.CreatedTime,
			Role:        uResp.Info.Role,
			Sex:         uResp.Info.Sex,
			IsAllData:   uResp.Info.IsAllData,
		},
		Token: types.JwtToken{
			AccessToken:  uResp.Token.AccessToken,
			AccessExpire: uResp.Token.AccessExpire,
			RefreshAfter: uResp.Token.RefreshAfter,
		},
	}, nil

	return
}
