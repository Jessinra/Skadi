package repositories

import (
	"fmt"
	"strings"

	"gitlab.com/trivery-id/skadi/utils/ptr"
)

type FindAllInput struct {
	Limit  *int
	Offset *int

	Filters map[string]interface{}
}

func (in *FindAllInput) FillDefault() {
	if in.Limit == nil {
		in.Limit = ptr.Int(100) // nolint
	}
	if in.Offset == nil {
		in.Offset = ptr.Int(0)
	}
}

// Where generate 'where' statement from filters.
func (in *FindAllInput) Where() string {
	statements := []string{}
	for k, v := range in.Filters {
		statements = append(statements, fmt.Sprintf("%s = %v", k, v))
	}

	return strings.Join(statements, " AND ")
}
