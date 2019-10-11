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

type restaurant struct {
	DisplayName string
	URL         string
}

func main() {
	var arg = 1
	if len(os.Args) > 1 {
		arg, _ = strconv.Atoi(os.Args[1])
	}
	getFood(arg)
}

func getFood(arg int) {
	restaurants := []restaurant{restaurant{
		DisplayName: "Express",
		URL:         expressen,
	}, restaurant{
		DisplayName: "Linsen",
		URL:         linsen,
	}}
	restaurantChannels := make([]chan []*Response, len(restaurants))
	restaurantResponses := make([][]*Response, len(restaurants))
	for i := range restaurantChannels {
		restaurantChannels[i] = make(chan []*Response)
	}
	for i := range restaurants {
		go getFoodFromRestaurant(arg, &(restaurants[i]), restaurantChannels[i])
	}
	for i, c := range restaurantChannels {
		restaurantResponses[i] = <-c
	}
	for i, r := range restaurants {
		PrintColor(Green, r.DisplayName)
		for _, res := range restaurantResponses[i] {
			swedishFirst := r.DisplayName == linsen
			PrintFood(res, swedishFirst)
		}
	}
}

func getFoodFromRestaurant(arg int, restaurant *restaurant, c chan []*Response) {
	channels := make([]chan *Response, arg)
	for i := range channels {
		channels[i] = make(chan *Response)
	}
	for i, channel := range channels {
		go getOneDay(restaurant, i, channel)
	}
	res := make([]*Response, arg)
	for i := 0; i < arg; i++ {
		res[i] = <-channels[i]
	}
	c <- res
}

func getOneDay(restaurant *restaurant, daysFromNow int, c chan *Response) {
	res, err := http.Get(getURL((*restaurant).URL, daysFromNow))
	if err != nil {
		panic("failed to get")
	}
	defer res.Body.Close()
	var result Response
	json.NewDecoder(res.Body).Decode(&result)
	c <- &result
}

func getURL(restaurant string, daysAgo int) string {
	formatedDate := daysAgoString(daysAgo)
	return fmt.Sprintf("http://carbonateapiprod.azurewebsites.net/api/v1/mealprovidingunits/%s-08d558129279/dishoccurrences?startDate=%s", restaurant, formatedDate)
}

func daysAgoString(daysFromNow int) string {
	t := time.Now().AddDate(0, 0, daysFromNow)
	return (t.Format(time.RFC3339)[:10])
}
