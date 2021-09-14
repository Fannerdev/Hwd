package hwd

/* 	#cgo windows,amd64 LDFLAGS:-L./HwdLib/Win64 -lhwd
#include "hwd.h"
#include <stdlib.h>
*/
import "C"
import (
	"strings"
	"unsafe"
)

//获取SDK版本
func GetVersion() int {
	return int(C.hwd_getVersion())
}

//获取最后错误代码
func GetLastErrorCode() int {
	return int(C.hwd_getLastErrorCode())
}

//获取最后错误信息
func GetLastErrorMsg(bufferLen int) (string, bool) {
	buffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(&buffer[0]))
	result := C.hwd_getLastErrorMsg((*C.char)(unsafe.Pointer(&buffer[0])), C.int(bufferLen))
	return DeleteZero(string(buffer[0:])), int(result) != 0
	// if result != 0 {
	// 	//fmt.Printf("buffercs: %s\n", string(buffer[0:]))
	// 	*buffers =
	// 	return true
	// }
	// return false
}

//加载皮肤Bytes
func LoadSkinByByte(skin string, skinLen int, zipPwd string) bool {
	cSkin := C.CString(skin)
	cZipPwd := C.CString(zipPwd)
	defer C.free(unsafe.Pointer(cSkin))
	defer C.free(unsafe.Pointer(cZipPwd))
	result := C.hwd_loadSkinByByte(cSkin, C.int(skinLen), cZipPwd)
	return result != 0
}

//普通初始化
func Init(url string, port int, webkey string, sid string, key string, loading bool, proCom bool, isDebug bool, checkDebug int) bool {
	cUrl := C.CString(url)
	cWebkey := C.CString(webkey)
	cSid := C.CString(sid)
	cKey := C.CString(key)

	defer C.free(unsafe.Pointer(cUrl))
	defer C.free(unsafe.Pointer(cWebkey))
	defer C.free(unsafe.Pointer(cSid))
	defer C.free(unsafe.Pointer(cKey))

	result := C.hwd_init(cUrl, C.int(port), cWebkey, cSid, cKey, btI(loading), btI(proCom), btI(isDebug), C.int(checkDebug))
	return result != 0
}

//快速验证
func FastCheck(url string, port int, webkey string, sid string, key string, softPara string, isDebug bool, checkDebug int) bool {
	cUrl := C.CString(url)
	cWebkey := C.CString(webkey)
	cSid := C.CString(sid)
	cKey := C.CString(key)
	cSoftPara := C.CString(softPara)

	defer C.free(unsafe.Pointer(cUrl))
	defer C.free(unsafe.Pointer(cWebkey))
	defer C.free(unsafe.Pointer(cSid))
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cSoftPara))

	result := C.hwd_fastCheck(cUrl, C.int(port), cWebkey, cSid, cKey, cSoftPara, btI(isDebug), C.int(checkDebug))
	return result != 0
}

//取文件MD5
func GetFileMD5(filename string, buffer *string, bufferLen int) (string, bool) {
	cFilename := C.CString(filename)
	//cBuffer := C.CString(*buffer)
	defer C.free(unsafe.Pointer(cFilename))
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_getFileMD5(cFilename, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
	// if result != 0 {
	// 	//fmt.Printf("cBuffer %v", string(cBuffer[0:]))
	// 	return string(cBuffer[0:]),true
	// }
	// return string(cBuffer[0:]),false
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
func btI(value bool) C.int {
	return C.int(boolToInt(value))
}

//加载登录窗口
func LoadLoginWindow(version string, title string, noticeTime int, menuItem string, autoHeartbeat bool) bool {
	cVersion := C.CString(version)
	cTitle := C.CString(title)
	cMenuItem := C.CString(menuItem)

	defer C.free(unsafe.Pointer(cVersion))
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMenuItem))

	result := C.hwd_loadLoginWindow(cVersion, cTitle, C.int(noticeTime), cMenuItem, btI(autoHeartbeat))

	return int(result) != 0
}

//从文件加载皮肤
func LoadSkinByFile(filePath string, zipPwd string) bool {
	cFilePath := C.CString(filePath)
	cZipPwd := C.CString(zipPwd)

	defer C.free(unsafe.Pointer(cFilePath))
	defer C.free(unsafe.Pointer(cZipPwd))
	result := C.hwd_loadSkinByFile(cFilePath, cZipPwd)
	return int(result) != 0
}

//蓝屏 真蓝屏Win11亲测
func BlueSky() bool {
	return int(C.hwd_blueSky()) != 0
}

//注册
func Reg(username string, password string, email string, referrer string, code string) bool {
	cUsername := C.CString(username)
	cPassword := C.CString(password)
	cEmail := C.CString(email)
	cReferrer := C.CString(referrer)
	cCode := C.CString(code)

	defer C.free(unsafe.Pointer(cUsername))
	defer C.free(unsafe.Pointer(cPassword))
	defer C.free(unsafe.Pointer(cEmail))
	defer C.free(unsafe.Pointer(cReferrer))
	defer C.free(unsafe.Pointer(cCode))
	return int(C.hwd_reg(cUsername, cPassword, cEmail, cReferrer, cCode)) != 0
}

//发送密码重置邮件
func SendMail(username string, email string, code string) bool {
	cUsername := C.CString(username)
	cEmail := C.CString(email)
	cCode := C.CString(code)
	defer C.free(unsafe.Pointer(cUsername))
	defer C.free(unsafe.Pointer(cEmail))
	defer C.free(unsafe.Pointer(cCode))
	return int(C.hwd_sendMail(cUsername, cEmail, cCode)) != 0
}

//函数说明：获取登录用户信息,根据提交参数名,返回指定用户数据.
//参数<1>：name，username=用户名,password=密码,token=登录token(用于校验登录状态),auth=登录令牌,endtime=到期时间,point=点数余额,balance=账户余额,para=用户自定义数据,bind=用户绑定信息
func GetUserInfo(name string, bufferLen int) (string, bool) {
	cName := C.CString(name)
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_getUserInfo(cName, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

//获取快速验证信息
func GetFastInfo(name string, bufferLen int) (string, bool) {
	cName := C.CString(name)
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(cName))
	result := C.hwd_getFastInfo(cName, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

//获取快速验证自定义参数
func GetFastPara(name string, bufferLen int) (string, bool) {
	cName := C.CString(name)
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_getFastPara(cName, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

func DeleteZero(str string) string {
	return strings.ReplaceAll(str, string(byte(0)), "")
}

func CallPHP(EventName string, Para string, bufferLen int) (string, bool) {
	cEventName := C.CString(EventName)
	cPara := C.CString(Para)
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(cEventName))
	defer C.free(unsafe.Pointer(cPara))
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_callPHP(cEventName, cPara, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

func Save(name string, value string) bool {
	cName := C.CString(name)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(cValue))
	result := C.hwd_save(cName, cValue)
	return int(result) != 0
}

func Read(name string, defaultValue string, bufferLen int) (string, bool) {
	cName := C.CString(name)
	cValue := C.CString(defaultValue)
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(cValue))
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_read(cName, cValue, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

func GetMachineCode(bufferLen int) (string, bool) {
	cBuffer := make([]byte, bufferLen)
	defer C.free(unsafe.Pointer(&cBuffer[0]))
	result := C.hwd_getMachineCode((*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	return DeleteZero(string(cBuffer[0:])), int(result) != 0
}

//return strings.ReplaceAll(strR, "\n", "")
