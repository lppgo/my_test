golang  简单封装日志包

usage :
    (1) 	location := time.FixedZone("Asia/Shanghai", 8*60*60)
    
    (2)配置日志文件
    
    	utils.InitLogger(config.Get("LOGPATH"), config.Get("DEBUG") != "0", location)、
        
    (3)使用
    
        utils.Logger.Infow("key", "xxxx", "999")
        
