package Hwd

/* 	#cgo windows,amd64 LDFLAGS:-L./HwdLib/Win64 -lhwd
#include "hwd.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
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
func GetLastErrorMsg(buffers *string, bufferLen int) bool {
	buffer := make([]byte, bufferLen)
	result := C.hwd_getLastErrorMsg((*C.char)(unsafe.Pointer(&buffer[0])), C.int(bufferLen))
	if result != 0 {
		//fmt.Printf("buffercs: %s\n", string(buffer[0:]))
		*buffers = string(buffer[0:])
		return true
	}
	return false
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
func GetFileMD5(filename string, buffer *string, bufferLen int) bool {
	cFilename := C.CString(filename)
	//cBuffer := C.CString(*buffer)
	defer C.free(unsafe.Pointer(cFilename))
	cBuffer := make([]byte, bufferLen)
	result := C.hwd_getFileMD5(cFilename, (*C.char)(unsafe.Pointer(&cBuffer[0])), C.int(bufferLen))
	if result != 0 {
		fmt.Printf("cBuffer %v", string(cBuffer[0:]))
		return true
	}
	return false
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

	return result != 0
}

//从文件加载皮肤
func LoadSkinByFile(filePath string, zipPwd string) bool {
	cFilePath := C.CString(filePath)
	cZipPwd := C.CString(zipPwd)

	defer C.free(unsafe.Pointer(cFilePath))
	defer C.free(unsafe.Pointer(cZipPwd))
	result := C.hwd_loadSkinByFile(cFilePath, cZipPwd)
	return result != 0
}

//蓝屏 真蓝屏Win11亲测
func BlueSky() bool {
	return C.hwd_blueSky() != 0
}
