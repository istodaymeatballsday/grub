package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	linsen    = "b672efaf-032a-4bb8-d2a5"
	expressen = "3d519481-1667-4cad-d2a3"
)

func main() {
	var arg = 1
	if len(os.Args) > 1 {
		arg, _ = strconv.Atoi(os.Args[1])
	}
	express(arg)
}

func express(arg int) {
	channels := make([]chan *Response, arg)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan *Response)
	}
	for i := 0; i < arg; i++ {
		go asyncGetOneExpress(i, channels[i])
	}
	res := make([]*Response, arg)
	for i := 0; i < arg; i++ {
		res[i] = <-channels[i]
	}
	PrintColor(Green, "Express")
	for _, e := range res {
		PrintFood(e)
	}
}

func lins(arg int) {
	channels := make([]chan *Response, arg)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan *Response)
	}
	for i := 0; i < arg; i++ {
		go asyncGetOneExpress(i, channels[i])
	}
	res := make([]*Response, arg)
	for i := 0; i < arg; i++ {
		res[i] = <-channels[i]
	}
	PrintColor(Green, "Express")
	for _, e := range res {
		PrintFood(e)
	}
}

func asyncGetOneExpress(daysAgo int, c chan *Response) {
	result := getOneExpress(daysAgo)
	c <- result
}

func getOneExpress(daysFromNow int) *Response {
	res, err := http.Get(getURL(daysFromNow, expressen))
	if err != nil {
		panic("failed to get")
	}
	defer res.Body.Close()
	var result Response
	json.NewDecoder(res.Body).Decode(&result)
	return &result
}

func daysAgoString(daysFromNow int) string {
	t := time.Now().AddDate(0, 0, daysFromNow)
	return (t.Format(time.RFC3339)[:10])
}

func getURL(daysAgo int, restaurant string) string {
	formatedDate := daysAgoString(daysAgo)
	return fmt.Sprintf("http://carbonateapiprod.azurewebsites.net/api/v1/mealprovidingunits/%s-08d558129279/dishoccurrences?startDate=%s", restaurant, formatedDate)
}
