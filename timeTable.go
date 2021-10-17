package main

type TimetableField struct {
	Train       string
	FromStation string
	ToStation   string
	FromTime    string
	ToTime      string
	WaitingTime string
	ElapsedTime int64
	Cost        float64
}

type Path struct {
	Stations    []TimetableField
	Price       float64
	ElapsedTime int64
}
