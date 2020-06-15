package main

import ("fmt"
"os"
log "github.com/sirupsen/logrus"
)

func init(){
	f, err := os.OpenFile("D:/vivek/work/go/goquickstart/app.log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
    Formatter := new(log.TextFormatter)
    // You can change the Timestamp format. But you have to use the same date and time.
    // "2006-02-02 15:04:06" Works. If you change any digit, it won't work
    // ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
    log.SetFormatter(Formatter)
    if err != nil {
        // Cannot open log file. Logging to stderr
        fmt.Println(err)
    }else{
        log.SetOutput(f)
    }
}

/*func main(){
	//why dependecy
	//what is dependency 
	//problem with dependency management in go
	//go module how helps
	log.Info("+++Process Started+++")
	log.Info("+++Process logs+++")
	fmt.Println("Hello Go!!!")
	log.Info("+++Process Completed+++")
}*/

