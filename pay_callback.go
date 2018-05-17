package wxpay

import (
	"encoding/xml"
	"errors"
)

func (this *Wxpay) payCallback(data []byte) (*ResPayNoticeStruct, error) {
	notice := new(ResPayNoticeStruct)

	err := xml.Unmarshal(data, notice)

	if err != nil {
		return nil, err
	}

	// 校验通知真实性
	param := make(map[string]interface{})

	sign, err := getSign(param)
	if err != nil || sign != notice.Sign {
		return nil, errors.New("error:invalid data.")
	}

	return notice, nil
}