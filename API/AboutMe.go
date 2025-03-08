package API

import (
	"net/http"

	"github.com/ColeFlenniken/personalsite/Templates"
)

func GetAbout(w http.ResponseWriter, r *http.Request) {
	Templates.AboutMe().Render(r.Context(), w)
}
