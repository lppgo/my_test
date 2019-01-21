golang  简单封装日志包

usage :
    (1)配置日志文件
    	utils.InitLogger(config.Get("LOGPATH"), config.Get("DEBUG") != "0", location)、
    (2)使用
        utils.Logger.Infow("key", "xxxx", "999")