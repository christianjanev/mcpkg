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

type GetResult struct {
    Slug                  string `json:"slug"`
    Title                 string `json:"title"`
    Description           string `json:"description"`
    Categories            []string `json:"categories"`
    Client_side           string `json:"client_side"`
    Server_side           string `json:"server_side"`
    Body                  string `json:"body"`
    Status                string `json:"status"`
    Requested_status      string `json:"requested_status"`
    Additional_categories []string `json:"additional_categories"`
    Issues_url            string `json:"issues_url"`
    Source_url            string `json:"source_url"`
    Wiki_url              string `json:"wiki_url"`
    Discord_url           string `json:"discord_url"`
    Donation_urls         []string `json:"donation_urls"`
    Project_type          string `json:"project_type"`
    Downloads             json.Number `json:"downloads"`
    Icon_url              string `json:"icon_url"`
    Color                 json.Number `json:"color"`
    Thread_id             string `json:"thread_id"`
    Monetization_status   string `json:"monetization_status"`
    Id                    string `json:"id"`
    Team                  string `json:"team"`
    Body_url              string `json:"body_url"`
    Moderator_message     string `json:"moderator_message"`
    Published             string `json:"published"`
    Updated               string `json:"updated"`
    Approved              string `json:"approved"`
    Queued                string `json:"queued"`
    Followers             json.Number `json:"followers"`
    License               struct {
        Id   string `json:"id"`
        Name string `json:"name"`
        Url  string `json:"url"`
    } `json:"license"`
    Versions       []string `json:"versions"`
    Game_versions  []string `json:"game_versions"`
    Loaders        []string `json:"loaders"`
    Gallery        []struct {
        Url         string `json:"url"`
        Featured    bool `json:"featured"`
        Title       string `json:"string"`
        Description string `json:"description"`
        Created     string `json:"created"`
        Ordering    json.Number `json:"ordering"`
    } `json:"gallery"`
}

type VersionResult struct {
    Name string
    Version_number string
    Changelog string
    Dependencies []struct {
        Version_id string
        Project_id string
        File_name string
        Dependency_type string
    }
    Game_versions []string
    Version_type string
    Loaders []string
    Featured bool
    Status string
    Requested_status string
    Id string
    Project_id string
    Author_id string
    Date_published string
    Downloads json.Number
    Changelog_url string
    Files []struct {
        Hashes struct {
            Sha512 string
            Sha1 string
        }
        Url string
        Filename string
        Primary bool
        Size json.Number
        File_type string
    }
}

func modrinthSearch(query string) SearchResults {
	resp, _ := get(MODRINTH_ENDPOINT + "/search?query=" + query)
	
    bodyObj, _ := io.ReadAll(resp.Body)
    bodyStr := string(bodyObj)
    body := SearchResults{}

    err := json.Unmarshal([]byte(bodyStr), &body)

    if err != nil {
		log.Println(err.Error())
		log.Fatalln("modrinthSearch() json unmarshal failed.")
	}

    return body
}

func modrinthGet(slug string) GetResult {
    resp, _ := get(MODRINTH_ENDPOINT + "/project/" + slug)

    bodyObj, _ := io.ReadAll(resp.Body)
    bodyStr := string(bodyObj)
    body := GetResult{}

    err := json.Unmarshal([]byte(bodyStr), &body)

    if err != nil {
        log.Println(err.Error())
		log.Fatalln("modrinthGet() json unmarshal failed.")
    }

    return body
}

func modrinthGetVersion(slug string, loaders string, version string) VersionResult {
    resp, _ := get(MODRINTH_ENDPOINT + "/project/" + slug + "/version?loaders=" + loaders + "?game_versions=" + version)

    bodyObj, _ := io.ReadAll(resp.Body)
    bodyStr := string(bodyObj)
    body := []VersionResult{}

    err := json.Unmarshal([]byte(bodyStr), &body)

    if err != nil {
        log.Println(err.Error())
		log.Fatalln("modrinthGetVersion() json unmarshal failed.")
    }

    return body[0]
}

func modInstall(args []string) {
	log.Println("Installing modrinth mod for " + args[0])
}

func modSearch(args []string) {
	log.Println("Searching modrinth mod for " + args[0])

    result := modrinthSearch(args[0])

    for index := 0; index < len(result.Hits); index++ {
        log.Println(result.Hits[index].Title + " - " + (result.Hits[index].Versions[0] + "-" + result.Hits[index].Versions[len(result.Hits[index].Versions)-1]) + " - " + result.Hits[index].Slug)
    }

}
