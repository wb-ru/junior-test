package chainofresponsibility

//We need to make an instrument to filter our incoming mail by categories
type Handler interface {
	SendRequest(message string) string
}

type HandlerWork struct {
	next Handler
}

func (h *HandlerWork) SendRequest(message string) (result string) {
	if message == "Job" {
		result = "Message sent to job folder."
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerShops struct {
	next Handler
}

func (h *HandlerShops) SendRequest(message string) (result string) {
	if message == "Shop" {
		result = "Message sent to shop folder."
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerCritical struct {
	next Handler
}

func (h *HandlerCritical) SendRequest(message string) (result string) {
	if message == "Monitoring" {
		result = "Message sent to critical folder."
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
