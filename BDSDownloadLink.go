/*
 * @Author: NyanCatda
 * @Date: 2021-12-28 20:54:59
 * @LastEditTime: 2022-07-30 23:51:36
 * @LastEditors: NyanCatda
 * @Description: 获取BDS最新版本下载链接和版本号
 * @FilePath: \Go-BDS-Download-Link\BDSDownloadLink.go
 */
package BDSDownloadLink

import (
	"bytes"
	"errors"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/nyancatda/HttpRequest"
)

/**
 * @description: 获取Windows版本下载链接
 * @param {*}
 * @return {string} 下载链接
 * @return {string} 版本号
 * @return {error}
 */
func GetWindows() (string, string, error) {
	Body, err := GetBDSDownloadPage()
	if err != nil {
		return "", "", err
	}

	//解析数据
	doc, err := htmlquery.Parse(bytes.NewReader(Body))
	if err != nil {
		return "", "", err
	}
	ALabel := htmlquery.Find(doc, `//*[@id="main-content"]/div/div/div[1]/div/div/div/div[2]/div/div/div/div[1]/div[3]/div/a/@href`)

	DownloadLink := htmlquery.InnerText(ALabel[0])

	//从文件名提取版本号
	_, fileName := filepath.Split(DownloadLink)
	fileExt := path.Ext(DownloadLink)
	fileName = strings.TrimRight(fileName, fileExt)
	FileNameSplit := strings.Split(fileName, "-")
	Version := FileNameSplit[2]

	return DownloadLink, Version, err
}

/**
 * @description: 获取Ubuntu版本下载链接
 * @param {*}
 * @return {string} 下载链接
 * @return {string} 版本号
 * @return {error}
 */
func GetUbuntu() (string, string, error) {
	Body, err := GetBDSDownloadPage()
	if err != nil {
		return "", "", err
	}

	//解析数据
	doc, err := htmlquery.Parse(bytes.NewReader(Body))
	if err != nil {
		return "", "", err
	}
	ALabel := htmlquery.Find(doc, `//*[@id="main-content"]/div/div/div[1]/div/div/div/div[1]/div[2]/div/div/div[2]/div[3]/div/a/@href`)

	DownloadLink := htmlquery.InnerText(ALabel[0])

	//从文件名提取版本号
	_, fileName := filepath.Split(DownloadLink)
	fileExt := path.Ext(DownloadLink)
	fileName = strings.TrimRight(fileName, fileExt)
	FileNameSplit := strings.Split(fileName, "-")
	Version := FileNameSplit[2]

	return DownloadLink, Version, err
}

/**
 * @description: 获取BDS下载页面内容
 * @return {*}
 */
func GetBDSDownloadPage() ([]byte, error) {
	// 伪造请求头
	Header := []string{
		"accept:text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"accept-encoding:deflate, br",
		"accept-language:zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"cache-control:max-age=0",
		"sec-ch-ua:'Not A;Brand';v='99', 'Chromium';v='96', 'Microsoft Edge';v='96'",
		"sec-ch-ua-mobile:?0",
		"sec-ch-ua-platform:'Windows'",
		"sec-fetch-dest:document",
		"sec-fetch-mode:navigate",
		"sec-fetch-site:none",
		"sec-fetch-user:?1",
		"upgrade-insecure-requests:1",
		"user-agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.62",
		`cookie:MicrosoftApplicationsTelemetryDeviceId=898ff9b2-df7e-481a-b645-dd92ec30a397; MSFPC=GUID=8295de03d4fd4632b21597e2e5b09f5b&HASH=8295&LV=202110&V=4&LU=1634923645680; _cs_c=0; WRUID=3583559246169403; _cs_id=f4ba85c0-90e1-aaf2-f9f8-85f1d64e535a.1637476669.15.1643894322.1643894322.1613561419.1671640669730; __CT_Data=gpv=19&ckp=tld&dm=minecraft.net&apv_1067_www32=19&cpv_1067_www32=19&rpv_1067_www32=18; MSGCC=granted; ApplicationGatewayAffinityCORS=1bfc026d9c4b7d17a636076dd33a8622; ApplicationGatewayAffinity=1bfc026d9c4b7d17a636076dd33a8622; AKA_A2=A; ak_bmsc=92FDF4B17DEC87715CF07F4B271F87C1~000000000000000000000000000000~YAAQlbHXF6Z1/TaCAQAAE7a9TxARFkKRf8XfxaxhxzVgl/tb4KeER5UW83B6UCwBtGWyRYCcAxjoSjagtmLyR7Xhy0oF1E5ZmRUp5VuL/zABULgJvWERftCRgrJunsBmLFfBiiOZAd8oJTl+3IW0V5FhUau3AGgNDxb/rgn33Mu8TItsK3Fe2Dn3t6dsLU+hZbDXcKT8i6XBonNqUbaO6SvXfcmLjq8+1ShEhyBqpakeOSMJqqRzob8yhjmH5tGkujcMT2/DJ3FWvftq8L71mAM3XYdY5MAEpyCbIa1xeKfsdC8gx0621xCeqRJ6brlX1DQNAChncRBBXlQWjcwAulFDD381b2a3dx6eOmjXSuoNdR/Xgipaq3fWYpvQFisxwJeCKGu6flRyQSuwfLI=; MSCC=NR; at_check=true; AMCVS_EA76ADE95776D2EC7F000101%40AdobeOrg=1; AMCV_EA76ADE95776D2EC7F000101%40AdobeOrg=1176715910%7CMCIDTS%7C19204%7CMCMID%7C18600876857001525325005227394747468072%7CMCAID%7CNONE%7CMCOPTOUT-1659202413s%7CNONE%7CvVersion%7C5.4.0%7CMCAAMLH-1659800013%7C11%7CMCAAMB-1659800013%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CMCCIDH%7C-698682135%7CMCSYNCSOP%7C411-19211; bm_sv=3C661B85B4D57B841D508FC87613AD3A~YAAQpLHXF7SlwzWCAQAA32LNTxCk2r1HfxibnDWDLyD8Iyz/0P1gIb1l3qUcIy2EN3ZQXnSt9LhU2+fVM/UUXCXgi2UFLYC/EHBSKhIr/Qm7S6TNy+4yNRurx9Cfj1aUF32YgVVpNz6yl6ATGzfpRiZAnxOq+dMKDV250NHA3Yju7dD5Y3/CE16/e4m+DdKN3g+1gddXcie1ZMIj8WWXVal2vC1erx/Ur01vztEqbEZAu6SBvGbxt+atw741/HOZ3zJ3~1; ai_session=hXdJuw8yuevDHJEcWRHmEv|1659195210626|1659196235453; mbox=PC#e1c90446d1cd4016a495e650818f051e.32_0#1693382945|session#8aa5e31f41ab47be80e41e47aef39985#1659198107; RT="z=1&dm=minecraft.net&si=81d11951-9dfc-49d3-83dd-c2bfb2b7ef0d&ss=l681zj49&sl=2&tt=dd3&bcn=%2F%2F684d0d45.akstat.io%2F&ld=m9le&ul=ma8w"`,
	}

	// 发送请求
	Body, HttpResponse, err := HttpRequest.GetRequest("https://www.minecraft.net/en-us/download/server/bedrock", Header)

	if HttpResponse.StatusCode != http.StatusOK {
		return nil, errors.New("HttpRequest Error" + strconv.Itoa(HttpResponse.StatusCode) + " " + HttpResponse.Status)
	}

	return Body, err
}
