package pattern

import "fmt"
type PayReq struct {
	OrderID string//订单号
}
type APayReq struct {
	PayReq
}
func (p *APayReq) Pay() string {
	fmt.Printf("支付成功-APayReq,%v",p.OrderID)
	return "支付成功"
}
type BPayReq struct {
	PayReq
	Uid int64 //用户ID
}
func (p *BPayReq) Pay() string {
	fmt.Printf("支付成功-BPayReq,%v",p.OrderID)
	return "支付成功"
}
