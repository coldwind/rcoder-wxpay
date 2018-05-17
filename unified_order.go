package wxpay

import (
	"encoding/xml"
)

func (this *Wxpay) unifiedOrder(param map[string]interface{}) (*ResUnifiedOrderStruct, error) {
	api := "https://api.mch.weixin.qq.com/pay/unifiedorder"

	param["appid"] = APPID
	param["mch_id"] = MCHID
	param["notify_url"] = NOTIFY_URL
	param["nonce_str"] = getNonceStr()

	// 最后生成sign
	sign, err := getSign(param)
	if err != nil {
		return nil, err
	}
	param["sign"] = sign

	unifiedData := new(UnifiedOrderStruct)
	unifiedData.AppId = APPID
	unifiedData.MchId = MCHID

	if _, ok := param["device_info"]; ok {
		unifiedData.DeviceInfo = param["device_info"].(string)
	}
	
	if _, ok := param["nonce_str"]; ok {
		unifiedData.NonceStr = param["nonce_str"].(string)
	}

	unifiedData.Sign = sign
	unifiedData.SignType = SIGN_TYPE

	if _, ok := param["body"]; ok {
		unifiedData.Body = param["body"].(string)
	}

	unifiedData.NotifyUrl = NOTIFY_URL

	if _, ok := param["out_trade_no"]; ok {
		unifiedData.OutTradeNo = param["out_trade_no"].(string)
	}

	if _, ok := param["SpbillCreateIp"]; ok {
		unifiedData.NotifyUrl = param["spbill_create_ip"].(string)
	}

	if _, ok := param["total_fee"]; ok {
		unifiedData.TotalFee = param["total_fee"].(int)
	}

	if _, ok := param["trade_type"]; ok {
		unifiedData.TradeType = param["trade_type"].(string)
	}

	data, err := xml.Marshal(unifiedData)
	if err != nil {
		return nil, err
	}

	res, err := streamRequest(api, data)
	if err != nil {
		return nil, err
	}

	resUnified := new(ResUnifiedOrderStruct)
	err = xml.Unmarshal(res, resUnified)

	if err != nil {
		return nil, err
	}

	return resUnified, nil
}