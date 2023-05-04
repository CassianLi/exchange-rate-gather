package utils

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// GetUserAgent Generative random User-Agent
func GetUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.63 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
		"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; AS; rv:11.0) like Gecko",
	}
	return userAgents[rand.Intn(len(userAgents))]
}

// DownloadHtml Download html from url
func DownloadHtml(url string) (body string, err error) {
	// 构造HTTP GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept-Language", "en-US")
	// 设置请求头中的User-Agent，模拟真实的浏览器访问
	req.Header.Set("User-Agent", GetUserAgent())

	// 发送HTTP请求，并获取响应
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 从响应中获取HTML源代码
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(htmlBytes), nil
}

// CheckElementExist check element exists
func CheckElementExist(ctx context.Context, sel string) bool {

	// 设置超时时间为10秒
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 创建一个变量存储元素是否存在的结果
	var res bool

	// 访问页面并等待元素加载
	err := chromedp.Run(ctx, chromedp.WaitVisible(sel))
	if err != nil {
		fmt.Println("Error waiting for element:", err)
		return false
	}

	// 在页面上查找元素
	err = chromedp.Run(ctx, chromedp.Evaluate(fmt.Sprintf(`document.querySelector("%s") !== null`, sel), &res))
	if err != nil {
		fmt.Println("Error checking element:", err)
		return false
	}

	return res
}

// ClickElement element click
func ClickElement(ctx context.Context, sel string) error {
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(sel),
		chromedp.Click(sel),
	)
	if err != nil {
		return fmt.Errorf("could not click on element: %v", err)
	}
	return nil
}

// ElementExist chromedp check element exists
func ElementExist(ctx context.Context, sel string) bool {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var res bool
	err := chromedp.Run(ctx, chromedp.Evaluate(fmt.Sprintf(`document.querySelector("%s") !== null`, sel), &res))
	if err != nil {
		return false
	}
	return res
}
