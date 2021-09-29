package datastore

import (
	"math"
	"strconv"
	"strings"

	"github.com/akanshat/phaidra/pkg/tree"
)

func hhmmToMinutes(time string) uint {
	str := strings.Split(time, ":")
	hh, _ := strconv.Atoi(str[0])
	mm, _ := strconv.Atoi(str[1])

	return uint(hh*60 + mm)
}

func minutesTohhmm(minutes uint) string {

	num := minutes
	hFloat := float64(num) / 60
	hInt := num / 60
	mins := (hFloat - float64(hInt)) * 60
	var mmInt = math.Round(float64(mins))
	hh := strconv.Itoa(int(hInt))
	mm := strconv.Itoa(int(mmInt))
	return hh + ":" + mm
}

func treeNodesToSensor(nodes []*tree.TreeNode) []SensorType {
	res := make([]SensorType, 0, len(nodes))

	for _, v := range nodes {
		res = append(res, SensorType{
			Id:         v.Name,
			Start_time: minutesTohhmm(v.MinTime),
			End_time:   minutesTohhmm(v.MaxTime),
		})
	}
	return res
}
