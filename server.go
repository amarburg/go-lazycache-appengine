package main

import (
  "github.com/amarburg/go-lazycache"
  "net/http"
  "fmt"
)

var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"

func init() {
  fmt.Println("Do init")
  //DefaultLogger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))

  lazycache.ViperConfiguration()
  lazycache.ConfigureImageStoreFromViper()

  http.HandleFunc("/", lazycache.IndexHandler)
  lazycache.AddMirror(OOIRawDataRootURL)
}
