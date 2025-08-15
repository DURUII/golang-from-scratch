package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	WELCOME = iota
	PLAYING
	WIN
	LOSE
)

type model struct {
	// 状态码
	statusCode int
	ticTime    time.Time
	tocTime    time.Time

	// 输入框相关
	errMsg     error
	isLegal    bool
	guessTrace []int
	target     int
	textInput  textinput.Model

	// 选项卡
	cursor        int
	selectedLevel level
}

// 新建模型（初始化 + 重开）
func newModel() model {
	model := model{
		statusCode: WELCOME,
		isLegal:    false,
		guessTrace: make([]int, 0, 100),
		target:     rand.Intn(99) + 1,
		textInput:  textinput.New(),
	}
	return model
}

// SaveToFile 序列化 txt 文件
func (m model) SaveToFile(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	var s strings.Builder
	var result = "LOSE"
	if m.guessTrace[len(m.guessTrace)-1] == m.target {
		result = "WIN"
	}
	layout := "2006-01-02 15:04:05"
	s.WriteString(fmt.Sprintln(m.ticTime.Format(layout), m.tocTime.Format(layout), m.tocTime.Sub(m.ticTime), m.target, m.guessTrace, result))
	if _, err := f.WriteString(s.String()); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
func (m model) Init() tea.Cmd { return nil }
