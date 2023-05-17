package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"

	"fmt"
	"os"
)

var configFile *string = flag.String("config", "config", "Config File")
var config map[string]string = make(map[string]string)

type ConfigSettings struct {
	xip               string
	workfolder        string
	sourcefolder      string
	destinationfolder string
	mapfolder         string
	schemafolder      string
	exefile           string
}

//----------------------------------------------------------------
func main() {
	flag.Parse()
	cfgset := &ConfigSettings{}
	cfgset.xip = "localhost"
	cfgset.workfolder = "/qa"
	cfgset.sourcefolder = "/sourcefiles"
	cfgset.destinationfolder = "/destinationfiles"
	cfgset.mapfolder = "/mapfiles"
	cfgset.schemafolder = "/schemafiles"
	cfgset.exefile = "/C1D0U484/C1D0U484.EXE"
	//------------------------------------------------------------
	fmt.Printf("EDI Tool - (c) Copyright Com1Software 1992-2023\n")
	fmt.Printf("Repository : github.com/Com1Software/EDI-Tool\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//--------------------------------------------------------- Read Config
	cfgFile, errRF := ioutil.ReadFile(*configFile)
	if errRF != nil {
		fmt.Printf("Error Reading Config %s\n", errRF)
	}
	//--------------------------------------------------------- Unmarshal Config
	errUM := json.Unmarshal(cfgFile, &config)
	if errUM != nil {
		fmt.Printf("Error UnMarshalling Config %s\n", errRF)
	}
	//--------------------------------------------------------- Start Logging
	logFile, errLF := os.OpenFile(config["logfile"], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if errLF != nil {
		fmt.Printf("Error Reading Config %s\n", errRF)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.Printf("Logging Started")
	//---------------------------------------------------------
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:
		fmt.Println("Command:")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on port 8080")
		fmt.Println("")
		//------------------------------------------------ Projects Page Handler
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := InitPage(*cfgset)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage(*cfgset)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Settings Page Handler
		http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
			xdata := SettingsPage(*cfgset)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Projects Page Handler
		http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
			xdata := ProjectsPage(*cfgset)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Examples Page Handler
		http.HandleFunc("/examples", func(w http.ResponseWriter, r *http.Request) {
			xdata := ExamplesPage(*cfgset)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------- Static Handler Handler
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
		//------------------------------------------------- Start Server
		Openbrowser("http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}
}

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

//----------------------------------------------------- Init Page
func InitPage(cfgset ConfigSettings) string {
	xip := cfgset.xip
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>EDI Tool</title>"
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>EDI Tool</H1>"
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xxip = fmt.Sprintf("%s", ipv4)
		}
	}
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<p> Host Port IP : " + xip
	xdata = xdata + "<BR> Machine IP : " + xxip + "</p>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/examples'> [ Examples ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/projects'> [ Projects ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/settings'> [ Settings ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/static/index.html'> [ Documentation ] </A>  "
	xdata = xdata + "<BR><BR><small>EDI Tool -  (c) Copyright Com1Software 1992-2023</small>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

//--------------------------------- Settings
func SettingsPage(cfgset ConfigSettings) string {
	xip := cfgset.xip
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>EDI Tool Projects</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>EDI Tool Projects</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<br><br>"
	xdata = xdata + "<A HREF='http://" + xip + ":8080'> [ Main Menu ] </A>  "
	xdata = xdata + "<BR><BR><small>EDI Tool -  (c) Copyright Com1Software 1992-2023</small>"
	xdata = xdata + "</body>"
	xdata = xdata + "</html>"
	return xdata

}

//--------------------------------- Examples
func ExamplesPage(cfgset ConfigSettings) string {
	xip := cfgset.xip
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>EDI Tool Projects</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>EDI Tool Projects</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<br><br>"
	xdata = xdata + "<A HREF='http://" + xip + ":8080'> [ Main Menu ] </A>  "
	xdata = xdata + "<BR><BR><small>EDI Tool -  (c) Copyright Com1Software 1992-2023</small>"
	xdata = xdata + "</body>"
	xdata = xdata + "</html>"
	return xdata

}

//--------------------------------- Projects
func ProjectsPage(cfgset ConfigSettings) string {
	xip := cfgset.xip
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>EDI Tool Projects</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>EDI Tool Projects</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<br><br>"
	xdata = xdata + "<A HREF='http://" + xip + ":8080'> [ Main Menu ] </A>  "
	xdata = xdata + "<BR><BR><small>EDI Tool -  (c) Copyright Com1Software 1992-2023</small>"
	xdata = xdata + "</body>"
	xdata = xdata + "</html>"
	return xdata

}

//--------------------------------- About
func AboutPage(cfgset ConfigSettings) string {
	xip := cfgset.xip
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>EDI Tool Projects</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>EDI Tool About</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<br><br>"
	xdata = xdata + "<A HREF='http://" + xip + ":8080'> [ Main Menu ] </A>  "
	xdata = xdata + "<BR><BR><small>EDI Tool -  (c) Copyright Com1Software 1992-2023</small>"
	xdata = xdata + "</body>"
	xdata = xdata + "</html>"
	return xdata

}

func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm.' : 'am.';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtdt').innerHTML = ' Date : '+day+', '+month+' '+dm+', '+yr+'.  - Time : '+h + ':' + m + ':' + s+' '+ampm;"
	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}
