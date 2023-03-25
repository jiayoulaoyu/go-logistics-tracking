package model

type Event struct {
	Date     string `json:"xDate"`
	Location string `json:"xPlace"`
	Message  string `json:"xInfo"`
}
