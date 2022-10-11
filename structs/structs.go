package structs

type Programmer struct {
	Email    string `json:"email" xml:"email"`
	Salary   int    `json:"salary" xml:"salary"`
	Team     int    `json:"team" xml:"team"`
	Position string `json:"position" xml:"position"`
}

type TechLeader struct {
	Programmer     `json:"programmer" xml:"programmer"`
	CurrentProject string `json:"current_project" xml:"currentProject"`
	Subordinates   int    `json:"subordinates" xml:"subordinates"`
}

type Junior struct {
	Programmer  `json:"programmer" xml:"programmer"`
	Resignation int `json:"resignation" xml:"resignation"`
}
