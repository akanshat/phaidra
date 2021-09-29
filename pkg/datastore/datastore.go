package datastore

import (
	"github.com/akanshat/phaidra/pkg/tree"
)

type DataStore struct {
	store map[string]*tree.Tree
}

func NewDataStore(equipmentsMappingRaw []byte) (*DataStore, error) {
	mapping, err := getMapping(equipmentsMappingRaw)

	if err != nil {
		return nil, err
	}

	d := &DataStore{
		store: make(map[string]*tree.Tree),
	}

	for key, value := range mapping.EquipmentsMapping {
		tr := tree.NewTree()

		for _, v := range value {
			tr.Insert(v.Id, hhmmToMinutes(v.Start_time), hhmmToMinutes(v.End_time))
		}
		d.store[key] = tr
	}
	return d, nil
}

func (d *DataStore) Query(equipment string, qminRaw, qmaxRaw string) []SensorType {
	qmin := hhmmToMinutes(qminRaw)
	qmax := hhmmToMinutes(qmaxRaw)

	tr, ok := d.store[equipment]

	if !ok {
		return []SensorType{}
	}

	nodes := tr.Search(qmin, qmax)

	return treeNodesToSensor(nodes)
}
