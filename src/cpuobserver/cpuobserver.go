package cpuobserver

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "time"
    "slacksender"
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
func checkCPUonTime() float64 {
    idle0, total0 := getCPUSample()
    time.Sleep(3 * time.Second)
    idle1, total1 := getCPUSample()

    idleTicks := float64(idle1 - idle0)
    totalTicks := float64(total1 - total0)
    cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

    fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n\n", cpuUsage, totalTicks-idleTicks, totalTicks)
    return cpuUsage;

}

func DoHealthCheckEvery(d time.Duration) {
    for x := range time.Tick(d) {
        fmt.Println("Time:",x)
        cpuUsage := checkCPUonTime();
        if(cpuUsage > 75){
            cpu:= strconv.FormatFloat(cpuUsage, 'g', 1, 64)
            fmt.Println("Time:",x)
            webhook_url := "https://hooks.slack.com/services/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"
            slacksender.SendTOSlack(cpu, webhook_url)
        }
       
    }
}
