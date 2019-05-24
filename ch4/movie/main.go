// package description
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color",omitempty`
	Actors []string `json:"actors"`
}

// type Movie struct {
// 	Title  string
// 	Year   int  `json:"released"`
// 	Color  bool `json:"color",omitempty`
// 	Actors []Actor
// }

// type Actor struct {
// 	Name string `json:"name"`
// }

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// var movies = []Movie{
	// 	{Title: "Casablanca", Year: 1942, Color: false,
	// 		Actors: []Actor{{"Humphrey Bogart"}, {"Ingrid Bergman"}}},
	// 	{Title: "Cool Hand Luke", Year: 1967, Color: true,
	// 		Actors: []Actor{{"Paul Newman"}}},
	// 	{Title: "Bullitt", Year: 1968, Color: true,
	// 		Actors: []Actor{{"Steve McQueen"}, {"Jacqueline Bisset"}}},
	// }

	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %v\n", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct {
		Title string `json:"title"`
	}

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %v\n", err)
	}
	fmt.Println(titles)

}