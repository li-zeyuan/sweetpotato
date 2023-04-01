package router

import (
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/config"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/handler"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.RecoveryWithWriter(mylogger.GetWriter(config.AppCfg.Logging.LoggingDir+"/error.log")))
	engine.Use(httptransfer.RequestIdMiddleware())

	engine.Use(httptransfer.NotStrictAuthorizationMiddleware(config.AppCfg.JwtSecret))
	engine.POST("/api/login/wechat", handler.LoginHandler.WechatLogin)
	engine.GET("/api/subject/list", handler.SubjectHandler.List)
	engine.GET("/api/subject/detail", handler.SubjectHandler.Detail)

	engine.Use(httptransfer.StrictAuthorizationMiddleware(config.AppCfg.JwtSecret))
	engine.GET("/api/user/detail", handler.UserHandler.Detail)
	engine.GET("/api/subject/study", handler.SubjectHandler.Study)
	engine.PUT("/api/user/study_num_edit", handler.UserHandler.StudyNumEdit)
	engine.GET("/api/study/record", handler.StudyHandler.RecordList)
	engine.GET("/api/study/knowledge", handler.StudyHandler.KnowledgeList)
	engine.POST("/api/study/doing", handler.StudyHandler.Doing)

	return engine
}
