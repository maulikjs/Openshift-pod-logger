/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 10/08/2018

File to define the s3 connectors and functions
*/


package main

import (
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
	// "sync/atomic"
	// "os"
	"fmt"
	// "log"
  "reflect"
  "strconv"
  "time"
  "strings"
)

func createSession(customEndpoint string) *session.Session {
  defaultResolver := endpoints.DefaultResolver()
  s3CustResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
    if service == "s3" {
      return endpoints.ResolvedEndpoint{
        URL:           customEndpoint,
        SigningRegion: "custom-signing-region",
      }, nil
    }

    return defaultResolver.EndpointFor(service, region, optFns...)
  }
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    Config: aws.Config{
      EndpointResolver: endpoints.ResolverFunc(s3CustResolverFn),
    },
  }))

  return sess
}


func pushToCeph(containerPtr []string, body []uint8, timestamp time.Time){

  fmt.Println(containerPtr)
  // fmt.Println(string(body))
  fmt.Println(timestamp)
  fmt.Println(reflect.TypeOf(S3session))

  key := strconv.Itoa(timestamp.Year()) +"-"+ strconv.Itoa(int(time.Now().Month())) +"-"+ strconv.Itoa(timestamp.Day()) +"/"+ containerPtr[0] +"/"+ containerPtr[2] +"/"+ timestamp.Format("15.04.05") + ".logs"

  fmt.Println(key)

	svc := s3.New(S3session)

  input := &s3.PutObjectInput{
      Body:                 aws.ReadSeekCloser(strings.NewReader(string(body))),
      Bucket:               aws.String("DH-STAGE-LOGS"),
      Key:                  aws.String(key),
      // ServerSideEncryption: aws.String("AES256"),
      // StorageClass:         aws.String("STANDARD_IA"),
  }

  result, err := svc.PutObject(input)
  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
      return
  }

  fmt.Println(result)
}
