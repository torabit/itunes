package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/torabit/itunes"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("")
	buf.WriteString("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	ctx := context.Background()

	client := itunes.New(&http.Client{})
	results, err := client.Search(
		ctx,
		itunes.Term("Helsinki Lambda Club"),
		itunes.Media(itunes.MediaMusic),
		itunes.Attribute(itunes.AttributeArtistTerm),
		itunes.Country(itunes.CountryJP),
		itunes.Entity(itunes.EntityMusicTrack),
		itunes.Limit(2),
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(results.Results) == results.ResultCount {
		for index, item := range results.Results {
			v := reflect.ValueOf(item)
			t := v.Type()

			if results.ResultCount > 1 && index > 0 {
				buf.WriteString("┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈\n")
			}

			for i := 0; i < v.NumField(); i++ {
				fieldName := t.Field(i).Name
				fieldValue := v.Field(i).Interface()
				buf.WriteString(fmt.Sprintf("┃ %v: %v\n", fieldName, fieldValue))
			}
		}
	}
	buf.WriteString("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Println(buf.String())
}
