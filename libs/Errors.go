package libs
import (
	"strconv"
	"strings"
)

//定义一个Errstr类型,用来做类型限制
type Errstr string

const (
	ERR_NONE = 0
	ERR_SEP = "=>"
	ERR_PARAM Errstr = "1=>参数错误"
	ERR_UNKNOWN Errstr = "2=>未知错误"
	ERR_SIGN Errstr = "3=>签名认证失败"
	ERR_SYSTEM Errstr = "4=>系统异常"
	ERR_OPERATE Errstr = "5=>操作失败"
	ERR_NO_RECORD Errstr = "11=>记录不存在"
)


//返回错误码
func (errstr Errstr ) GetErrno() int {
	s := strings.Split(string(errstr), ERR_SEP)
	errno, _ := strconv.Atoi(s[0])
	return errno
}


//返回错误信息
func (errstr Errstr ) GetErrmsg() string {
	return strings.Split(string(errstr), ERR_SEP)[1]
}

//返回错误信息
func (errstr Errstr ) Error() string {
	return string(errstr)
}
