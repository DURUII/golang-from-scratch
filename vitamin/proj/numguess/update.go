package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strconv"
	"time"
)

const filepath string = "./game.txt"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.statusCode {
	case WELCOME:
		return m.UpdateWelcome(msg)
	case PLAYING:
		return m.UpdatePlaying(msg)
	case WIN, LOSE:
		return m.UpdateResult(msg)
	}
	return m, nil
}

func (m model) UpdateWelcome(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			// 更新光标
			m.cursor = (m.cursor - 1 + len(levels)) % len(levels)
		case "down":
			m.cursor = (m.cursor + 1 + len(levels)) % len(levels)
		case "enter":
			// 初始化难度设置，却换状态
			m.selectedLevel = levels[m.cursor]
			m.statusCode = PLAYING
			m.ticTime = time.Now()
			m.cursor = 0
		}
	case error:
		return m, nil
	}
	return m, nil
}

func (m model) UpdatePlaying(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// 清空错误信息
	m.errMsg = nil
	m.textInput.Update(m.textInput.Focus())
	m.textInput.Update(textinput.Blink())
	m.textInput, cmd = m.textInput.Update(msg)
	// 验证数字是否正确，即时反馈
	guessVal, err := validateInput(m.textInput.Value())
	// 数据校验不通过
	if err != nil {
		m.isLegal = false
	} else {
		// 数据校验通过
		m.isLegal = true
		m.errMsg = nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			// 只有当用户按 Enter 之后，错误信息才会显示
			if err != nil {
				m.errMsg = err
			} else {
				// 更新原子状态
				m.guessTrace = append(m.guessTrace, guessVal)
				// 判断输赢
				if guessVal == m.target {
					m.statusCode = WIN
				}
				if len(m.guessTrace) == m.selectedLevel.numChance {
					m.statusCode = LOSE
				}
				// 还原初始状态
				m.textInput.Reset()
				m.textInput.CursorStart()
				m.textInput.Cursor.Focus()
				m.errMsg = nil
				m.isLegal = false
				// 更新时间
				m.tocTime = time.Now()
			}
		}
	case error:
		return m, nil
	}
	return m, cmd
}

func validateInput(input string) (int, error) {
	val, err := strconv.Atoi(input)
	if err != nil {
		return -1, ErrNaN
	}
	if 1 <= val && val <= 100 {
		return val, nil
	}
	return -1, ErrOutOfBounds
}

func (m model) UpdateResult(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if err := m.SaveToFile(filepath); err != nil {
				return m, tea.Quit
			}
			return m, tea.Quit
		case "r":
			if err := m.SaveToFile(filepath); err != nil {
				return m, tea.Quit
			}
			m = newModel()
			return m, nil
		}
	}
	return m, nil
}
