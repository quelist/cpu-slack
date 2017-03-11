package main

import (
    "os"
    "strconv"
    "time"
    "cpuobserver"
)


func main() {
    var DEFAULT_INTEVERAL int16 =1000

    if len(os.Args) > 1 {
        CUSTOM_INTEVERAL,_ := strconv.Atoi(os.Args[1])
        //sending health check request with custom interval
        cpuobserver.DoHealthCheckEvery( time.Duration( CUSTOM_INTEVERAL) * time.Millisecond )
    } else {
        //sending health check request with default interval
        cpuobserver.DoHealthCheckEvery( time.Duration( DEFAULT_INTEVERAL ) * time.Millisecond )
    }
    
}