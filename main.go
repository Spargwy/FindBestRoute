package main

import (
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
		log.Fatal("Cant parse csv file: ", err)
	}
	fmt.Println("\nCalculating...")
	findBestPathAndPrintIt(timetable, filter)

}

func findBestPathAndPrintIt(timetable []TimetableField, filter string) error {
	var path Path
	var result []Path
	graph, uniqueStations := setGraphAndUniqueStations(timetable)
	for i := range uniqueStations {
		for j := range graph[uniqueStations[i]] {
			findAllPaths(graph, graph[uniqueStations[i]][j], path, &result)
		}
	}
	fmt.Printf("Total %d routes\n\n", len(result))
	bestPaths, err := CalculateBestPath(result, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	resultMessages(bestPaths, filter)

	return nil
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
					"Time:", station.FromTime, station.ToTime,
					"=", station.ElapsedTime, "minutes, price:", station.Cost)
			}
		}
	}
}
