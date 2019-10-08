package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"tessa/config"
)

//SendSMSLive247 ...
func SendSMSLive247(sender, recipient, message string) bool {
	if recipient == "" || message == "" || config.Get().Smslive247 == "" || sender == "" {
		return false
	}

	sendSMSURL := "http://www.smslive247.com/http/index.aspx?cmd=sendmsg&sessionid=%s&message=%s&sender=%s&sendto=%s&msgtype=0"
	sendSMSURL = fmt.Sprintf(sendSMSURL, config.Get().Smslive247, url.QueryEscape(message), sender, url.QueryEscape(recipient))
	httpReq, _ := http.NewRequest("GET", sendSMSURL, nil)
	client := &http.Client{}
	client.Do(httpReq)
	return true
}

//SendSMSTwilio ...
func SendSMSTwilio(sender, recipient, message string) bool {

	twilio := config.Get().Twilio
	if twilio["authtoken"] == "" || twilio["accountsid"] == "" || twilio["sender"] == "" {
		return false
	}

	if sender == "" {
		sender = twilio["sender"]
	}

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + twilio["accountsid"] + "/Messages.json"
	msgData := url.Values{}
	msgData.Set("To", recipient)
	msgData.Set("From", sender)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(twilio["accountsid"], twilio["authtoken"])
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client.Do(req)
	// resp, _ := client.Do(req)
	// var data map[string]interface{}
	// decoder := json.NewDecoder(resp.Body)
	// err := decoder.Decode(&data)
	// if err == nil {
	// 	fmt.Println(data)
	// }
	return true
}
