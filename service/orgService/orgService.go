package orgService

import (
	"github.com/gin-gonic/gin"
	"sauth/model/orgModel"
	"sauth/db"
	"sauth/util"
	"database/sql"
)

func Find(ctx *gin.Context) ([]orgModel.Org, error) {
	model := make([]orgModel.Org, 0)
	err := db.Engine.Find(&model)
	return model, err
}

func Save(ctx *gin.Context) (int64, error) {
	var model orgModel.Org
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return -1, err
	}
	if "" != model.Uuid {
		res, err = db.Engine.Update(&model, &orgModel.Org{Uuid: model.Uuid})
	} else {
		model.Uuid = util.GetUuid()
		model.Id = model.TenantId + `-` + model.OrgId
		res, err = db.Engine.Insert(&model)
	}
	return res, err
}

func Delete(ctx *gin.Context) (sql.Result, error) {
	var model orgModel.Org
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}
	res, err := orgModel.Delete(model.Uuid)
	return res, err
}
