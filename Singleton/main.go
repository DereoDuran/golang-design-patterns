package main

import "sync"

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		capitals := map[string]int{
			"Seoul":   9708483,
			"Busan":   3403135,
			"Incheon": 2632035,
			"Daegu":   2466052,
			"Gwangju": 1502881,
		}
		db.capitals = capitals
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	println("Population of Seoul = ", pop)

	names := []string{"Seoul", "Busan"}
	tp := GetTotalPopulation(names)
	println("TP = ", tp)

	tp = GetTotalPopulationEx(GetSingletonDatabase(), names)
	println("TP = ", tp)

	db2 := DummyDatabase{}
	tp = GetTotalPopulationEx(&db2, []string{"alpha", "gamma"})
	println("TP = ", tp)
}
