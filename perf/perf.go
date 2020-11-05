/**
 * @author [kevinyang]
 * @email [yangchujie6@mail.com]
 * @create date 2020-11-04 17:05
 * @modify date 2020-11-04 17:05
 * @desc [进程数据采集服务]
 */
package perf

import (
	"errors"
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func StartCollect(pid int32) (err error) {
	var processExist bool
	var currentProcess *process.Process
	if processExist, err = process.PidExists(pid); err != nil {
		return
	}
	if !processExist {
		err = errors.New(fmt.Sprintf("指定的进程ID: %d 不存在", pid))
		return
	}

	if currentProcess, err = process.NewProcess(pid); err != nil {
		return
	}
	if processExe, e := currentProcess.Exe(); e != nil {
		err = errors.New(fmt.Sprintf("获取进程执行路径异常: %s", e))
		return
	} else {
		fmt.Printf("开始进行进程资源占用数据采集. process: %s \n", processExe)
	}
	context := &Context{
		p: currentProcess,
	}
	context.collect()
	return
}
