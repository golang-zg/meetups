package main

type Person struct {
	Name             string `json:"name"`
	AttendingMeetups bool
	// Collect all unknown Person fields
	Data map[string]any `json:",unknown"`
}
