package wxpay

import (
	"math/rand"
	"fmt"
	"time"
	"crypto/md5"
	"strconv"
	"errors"
	"sort"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
)

func getNonceStr() string {
	nowTime := time.Now().Unix()
	rand.Seed(nowTime)
	randNum := rand.Intn(2000000000) + 20000000
	randString := fmt.Sprintf("%X-%d", randNum, nowTime)

	hasMd5 := md5.Sum([]byte(randString))
	md5String := fmt.Sprintf("%X", hasMd5)

	return md5String
}

func getSign(param map[string]interface{}) (string, error) {
	keys := make([]string, 0, 8)
	kvSet := make(map[string]string)
	for k, v := range param {
		keys = append(keys, k)
		switch v.(type) {
		case string:
			kvSet[k] = v.(string)
		case int:
			kvSet[k] = strconv.Itoa(v.(int))
		case int64:
			kvSet[k] = strconv.FormatInt(v.(int64), 10)
		case float32:
			kvSet[k] = strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32)
		case float64:
			kvSet[k] = strconv.FormatFloat(v.(float64), 'E', -1, 32)
		default:
			return "", errors.New("error: the key must be string|int|int64|float32|float64 (" + k + ")")
		}
	}

	// 排序KEY
	sort.Strings(keys)

	// 组成字串
	paramSlice := make([]string, 0, 16)

	for _, key := range keys {
		if _, ok := kvSet[key]; ok {
			continue
		}
		paramSlice = append(paramSlice, key + "=" + kvSet[key])
	}

	// import
	paramString := strings.Join(paramSlice, "&")
	paramString = paramString + "&key=" + APPKEY

	var cryptoData string
	if SIGN_TYPE == "MD5" {
		hasMd5 := md5.Sum([]byte(paramString))
		cryptoData = fmt.Sprintf("%X", hasMd5)
	} else {
		key := []byte(APPKEY)
		h := hmac.New(sha256.New, key)
		h.Write([]byte(paramString))
		cryptoData = fmt.Sprintf("%X", h.Sum(nil))
	}

	return cryptoData, nil
}

func request(api string, param map[string]interface{}, method string) ([]byte, error) {

	data := make(url.Values)

	for k, v := range param {
		switch v.(type) {
		case string:
			data[k] = []string{v.(string)}
		case int:
			data[k] = []string{strconv.Itoa(v.(int))}
		case int64:
			data[k] = []string{strconv.FormatInt(v.(int64), 10)}
		case float32:
			data[k] = []string{strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32)}
		case float64:
			data[k] = []string{strconv.FormatFloat(v.(float64), 'E', -1, 32)}
		default:
			return []byte{}, errors.New("error: the key must be string|int|int64|float32|float64 (" + k + ")")
		}
	}

	var request *http.Response
	var err error
	if method == "POST" {
		request, err = http.PostForm(api, data)
		if err != nil {
			return []byte{}, err
		}
	} else {
		getParam := make([]string, 0, 16)
		for k, v := range data {
			getParam = append(getParam, k + "=" + v[0])
		}

		paramString := strings.Join(getParam, "&")
		request, err = http.Get(api + "?" + paramString)
	}
	defer request.Body.Close()

	reqByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return []byte{}, err
	}

	return reqByte, nil
}

func streamRequest(api string, data []byte) ([]byte, error) {
	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", api, body)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	
	reqByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return reqByte, nil
}