package chatModel

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// PostOpenaiChat 向openai发送post请求
func PostOpenaiChat(params ChatParams, apikey string) (ChatCompletion, error) {
	// 将聊天对话参数转换为JSON格式的字节切片
	data, err := json.Marshal(params)
	if err != nil {
		return ChatCompletion{}, err
	}

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个HTTP请求对象，指定URL，方法和数据
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(data))
	if err != nil {
		return ChatCompletion{}, err
	}

	// 设置请求头部，指定内容类型和授权信息
	req.Header.Set("Content-Type", "application/json")
	OpenaiApiKey := "Bearer " + apikey
	req.Header.Set("Authorization", OpenaiApiKey)

	// 发送请求，并获取响应对象
	resp, err := client.Do(req)
	if err != nil {
		return ChatCompletion{}, err
	}

	// 延迟关闭响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// 读取响应体中的数据，并转换为字节切片
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatCompletion{}, err
	}

	// 定义一个聊天对话补全结果变量
	var completion ChatCompletion
	//fmt.Println(string(bodyText))
	// 将响应体中的数据解析为聊天对话补全结果结构体
	err = json.Unmarshal(bodyText, &completion)
	if err != nil {
		return ChatCompletion{}, err
	}

	// 返回聊天对话补全结果和错误
	return completion, nil
}
