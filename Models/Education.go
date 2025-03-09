package models

type Education struct {
	Id     string   `json:"id"`
	Date   string   `json:"date"`
	School string   `json:"school"`
	Degree string   `json:"degree"`
	GPA    string   `json:"gpa"`
	Points []string `json:"points"`
}

var Educations = []Education{
	Education{
		Id:     "Education-0",
		Date:   "2024 - Present",
		School: "Merrimack College",
		Degree: "Masters of Computer Science",
		GPA:    "4.0",
		Points: []string{
			"Software Engineering Concentration",
			"Gained a deep understanding of how to design and implement software systems",
			"Focus on building scalable, efficient and maintainable software systems",
		},
	},
	Education{
		Id:     "Education-1",
		Date:   "2021-2024",
		School: "Baylor University",
		Degree: "Management Information Systems + CS Minor",
		GPA:    "4.0",
		Points: []string{
			" Accolades: Wacode Hackathon 2nd place, LA CodeSprints 97th place, PwC case competition finalist",
			"President and Founder of T.I.D.E, a club focused on helping computer science students prepare for technical interviews",
			"Member of the  Engineering and Computer Science Leadership Council",
			"Inaugural winner of the Hankamer Scholars Outstanding Student award",
			"Relevant Coursework: Software Engineering, Databases, Systems Analysis & Design, Cyber Security, Data Structures, Algorithms, Data Visualization, Discrete Mathematics",
		},
	},
}
