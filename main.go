package main

import (
  "github.com/amarburg/go-lazycache"
//  kitlog "github.com/go-kit/kit/log"
//  "os"
  "net/http"
)

var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"


func init() {
  //defaultLogger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))

	http.HandleFunc("/", lazycache.IndexHandler)

//  ConfigureImageStore(*image_store, *google_bucket, defaultLogger)
  // server := StartLazycacheServer(*bind, *port)
  // defer server.Stop()

  lazycache.AddMirror(OOIRawDataRootURL)
}
