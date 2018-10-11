/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 09/27/2018

File to fetch the pod logs and serve them over a endpoint given some parameters
*/

package main

import (
  "net/http"
  "fmt"
  // "net/url"
  "crypto/tls"
  "bufio"
  "io/ioutil"
  "time"
  "strings"
  // "github.com/aws/aws-sdk-go/service/s3"
  // "github.com/aws/aws-sdk-go/aws/endpoints"
  "github.com/aws/aws-sdk-go/aws/session"
  )

func getLogs(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  name := r.URL.RawQuery
  // TODO: 
  w.Write([]byte(name))
}


func scrapeLogs(containerPtr []string, s3session *session.Session ) {
  url := "https://"+ *cluster + ":" + *clusterPort + "/api/v1/namespaces/"+containerPtr[0]+"/pods/" +containerPtr[1]+ "/log?timestamps=true&container="+containerPtr[2]
  lastTimeStamp := time.Now()
  for {
    lastTimeStamp = fetchLogs(containerPtr, url, lastTimeStamp)
    time.Sleep(120 * time.Second)
  }
}


func fetchLogs(containerPtr []string, url string, sinceTime time.Time) time.Time{
  errorFound := false
  RFC3339Nano := "2006-01-02T15:04:05.999999999Z"

  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
  utc,_ := time.LoadLocation("Etc/UTC")
  requesturl := url + "&sinceTime="+sinceTime.In(utc).Format(RFC3339Nano)
  fmt.Println(requesturl)
  req, _ := http.NewRequest("GET", requesturl , nil)
  req.Header.Add("Accept","application/json")
  req.Header.Add("Authorization", "Bearer "+*token)
  resp, _ := http.DefaultClient.Do(req)
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return sinceTime.In(utc)
  }

  reader := bufio.NewReader(resp.Body)
  lastTimeStamp := sinceTime.In(utc)

  for {
    line, err := reader.ReadString('\n')
    timeStampPosition := strings.Index(line,"Z")
    if timeStampPosition > 0 && len(line) > 0{
      lastTimeStamp,_ = time.Parse(RFC3339Nano, line[0:timeStampPosition+1])
    }


    if CaseInsensitiveContains(line, "error") || CaseInsensitiveContains(line, "warn") {
        errorFound = true
    }

    if err != nil {
      if errorFound {
        go writeLogsToCeph(url,lastTimeStamp)
      }
      return lastTimeStamp.Add(time.Nanosecond)
    }
  }

  return lastTimeStamp
}

func writeLogsToCeph(url string, sinceTime time.Time) {

  utc,_ := time.LoadLocation("Etc/UTC")
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
  RFC3339Nano := "2006-01-02T15:04:05.999999999Z"
  requesturl := url + "&sinceTime="+sinceTime.Add(-5*time.Minute).In(utc).Format(RFC3339Nano)

  req, _ := http.NewRequest("GET", requesturl , nil)
  req.Header.Add("Accept","application/json")
  req.Header.Add("Authorization", "Bearer "+*token)
  resp, _ := http.DefaultClient.Do(req)
  defer resp.Body.Close()

  body,_ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))

  _ = body
  if resp.StatusCode != http.StatusOK {
    fmt.Println("Error getting data to write to ceph")

  }

}
