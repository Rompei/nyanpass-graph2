package nyanpass

import (
	"fmt"
	"github.com/gonum/plot"
)

type RelabelTicks struct{}

func (RelabelTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		if t.Label == "" { // Skip minor ticks, they are fine.
			continue
		}
		tks[i].Label = convertShortNumber(t.Value, 0) // %g uses scientific notation if it is more compact
		// or tks[i].Label = withCommas(t.Value)
		// or whatever else you want
	}
	return tks
}

func addCommas(s string) string {
	rev := ""
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
		n++
		if n%3 == 0 {
			rev += ","
		}
	}
	s = ""
	for i := len(rev) - 1; i >= 0; i-- {
		s += string(rev[i])
	}
	return s
}

func convertShortNumber(src float64, it int) string {
	c := []string{"K", "M", "B", "T"}
	if src < 1000 {
		return fmt.Sprintf("%.0f", src)
	}
	f := float64((int64(src) / 100) / 10.0)
	isRound := (int64(f)*10)%10 == 0
	if f < 1000 {
		if f > 99.9 || isRound || (!isRound && f > 99.9) {
			return fmt.Sprintf("%d%s", int(f)*10/10, c[it])
		}
		return fmt.Sprintf("%.0f%s", f, c[it])
	}
	return convertShortNumber(f, it+1)

}
