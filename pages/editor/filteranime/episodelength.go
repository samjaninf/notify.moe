package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// EpisodeLength ...
func EpisodeLength(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without an episode length",
		func(anime *arn.Anime) bool {
			return anime.EpisodeLength == 0
		},
		nil,
	)
}
