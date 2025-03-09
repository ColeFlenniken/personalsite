package API

import (
	"net/http"

	"github.com/ColeFlenniken/personalsite/Templates"
	"github.com/ColeFlenniken/personalsite/models"
)

func GetResume(w http.ResponseWriter, r *http.Request) {
	Templates.ResumePage().Render(r.Context(), w)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	Templates.Job(models.Jobs).Render(r.Context(), w)
}

func GetEducation(w http.ResponseWriter, r *http.Request) {
	Templates.Education(models.Educations).Render(r.Context(), w)
}
