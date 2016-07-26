package bardo

import (
  "regexp"
)

var dbNameR = regexp.MustCompile(`dbname\=([^\s]+?)\s`)

// GetDBNameFromURL
// Get the database name from the open string
func GetDBNameFromURL(url string) string {
	// Find database name (that we will create)
	m := dbNameR.FindStringSubmatch(url)
	return m[1]
}

// ReplaceDBNameInURL
// Replace the database name in the open string
func ReplaceDBNameInURL(url string, dbname string) string {
	// Find database name (that we will create)
	return dbNameR.ReplaceAllString(url, "dbname="+dbname+" ")
}
