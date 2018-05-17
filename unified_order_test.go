package wxpay

import (
	"testing"
)

func Test_unifiedOrder(t *testing.T) {

	param := make(map[string]interface{})
	param["out_trade_no"] = "20091227091010"
	param["total_fee"] = 1
	param["spbill_create_ip"] = "127.0.0.1"
	
	wx := new(Wxpay)
	res, err := wx.unifiedOrder(param)
	t.Log(res, err)
}