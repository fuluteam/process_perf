package perf

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/labstack/gommon/log"
	"github.com/shirou/gopsutil/process"
)

const (
	INSPECT_INTERVAL      = 2 * time.Second // 采样时间间隔
	TIME_FORMAT           = "15:04:05"      // 采样时间格式
	OUTPUT_CHART_FILENAME = "chart.html"    // 图表输出文件名
)

type Context struct {
	p                *process.Process
	memoryStat       []map[string]interface{}
	cpuStat          []map[string]interface{}
	threadStat       []map[string]interface{}
	ioReadCountStat  []map[string]interface{}
	ioWriteCountStat []map[string]interface{}
}

func (c *Context) collect() {
	var wg sync.WaitGroup
	wg.Add(1)
	collectTicker := time.NewTicker(INSPECT_INTERVAL)
	defer func() {
		collectTicker.Stop()
	}()
	userSignal := make(chan os.Signal)
	signal.Notify(userSignal, os.Interrupt)
	go func(t *time.Ticker) {
		for {
			<-t.C
			fmt.Println("采集进程数据: ", time.Now().Format("15:04:05"))
			c.startAllInspect()
		}
	}(collectTicker)
	select {
	case <-userSignal:
		fmt.Println("停止进程资源采集。 开始生成进程资源占用图表...")
		c.renderCharts()
		fmt.Println("资源占用图表生成完毕。你可以点击当前目录下「chart.html」进行查看。")
		os.Exit(1)
	}
	wg.Wait()
}

func (c *Context) renderCharts() {
	p := charts.NewPage()
	p.Add(
		renderMemoryChart(c.memoryStat),
		renderCpuChart(c.cpuStat),
		renderThreadChart(c.threadStat),
		// renderIoCountChart(process.ioReadCountStat, process.ioWriteCountStat),
	)
	f, err := os.Create(OUTPUT_CHART_FILENAME)
	if err != nil {
		log.Error(err)
	}
	p.Render(f)
}

func (c *Context) startAllInspect() {
	c.memoryInspect()
	c.cpuInspect()
	c.threadInspect()
}

func (c *Context) memoryInspect() {
	memInfo, _ := c.p.MemoryInfo()
	memValue := float32(memInfo.RSS) / 1024 / 1024
	c.memoryStat = append(c.memoryStat, map[string]interface{}{
		"time":  time.Now().Format(TIME_FORMAT),
		"value": memValue,
	})
}

func (c *Context) cpuInspect() {
	cpuPercent, err := c.p.CPUPercent()
	if err != nil {
		log.Error(err)
		return
	}
	c.cpuStat = append(c.cpuStat, map[string]interface{}{
		"time":  time.Now().Format(TIME_FORMAT),
		"value": cpuPercent,
	})
}

func (c *Context) threadInspect() {
	threadNums, err := c.p.NumThreads()
	if err != nil {
		log.Error(err)
		return
	}
	c.threadStat = append(c.threadStat, map[string]interface{}{
		"time":  time.Now().Format(TIME_FORMAT),
		"value": threadNums,
	})
}

func (c *Context) ioCountInspect() {
	ioStat, err := c.p.IOCounters()
	if err != nil {
		log.Error(err)
		return
	}
	c.ioReadCountStat = append(c.ioReadCountStat, map[string]interface{}{
		"time":  time.Now().Format(TIME_FORMAT),
		"value": ioStat.ReadCount,
	})
	c.ioWriteCountStat = append(c.ioWriteCountStat, map[string]interface{}{
		"time":  time.Now().Format(TIME_FORMAT),
		"value": ioStat.WriteCount,
	})
}
