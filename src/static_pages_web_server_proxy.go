package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "io"
        "io/ioutil"
        "encoding/json"
)

var JSON_FILE = "/root/go_web_server_proxy/conf/file_assignment.json"
//JSON format : 
//[{
//                "source": "get1.mpd",
//                "destination": "/root/go_web_server_proxy/conf/files/get1.mpd"
//        },
//        {
//                "source": "2700000.ts",
//                "destination": "/root/go_web_server_proxy/conf/files/2700000.ts"
//        }
//]


var assignment_map map[string]string

type SrcDestPair struct {
        Source string `json:"source"`
        Destination string `json:"destination"`
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path[strings.LastIndex(r.URL.Path, "/") + 1:])
        if val, ok := assignment_map[r.URL.Path[strings.LastIndex(r.URL.Path, "/") + 1:]]; ok {

                Openfile, err := os.Open(val)
                defer Openfile.Close()
                if err != nil {
                        http.Error(w, "File not found.", 404)
                        return
                }
                FileHeader := make([]byte, 512)
                Openfile.Read(FileHeader)
                Openfile.Seek(0, 0)
                io.Copy(w, Openfile)
        } else {
                fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
        }
}

func loadJson() {
        jsonFile, err := os.Open(JSON_FILE)
        if err != nil {
                fmt.Println("Getting error")
                fmt.Println(err)
        }
        fmt.Println("Successfully Opened file_assignment.json")
        defer jsonFile.Close()
        byteValue, _ := ioutil.ReadAll(jsonFile)
        assignment_map = make(map[string]string)
        var assignment []SrcDestPair
        json.Unmarshal(byteValue, &assignment)
        for i := 0; i < len(assignment); i++ {
                fmt.Println("Source " + assignment[i].Source)
                fmt.Println("Destination " + assignment[i].Destination)
                assignment_map[assignment[i].Source] = assignment[i].Destination
        }
        fmt.Println("map : ", assignment_map)
}

func main() {
        loadJson()
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":12345", nil))
}
