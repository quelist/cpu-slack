package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "time"
    "net/http"
    "encoding/json"
    "bytes"
)

func getCPUSample() (idle, total uint64) {
    contents, err := ioutil.ReadFile("/proc/stat")
    if err != nil {
        return
    }
    lines := strings.Split(string(contents), "\n")
    for _, line := range(lines) {
        fields := strings.Fields(line)
        if fields[0] == "cpu" {
            numFields := len(fields)
            for i := 1; i < numFields; i++ {
                val, err := strconv.ParseUint(fields[i], 10, 64)
                if err != nil {
                    fmt.Println("Error: ", i, fields[i], err)
                }
                total += val // tally up all the numbers to get total ticks
                if i == 4 {  // idle is the 5th field in the cpu line
                    idle = val
                }
            }
            return
        }
    }
    return
}

func sendToSlack(cpuUsage float64){
   msg:= strconv.FormatFloat(cpuUsage, 'f', 6, 64)
   fmt.Println(msg)
  url := "https://hooks.slack.com/services/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"
  //exp := `{"text":"Buy cheese and bread for breakfast."}`
  //var message = []byte(exp)

  values := map[string]string{"text": msg}
  message, _ := json.Marshal(values)

  req, _ := http.NewRequest("POST", url, bytes.NewBuffer(message))

  req.Header.Add("content-type", "application/x-www-form-urlencoded")
  req.Header.Add("cache-control", "no-cache")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  fmt.Println(res)
  fmt.Println(string(body))
}

func main() {
    idle0, total0 := getCPUSample()
    time.Sleep(3 * time.Second)
    idle1, total1 := getCPUSample()

    idleTicks := float64(idle1 - idle0)
    totalTicks := float64(total1 - total0)
    cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
    sendToSlack(cpuUsage)

    fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}