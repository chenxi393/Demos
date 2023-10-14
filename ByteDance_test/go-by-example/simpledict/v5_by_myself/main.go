package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type Dictionary_Response_caiyun struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query_with_caiyunai(word string, wg *sync.WaitGroup) {
	client := &http.Client{}
	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	// 返回的是 buf []byte 需要
	var data = bytes.NewReader(buf)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "36d096509ce7142cf4652a2c0c78c720")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad Status Code: ", resp.StatusCode, "body ", string(bodyText))
	}

	var dic_Respons Dictionary_Response_caiyun
	err = json.Unmarshal(bodyText, &dic_Respons)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dic_Respons)
	fmt.Println(word, "UK:", dic_Respons.Dictionary.Prons.En,
		"US:", dic_Respons.Dictionary.Prons.EnUs)
	for _, item := range dic_Respons.Dictionary.Explanations {
		fmt.Println(item)
	}

	wg.Done()
}

type dict_response_with_youdao struct {
	VideoSents struct {
		WordInfo struct {
			Sense []string `json:"sense"`
		} `json:"word_info"`
	} `json:"video_sents"`
	Simple struct {
		Word []struct {
			Usphone string `json:"usphone"`
			Ukphone string `json:"ukphone"`
		} `json:"word"`
	} `json:"simple"`
}

type Url_Params struct {
	Q       string `url:"q"`
	Le      string `url:"le"`
	T       int    `url:"t"`
	Client  string `url:"client"`
	Sign    string `url:"sign"`
	KeyFrom string `url:"keyfrom"`
}

func query_with_youdao(word string, wg *sync.WaitGroup) {
	client := &http.Client{}
	Request := Url_Params{
		Q:       word,
		Le:      "en",
		T:       2,
		Client:  "web",
		Sign:    "3c71569a04e3231adce6ef811c67148a",
		KeyFrom: "webdict",
	}

	// gpt给的 将结构体变成url参数
	// 使用 url.Values 类型来保存转换后的 URL 参数
	values := url.Values{}

	// 将字段按照指定顺序添加到 url.Values 中
	values.Add("q", Request.Q)
	values.Add("le", Request.Le)
	values.Add("t", fmt.Sprintf("%d", Request.T))
	values.Add("client", Request.Client)
	values.Add("sign", Request.Sign)
	values.Add("keyfrom", Request.KeyFrom)

	// 将 url.Values 转换为 URL 编码字符串
	encoded := values.Encode()
	//fmt.Println(encoded) // 输出: "client=web&keyfrom=webdict&le=en&q=row&sign=3c71569a04e3231adce6ef811c67148a&t=2"
	//这个顺序和原来的顺序不一样好像也没关系

	var data = strings.NewReader(encoded)
	req, err := http.NewRequest("POST", "https://dict.youdao.com/jsonapi_s?doctype=json&jsonversion=4", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID_NCOO=1644065916.859692; OUTFOX_SEARCH_USER_ID=-1797365482@58.20.162.176; __yadk_uid=MDP29jp0W6jdj4ucEBFYlv1Uqnm1Zw02; rollNum=true; advertiseCookie=advertiseValue; ___rl__test__cookies=1691264812219")
	req.Header.Set("Origin", "https://dict.youdao.com")
	req.Header.Set("Referer", "https://dict.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var dict_re dict_response_with_youdao
	err = json.Unmarshal(bodyText, &dict_re)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", dict_re)
	fmt.Println(word, "UK:", "["+dict_re.Simple.Word[0].Ukphone+"]",
		"US:", "["+dict_re.Simple.Word[0].Usphone+"]")
	for _, item := range dict_re.VideoSents.WordInfo.Sense {
		fmt.Println(item)
	}

	wg.Done()
}

func main() {
	//	使用原始字符串字面量``时，字符串中的特殊字符（如转义字符）将被原样保留，
	//不会进行转义。这对于需要包含大量特殊字符的字符串非常有用，因为不需要手动进行转义。
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDIct WORD
example: simpleDict hello
`)
		os.Exit(1)
	}
	word := os.Args[1]
	var cnt sync.WaitGroup
	cnt.Add(2)
	go query_with_caiyunai(word, &cnt)
	go query_with_youdao(word, &cnt)
	cnt.Wait()
}
