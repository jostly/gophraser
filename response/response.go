package response

import (
	"net/http"
	"strings"

	"github.com/hoisie/mustache"
)

type Phrase struct {
	Adjective string
	Animal    string
}

func (p Phrase) String() string {
	return p.Adjective + " " + p.Animal
}

func BuildResponse(req *http.Request, phrase Phrase, w http.ResponseWriter) string {
	accept := strings.Join(req.Header["Accept"], ",")
	switch {
	case strings.Contains(accept, "text/html"):
		return htmlResponseBuilder(phrase, w)
	case strings.Contains(accept, "application/json"):
		return jsonResponseBuilder(phrase, w)
	default:
		return textResponseBuilder(phrase, w)
	}
}

func htmlResponseBuilder(phrase Phrase, w http.ResponseWriter) string {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return mustache.RenderFile("templates/phrase.mustache",
		map[string]string{"phrase": phrase.String()})
}

func jsonResponseBuilder(phrase Phrase, w http.ResponseWriter) string {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return "{\"adjective\": \"" + phrase.Adjective + "\", \"animal\": \"" + phrase.Animal + "\"}"
}

func textResponseBuilder(phrase Phrase, w http.ResponseWriter) string {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	return phrase.String()
}
