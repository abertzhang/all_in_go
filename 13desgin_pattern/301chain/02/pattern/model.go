package pattern

import (
	"fmt"
	"runtime"
)
type Context struct {}
//实现IHandler接口
type Next struct {
	nextHandler IHandler
}
func (n *Next) SetNext(next IHandler) IHandler {
	n.nextHandler = next
	return next
}
func (n *Next) Run(c *Context) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(c); err != nil {
			return err
		}
		return (n.nextHandler).Run(c)
	}
	return nil
}
//
type NullHandler struct {
	Next
}
func (h *NullHandler) Do(c *Context) error {
	return nil
}
//
type ArgumentHandler struct {
	//合成复用Next
	Next
}
func (h *ArgumentHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "校验参数成功...")
	return nil
}
//
type AddressInfoHandler struct {
	//合成复用Next
	Next
}
func (h *AddressInfoHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取地址信息...")
	fmt.Println(runFuncName(), "地址信息校验...")
	return nil
}
//
type CartInfoHandler struct {
	//合成复用Next
	Next
}
func (h *CartInfoHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取购物车数据...")
	return nil
}
//
type StockInfoHandler struct {
	//合成复用Next
	Next
}
func (h *StockInfoHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取商品库存信息...")
	fmt.Println(runFuncName(), "库存信息校验...")
	return nil
}
//
type PromotionInfoHandler struct {
	//合成复用Next
	Next
}
func (h *PromotionInfoHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取促销信息...")
	fmt.Println(runFuncName(), "促销信息校验...")
	return nil
}
//
type ShipmentInfoHandler struct {
	//合成复用Next
	Next
}
func (h *ShipmentInfoHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取物流信息...")
	return nil
}
//
type PromotionUseHandler struct {
	//合成复用Next
	Next
}
func (h *PromotionUseHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "获取优惠信息...")
	return nil
}

//
type StockSubtractHandler struct {
	//合成复用Next
	Next
}
func (h *StockSubtractHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "扣减库存...")
	return nil
}
//
type CartDelHandler struct {
	Next
}
func (h *CartDelHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "删除购物车商品...")
	return nil
}
//
type DBTableOrderHandler struct {
	Next
}
func (h *DBTableOrderHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "写订单表...")
	return nil
}

//
type DBTableOrderSkusHandler struct {
	Next
}
func (h *DBTableOrderSkusHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "写订单SKU表...")
	return nil
}

type DBTableOrderPromotionsHandler struct {
	Next
}
func (h *DBTableOrderPromotionsHandler) Do(c *Context) error {
	fmt.Println(runFuncName(), "写订单优惠信息表...")
	return nil
}

func runFuncName()string {
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
