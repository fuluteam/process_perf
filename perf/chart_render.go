package perf

import "github.com/go-echarts/go-echarts/charts"

func renderMemoryChart(data []map[string]interface{}) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "内存占用", Subtitle: "数据单位: M"}, charts.InitOpts{PageTitle: "进程资源占用监控"})
	var timeItems []string
	var valuePoints []float32
	for _, item := range data {
		timeItems = append(timeItems, item["time"].(string))
		valuePoints = append(valuePoints, item["value"].(float32))
	}
	line.AddXAxis(timeItems)
	line.AddYAxis("内存", valuePoints, charts.LineOpts{Smooth: true})
	return line
}

func renderCpuChart(data []map[string]interface{}) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "CPU占用", Subtitle: "数据单位: %"})
	var timeItems []string
	var valuePoints []float64
	for _, item := range data {
		timeItems = append(timeItems, item["time"].(string))
		valuePoints = append(valuePoints, item["value"].(float64))
	}
	line.AddXAxis(timeItems)
	line.AddYAxis("CPU", valuePoints, charts.LineOpts{Smooth: true})
	return line
}

func renderThreadChart(data []map[string]interface{}) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Thread数量", Subtitle: "数据单位: 个"})
	var timeItems []string
	var valuePoints []int32
	for _, item := range data {
		timeItems = append(timeItems, item["time"].(string))
		valuePoints = append(valuePoints, item["value"].(int32))
	}
	line.AddXAxis(timeItems)
	line.AddYAxis("线程数", valuePoints, charts.LineOpts{Smooth: true})
	return line
}

func renderIoCountChart(readData []map[string]interface{}, writeData []map[string]interface{}) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "IO次数", Subtitle: "数据单位: 次"})
	var readTimeItems []string
	var readValuePoints []uint64
	for _, item := range readData {
		readTimeItems = append(readTimeItems, item["time"].(string))
		readValuePoints = append(readValuePoints, item["value"].(uint64))
	}
	line.AddXAxis(readTimeItems)
	line.AddYAxis("IO读", readValuePoints, charts.LineOpts{Smooth: true})
	var writeTimeItems []string
	var writeValuePoints []uint64
	for _, item := range writeData {
		writeTimeItems = append(writeTimeItems, item["time"].(string))
		writeValuePoints = append(writeValuePoints, item["value"].(uint64))
	}
	line.AddXAxis(writeTimeItems)
	line.AddYAxis("IO写", writeValuePoints, charts.LineOpts{Smooth: true})
	return line
}
