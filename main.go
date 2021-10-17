package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Choose filter: price or time")
	}
	filter := args[1]
	timetable, err := ParseCSVFile()
	if err != nil {
		log.Println("Cant parse csv file: ", err)
	}
	findBestPathAndPrintIt(timetable, filter)

}

func findBestPathAndPrintIt(timetable []TimetableField, filter string) {
	var path Path
	var result []Path
	graph, uniqueStations := setGraphAndUniqueStations(timetable)
	for i := range uniqueStations {
		for j := range graph[uniqueStations[i]] {
			findAllPaths(graph, graph[uniqueStations[i]][j], path, &result)
		}
	}
	bestPaths, err := CalculateBestPath(result, filter)
	if err != nil {
		log.Println(err)
	}
	resultMessages(bestPaths, filter)

}

func setGraphAndUniqueStations(timetable []TimetableField) (graph map[string][]TimetableField, uniqueStations []string) {
	graph = make(map[string][]TimetableField)
	for i := range timetable {
		graph[timetable[i].FromStation] = append(graph[timetable[i].FromStation], timetable[i])
	}
	for key := range graph {
		uniqueStations = append(uniqueStations, key)
	}
	return graph, uniqueStations
}

func findAllPaths(graph map[string][]TimetableField, fromStation TimetableField, path Path, result *[]Path) {
	path.Stations = append(path.Stations, fromStation)

	availableRoutes := graph[fromStation.ToStation]
	for i := range availableRoutes {
		if !ElementInArray(availableRoutes[i].ToStation, path.Stations) {
			findAllPaths(graph, availableRoutes[i], path, result)
		}
	}

	if len(path.Stations) == len(graph)-1 {
		*result = append(*result, path)
	}
}

func ParseCSVFile() (timetable []TimetableField, err error) {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Println("Cant open csv file: ", err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return timetable, err
	}
	for i := range records {
		timetableField, err := ParseTimetableField(records[i][0])
		if err != nil {
			return timetable, err
		}
		timetable = append(timetable, timetableField)
	}
	return timetable, nil
}

func resultMessages(bestPaths []Path, filter string) {
	if filter == "price" {
		fmt.Println("Cheapest paths: ")
		for i := range bestPaths {
			fmt.Printf("\n\n Total time: %d\n Total price: %f\n", bestPaths[i].ElapsedTime, bestPaths[i].Price)
			for j := range bestPaths[i].Stations {
				station := bestPaths[i].Stations[j]
				fmt.Println(station.Train, ":", station.FromStation, "->", station.ToStation,
					"Time:", station.FromTime, station.ToTime, "=", station.ElapsedTime, "minutes, price:", station.Cost)
			}
		}
	} else if filter == "time" {
		fmt.Println("Fastest pathes")
		for i := range bestPaths {
			fmt.Printf("\n\n Total time: %d\n Total price: %f\n", bestPaths[i].ElapsedTime, bestPaths[i].Price)
			for j := range bestPaths[i].Stations {
				station := bestPaths[i].Stations[j]
				fmt.Println(station.Train, ":", station.FromStation, "->", station.ToStation,
					"Time:", station.FromTime, station.ToTime, "=", station.ElapsedTime, "minutes, price:", station.Cost)
			}
		}
	}
}
