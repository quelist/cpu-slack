package slacksender

import (
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func SendTOSlack(msg, url string)  string{
  values := map[string]string{"text": msg}
  message, _ := json.Marshal(values)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
  if err != nil {
  return "Some error occured"
  }

  req.Header.Add("content-type", "application/x-www-form-urlencoded")
  req.Header.Add("cache-control", "no-cache")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  return string(body)
}