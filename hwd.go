package hwd

/* 	#cgo windows,amd64 LDFLAGS:-L./HwdLib/Win64 -lhwd
#include "hwd.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func GetVersion() int {
	return int(C.hwd_getVersion())
}

func GetLastErrorCode() int {
	return int(C.hwd_getLastErrorCode())
}

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

func LoadSkinByByte(skin string, skinLen int, zipPwd string) bool {
	cSkin := C.CString(skin)
	cZipPwd := C.CString(zipPwd)
	defer C.free(unsafe.Pointer(cSkin))
	defer C.free(unsafe.Pointer(cZipPwd))
	result := C.hwd_loadSkinByByte(cSkin, C.int(skinLen), cZipPwd)
	return result != 0
}

func Init(url string, port int, webkey string, sid string, key string, loading bool, proCom bool, isDebug bool, checkDebug int) bool {
	cUrl := C.CString(url)
	cWebkey := C.CString(webkey)
	cSid := C.CString(sid)
	cKey := C.CString(key)

	defer C.free(unsafe.Pointer(cUrl))
	defer C.free(unsafe.Pointer(cWebkey))
	defer C.free(unsafe.Pointer(cSid))
	defer C.free(unsafe.Pointer(cKey))

	result := C.hwd_init(cUrl, C.int(port), cWebkey, cSid, cKey, C.int(BoolToInt(loading)), C.int(BoolToInt(proCom)), C.int(BoolToInt(isDebug)), C.int(checkDebug))
	return result != 0
}
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

	result := C.hwd_fastCheck(cUrl, C.int(port), cWebkey, cSid, cKey, cSoftPara, C.int(BoolToInt(isDebug)), C.int(checkDebug))
	return result != 0
}

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

func BoolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
func BtI(value bool) C.int {
	return C.int(BoolToInt(value))
}
func LoadLoginWindow(version string, title string, noticeTime int, menuItem string, autoHeartbeat bool) bool {
	cVersion := C.CString(version)
	cTitle := C.CString(title)
	cMenuItem := C.CString(menuItem)

	defer C.free(unsafe.Pointer(cVersion))
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMenuItem))

	result := C.hwd_loadLoginWindow(cVersion, cTitle, C.int(noticeTime), cMenuItem, BtI(autoHeartbeat))

	return result != 0
}
func LoadSkinByFile(filePath string, zipPwd string) bool {
	cFilePath := C.CString(filePath)
	cZipPwd := C.CString(zipPwd)

	defer C.free(unsafe.Pointer(cFilePath))
	defer C.free(unsafe.Pointer(cZipPwd))
	result := C.hwd_loadSkinByFile(cFilePath, cZipPwd)
	return result != 0
}

func BlueSky() bool {
	return C.hwd_blueSky() != 0
}
