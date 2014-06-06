package main

import (
	"flag"
	"fmt"
	"github.com/hoisie/web"
	"github.com/mattn/go-lingr"
	"math/rand"
	"os"
	"regexp"
)

const pngUrl = "http://mattn.tonic-water.com/sushiyuki/sushiyuki_images/%02d.png"

var re1 = regexp.MustCompile(`(sushi|寿司)`)
var re2 = regexp.MustCompile(`^(sushi|寿司) (.+)$`)

var m = map[string]int{
	"yes":    1,
	"no":     2,
	"ok":     3,
	"thanks": 4, "thank you": 4, "gyoku": 4,
	"sorry":      5,
	"sigh":       6,
	"angry":      7,
	"no comment": 8,
	"cool":       9,
	"help":       11,
	"what":       12, "question": 12,
	"sleep": 13, "sleeply": 13,
	"oh no": 14,
	"love":  15,
	"grin":  16,
	"bye":   17,
	"sneak": 18,
	"hide":  19,
	"peel":  20,
	"hot":   21,
	"fail":  22, "dip": 22,
	"too much": 23, "ikura": 23,
	"happy": 24,
	"smile": 25, "boom": 25,
	"wat": 26, "anago": 26,
	"tea": 27, "content": 27, "agari": 27,
	"gari": 28, "don't forget": 28,
	"wasabi": 29, "sabi": 29,
	"come on": 30, "c'mon": 30,
	"sparkles":  31,
	"sweat":     32,
	"cry":       33,
	"surprised": 34,
	"idea":      35,
	"sad":       36, "sob": 36,
	"chat":  37,
	"phone": 38, "call": 38,
	"hello":   39,
	"see you": 40,
}

func defaultAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":80"
	}
	return ":" + port
}

var addr = flag.String("addr", defaultAddr(), "server address")

func main() {
	flag.Parse()

	web.Post("/", func(ctx *web.Context) string {
		status, err := lingr.DecodeStatus(ctx.Request.Body)
		if err != nil {
			ctx.Abort(500, err.Error())
			return err.Error()
		}
		for _, event := range status.Events {
			if message := event.Message; message != nil {
				if re2.MatchString(message.Text) {
					if sushi, ok := m[re2.FindStringSubmatch(message.Text)[2]]; ok {
						return fmt.Sprintf(pngUrl, sushi)
					}
				} else if re1.MatchString(message.Text) {
					return fmt.Sprintf(pngUrl, rand.Int()%40+1)
				}
			}
		}
		return ""
	})
	web.Run(*addr)
}
