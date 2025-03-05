package models

type Job struct {
    Id      string   `json:"id"`
    Date    string   `json:"date"`
    Role    string   `json:"role"`
    Company string   `json:"company"`
    Points  []string `json:"points"`
}

var Jobs = []Job{
    {
		Id: "0",
		Date: "2024 - Present",
        Role: "Software Developer / AI Engineer",
		Company: "Textron",
        Points: []string{
            "Led development for the Enterprise flagship AI solution",
            "Identified and developed a solution to a critical bug causing the company website to crash daily",
        },
    },
    {
		Id: "1",
        Role: "Software Developer Intern",
		Date: "2023",
		Company: "Textron Systems",
        Points: []string{
            "Planned, developed, and integrated an inventory management system to improve peripheral tracking and enable inventory data analysis across 4 locations",
            "Automated new hire profile error detection and incident creation",
            "Implemented a badge scanner to support tracking of employees entering the Tech Bar",
        },
    },
    {
		Id: "2",
        Role: "Research Assistant",
		Date: "2022-2024",
		Company: "Baylor University",
        Points: []string{
            "Determined governance structure and execution of 11 different blockchains",
            "Coded decentralization measures to identify strengths of differing blockchains",
            "Wrote scripts to transform data pulled from an external API to prepare for statistical analysis",
        },
    }
}

