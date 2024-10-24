package main

type Resume struct {
	Name         string        `yaml:"name,omitempty"`
	Title        string        `yaml:"title,omitempty"`
	ProfileImage string        `yaml:"profile_image,omitempty"`
	Contact      *Contact      `yaml:"contact,omitempty"`
	Summary      string        `yaml:"summary,omitempty"`
	Experience   []Experience  `yaml:"professional_experience,omitempty"`
	Skills       *Skills       `yaml:"skills,omitempty"`
	Projects     []Project     `yaml:"projects,omitempty"`
	Education    []Education   `yaml:"education,omitempty"`
	Languages    []Language    `yaml:"languages,omitempty"`
	Certificates []Certificate `yaml:"certificates,omitempty"`
	Awards       []string      `yaml:"awards,omitempty"`
}

type Contact struct {
	Email    string `yaml:"email,omitempty"`
	Phone    string `yaml:"phone,omitempty"`
	Website  string `yaml:"website,omitempty"`
	Github   string `yaml:"github,omitempty"`
	LinkedIn string `yaml:"linkedin,omitempty"`
}

type Experience struct {
	Title        string   `yaml:"title,omitempty"`
	Company      string   `yaml:"company,omitempty"`
	Location     string   `yaml:"location,omitempty"`
	StartDate    string   `yaml:"start_date,omitempty"`
	EndDate      string   `yaml:"end_date,omitempty"`
	Achievements []string `yaml:"achievements,omitempty"`
}

type Project struct {
	Name         string   `yaml:"name,omitempty"`
	URL          string   `yaml:"url,omitempty"`
	Description  string   `yaml:"description,omitempty"`
	Technologies []string `yaml:"technologies,omitempty"`
	Points       []string `yaml:"points,omitempty"`
}

type Certificate struct {
	Name string `yaml:"name,omitempty"`
	URL  string `yaml:"url,omitempty"`
}

type Skills struct {
	Programming []string `yaml:"programming_languages,omitempty"`
	Security    []string `yaml:"security,omitempty"`
	DevSecOps   []string `yaml:"devsecops,omitempty"`
	Engineering []string `yaml:"software_engineering,omitempty"`
}

type Education struct {
	Degree      string   `yaml:"degree,omitempty"`
	Institution string   `yaml:"institution,omitempty"`
	Period      string   `yaml:"period,omitempty"`
	Location    string   `yaml:"location,omitempty"`
	Courses     []string `yaml:"courses,omitempty"`
}

type Language struct {
	Name  string `yaml:"name,omitempty"`
	Level string `yaml:"level,omitempty"`
}