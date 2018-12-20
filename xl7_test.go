package xl7

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	t.Log("测试开始")
	if !InitDownloadEngine() {
		t.Fatal("初始引擎失败")
	}
	defer UninitDownloadEngine()

	errID, taskID := URLDownloadToFile("./QQ9.0.8_2.exe", "http://qd.myapp.com/myapp/qqteam/pcqq/QQ9.0.8_2.exe", "")
	if errID == XL_SUCCESS {
		t.Log("任务ID:", taskID)

		errID, status, pullFileSize, pullRecvSize := QueryTaskInfo(taskID)
		if errID == XL_SUCCESS {
			t.Logf("status=%d, pullFileSize=%d, pullRecvSize=%d", status, pullFileSize, pullRecvSize)
			// 开始下载
			switch status {
			case Connect:
				t.Log("已建立连接")
			case Download:
				t.Log("开始下载")
			case Pause:
				t.Log("暂停")
			case Success:
				t.Log("下载完成")
				break
			case Fail:
				t.Log("下载失败")
				break
			}
		} else {
			t.Log("下载错误：", GetErrorMsg(errID))
			break
		}
		time.Sleep(time.Second * 1)

	} else {
		t.Log("下载错误：", GetErrorMsg(errID))
	}
}
