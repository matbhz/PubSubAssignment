package Models

type Subscriber struct {
	Subscriptions map[string]chan *Message
}

func (s *Subscriber) HasSubscription(topic string) bool{
	return s.Subscriptions[topic] != nil
}

func NewSubscriber() *Subscriber {
	return &Subscriber{Subscriptions: make(map[string]chan *Message)}
}

func (s *Subscriber) Poll(topic string) *Message {
	select {
		case message := <-s.Subscriptions[topic]:
				return message
		default:
			return nil
	}
}