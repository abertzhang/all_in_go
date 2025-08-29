package main

import (
	"fmt"
	"studyGo/13desgin_pattern/301chain/02/pattern"
)

func main() {
//初始化空的handler
nullHandler := &pattern.NullHandler{}
//链式调用
nullHandler.SetNext(&pattern.ArgumentHandler{}).
	SetNext(&pattern.AddressInfoHandler{}).
	SetNext(&pattern.CartInfoHandler{}).
	SetNext(&pattern.StockInfoHandler{}).
	SetNext(&pattern.PromotionInfoHandler{}).
	SetNext(&pattern.ShipmentInfoHandler{}).
	SetNext(&pattern.PromotionUseHandler{}).
	SetNext(&pattern.StockSubtractHandler{}).
	SetNext(&pattern.CartDelHandler{}).
	SetNext(&pattern.DBTableOrderHandler{}).
	SetNext(&pattern.DBTableOrderSkusHandler{}).
	SetNext(&pattern.DBTableOrderPromotionsHandler{})

	//开始执行业务
	if err:=nullHandler.Run(&pattern.Context{});err!=nil{
		//处理错误
		fmt.Println("Fail | Error",err.Error())
		return
	}
	fmt.Println("Success | 订单创建成功")
}