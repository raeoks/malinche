package malinche

import "time"

var (
	// Version - Malinche Version
	Version = "0.0.1"
	// Build - Malinche Build Date
	Build = time.Now().UTC().Format(time.UnixDate)
)
