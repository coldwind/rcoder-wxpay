package wxpay

import (
	"encoding/xml"
)

/**************统一下单**************
<xml>
	<appid>wx2421b1c4370ec43b</appid>
	<attach>支付测试</attach>
	<body>APP支付测试</body>
	<mch_id>10000100</mch_id>
	<nonce_str>1add1a30ac87aa2db72f57a2375d8fec</nonce_str>
	<notify_url>http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php</notify_url>
	<out_trade_no>1415659990</out_trade_no>
	<spbill_create_ip>14.23.150.211</spbill_create_ip>
	<total_fee>1</total_fee>
	<trade_type>APP</trade_type>
	<sign>0CB01533B8C1EF103065174F50BCA001</sign>
</xml>
************************************/
type UnifiedOrderStruct struct {
	XMLName xml.Name `xml:"xml"`
	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr string `xml:"nonce_str"`
	NotifyUrl string `xml:"notify_url"`
	OutTradeNo string `xml:"out_trade_no"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	TotalFee int `xml:"total_fee"`
	TradeType string `xml:"trade_type"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type"`
	Body string `xml:"body"`
}

/**************下单返回**************
<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[APP]]></trade_type>
</xml>
************************************/
type ResUnifiedOrderStruct struct {
	XMLName xml.Name `xml:"xml"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	DeviceInfo string `xml:"device_info"`
	Sign string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId string `xml:"prepay_id"`
	TradeType string `xml:"trade_type"`
	ErrCode string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

type ResPayNoticeStruct struct {
	XMLName xml.Name `xml:"xml"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	DeviceInfo string `xml:"device_info"`
	Sign string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId string `xml:"prepay_id"`
	TradeType string `xml:"trade_type"`
	ErrCode string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	OpenId string `xml:"openid"`
	IsSubscribe string `xml:"is_subscribe"`
	BankType string `xml:"bank_type"`
	TotalFee int `xml:"total_fee"`
	FeeType string `xml:"fee_type"`
	CashFee string `xml:"int cash_fee"`
	CashFeeType string `xml:"cash_fee_type"`
	CouponFee int `xml:"coupon_fee"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo string `xml:"out_trade_no"`
	Attach string `xml:"attach"`
	TimeEnd string `xml:"time_end"`
}