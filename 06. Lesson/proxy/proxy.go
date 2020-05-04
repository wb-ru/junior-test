/*We need to make the spam filter for our exchange server.
In this example we want to filter out messages with text including "Political news"
before it reacheas our user*/
package proxy

import (
	"strings"
)

type incomingMail interface {
	Send(body string) string
}

type proxy struct {
	realRecipient incomingMail
}

func (p *proxy) Send(body string) string {
	if strings.Contains(body, "Political news") {
		return "Sorry you message has been filtered by our strong anti-political-spam-security."
	}
	if p.realRecipient == nil {
		p.realRecipient = &realRecipient{}
	}
	return p.realRecipient.Send(body)
}

type realRecipient struct {
}

func (s *realRecipient) Send(body string) string {
	return "Your message has successfully reached our exchange server!"
}
