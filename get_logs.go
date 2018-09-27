/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 09/27/2018
Date last updated: 09/27/2018

File to fetch the pod logs and serve them over a endpoint given some parameters
*/

package main

import "net/http"


func getlogs(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  name := r.URL.RawQuery
  w.Write([]byte(name))
}
