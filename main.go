package main


import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Sensor struct {
	ID int `json:"id"`
	Planet string `json:"planet"`
	Mood string `json:"mood"`
}

type Measurements struct {
	Sensor sensors
}


func handler(w http.ResponseWriter, r *http.Request) {
	
	//conrtol the mood sniffing algorithm intensity
	beHappy := false
	
	sensorsWriteAPI := "http://mood-sensors.dev.dekt.io/activate"
	sensorsReadAPI := "http://mood-sensors.dev.dekt.io/measure"
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
        fmt.Fprintf(w, "<H1><font color='navy'>Welcome to the DevX Mood Analyzer </font></H1><H2>")

	if !beHappy { 
		fmt.Fprintf(w, "<font color='red'>")
		fmt.Fprintf(w,"Your overall mood is not great. We hope it will get better.")
		fmt.Fprintf(w, "</font>")
		fmt.Fprintf(w, "<BR><BR><img src='https://raw.githubusercontent.com/dektlong/devx-mood/main/sad-dog.jpg' alt=''>")
		fmt.Fprintf(w, "</H2>")
		fmt.Fprintf(w, "<BR><font color='brown'>Aggressive mood sniffing algorithm</font><BR>")
	} else { //always happy
		fmt.Fprintf(w, "<font color='green'>")
		fmt.Fprintf(w,"Your mood is always happy. Good for you!")
		fmt.Fprintf(w, "</font>")
		fmt.Fprintf(w, "<BR><BR><img src='https://raw.githubusercontent.com/dektlong/devx-mood/main/happy-dog.jpg' alt=''>")
		fmt.Fprintf(w, "</H2>")
		fmt.Fprintf(w, "<BR><font color='brown'>Mild mood sniffing algorithm</font><BR>")

	}	
	
	//activate sensors
	fmt.Fprintf(w, "<BR><BR>")
	fmt.Fprintf(w, "<font color='purple'>/activate:	</font><BR>")
	fmt.Fprintf(w, "<font color='gray'>")
	for i := 1; i < 11; i++ {
		http.Get(sensorsWriteAPI)
	}
	fmt.Fprintf(w, "sensors activated on all planets")
	//fmt.Fprintf(w, "[{\"sensorsStatus\":\"activated\"}]")
	fmt.Fprintf(w, "</font><BR>")
	
	//collect measurements from sensors
	fmt.Fprintf(w, "<font color='purple'>/measure: </font><BR>")
	fmt.Fprintf(w, "<font color='gray'>")
	response, err := http.Get(sensorsReadAPI)
	if err != nil {
		fmt.Fprintf(w,"ERROR! in calling measure API")
	} 

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	
	var measurements Measurements
	if err := json.Unmarshal(responseData, &measurements); err != nil {  // Parse []byte to the go struct pointer
        fmt.Fprintf(w,"Can not unmarshal JSON")
    }
		
	fmt.Fprintf(w,"%+v\n", measurements)
    
			//fmt.Fprintf(w, sensor.mood)
	//fmt.Fprintf(w,string(responseData))
	fmt.Fprintf(w, "</font>")
}

func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}


