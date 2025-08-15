package main

import "fmt"

// level 定义难度，更好地支持新增难度设置/更改难度命名和尝试次数
type level struct {
	id        int
	label     string
	numChance int
}

// String 定义相应方法，便于以后生成选项列表
func (l *level) String() string {
	return fmt.Sprintf("%s（%d次机会）", l.label, l.numChance)
}

// IsZero 空值判断
func (l *level) IsZero() bool {
	return *l == level{}
}

// 难易程度，应用 Go iota 自增定义方式
const (
	EASY = iota
	MEDIUM
	HARD
)

// levels 更好的支持添加删除
var levels = []level{
	{EASY, "简单", 10},
	{MEDIUM, "中等", 5},
	{HARD, "困难", 3},
}

var LevelChoices []string

// init 初始化难度选择页面的列表
func init() {
	for _, l := range levels {
		LevelChoices = append(LevelChoices, l.String())
	}
}
