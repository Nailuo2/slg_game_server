package middleware

import (
	"SLG_sg_server/net"
	"fmt"
	"log"
)

func Log() net.MiddlewareFunc {
	return func(next net.HandlerFunc) net.HandlerFunc {
		return func(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
			log.Println("请求路由", req.Body.Name)
			log.Println("请求参数", fmt.Sprintf("%v", req.Body.Msg))
			next(req, rsp)
		}
	}
}
