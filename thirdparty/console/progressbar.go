package console

import (
	"fmt"
	"strings"
)

// ProgressBar is a way of visualizing work progress in the console
type ProgressBar struct {
	Prefix string
	Suffix string
	Fill   string
	Length int
}

// NewProgressBar initializes and returns a new ProgressBar
func NewProgressBar() ProgressBar {
	return ProgressBar{
		Prefix: "Progress",
		Suffix: "Complete",
		Fill:   "=",
		Length: 25,
	}
}

// Update refreshes the progress
func (p ProgressBar) Update(iteration int, total int) {
	percent := float64(iteration) / float64(total)
	filledLength := p.Length * iteration / total
	end := ">"

	if iteration == total {
		end = "="
	}
	bar := strings.Repeat(p.Fill, filledLength) + end + strings.Repeat("-", p.Length-filledLength)
	fmt.Printf("\r%s [%s] %.2f%% %s", p.Prefix, bar, percent*100, p.Suffix)
	if iteration == total {
		fmt.Println()
	}
}
