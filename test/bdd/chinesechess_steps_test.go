package bdd

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestChineseChess(t *testing.T) {
	status := godog.TestSuite{
		Name:                "chinesechess",
		ScenarioInitializer: InitializeChineseChessScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"../../features/chinesechess.feature"},
		},
	}.Run()
	if status != 0 {
		t.Fatalf("godog failed with status %d", status)
	}
}
