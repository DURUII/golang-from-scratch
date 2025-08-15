package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	ErrNaN         = errors.New("只对数值产生反应，谢谢")
	ErrOutOfBounds = errors.New("大胆，警告你踩过界了")
	gap            = "\n"
	question       = "? "
	cursor         = "▸ "
	checkmark      = "✔ "
	cross          = "✗ "
	welcomeFooter  = "↑ ↓ → ←: [选择] enter: [确认] q: [退出]"
	welcomeHint    = "意外闯入者，请解开数字宝藏的密码锁。好运！ "
	playingFooter  = "enter: [确认] q: [退出]"
	resultFooter   = "q: [退出] r: [重开]"
	hintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	questionStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	cursorStyle    = lipgloss.NewStyle()
	checkmarkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	crossStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

func (m model) View() string {
	footer := hintStyle.Render(welcomeFooter)
	switch m.statusCode {
	case WELCOME:
		header := hintStyle.Render(welcomeHint)
		content := welcomeContent(m.cursor)
		return header + gap + content + footer
	case PLAYING:
		footer = hintStyle.Render(playingFooter)
		header, warning := questionPrompt(m.isLegal, m.errMsg, len(m.guessTrace), m.selectedLevel.numChance)
		hint := ""
		if len(m.guessTrace) > 0 {
			guessVal := m.guessTrace[len(m.guessTrace)-1]
			hint = playingHint(guessVal, m.target)
		}
		if warning != "" {
			hint = warning
		}
		hint = questionStyle.Render(hint)
		return header + gap + m.textInput.View() + gap + hint + gap + footer
	case WIN, LOSE:
		footer = resultFooter
		header := "最后一次尝试失败，密门沉默不语，谜题仍未破解。"
		content := fmt.Sprintf("耗时 %d 分 %d 秒", int(m.tocTime.Sub(m.ticTime).Minutes()), int(m.tocTime.Sub(m.ticTime).Seconds())%60)
		if m.statusCode == WIN {
			header = "机关轻响，密码锁应声而开。谜题已解！"
			content = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Render(fmt.Sprintf("总共猜测次数：%d 次", len(m.guessTrace)) + gap + content)
		} else {
			content = lipgloss.NewStyle().Foreground(lipgloss.Color("125")).Render(content)
		}
		return hintStyle.Render(header) + gap + gap + content + gap + gap + hintStyle.Render(footer)
	}
	return ""
}

func welcomeContent(selected int) string {
	var s strings.Builder
	s.WriteString(questionStyle.Render(question) + "要接受怎样强度的挑战: " + gap)
	for i, level := range levels {
		prefix := "  "
		var isSelected bool
		if isSelected = i == selected; isSelected {
			prefix = cursorStyle.Render(cursor)
		}
		s.WriteString(fmt.Sprintf("%s%s%s", prefix,
			lipgloss.NewStyle().Underline(isSelected).Render(level.String()), gap),
		)
	}
	return s.String()
}

func playingHint(guessVal, target int) (hint string) {
	if guessVal > target {
		hint = fmt.Sprintf("%d吗，可惜猜的太大了——再保守些", guessVal)
	} else if guessVal < target {
		hint = fmt.Sprintf("%d？数字太小了，连门锁都懒得响。", guessVal)
	} else {
		hint = "机关轻响，密码锁应声而开。你成功解开了谜题。"
	}
	return
}

func questionPrompt(isLegal bool, errMsg error, numHasGuessed int, numTotalChance int) (string, string) {
	var header, warning strings.Builder
	mark := crossStyle.Render(cross)
	if isLegal {
		mark = checkmarkStyle.Render(checkmark)
	}
	if errMsg != nil {
		warning.WriteString(crossStyle.Render(fmt.Sprint(errMsg)))
	}
	header.WriteString("开始解谜，剩余次数：[")
	for i := 0; i < numTotalChance; i++ {
		if i >= numTotalChance-numHasGuessed {
			header.WriteString("▱")
		} else {
			header.WriteString("▰")
		}
	}
	header.WriteString(fmt.Sprintf("] [%d/%d]\n", numTotalChance-numHasGuessed, numTotalChance))
	header.WriteString(mark)
	header.WriteString("输入你心中的猜想整数（1-100）")
	return hintStyle.Render(header.String()), warning.String()
}
