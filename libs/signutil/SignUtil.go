package signutil
import (
	"net/url"
	"hulk_salthttp/libs"
	"github.com/astaxie/beego"
	"strings"
	"sort"
	"hulk_salthttp/libs/logger"
)

/**
 * 生成签名
 * @param vals必须包含有app_key字段, 因为需要用app_key获取对应的公钥
 */
func GenSignatureByValues(vals url.Values) (string, error) {
	// 验证app_key
	if vals.Get("app_key") == "" {
		return "", libs.ERR_SIGN_APP_KEY
	}

	appkey := vals.Get("app_key")
	keyMaps := getAppKeys()
	if _, ok := keyMaps[appkey]; !ok {
		return "", libs.ERR_SIGN_NO_SECRET_KEY
	}

	// 生成签名
	// 1. 将所有参数按字段升序排序所有值
	fields := make([]string, 0)
	for k, _ := range vals {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	// 2. php: http_build_query()的效果
	rawStr := ""
	for _, key := range fields {
		rawStr += url.QueryEscape(key) + "=" + url.QueryEscape(vals.Get(key)) + "&"
	}
	rawStr = strings.TrimRight(rawStr, "&")

	// 3. md5(md5(rawStr) + public_key)
	signStr := libs.Md5Str(libs.Md5Str(rawStr) + keyMaps[appkey])

	return signStr, nil
}

/**
 * 解析配置文件, 获取已配置的app_key对应的public_key的Map
 */
func getAppKeys() map[string]string {
	keys := beego.AppConfig.String("sign::AppKeys")
	arrKeys := strings.Split(keys, ";")
	keyMaps := map[string]string{}
	for _, item := range arrKeys {
		arr := strings.Split(item, "|")
		if len(arr) < 2 {
			continue
		}

		keyMaps[arr[0]] = arr[1]
	}

	return keyMaps
}

/**
 * 验证签名
 * @param vals
 * 备注: vals中必须包含有: app_key , sign 字段
 */
func CheckSignature(vals url.Values) (bool, error) {
	if vals.Get("sign") == "" {
		return false, libs.ERR_SIGN
	}

	genSign, err := GenSignatureByValues(vals)
	if err != nil {
		return false, err
	}

	if vals.Get("sign") != genSign {
		logger.Error("checksigniture", map[string]string{
			"req_sign": vals.Get("sign"),
			"gen_sign": genSign,
		})
		return false, libs.ERR_SIGN
	}

	return true, nil
}
