package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/torabit/itunes"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Look up Kaho Nakamura's NIA by iTunes ID: 1611115294\n")
	buf.WriteString("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	ctx := context.Background()

	client := itunes.New()
	results, err := client.Lookup(ctx, itunes.ID(1611115294))

	if err != nil {
		log.Fatal(err)
	}

	if len(results.Results) == results.ResultCount {
		for _, item := range results.Results {
			v := reflect.ValueOf(item)
			t := v.Type()

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
