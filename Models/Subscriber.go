package Models


// e.g.:
// Subscriber named S1, has a News message channel and a Sports message channel

type Subscriber struct {
	Name string
	Subscriptions map[string]chan *Message
}

func (s *Subscriber) HasSubscription(topic string) bool{
	return s.Subscriptions[topic] != nil
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{Name: name, Subscriptions: make(map[string]chan *Message)}
}

func (s *Subscriber) Poll(topic string) *Message {
	select {
		case message := <-s.Subscriptions[topic]:
				return message
		default:
			return nil
	}
}