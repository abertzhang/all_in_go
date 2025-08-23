package pattern
//实现ISaveArticle接口
type SaveRedis struct {
}

func (s *SaveRedis) CreateProse() IProse {
	return &RedisArticle{}
}

func (s *SaveRedis) CreateAncientPoetry() IAncientPoetry {
	return &MySQLArticle{}
}
//实现ISaveArticle接口
type SaveMySQL struct {
}

func (s *SaveMySQL) CreateProse() IProse {
	return &MySQLArticle{}
}

func (s *SaveMySQL) CreateAncientPoetry() IAncientPoetry {
	return &MySQLArticle{}
}

// 用于Redis的文章
//实现两类文章IProse和IAncientPoetry接口
type RedisArticle struct{}
func (r *RedisArticle) SaveProse() {
	println("save prose to redis")
}
func (r *RedisArticle) SaveAncientPoetry() {
	println("save ancient poetry to redis")
}
//用于保存mysql的文章
//实现两类文章IProse和IAncientPoetry接口
type MySQLArticle struct{}
func (m *MySQLArticle) SaveProse() {
	println("save prose to mysql")
}
func (m *MySQLArticle) SaveAncientPoetry() {
	println("save ancient poetry to mysql")
}