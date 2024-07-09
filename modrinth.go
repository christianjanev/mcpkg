package main

import (
	"encoding/json"
	"io"
	"log"
)

const MODRINTH_ENDPOINT string = "https://api.modrinth.com/v2" 

type SearchResult struct {
    Project_id			string
    Project_type		string
    Slug				string
    Author				string
    Title				string
    Description			string
    Categories			[]string
    Display_categories  []string
    Versions			[]string
    Downloads			json.Number
    Follows				json.Number
    Icon_url			string
    Date_created		string
    Date_modified		string
    Latest_version		string
    License				string
    Client_side			string
    Server_side			string
    Gallery				[]string
    Featured_gallery	string
    Color				json.Number
}

type SearchResults struct {
	Hits	  []SearchResult
	Offset	  json.Number
	Limit 	  json.Number
	Total_hit json.Number
}

func modrinthSearch(query string) (SearchResults, error) {
	resp, _ := get(MODRINTH_ENDPOINT + "/search?query=" + query)
	
    bodyObj, _ := 
}

func modInstall(args []string) {
	log.Println("Installing modrinth mod for " + args[0])
}

func modSearch(args []string) {
	log.Println("Searching modrinth mod for " + args[0])
}
