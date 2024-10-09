package controller

import (
	"SLG_sg_server/constant"
	"SLG_sg_server/net"
	"SLG_sg_server/server/common"
	"SLG_sg_server/server/game/logic"
	"SLG_sg_server/server/game/middleware"
	"SLG_sg_server/server/game/model"
	"SLG_sg_server/server/game/model/data"
)

var SkillController = &skillController{}

type skillController struct {
}

func (sh *skillController) Router(r *net.Router) {
	g := r.Group("skill")
	g.Use(middleware.Log())
	g.AddRouter("list", sh.list, middleware.CheckRole())
}
func (s *skillController) list(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//查找战报表 得出数据
	rspObj := &model.SkillListRsp{}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}
	rid := role.(*data.Role).RId
	skills, err := logic.SkillService.GetSkills(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rspObj.List = skills
	rsp.Body.Msg = rspObj
	rsp.Body.Code = constant.OK
}
