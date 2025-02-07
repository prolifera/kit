package xproxy

import (
	"github.com/prolifera/kit/log"
	"golang.org/x/net/proxy"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// proxyList 存储代理列表
var proxyList []string
var mu sync.Mutex

// InitProxyByInput 初始化代理列表
func InitProxyByInput(proxyStr string, proxys []string) {
	mu.Lock()
	defer mu.Unlock()

	if proxyStr != "" {
		proxyList = strings.Split(proxyStr, ",")
	} else {
		proxyList = proxys
	}

}

// GetSequentialProxyClient 获取顺序的 http.Client
func GetSequentialProxyClient(index *int) (*http.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if len(proxyList) == 0 {
		return &http.Client{}, nil
	}

	if *index >= len(proxyList) {
		*index = 0
	}

	p := proxyList[*index]
	*index = (*index + 1) % len(proxyList)
	return NewHTTPClient(p)
}

// GetRandomProxyClient 获取随机的 http.Client
func GetRandomProxyClient() (*http.Client, error) {
	mu.Lock()
	defer mu.Unlock()
	if len(proxyList) == 0 {
		return &http.Client{}, nil
	}
	p := proxyList[rand.Intn(len(proxyList))]
	return NewHTTPClient(p)
}

func GetClientByProxy(proxyUrl string) (*http.Client, error) {
	return NewHTTPClient(proxyUrl)
}

// newHTTPClient 根据 proxyUrl 返回一个 http.Client
func NewHTTPClient(proxyUrl string) (*http.Client, error) {

	u, err := url.Parse(proxyUrl)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "parse proxy error", "error": err.Error()}).Send()
		return nil, err
	}

	dialer, err := proxy.FromURL(u, proxy.Direct)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "parse proxy error", "error": err.Error()}).Send()
		return nil, err
	}
	transport := &http.Transport{
		Dial: dialer.Dial,
	}
	client := &http.Client{
		Transport: transport,
	}

	return client, nil
}
