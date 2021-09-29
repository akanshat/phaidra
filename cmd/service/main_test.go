package main

import (
	"os"
	"testing"

	"github.com/akanshat/phaidra/pkg/datastore"
)

func TestRunService(t *testing.T) {
	file, err := os.Open("./data.json")
	if err != nil {
		return
	}

	qch := make(chan query)
	resch := make(chan []datastore.SensorType)

	testcases := []struct {
		qu       query
		expected []datastore.SensorType
	}{
		{
			qu: query{
				name: "Equipment2",
				qmin: "10:30",
				qmax: "19:15",
			},
			expected: []datastore.SensorType{
				{
					Id:         "Sensor2",
					Start_time: "11:00",
					End_time:   "14:30",
				},
				{
					Id:         "Sensor4",
					Start_time: "07:00",
					End_time:   "14:30",
				},
				{
					Id:         "Sensor5",
					Start_time: "13:45",
					End_time:   "20:45",
				},
			},
		},
		{
			qu: query{
				name: "Equipment2",
				qmin: "11:45",
				qmax: "19:50",
			},
			expected: []datastore.SensorType{
				{
					Id:         "Sensor2",
					Start_time: "11:45",
					End_time:   "19:50",
				},
				{
					Id:         "Sensor4",
					Start_time: "07:00",
					End_time:   "14:30",
				},
				{
					Id:         "Sensor5",
					Start_time: "13:45",
					End_time:   "20:45",
				},
				{
					Id:         "Sensor6",
					Start_time: "11:45",
					End_time:   "19:50",
				},
			},
		},
		{
			qu: query{
				name: "Equipment4",
				qmin: "14:36",
				qmax: "16:10",
			},
			expected: []datastore.SensorType{
				{
					Id:         "Sensor2",
					Start_time: "11:45",
					End_time:   "19:50",
				},
			},
		},
	}
	go runService(file, qch, resch)
	for _, tc := range testcases {
		qch <- tc.qu
		res := <-resch
		for i, v := range res {
			if v.Id != tc.expected[i].Id || v.Start_time != tc.expected[i].Start_time || v.End_time != tc.expected[i].End_time {
				t.Logf("expected %+v, actual %+v, for tc: %+v", tc.expected, v, tc)
				t.Fail()
				break
			}
		}
	}

	close(qch)
	close(resch)

}
