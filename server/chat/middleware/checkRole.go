package middleware

import (
	"SLG_sg_server/constant"
	"SLG_sg_server/net"
	"log"
)

func CheckRoleId() net.MiddlewareFunc {
	return func(next net.HandlerFunc) net.HandlerFunc {
		return func(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
			log.Println("进入到角色检测....")
			_, err := req.Conn.GetProperty("rid")
			if err != nil {
				rsp.Body.Code = constant.SessionInvalid
				return
			}
			next(req, rsp)
		}
	}
}
