package main

import "fmt"

func CalculateBestPath(routes []Path, filter string) (bestPaths []Path, err error) {
	switch filter {
	case "price":
		bestPaths, err = CalculateCheapestPath(routes, filter)
		if err != nil {
			return
		}
	case "time":
		bestPaths, err = CalculateFasterPath(routes, filter)
		if err != nil {
			return
		}
	default:
		fmt.Println("choose correct filter")
		return
	}
	return
}

func CalculateCheapestPath(routes []Path, filter string) (cheapestPaths []Path, err error) {
	for i := range routes {
		routes[i].ElapsedTime, err = CalculatePathTime(routes[i].Stations)
		if err != nil {
			return cheapestPaths, err
		}

	}
	for i := range routes {
		routes[i].Price, err = CalculatePathPrice(routes[i].Stations)
		if err != nil {
			return cheapestPaths, err
		}
	}
	sorted := MergeSort(routes, filter)
	cheapestPaths = append(cheapestPaths, sorted[0], sorted[1], sorted[2])
	return cheapestPaths, nil
}

func CalculateFasterPath(routes []Path, filter string) (fasterPaths []Path, err error) {
	for i := range routes {
		routes[i].Price, err = CalculatePathPrice(routes[i].Stations)
		if err != nil {
			return fasterPaths, err
		}
	}
	for i := range routes {
		routes[i].ElapsedTime, err = CalculatePathTime(routes[i].Stations)
		if err != nil {
			return fasterPaths, err
		}

	}
	sorted := MergeSort(routes, filter)
	fasterPaths = append(fasterPaths, sorted[0], sorted[1], sorted[2])
	return fasterPaths, nil
}

func CalculatePathTime(path []TimetableField) (pathTime int64, err error) {
	for i := range path {
		time, err := ParseTimeToMinutes(path[i].FromTime, path[i].ToTime)
		if err != nil {
			return pathTime, err
		}
		pathTime += time
		if i > 0 {
			waitingTime, err := ParseTimeToMinutes(path[i-1].ToTime, path[i].FromTime)
			if err != nil {
				return pathTime, err
			}
			pathTime += waitingTime
		}
	}
	return
}

func CalculatePathPrice(path []TimetableField) (pathPrice float64, err error) {
	for i := range path {
		pathPrice += path[i].Cost
	}
	return
}
