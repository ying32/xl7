package xl7

import (
	"syscall"
	"unsafe"
)

/*
------------------导出函数--------------
*/

var (
	_XLDownloadDLl            = syscall.NewLazyDLL("XLDownload.dll")
	_XLInitDownloadEngine     = _XLDownloadDLl.NewProc("XLInitDownloadEngine")
	_XLURLDownloadToFile      = _XLDownloadDLl.NewProc("XLURLDownloadToFile")
	_XLQueryTaskInfo          = _XLDownloadDLl.NewProc("XLQueryTaskInfo")
	_XLPauseTask              = _XLDownloadDLl.NewProc("XLPauseTask")
	_XLContinueTask           = _XLDownloadDLl.NewProc("XLContinueTask")
	_XLContinueTaskFromTdFile = _XLDownloadDLl.NewProc("XLContinueTaskFromTdFile")
	_XLStopTask               = _XLDownloadDLl.NewProc("XLStopTask")
	_XLUninitDownloadEngine   = _XLDownloadDLl.NewProc("XLUninitDownloadEngine")
	_XLGetErrorMsg            = _XLDownloadDLl.NewProc("XLGetErrorMsg")
)

func toStrPtr(str string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

// InitDownloadEngine 初始下载引擎
func InitDownloadEngine() bool {
	r, _, _ := _XLInitDownloadEngine.Call()
	return r != 0
}

// URLDownloadToFile 从URL下载文件
func URLDownloadToFile(fileName, URL, refURL string) (errID uint32, taskID int) {
	r, _, _ := _XLURLDownloadToFile.Call(toStrPtr(fileName), toStrPtr(URL), toStrPtr(refURL), uintptr(unsafe.Pointer(&taskID)))
	errID = uint32(r)
	return
}

// QueryTaskInfo 查询指定下载任务的信息
func QueryTaskInfo(taskID int) (errID uint32, status int, pullFileSize, pullRecvSize uint64) {
	status = -1
	r, _, _ := _XLQueryTaskInfo.Call(uintptr(taskID), uintptr(unsafe.Pointer(&status)), uintptr(unsafe.Pointer(&pullFileSize)), uintptr(unsafe.Pointer(&pullRecvSize)))
	errID = uint32(r)
	return
}

// PauseTask 暂停指定任务
func PauseTask(taskID int, newTaskID *int) uint32 {
	r, _, _ := _XLPauseTask.Call(uintptr(taskID), uintptr(unsafe.Pointer(newTaskID)))
	return uint32(r)
}

// ContinueTask 继续下载指定任务
func ContinueTask(taskID int) uint32 {
	r, _, _ := _XLContinueTask.Call(uintptr(taskID))
	return uint32(r)
}

// ContinueTaskFromTdFile 从.td文件继续下载, 并返回任务ID
func ContinueTaskFromTdFile(tdFileFullPath string) (errID uint32, taskID int) {
	r, _, _ := _XLContinueTaskFromTdFile.Call(toStrPtr(tdFileFullPath), uintptr(unsafe.Pointer(&taskID)))
	errID = uint32(r)
	return
}

// StopTask 停止下载指定任务
func StopTask(taskID int) {
	_XLStopTask.Call(uintptr(taskID))
}

// UninitDownloadEngine 反向初始下载引擎
func UninitDownloadEngine() bool {
	r, _, _ := _XLUninitDownloadEngine.Call()
	return r != 0
}

// GetErrorMsg 从错误ID转为错误文本信息
func GetErrorMsg(errorID uint32) string {
	var ptr, size uintptr
	r, _, _ := _XLGetErrorMsg.Call(uintptr(errorID), uintptr(unsafe.Pointer(&ptr)), uintptr(unsafe.Pointer(&size)))
	if r == XL_ERROR_BUFFER_TOO_SMALL {
		str := make([]uint16, size+1)
		r, _, _ := _XLGetErrorMsg.Call(uintptr(errorID), uintptr(unsafe.Pointer(&str[0])), uintptr(unsafe.Pointer(&size)))
		if r == XL_SUCCESS {
			return syscall.UTF16ToString(str)
		}
	}
	return ""
}

// 下载状态定义
// type TaskStatus uint32

const (
	Connect  = 0  // 已经建立连接
	Download = 2  // 开始下载
	Pause    = 10 // 暂停
	Success  = 11 // 成功下载
	Fail     = 12 // 下载失败
)
