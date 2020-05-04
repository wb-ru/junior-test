package mediator

//We want to create conference manager to get rid of conference queue chaos.

import (
	"fmt"
	"sync"
)

type mediator interface {
	canStart(conference) bool
	notifyFree()
}

type conferenceManager struct {
	isConferenceRoomFree bool
	lock                 *sync.Mutex
	roomQueue            []conference
}

type conference interface {
	requestConference()
	freeConference()
	quitConference()
}

type teamA struct {
	mediator mediator
}

func (tm *teamA) requestConference() {
	if tm.mediator.canStart(tm) {
		fmt.Println("TeamA: Starting conference. Do not disturb us!")
	} else {
		fmt.Println("TeamA: Waiting...")
	}
}

func (tm *teamA) quitConference() {
	fmt.Println("teamA: Leaving conference room.")
	tm.mediator.notifyFree()
}

func (tm *teamA) freeConference() {
	fmt.Println("teamA: Room is free, you can come in.")
}

type teamB struct {
	mediator mediator
}

func (tm *teamB) requestConference() {
	if tm.mediator.canStart(tm) {
		fmt.Println("TeamB: Starting conference. Do not disturb us!")
	} else {
		fmt.Println("TeamB: Waiting...")
	}
}

func (tm *teamB) quitConference() {
	fmt.Println("teamB: Leaving conference room.")
	tm.mediator.notifyFree()
}

func (tm *teamB) freeConference() {
	fmt.Println("teamB: Room is free, you can come in.")
}

func initializeConferenceManager() *conferenceManager {
	return &conferenceManager{
		isConferenceRoomFree: true,
		lock:                 &sync.Mutex{},
	}
}

func (c *conferenceManager) canStart(cnf conference) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isConferenceRoomFree {
		c.isConferenceRoomFree = false
		return true
	}
	c.roomQueue = append(c.roomQueue, cnf)
	return false
}

func (c *conferenceManager) notifyFree() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.isConferenceRoomFree {
		c.isConferenceRoomFree = true
	}
	if len(c.roomQueue) > 0 {
		firstOrgInQueue := c.roomQueue[0]
		c.roomQueue = c.roomQueue[1:]
		firstOrgInQueue.freeConference()
	}
}
