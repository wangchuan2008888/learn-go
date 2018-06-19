package statistic

import (
	"sort"
	"fmt"
	"strings"
	"strconv"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) (median float64) {
	_len := len(numbers)
	middle := _len / 2
	rslt := numbers[middle]
	if len(numbers) %2 == 0 {
		rslt = (rslt + numbers[middle-1])/2
	}
	return rslt
}
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(stats.numbers)
	return stats
}

func (stat *statistics) String() string {
	var str_numbers = make([] string, len(stat.numbers))
	for idx, number := range stat.numbers {
		str_numbers[idx] = strconv.FormatFloat(number, 'f', 2, 64)
	}
	return fmt.Sprintf("numbers: %s\nmedia:%.2f\nmedia: %.2f\n", strings.Join(str_numbers, ","), stat.mean, stat.median)
}

