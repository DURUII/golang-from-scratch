package main

type MinStack struct {
	items   []int
	histMin []int // prefix min
}

func Constructor() MinStack {
	return MinStack{
		items:   make([]int, 0, 30_000),
		histMin: make([]int, 0, 30_000),
	}
}

// 风格建议简短且语义明确的小写名词或首字母
func (s *MinStack) Push(val int) {
	s.items = append(s.items, val)
	if len(s.histMin) == 0 || val < s.histMin[len(s.histMin)-1] {
		s.histMin = append(s.histMin, val)
	} else {
		s.histMin = append(s.histMin, s.histMin[len(s.histMin)-1])
	}
}

func (s *MinStack) Pop() {
	s.items = s.items[:len(s.items)-1]
	s.histMin = s.histMin[:len(s.histMin)-1]
}

func (s *MinStack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *MinStack) GetMin() int {
	return s.histMin[len(s.histMin)-1]
}
