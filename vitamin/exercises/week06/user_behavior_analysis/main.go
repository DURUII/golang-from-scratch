package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Record struct {
	t    time.Time
	u    string
	op   string
	info string
}

// 用户统计信息
type T1 struct {
	u       string
	countOp int
	firstOp time.Time
	lastOp  time.Time
}

// 行为统计信息
type T2 struct {
	op      string
	countOp int
}

func parseRecords(filepath string) []Record {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	layout := "2006-01-02 15:04:05" // Define the layout
	records := make([]Record, 0, 100)
	for _, line := range lines {
		items := strings.Split(line, ",")
		temp := items[0]
		t, _ := time.Parse(layout, temp)
		u := items[1]
		op := items[2]
		info := items[3]
		records = append(records, Record{t, u, op, info})
	}
	return records
}

func calcUserStatistics(records []Record) map[string]T1 {
	t1 := map[string]T1{}
	layout := "2006-01-02 15:04:05" // Define the layout
	for _, record := range records {
		if _, ok := t1[record.u]; !ok {
			t1[record.u] = T1{
				u:       record.u,
				countOp: 1,
				firstOp: record.t,
				lastOp:  record.t,
			}
		} else {
			temp := t1[record.u]
			temp.countOp += 1
			temp.lastOp = record.t
			t1[record.u] = temp
		}
	}
	var str strings.Builder
	file, err := os.OpenFile("user_statistics.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开或创建文件:", err)
		return nil
	}
	defer file.Close()
	for _, value := range t1 {
		str.WriteString(fmt.Sprintf("%s %d %s %s\n", value.u, value.countOp, value.firstOp.Format(layout), value.lastOp.Format(layout)))
	}
	n, err := file.WriteString(str.String())
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return nil
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
	return t1
}

func calcActionStatistics(records []Record) map[string]T2 {
	t2 := map[string]T2{}
	for _, record := range records {
		if _, ok := t2[record.op]; !ok {
			t2[record.op] = T2{
				op:      record.op,
				countOp: 1,
			}
		} else {
			temp := t2[record.op]
			temp.countOp += 1
			t2[record.op] = temp
		}
	}
	var str strings.Builder
	file, err := os.OpenFile("action_statistics.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开或创建文件:", err)
		return nil
	}
	defer file.Close()
	for _, value := range t2 {
		str.WriteString(fmt.Sprintf("%s %d\n", value.op, value.countOp))
	}
	n, err := file.WriteString(str.String())
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return nil
	}
	fmt.Printf("成功写入 %d 个字节\n", n)
	return t2
}

//func calcTimeStatistics(records []Record) map[string]T3 {
//	return nil
//}

func main() {
	// 文件读取与解析
	records := parseRecords("user_actions.log")

	// 用户统计信息
	calcUserStatistics(records)

	// 行为统计信息
	calcActionStatistics(records)

	// 时间统计信息
	//calcTimeStatistics(records)
}
