package common

const API_Key = "anAPIKey"

const (
	STATUS_UNASSIGNED = 0
	STATUS_TAKEN      = 1
	STATUS_FINISHED   = 2
)

var statusText = map[int]string{
	STATUS_UNASSIGNED: "UNASSIGNED",
	STATUS_TAKEN:      "TAKEN",
	STATUS_FINISHED:   "FINISHED",
}
