package pattern
//保存文章接口
type ISaveArticle interface {
	CreateProse() IProse
	CreateAncientPoetry() IAncientPoetry
}
//散文接口
type IProse interface {
	SaveProse() 
}
//古诗接口
type IAncientPoetry interface {
	SaveAncientPoetry()
}