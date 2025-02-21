package API

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ColeFlenniken/personalsite/Templates"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
	Templates.TestCard().Render(r.Context(), w)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	Templates.Index().Render(r.Context(), w)

}

func GetResume(w http.ResponseWriter, r *http.Request) {
	Templates.ResumePage().Render(r.Context(), w)
}

func GetNextLetter(w http.ResponseWriter, r *http.Request) {
	arr := strings.Split("Hi I am Cole. A software developer based out of Providence Rhode Island. My expertise focuses on building software ...", " ")
	ndx, err := strconv.Atoi(r.URL.Query().Get("ndx"))
	fmt.Println(err)
	if err != nil || ndx >= len(arr) {
		return
	}

	str := strconv.Itoa(ndx + 1)
	animationNames := [2]string{"appearing-letters-down", "appearing-letters-up"}

	Templates.NextLetter(arr[ndx]+" ", "/Index/Letters?ndx="+str, animationNames[ndx%2]).Render(r.Context(), w)

}

func GetJobCard(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	query := r.URL.Query()
	//need to add error handling
	id, err := strconv.Atoi(query.Get("card"))
	if err != nil {
		fmt.Println(err)
	}
	titles := [3]string{"Software Developer / AI Engineer", "Software Developer Intern", "Research Assistant"}
	descs := [3]string{`<ul style="color:white;">
							<li>• Led development for the Enterprise flagship AI solution</li>
							<li>• Identified and developed a solution to a critical bug causing the company website to crash daily</li>
						</ul>`,
		`<ul>
							<li>• Planned, developed, and integrated an inventory management system to improve peripheral tracking and enable inventory data analysis across 4 locations</li>
							<li>• Automated new hire profile error detection and incident creation</li>
							<li>• Implemented a badge scanner to support tracking of employees entering the Tech Bar</li>
						</ul>`,
		`<ul>
							<li>• Determined governance structure and execution of 11 different blockchains</li>
							<li>• Coded decentralization measures to identify strengths of differing blockchains</li>
							<li>• Wrote scripts to transform data pulled from an external API to prepare for statistical analysis</li>
						</ul>`}
	Templates.FrontPageCard(titles[id], descs[id]).Render(r.Context(), w)
}
