package notify

import "supervisor-event-listener/event"

type Notifiable interface {
	Send(*event.Message) error
}
