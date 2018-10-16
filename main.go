/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 09/27/2018

Main entrypoint file for the service
*/


package main

import (
	"fmt"
	// "log"
	"flag"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	)

// var logger = log.GetLogger()
var token *string
var fileptr *string
var cluster *string
var clusterPort *string
var valsFromCSV [][]string
var s3instance *string
var bucketName *string
var S3session *session.Session


func main() {

	// logger.Log("Starting Service...")

	// Fetch Command Line arguements (Format: argName, Default, description)
	token = flag.String("token","", "Service Account Token to read cluster data")
	fileptr = flag.String("fileptr", "./containers.csv", "Pointer to file with the list of pods to monitor")
	cluster = flag.String("cluster", "datahub.upshift.redhat.com", "Pointer to the Openshift cluster")
	clusterPort = flag.String("clusterport", "443", "Port on which to connect the openshift cluster")
	bucketName = flag.String("s3bucket", "DH-STAGE-LOGS", "S3 bucket to push logs to")
	// Only works if you have run "aws configure" and have set the access keys
	s3instance = flag.String("s3instance", "https://s3.upshift.redhat.com", "URL of the s3 s3 (CEPH) instance on which to store logs")
	// Parse the command line arguements
	flag.Parse()

	S3session = createSession(*s3instance)



	valsFromCSV := readCSV(*fileptr)
	for _, element := range valsFromCSV {
		fmt.Println(element)
		go scrapeLogs(element)
	}

	//Start the service and service REST endpoints
	request_handler_main()
}


//dh-prod-elastalert,elastalertcore-1-lfdpm,elastalertcore
