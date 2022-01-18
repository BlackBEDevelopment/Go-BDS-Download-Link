//获取BDS最新版本下载链接和版本号
package BDSDownloadLink

import (
	"compress/gzip"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/antchfx/htmlquery"
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
	defer Body.Close()

	//解析数据
	doc, err := htmlquery.Parse(Body)
	if err != nil {
		return "", "", err
	}
	ALabel := htmlquery.Find(doc, `//*[@id="main-content"]/div/div/div/div/div/div/div[1]/div[2]/div/div/div[1]/div[3]/div/a/@href`)

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
	defer Body.Close()

	//解析数据
	doc, err := htmlquery.Parse(Body)
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

//获取BDS下载页面内容
func GetBDSDownloadPage() (io.ReadCloser, error) {
	var Body io.ReadCloser

	req, err := http.NewRequest("GET", "https://www.minecraft.net/en-us/download/server/bedrock", nil)
	if err != nil {
		return Body, err
	}

	//伪造请求头
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-ch-ua", "'Not A;Brand';v='99', 'Chromium';v='96', 'Microsoft Edge';v='96'")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "'Windows'")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.62")
	req.Header.Set("cookie", "MSGCC=granted; MicrosoftApplicationsTelemetryDeviceId=898ff9b2-df7e-481a-b645-dd92ec30a397; MSFPC=GUID=8295de03d4fd4632b21597e2e5b09f5b&HASH=8295&LV=202110&V=4&LU=1634923645680; _cs_c=0; _CT_RS_=Recording; WRUID=3583559246169403; ApplicationGatewayAffinityCORS=1bfc026d9c4b7d17a636076dd33a8622; ApplicationGatewayAffinity=1bfc026d9c4b7d17a636076dd33a8622; AKA_A2=A; ak_bmsc=B28E14B5C090F869CA01770B94A225BB~000000000000000000000000000000~YAAQdYEyF2wCNyx9AQAAtFtK/A4wb6pCrMZpWwVwC2Tirg7wrBRXJK567Df94nK9WZ9opQaLhNbFhtvyu9DtgQWF/vMVCAtmZUXiuykiHPs7rLzN8WPD1gRUscuFWTyu3yRT/83wKRguF8JyGPspE4XbhfdrDJBKDZ02CLlxpQn1dTM6vUw1s41kHKd2rCsVsQb5W1sg9jmIEkB7Mkm7bP9LgRwJ3TG6Bsvy3F3KFAuzzlw+VUSMSYhj8ED3QC+icnYiRiz3NqPGNmsHnP3zTIBA8F3MJS6FnuoKthWJPNLY9AYQ1xL2Y/Jcpk3UxsmbC5orHmo5Hts4acV60nPqlHw5Ns8cc0F9iabUE2m/99WhdfEm4qbqS3yg/ncSzMWUWO3AFeEEfo/eLTXFjgM=; MSCC=NR; _cs_cvars=%7B%7D; _cs_id=f4ba85c0-90e1-aaf2-f9f8-85f1d64e535a.1637476669.7.1640615279.1640615279.1613561419.1671640669730; _cs_s=1.1.0.1640617079675; __CT_Data=gpv=4&ckp=tld&dm=minecraft.net&apv_1067_www32=4&cpv_1067_www32=4&rpv_1067_www32=4; ai_session=30UVW2j3scggQM+p4yAbxn|1640615276066|1640615526043; bm_sv=F93FAA2EEE6B83B0B880205CFFCE26B5~fEoe2z/JwmZ0dH9LjWggry4SPgNUoJDX0mIEiY9mf0AtRuAKVlPQU9b3XMPYPnU4Pu4Q9iz1pNSK2SgR4uDQDLIOGbPyp6q1GUW1hFd/AwLlYkUHmR2GgMAffDyNHHxnzPe4QFLKBzw/Vd2lOAeNlcmcHlATuuWsEay1ML7Ohc0=; RT='z=1&dm=minecraft.net&si=c9c49742-58c9-41b0-8b50-0cc2b3742a0f&ss=kxorz0nf&sl=2&tt=z5v&bcn=%2F%2F684d0d46.akstat.io%2F&ld=z6l&ul=5ohj'")

	//发起请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Body, err
	}

	//解压数据
	Body, err = gzip.NewReader(res.Body)
	if err != nil {
		return Body, err
	}

	return Body, err
}
