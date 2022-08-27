package main

import (
	"fmt"
	"log"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/inputs"
)

var client *goobs.Client

func main() {
	var err error
	for {
		client, err = goobs.New("localhost:4455",
			goobs.WithPassword("from-websocket-ui"),
		)
		if err != nil {
			log.Println("goobs.New() err:", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}
	go client.Listen(obsEventsErrs)

	log.Println(play("Media Source", "/home/username/Videos/file.mp4", true))
	log.Println(play("VLC Video Source", "/home/username/Videos/file.mp4", true))

	time.Sleep(10 * time.Second)

	log.Println(play("Media Source", "/home/username/Videos/file.mp4", true))
	log.Println(play("VLC Video Source", "/home/username/Videos/file.mp4", true))

	time.Sleep(10 * time.Minute)
}

func obsEventsErrs(e interface{}) {
	log.Printf("events/errs %T %v\n", e, e)
}

func play(sourceName, input string, local bool) error {
	gisRes, err := client.Inputs.GetInputSettings(&inputs.GetInputSettingsParams{
		InputName: sourceName,
	})
	if err != nil {
		return err
	}

	settings := gisRes.InputSettings
	switch gisRes.InputKind {
	case "ffmpeg_source":
		setFFmpeg(settings, input, local)
	case "vlc_source":
		setVLC(settings, input)
	default:
		err := fmt.Errorf("unknown source type %q for %q",
			gisRes.InputKind, sourceName)
		return err
	}

	_, err = client.Inputs.SetInputSettings(&inputs.SetInputSettingsParams{
		InputName:     sourceName,
		InputSettings: settings,
	})
	if err != nil {
		return err
	}
	return nil
}

func setVLC(settings map[string]interface{}, input string) {
	settings["playlist"] = []map[string]interface{}{
		{
			"hidden":   false,
			"selected": true,
			"value":    input,
		},
	}
}

func setFFmpeg(settings map[string]interface{}, input string, local bool) {
	settings["is_local_file"] = local
	if local {
		settings["local_file"] = input
	} else {
		settings["input"] = input
	}
}
