package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	conferenceManager := initializeConferenceManager()
	teamA := &teamA{
		mediator: conferenceManager,
	}
	teamB := &teamB{
		mediator: conferenceManager,
	}
	teamA.requestConference()
	teamB.requestConference()
	teamA.quitConference()
}
