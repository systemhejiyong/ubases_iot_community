package sms

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type HuyiClient struct {
	appId  string
	appKey string
}

func GetHuyiClient(appId string, appKey string) (*HuyiClient, error) {
	return &HuyiClient{
		appId:  appId,
		appKey: appKey,
	}, nil
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (hc *HuyiClient) SendMessage(template string, param map[string]string, targetPhoneNumber ...string) error {
	code, ok := param["code"]
	if !ok {
		return fmt.Errorf("missing parameter: msg code")
	}

	if len(targetPhoneNumber) < 1 {
		return fmt.Errorf("missin parer: trgetPhoneNumber")
	}

	_now := strconv.FormatInt(time.Now().Unix(), 10)
	smsContent := fmt.Sprintf(template, code)
	v := url.Values{}
	v.Set("account", hc.appId)
	v.Set("content", smsContent)
	v.Set("time", _now)
	passwordStr := hc.appId + hc.appKey + "%s" + smsContent + _now
	for _, mobile := range targetPhoneNumber {
		password := fmt.Sprintf(passwordStr, mobile)
		v.Set("password", GetMd5String(password))
		v.Set("mobile", mobile)

		body := strings.NewReader(v.Encode()) //encode form data
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
