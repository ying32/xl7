package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ying32/xl7"
)

func main() {
	fmt.Println("测试开始")
	if !xl7.InitDownloadEngine() {
		panic("初始引擎失败")
	}
	defer xl7.UninitDownloadEngine()

	fileName := filepath.Dir(os.Args[0]) + "\\QQ9.0.8_2.exe"
	tdFileName := fileName + ".td"

	// 文件存在，则从TD文件加载
	var errID uint32
	var taskID int
	if fileExists(tdFileName) {
		errID, taskID = xl7.ContinueTaskFromTdFile(tdFileName)
	} else {
		errID, taskID = xl7.URLDownloadToFile(fileName, "http://qd.myapp.com/myapp/qqteam/pcqq/QQ9.0.8_2.exe", "")
	}
	if errID == xl7.XL_SUCCESS {
		fmt.Println("任务ID:", taskID)
		for {
			errID, status, pullFileSize, pullRecvSize := xl7.QueryTaskInfo(taskID)
			if errID == xl7.XL_SUCCESS {
				fmt.Printf("status=%d, pullFileSize=%d, pullRecvSize=%d\n", status, pullFileSize, pullRecvSize)
				// 开始下载
				switch status {
				case xl7.Connect:
					fmt.Println("已建立连接")
				case xl7.Download:
					fmt.Println("开始下载")
				case xl7.Pause:
					fmt.Println("暂停")
				case xl7.Success:
					fmt.Println("下载完成")
					break
				case xl7.Fail:
					fmt.Println("下载失败")
					break
				default:
				}
			} else {
				fmt.Println("下载错误：", xl7.GetErrorMsg(errID))
				break
			}
			time.Sleep(time.Second * 1)
		}
	} else {
		fmt.Println("下载错误：", xl7.GetErrorMsg(errID))
	}
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
