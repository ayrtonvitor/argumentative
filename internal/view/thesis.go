package view

import (
	"github.com/ayrtonvitor/argumentative/internal/database"
	"github.com/google/uuid"
)

type Thesis struct {
	ID          uuid.UUID
	Title       string
	Description string
	Arguments   []Argument
}

type Argument struct {
	ID          uuid.UUID
	Brief       string
	Description string
	Sources     []Source
}

type Source struct {
	ID      uuid.UUID
	Content string
}

func GetTheses(
	theses []database.Thesis,
	arguments []database.GetArgumentFromThesisIdRow,
	sources []database.Argumentsource,
) []Thesis {
	sourceMap := make(map[uuid.UUID][]Source)
	argumentMap := make(map[uuid.UUID][]Argument)
	res := make([]Thesis, 0)

	for _, src := range sources {
		sourceMap[src.ArgumentID] = append(sourceMap[src.ArgumentID], Source{
			ID:      src.ID,
			Content: src.Content,
		})
	}
	for _, arg := range arguments {
		argumentMap[arg.ThesisID] = append(argumentMap[arg.ThesisID], Argument{
			ID:          arg.ID,
			Brief:       arg.Brief,
			Description: arg.Description.String,
			Sources:     sourceMap[arg.ID],
		})
	}

	for _, thesis := range theses {
		res = append(res, Thesis{
			ID:          thesis.ID,
			Title:       thesis.Title,
			Description: thesis.Description.String,
			Arguments:   argumentMap[thesis.ID],
		})
	}
	return res
}
