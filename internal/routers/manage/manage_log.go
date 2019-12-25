package manage

import (
	"github.com/infamt/gocron/internal/models"
	"github.com/infamt/gocron/internal/modules/logger"
	"github.com/infamt/gocron/internal/modules/utils"
	"github.com/infamt/gocron/internal/routers/base"
	macaron "gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) string {
	actionLogModel := new(models.ActionLog)
	params := models.CommonMap{}
	base.ParsePageAndPageSize(ctx, params)
	total, err := actionLogModel.Total()
	if err != nil {
		logger.Error(err)
	}
	actionLogs, err := actionLogModel.List(params)
	if err != nil {
		logger.Error(err)
	}

	jsonResp := utils.JsonResponse{}

	return jsonResp.Success(utils.SuccessContent, map[string]interface{}{
		"total": total,
		"data":  actionLogs,
	})
}
