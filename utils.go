package goku_plugin

var (
	_redisManager RedisManager
	_logGeneral   LoggerGeneral
)

func SetRedisManager(manager RedisManager) {
	if _redisManager != nil {
		panic("repeat set RedisManager")
	}
	_redisManager = manager
}

func SetLog(general LoggerGeneral) {
	if _logGeneral != nil {
		panic("repeat set LoggerGeneral")
	}
	_logGeneral = general
}
func GenAccessLogger(dirName, fileName string, period LogPeriod) Logger {
	if _logGeneral == nil {
		return nil
	}

	return _logGeneral.GenAccessLogger(dirName, fileName, period)
}
func GetRedis() Redis {
	if _redisManager == nil {
		return nil
	}
	return _redisManager.Default()
}
func GetRedisByName(name string) (Redis, bool) {
	return _redisManager.Get(name)
}
