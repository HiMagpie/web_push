package httptool

import (
	"github.com/astaxie/beego/httplib"
	"time"
	"encoding/json"
	"net/url"
	"strings"
	"net/http"
	"hulk_salthttp/libs/logger"
)

/**
 * 通过proxy执行的post请求
 *
 * @param uri string 请求的url eg:http:zjl.qihoo.net:8360/test/index
 * @param data interface{} post的参数会放到data字段里 eg: map[string]string{"jid":"xxxx"}
 * @param proxy string 代理proxy eg: 10.16.57.219
 * @return string 返回的Body字串
 * @return error http请求发送的错误
 *
 */
func PostByProxy(uri string, data interface{}, proxy string) (string, error) {
	req := httplib.Post(uri)
	//设置超时时间
	req.SetTimeout(10 * time.Second, 30 * time.Second)
	dataStr, _ := json.Marshal(data)
	req.Param("data", string(dataStr))
	if len(proxy) > 0 {
		u, _ := url.ParseRequestURI(uri)
		logger.Debug("postbyproxy", u)
		//组装proxy的url
		proxy_url := u.Scheme + "://" + proxy
		host_port := strings.Split(u.Host, ":")
		if len(host_port) == 2 {
			proxy_url = proxy_url + ":" + host_port[1]
		}
		logger.Debug("postbyproxy", "proxy_url:", proxy_url)
		req.SetProxy(func(req *http.Request) (*url.URL, error) {
			u, _ := url.ParseRequestURI(proxy_url)
			logger.Debug("postbyproxy", "proxy_struct", u)
			return u, nil
		})
	}
	res, err := req.String()
	loginfo := map[string]interface{}{
		"error":err,
		"uri":uri,
		"proxy":proxy,
		"data":data,
		"res":res,
	}
	if err != nil {
		logger.Error("postbyproxy", loginfo)
	}else {
		logger.Info("postbyproxy", loginfo)
	}
	return res, err
}

