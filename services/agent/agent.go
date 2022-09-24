// Package agent holds on services
// that build all together the business flow,
// agent only depends on package core.
package agent

import (
	"github.com/delveper/heroes/core"
)

type Agent struct {
	Repo core.Repository
	// Log
}

func NewUserAgent(ur core.Repository) *Agent {
	return &Agent{Repo: ur}
}
