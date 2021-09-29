package datastore

import "encoding/json"

type SensorType struct {
	Id         string
	Start_time string
	End_time   string
}

type EquipmentsMappingType map[string][]SensorType

type PlantsData struct {
	Plants            []string
	Equipments        []string
	Sensors           []string
	PlantsMapping     map[string][]string
	EquipmentsMapping EquipmentsMappingType
}

func getMapping(raw []byte) (PlantsData, error) {
	result := PlantsData{}
	err := json.Unmarshal(raw, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
