package pattern
import "testing"
func TestPay(t *testing.T) {
	var pay IPay
	pay = &APayReq{
		PayReq: PayReq{
			OrderID: "123456",
			Uid:     1,
		},
	}
	if pay.Pay() != "支付成功" {
		t.Errorf("期望支付成功，实际支付失败")
	}
	pay = &BPayReq{
		PayReq: PayReq{
			OrderID: "654321",
			Uid:     2,
		},
	}
	if pay.Pay() != "支付成功" {
		t.Errorf("期望支付成功，实际支付失败")
	}
}