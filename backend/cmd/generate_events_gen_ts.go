package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"media-organizer/backend/events"
)

func main() {
	f, err := os.Create("../../frontend/src/lib/events-gen.ts")

	if err != nil {
		log.Fatalf("could not create file: %v", err)
	}

	defer f.Close()

	// Marshal Go struct to JSON
	data, err := json.MarshalIndent(events.EventNames, "", "  ")

	if err != nil {
		log.Fatalf("failed to marshal event names: %v", err)
	}

	// Write to TS file as a const
	tmpl := template.Must(template.New("ts").Parse(`// This file is auto-generated from Go. Do not edit manually.
export const EventNames = {{ . }};
export type EventName = keyof typeof EventNames;
`))

	err = tmpl.Execute(f, string(data))
	if err != nil {
		log.Fatalf("failed to write template: %v", err)
	}
}
