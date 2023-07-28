package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"io/ioutil"
	"net/http"
	"strings"
)

func CheckRequireValid(ob interface{}) error {
	validator := validation.Validation{RequiredFirst: true}
	passed, err := validator.Valid(ob)
	if err != nil {
		return err
	}
	if !passed {
		var err string
		for _, e := range validator.Errors {
			err += fmt.Sprintf("[%s: %s] ", e.Field, e.Message)
		}
		return fmt.Errorf(err)
	}
	return nil
}

func GetMediaDirectory(mediaType, creatorId, filename string) (string, error) {
	// Get media upload directory
	// Folder: Audio, video, image, data, document
	directoryMap := map[string]string{
		"JPEG":    MediaTypeImage,
		"PNG":     MediaTypeImage,
		"JPG":     MediaTypeImage,
		"GIF":     MediaTypeImage,
		"MP4":     MediaTypeVideo,
		"YOUTUBE": MediaTypeVideo,
		"MP3":     MediaTypeAudio,
		"TXT":     MediaTypeData,
		"DOC":     MediaTypeData,
		"ZIP":     MediaTypeData,
		"CSV":     MediaTypeData,
		"PDF":     MediaTypeData,
		"SVG":     MediaTypeImage,
		"WEBM":    MediaTypeAudio,
	}

	mediaDir, ok := directoryMap[strings.ToUpper(mediaType)]
	if !ok {
		return "", fmt.Errorf("Media type not support")
	}
	return fmt.Sprintf("%s/%s/%s/%s", creatorId, mediaDir, filename), nil
}

func GetJarvanResponse(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
