package curl

import (
	"errors"
	"github.com/mikemintang/go-curl"
	"log"
)

/**
 *  get请求
 *  @param url string  请求地址
 *  @param queries map[string]string 请求参数
 *  @return string，error
 */
func HttpGet(url string, queries map[string]string) (string,error) {
	headers := map[string]string{
		"Content-Type":  "application/json",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetQueries(queries).
		Get()
	if err != nil {
		return "" ,err
	} else {
		if resp.IsOk() {
			return resp.Body,nil
		} else {
			log.Printf("%v\n",resp.Raw)
			return "" ,errors.New("请求失败")
		}
	}
}

/**
 *  get请求
 *  @param url string  请求地址
 *  @param queries map[string]string 请求参数
 *  @return string，error
 */
func HttpPost(url string, queries map[string]string,postData map[string]interface{}) (string,error) {
	headers := map[string]string{
		//"User-Agent":    "Sublime",
		//"Authorization": "Bearer access_token",
		"Content-Type":  "application/json",
	}

	cookies := map[string]string{
		//"userId":    "12",
		//"loginTime": "15045682199",
	}

	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetCookies(cookies).
		SetQueries(queries).
		SetPostData(postData).
		Post()

	if err != nil {
		return "" ,err
	} else {
		if resp.IsOk() {
			return resp.Body,nil
		} else {
			log.Printf("%v\n",resp.Raw)
			return "" ,errors.New("请求失败")
		}
	}
}