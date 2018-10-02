/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 10/01/2018

To write all the small helper functions
*/


package main

import (
	// "fmt"
	"log"
  "bufio"
  "os"
  "strings"
	)

// Read the .csv file with the namespace,pod,container name from the fileptr
// and split it into a 2d string array
func readCSV(fileptr string) [][]string {

  var values [][]string

  file, err := os.Open(fileptr)
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan()  {
    vals := strings.Split(scanner.Text(),",")
    values = append(values,vals)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return values
}


func CaseInsensitiveContains(s, substr string) bool {
    s, substr = strings.ToUpper(s), strings.ToUpper(substr)
    return strings.Contains(s, substr)
}
