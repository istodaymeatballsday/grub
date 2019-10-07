package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var arg = 1
	if len(os.Args) > 1 {
		arg, _ = strconv.Atoi(os.Args[1])
	}
	express(arg)
}

func express(arg int) {
	channels := make([]chan *ExpressResponse, arg)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan *ExpressResponse)
	}
	for i := 0; i < arg; i++ {
		go asyncGetOneExpress(i, channels[i])
	}
	res := make([]*ExpressResponse, arg)
	for i := 0; i < arg; i++ {
		res[i] = <-channels[i]
	}
	PrintColor(Green, "Express")
	for _, e := range res {
		PrintFood(e)
	}
}

func asyncGetOneExpress(daysAgo int, c chan *ExpressResponse) {
	result := getOneExpress(daysAgo)
	c <- result
}

func getOneExpress(daysFromNow int) *ExpressResponse {
	res, err := http.Get(getURL(daysFromNow))
	if err != nil {
		panic("failed to get")
	}
	defer res.Body.Close()
	var result ExpressResponse
	json.NewDecoder(res.Body).Decode(&result)
	return &result
}

func daysAgoString(daysFromNow int) string {
	t := time.Now().AddDate(0, 0, daysFromNow)
	return (t.Format(time.RFC3339)[:10])
}

func getURL(daysAgo int) string {
	formatedDate := daysAgoString(daysAgo)
	return fmt.Sprintf("http://carbonateapiprod.azurewebsites.net/api/v1/mealprovidingunits/3d519481-1667-4cad-d2a3-08d558129279/dishoccurrences?startDate=%s", formatedDate)
}
