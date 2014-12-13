package main

import (
	"flag"
	"time"
	"fmt"
	"os"
)


var options = &struct {
	configFile           string
	spoolSize            uint64
	dprofTimeout         time.Duration
	debug                bool
	ipPort               string
	logRedis             string
	logDynomite          string
}{
	spoolSize:            1,
	dprofTimeout:         time.Second * 15,
	debug:                false,
	ipPort:               "127.0.0.1:4000",
	//Defaults - unless set at commandline
        logRedis:             "/home/manish/dprof-collector/log.redis",
        logDynomite:          "/home/manish/dprof-collector/log.dynomite",
}


func main() {

	flag.StringVar(&options.configFile, "config", options.configFile,"");
	flag.Uint64Var(&options.spoolSize, "spool-size", options.spoolSize,"");
	flag.BoolVar(&options.debug, "debug", options.debug, "")
	flag.StringVar(&options.ipPort, "ipport", options.ipPort,"");
	flag.StringVar(&options.logRedis, "logredis", options.logRedis,"");
	flag.StringVar(&options.logDynomite, "logdynomite", options.logDynomite,"");

	flag.Parse()

	if options.configFile == "" {
		fmt.Println("fatal: Missing config file option -config=config-file-name")		
		os.Exit(1)
	}

	if options.debug {
		fmt.Println(options.configFile)
		fmt.Println(options.spoolSize)
		fmt.Println(options.dprofTimeout)
		fmt.Println(options.debug)
		fmt.Println(options.ipPort)
		fmt.Println(options.logRedis)
		fmt.Println(options.logDynomite)
	}

	/*
	fi, err := os.Stat(options.configFile)
	if err != nil {
		fmt.Println("fatal: Unable to find config file")
		os.Exit(1)
	}
         */

	//config := ReadConfig(options.configFile)

	//Upon restart - just read all log files entries with time stamps greater than present time
	//               ignore past entries
	//Steady state model - for each file in config, start a new logreader process
	//logreader process reads upto options.spoolSize entries and writes to the network
	//Future enhancement -- persist info so that can determine where to restart logreader for 
	//a log file.
	//For initial test - hardwire file names
	
	go LogSender(options.logDynomite,0)
	go LogSender(options.logRedis,1)
	var inp int
	fmt.Scanln(&inp)
	fmt.Println("done")
}

func debug_print(msgfmt string, args ...interface{}) {
	if !options.debug {
		return
	}
	fmt.Printf(msgfmt, args...)
}

