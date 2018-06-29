package main

import(
  "fmt"
  "log"
  "net/http"
  "net/url"
  "encoding/json"
)

type iptolocationinfo struct {
  As string `json:"as"`
  City string `json:"city"`
  Country string `json:"country"`
  CountryCode string `json:"countryCode"`
  Isp string `json:"isp"`
  Lat float64 `json:"lat"`
  Lon float64 `json:"lon"`
  Org string `json:"org"`
  Query string `json:"query"`
  Region string `json:"region"`
  RegionName string `json:"regionName"`
  Status string `json:"status"`
  Timezone string `json:"timezone"`
  Zip string `json:"zip"`
}

func main(){
  ip := "8.8.8.8"
  //escape the ip
  safeip := url.QueryEscape(ip)

  //build a url string
  url := fmt.Sprintf("http://ip-api.com/json/%s",safeip)
  fmt.Println("URL :\n",url)

  //build request
  req,err := http.NewRequest("GET",url,nil)
  if err != nil {
      log.Fatal("New Request: ",err)
      return
  }

  //create client
    client := &http.Client{}

  //send request via client
     response, err := client.Do(req)
     if err != nil{
       log.Fatal("Response Error: ",err)
       return
     }

  //clsoe the request
    defer response.Body.Close()

  //store response to iptolocationinfotype
    var ipinfo iptolocationinfo

   //create a json decoder
    jsondecoder := json.NewDecoder(response.Body)

    //docode json to your defined struct
      err = jsondecoder.Decode(&ipinfo)
      if err != nil{
        log.Fatal("Decoding Error: ",err)
      }

      fmt.Println("City: ",ipinfo.City)
      fmt.Println("Time Zone: ",ipinfo.Timezone)

}
