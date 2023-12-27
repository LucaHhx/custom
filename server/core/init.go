package core

import (
	"custom-server/global"
	"custom-server/initialize"
	"go.uber.org/zap"
)

func Init() {
	global.GVA_VP = Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		if global.GVA_CONFIG.System.Migrate {
			initialize.RegisterTables() // 初始化表
		}

		// 程序结束前关闭数据库链接
		//db, _ := global.GVA_DB.DB()
		//defer db.Close()
	}
}
