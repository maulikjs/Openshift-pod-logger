/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 09/27/2018
Date last updated: 09/27/2018

Main entrypoint file for the service
*/


package main

import (
	// "fmt"
	// "log"
	"flag"
	)

// var logger = log.GetLogger()
var token *string
var fileptr *string
var cluster *string

func main() {

	// logger.Log("Starting Service...")

	// Fetch Command Line arguements (Format: argName, Default, description)
	token := flag.String("token","1234", "Service Account Token to read cluster data")
	fileptr := flag.String("fileptr", "./containers.csv", "Pointer to file with the list of pods to monitor")
	cluster := flag.String("cluster", "https://datahub.upshift.redhat.com", "Pointer to the Openshift cluster")
	clusterPort := flag.String("clusterport", "443", "Port on which to connect the openshift cluster")
	// Parse the command line arguements
	flag.Parse()

	//Start the service and service REST endpoints
	request_handler_main()
}
