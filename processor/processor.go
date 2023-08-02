package processor

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Processor struct {
	file *os.File
}

type StudenReport struct {
	StudentId    string  `json:"student_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MathScore    int     `json:"math_score"`
	EnglishScore int     `json:"english_score"`
	ScienceScore int     `json:"science_score"`
	TotalScore   int     `json:"total_score"`
	AverageScore float64 `json:"average_score"`
}

func NewProcessor(filename string) (*Processor, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("Error creating file: %w", err)
	}
	processor := &Processor{file}

	return processor, nil
}

func (p *Processor) WriteHeaders() error {
	reportId := uuid.New().String()
	generatedDate := time.Now().String()
	headers := []byte(fmt.Sprintf(`{
  "report_id": "%s",
  "generated_date": "%s",
  "students": [
  `, reportId, generatedDate))
	p.file.Write(headers)

	return nil
}

func (p *Processor) Write(data string, first bool) error {
	dataSplit := strings.Split(data, ",")

	studentId := dataSplit[0]
	firstName := dataSplit[1]
	lastName := dataSplit[2]

	mathScore, err := strconv.Atoi(dataSplit[3])
	if err != nil {
		panic(err)
	}

	englishScore, err := strconv.Atoi(dataSplit[4])
	if err != nil {
		panic(err)
	}

	scienceScore, err := strconv.Atoi(dataSplit[5])
	if err != nil {
		panic(err)
	}
	totalScore := mathScore + englishScore + scienceScore

	averageScore := float64(totalScore) / float64(3)

	studentReport := StudenReport{
		StudentId:    studentId,
		FirstName:    firstName,
		LastName:     lastName,
		MathScore:    mathScore,
		EnglishScore: englishScore,
		ScienceScore: scienceScore,
		TotalScore:   totalScore,
		AverageScore: averageScore,
	}

	content, err := json.MarshalIndent(studentReport, "    ", "      ")
	if err != nil {
		return fmt.Errorf("Error marshalling to json  to file :%w", err)
	}
	if first == false {
		_, err = p.file.WriteString(", \n")
		if err != nil {
			return fmt.Errorf("Error writing to file :%w", err)
		}
	}

	_, err = p.file.WriteString("    ")
	if err != nil {
		return fmt.Errorf("Error writing to file :%w", err)
	}

	_, err = p.file.Write(content)
	if err != nil {
		return fmt.Errorf("Error writing to file :%w", err)
	}

	return nil
}

func (p *Processor) Flush() error {
	footer := []byte(`
  ]
}
  `)
	_, err := p.file.Write(footer)
	if err != nil {
		return fmt.Errorf("Error writing to file :%w", err)
	}
	err = p.file.Close()
	if err != nil {
		return fmt.Errorf("Error closing file :%w", err)
	}
	return nil
}
