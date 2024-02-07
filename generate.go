package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

type ExperienceSection struct {
	Langs     []string `yaml:"langs"`
	Libs      []string `yaml:"libs"`
	Platforms []string `yaml:"platforms"`
	Others    []string `yaml:"others"`
}

type PersonalData struct {
	Experienced *ExperienceSection `yaml:"experienced"`
	Handy       *ExperienceSection `yaml:"handy"`
}

type Template struct {
	Age              float64
	ExperiencedIcons string
	HandyIcons       string
}

func main() {
	data, err := os.ReadFile("data.yml")
	if err != nil {
		log.Fatal("Failed to open data file")
	}

	var personalData PersonalData
	err = yaml.Unmarshal(data, &personalData)
	if err != nil {
		log.Fatal("Failed to parse data file as yml")
	}

	file, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to open README.md")
	}
	defer file.Close()

	tmpl, err := template.New("README.tmpl.md").ParseFiles("README.tmpl.md")
	if err != nil {
		log.Fatal("Failed to load template")
	}

	err = tmpl.Execute(file, Template{
		Age:              calculateAge(),
		ExperiencedIcons: generateIconsForSection(*personalData.Experienced),
		HandyIcons:       generateIconsForSection(*personalData.Handy),
	})

	if err != nil {
		panic(err)
	}

	log.Println("Done!")

}

func calculateAge() float64 {
	dateOfBirth, err := time.Parse("2006-01-02", "2002-03-30") // Go Time Template, My Birthday
	if err != nil {
		log.Fatal("Failed to parse Date of Birth")
	}

	now := time.Now().Truncate(time.Hour)

	elapsedSeconds := now.Sub(dateOfBirth).Seconds()

	daysInYear := 365.2425 // Average number of days in a year
	secondsInYear := daysInYear * 24 * 60 * 60

	age := float64(elapsedSeconds) / secondsInYear
	age = math.Floor(age*1000) / 1000
	return age
}

func generateIconsForSection(section ExperienceSection) string {
	var ExperiencedIcons strings.Builder

	for _, item := range section.Langs {
		ExperiencedIcons.WriteString(item)
		ExperiencedIcons.WriteString(" ")
	}

	for _, item := range section.Libs {
		ExperiencedIcons.WriteString(item)
		ExperiencedIcons.WriteString(" ")
	}

	for _, item := range section.Platforms {
		ExperiencedIcons.WriteString(item)
		ExperiencedIcons.WriteString(" ")
	}

	for _, item := range section.Others {
		ExperiencedIcons.WriteString(item)
		ExperiencedIcons.WriteString(" ")
	}

	result := strings.Trim(ExperiencedIcons.String(), " ")
	return strings.ReplaceAll(result, " ", "%20")

}
