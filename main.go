package amootsms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type Api struct {
	Username string
	Password string
}

const baseUrl string = "https://portal.SMS.com/webservice2.asmx/"

func _init(Endpoint string, FormData map[string]interface{}) (string, string) {
	var params string = ""
	var _EndPoint string = Endpoint
	_EndPoint = strings.Replace(_EndPoint, "_REST", "", -1)
	_EndPoint = strings.Replace(_EndPoint, "/", "", -1)
	_EndPoint = _EndPoint + "_REST"
	for key, element := range FormData {
		params += fmt.Sprintf("&%s=%s", key, element)
	}
	return _EndPoint, params
}

func (receiver Api) Call(Endpoint string, FormData map[string]interface{}) map[string]interface{} {
	_EndPoint, params := _init(Endpoint, FormData)
	data := receiver.makeCall(_EndPoint, params)
	return data
}

func (receiver Api) makeCall(_EndPoint string, params string) map[string]interface{} {
	url := fmt.Sprintf("%s%s?UserName=%s&Password=%s", baseUrl, _EndPoint, receiver.Username, receiver.Password)
	url = url + params
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("ss", "sss")
	err := writer.Close()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Fatal(err)
	}
	return dat
}

func (receiver Api) AccountStatus() map[string]interface{} {
	data := receiver.Call("AccountStatus", map[string]interface{}{})
	return data
}

func (receiver Api) SendSimple(Message string, lineNumber string, SendDateTime string, Mobiles []string) map[string]interface{} {
	if lineNumber == "" {
		lineNumber = "public"
	}
	if SendDateTime == "" {
		SendDateTime = time.Now().Format(time.RFC3339)
	}
	_mobiles := strings.Join(Mobiles, ",")
	data := receiver.Call("SendSimple", map[string]interface{}{
		"SendDateTime":   SendDateTime,
		"SMSMessageText": Message,
		"LineNumber":     lineNumber,
		"Mobiles":        _mobiles,
	})
	return data
}

func (receiver Api) SendWithPattern(PatternCodeID string, PatternValues string, Mobile string) map[string]interface{} {
	data := receiver.Call("SendWithPattern", map[string]interface{}{
		"PatternCodeID": PatternCodeID,
		"PatternValues": PatternValues,
		"Mobile":        Mobile,
	})
	return data
}

func (receiver Api) SendQuickOTP(CodeLength string, OptionalCode string, Mobile string) map[string]interface{} {
	data := receiver.Call("SendQuickOTP", map[string]interface{}{
		"CodeLength":   CodeLength,
		"OptionalCode": OptionalCode,
		"Mobile":       Mobile,
	})
	return data
}
