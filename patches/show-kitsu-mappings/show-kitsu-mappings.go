package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	for mapping := range arn.StreamKitsuMappings() {
		fmt.Printf(
			"Kitsu Anime %s: %s mapped to %s\n",
			mapping.Relationships.Item.Data.ID,
			color.YellowString(mapping.Attributes.ExternalSite),
			color.GreenString(mapping.Attributes.ExternalID),
		)
	}
}
