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
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
	// "sync/atomic"
	// "os"
	// "fmt"
	// "log"
	// "strings"
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
