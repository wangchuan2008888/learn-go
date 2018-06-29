package statistic

import (
	"testing"
	"fmt"
	"net/http"
	"log"
	"strings"
	"strconv"
)

func Test_statistic(t *testing.T) {
	data := []float64{1.2, 1.3, 1.4, 1.5, 1.2}
	stat := getStats(data)
	fmt.Println(stat.String())
}

var header string = "<html>"
var bottom string = "</html>"
var form string = `<form action='/' method='post'>
<label for="numbers"> Numbers:</label><br/>
<input type="text" name="numbers"/> <br/>
<input type="submit" value="Calculate"/>
</form>`

func processRequest(requst *http.Request) (numbers []float64, ok bool) {
	if slice, fond := requst.Form["numbers"]; fond {
		var numbers []float64
 		fmt.Println(slice)
 		fmt.Println(slice[0])
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			value, _ := strconv.ParseFloat(field,64)
			numbers = append(numbers, value)
		}
		return numbers, true
	}
	return nil, false
}
func formatStats(state *statistics) string {
	return state.String()
}
func homePage(writer http.ResponseWriter, requst *http.Request) {
	err := requst.ParseForm()
	fmt.Fprint(writer, header)

	if err != nil {
		fmt.Fprint(writer, err)
	} else {
		if numbers, ok := processRequest(requst); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(&stats))
		} else {
			fmt.Fprint(writer, form)
		}
	}
	fmt.Fprint(writer, bottom)
}

func Test_server(t *testing.T) {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("fail to start server", err)
	}
}
