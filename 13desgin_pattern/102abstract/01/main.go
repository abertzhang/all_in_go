package main

import (
	"studyGo/13desgin_pattern/102abstract/01/pattern"
)
func main() {	
	ExampleSaveRedis()
	ExampleSaveMySQL()
}
//
func Save(saveArticle pattern.ISaveArticle) {
	saveArticle.CreateProse().SaveProse()
	saveArticle.CreateAncientPoetry().SaveAncientPoetry()
}
func ExampleSaveRedis() {
	factory := &pattern.SaveRedis{}
	Save(factory)
}
func ExampleSaveMySQL() {
	factory := &pattern.SaveMySQL{}
	Save(factory)	
}