package editor

import (
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxImageEntries = 70

// LowResolutionAnimeImages ...
func LowResolutionAnimeImages(ctx *aero.Context) string {
	basePath := path.Join(arn.Root, "images/anime/original/")
	files, err := ioutil.ReadDir(basePath)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error reading anime images directory", err)
	}

	lowResAnime := []*arn.Anime{}

	for _, file := range files {
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fullPath := path.Join(basePath, file.Name())
		width, height, _ := getImageDimensions(fullPath)

		if width < arn.AnimeImageLargeWidth*2 || height < arn.AnimeImageLargeHeight*2 {
			animeID := file.Name()
			animeID = strings.TrimSuffix(animeID, filepath.Ext(animeID))

			anime, err := arn.GetAnime(animeID)

			if err == nil {
				lowResAnime = append(lowResAnime, anime)
			}
		}
	}

	count := len(lowResAnime)

	if count > maxImageEntries {
		lowResAnime = lowResAnime[:maxImageEntries]
	}

	return ctx.HTML(components.AnimeEditorListFull(
		"Anime with low resolution images",
		lowResAnime,
		count,
		"/editor/anime/missing/hiresimage",
		func(anime *arn.Anime) string {
			return "https://www.google.com/search?q=" + anime.Title.Canonical + "&tbm=isch"
		},
	))
}

// getImageDimensions retrieves the dimensions for the given file path.
func getImageDimensions(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	defer file.Close()

	if err != nil {
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)

	if err != nil {
		return 0, 0, err
	}

	return image.Width, image.Height, nil
}
