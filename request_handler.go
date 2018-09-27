/*
Created by: Maulik Shah (mshah@redhat.com)
Date Created: 09/27/2018
Date last updated: 09/27/2018

File to define how various API endpoints are handled
*/


package main

import (
  "net/http"
  //"fmt"
)
func request_handler_main() {
  requestEndpoint := http.NewServeMux()

  /// Redirect requests to this endpoint to the associated function
  requestEndpoint.HandleFunc("/getlogs",getLogs)

  /// Port: 3000 to be exposed to handle all incoming requests
  http.ListenAndServe(":3000", requestEndpoint)
}
