package main

import (
	"log"
	"math"
	"os"
	"slices"
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
	Age                  float64
	ExperiencedIcons     string
	ExperiencedIconCount float64
	HandyIcons           string
	HandyIconCount       float64
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

	experienced, experiencedCount := generateIconsForSection(*personalData.Experienced)
	handy, handyCount := generateIconsForSection(*personalData.Handy)

	err = tmpl.Execute(file, Template{
		Age:                  calculateAge(),
		ExperiencedIcons:     experienced,
		ExperiencedIconCount: math.RoundToEven(float64(experiencedCount)),
		HandyIcons:           handy,
		HandyIconCount:       math.RoundToEven(float64(handyCount)),
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

func generateIconsForSection(section ExperienceSection) (string, int) {
	icons := []string{}

	icons = append(icons, section.Langs...)
	icons = append(icons, section.Libs...)
	icons = append(icons, section.Platforms...)
	icons = append(icons, section.Others...)

	slices.Sort(icons)

	result := strings.Join(icons, "%2C")
	return result, len(icons)

}
