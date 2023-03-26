package ups

import (
	"encoding/json"
	"errors"
	"go-logistics-tracking/consts"
	"go-logistics-tracking/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var token Token

var lock = sync.Mutex{}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: false,
}

var client = &http.Client{Transport: tr}

type Token struct {
	Token_st        string
	Token           string
	TokenExpireTime int64
}

type Response struct {
	StatusCode  string        `json:"statusCode"`
	StatusText  string        `json:"statusText"`
	TrackDetail []TrackDetail `json:"trackDetails"`
}
type TrackDetail struct {
	ErrorCode                  string   `json:"errorCode"`
	ErrorText                  string   `json:"errorText"`
	TrackingNumber             string   `json:"trackingNumber"`
	ShipmentProgressActivities []Events `json:"shipmentProgressActivities"`
}

type Events struct {
	Location     string `json:"location"`
	GmtDate      string `json:"gmtDate"`
	GmtOffset    string `json:"gmtOffset"`
	GmtTime      string `json:"gmtTime"`
	ActivityScan string `json:"activityScan"`
}

func GetToken() (t Token, e error) {
	if token.Token != "" && token.TokenExpireTime <= time.Now().UnixMilli() {
		return token, nil
	}
	lock.Lock()
	if token.Token != "" && token.TokenExpireTime <= time.Now().UnixMilli() {
		return token, e
	}
	response, err := client.Get("https://www.ups.com/track?loc=en_US")
	if err != nil {
		return token, err
	}
	defer func() {
		response.Body.Close()
		lock.Unlock()
	}()
	cookies := response.Header.Values("Set-Cookie")
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie, "X-CSRF-TOKEN") {
			cookie = strings.Split(cookie, ";")[0]
			csrfToken := cookie[strings.Index(cookie, "=")+1:]
			token.Token = csrfToken
		} else if strings.HasPrefix(cookie, "X-XSRF-TOKEN-ST") {
			cookie = strings.Split(cookie, ";")[0]
			stToken := cookie[strings.Index(cookie, "=")+1:]
			token.Token_st = stToken
		}
	}
	token.TokenExpireTime = time.Now().UnixMilli() + 10*60*1000
	return token, nil
}

// 抓取ups
func grab(logisticsNo string) ([]byte, error) {
	token, err := GetToken()
	if err != nil {
		return nil, err
	}
	params := make(map[string]any)
	params["Locale"] = "en_US"
	params["TrackingNumber"] = []string{logisticsNo}
	b, _ := json.Marshal(params)
	requestJson := string(b)
	request, _ := http.NewRequest("POST", "https://www.ups.com/track/api/Track/GetStatus?loc=en_US", strings.NewReader(requestJson))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-xsrf-token", token.Token_st)
	request.Header.Add("cookie", "X-CSRF-TOKEN="+token.Token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("请求错误,状态码" + strconv.Itoa(response.StatusCode))
	}
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		Body.Close()
	}(response.Body)
	return res, nil
}
func Track() func(request []model.TrackRequest) []model.TrackResponse {
	return execute
}

func execute(request []model.TrackRequest) []model.TrackResponse {
	resultList := make([]model.TrackResponse, len(request))
	for i, v := range request {
		res, err := grab(v.LogisticsNo)
		if err != nil {
			resultList[i] = model.BuildEmptyTrackResponse(v.LogisticsNo, err.Error(), consts.NO_RECORD)
			continue
		}
		response := new(Response)
		//str := string(res)
		err = json.Unmarshal(res, response)
		if err != nil {
			resultList[i] = model.BuildEmptyTrackResponse(v.LogisticsNo, err.Error(), consts.NO_RECORD)
			continue
		}
		if response.StatusCode != "200" {
			resultList[i] = model.BuildEmptyTrackResponse(v.LogisticsNo, response.StatusText, consts.NO_RECORD)
			continue
		}
		if len(response.TrackDetail) == 0 {
			resultList[i] = model.BuildEmptyTrackResponse(v.LogisticsNo, "", consts.NO_RECORD)
			continue
		}
		trackDetail := response.TrackDetail[0]
		if trackDetail.TrackingNumber != v.LogisticsNo {
			resultList[i] = model.BuildEmptyTrackResponse(v.LogisticsNo, "追踪单号与UPS接口返回单号不一致", consts.NO_RECORD)
			continue
		}
		resModel := model.BuildEmptyTrackResponse(v.LogisticsNo, "", consts.TRANSIT)
		events := make([]model.Event, len(trackDetail.ShipmentProgressActivities))

		for idx, v := range trackDetail.ShipmentProgressActivities {
			location, _ := time.LoadLocation("GMT")
			t, _ := time.ParseInLocation("2006010215:04:05", v.GmtDate+v.GmtTime, location)
			offsetStr := v.GmtOffset
			var offset int
			arr := strings.Split(offsetStr, ":")
			if len(arr) == 2 {
				hour, _ := strconv.Atoi(arr[0])
				minutes, _ := strconv.Atoi(arr[1])
				offset = hour*3600 + minutes*60
			}
			utcZone := time.FixedZone("UTC", offset)
			t = t.In(utcZone)
			eventTime := t.Format("2006-01-02 15:04:05")
			event := model.Event{Date: eventTime, Location: v.Location, Message: v.ActivityScan}
			events[idx] = event
		}
		resModel.EventListEn = events
		resModel.EventListZh = events
		resultList[i] = resModel
	}
	return resultList
}
