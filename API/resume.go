package API

import (
	"net/http"

	"github.com/ColeFlenniken/personalsite/Templates"
	"github.com/ColeFlenniken/personalsite/models"
)

func GetResume(w http.ResponseWriter, r *http.Request) {
	Templates.ResumePage().Render(r.Context(), w)
}

func GetJobCardHTML(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("card")
	for i := 0; i < len(models.Jobs); i++ {
		if id == models.Jobs[i].Id {
			Templates.JobCardHTML(models.Jobs[i]).Render(r.Context(), w)
		}
	}
	for i := 0; i < len(models.Educations); i++ {
		if id == models.Educations[i].Id {
			Templates.EducationCardHTML(models.Educations[i]).Render(r.Context(), w)
		}
	}
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	Templates.Job(models.Jobs).Render(r.Context(), w)
}

func GetEducation(w http.ResponseWriter, r *http.Request) {
	Templates.Education(models.Educations).Render(r.Context(), w)
}
