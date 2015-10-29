package Tests
import "testing"
import (
	"github.com/matbhz/CitrixAssignment/Models"
	"time"
)

func TestPollingForAnNonExistingTopic_ShouldReturnNull(t *testing.T) {
	const TOPIC = "TestTopic"
	subscriber := Models.NewSubscriber("Test Subscriber")

	valuePolled := subscriber.Poll(TOPIC)

	if (valuePolled != nil) {
		t.Error()
	}
}

func TestPollingForAnExistingTopicWithEmptyMessages_ShouldReturnNull(t *testing.T) {
	const TOPIC = "TestTopic"
	subscriber := Models.NewSubscriber("Test Subscriber")

	subscriber.Subscriptions[TOPIC] = make(chan *Models.Message)


	valuePolled := subscriber.Poll(TOPIC)

	if (valuePolled != nil) {
		t.Error()
	}
}

func TestPollingForAnExistingTopicWithMessages_ShouldNotReturnNull(t *testing.T) {

	const TOPIC = "TestTopic"
	subscriber := Models.NewSubscriber("Test Subscriber")

	subscriber.Subscriptions[TOPIC] = make(chan *Models.Message)
	go func() { subscriber.Subscriptions[TOPIC] <- &Models.Message{Message: "A message", PublishedAt: time.Now().Format(time.RFC850)} }()
	time.Sleep(2 * time.Second) // TODO: Find better way to wait for blocking channels

	valuePolled := subscriber.Poll(TOPIC)

	if (valuePolled == nil) {
		t.Error()
	}

}

func TestPollingForAnExistingTopicUntilNoMoreMessagesAreLeft_ShouldReturnNull(t *testing.T) {

	const TOPIC = "TestTopic"
	subscriber := Models.NewSubscriber("Test Subscriber")

	subscriber.Subscriptions[TOPIC] = make(chan *Models.Message)
	go func() { subscriber.Subscriptions[TOPIC] <- &Models.Message{Message: "I can haz cheezburguer", PublishedAt: time.Now().Format(time.RFC850)} }()
	go func() { subscriber.Subscriptions[TOPIC] <- &Models.Message{Message: "My e-mail is lol@cats.com", PublishedAt: time.Now().Format(time.RFC850)} }()
	go func() { subscriber.Subscriptions[TOPIC] <- &Models.Message{Message: "Please read my msg!", PublishedAt: time.Now().Format(time.RFC850)} }()
	time.Sleep(2 * time.Second) // TODO: Find better way to wait for blocking channels

	valuePolled := subscriber.Poll(TOPIC)
	valuePolled = subscriber.Poll(TOPIC)
	valuePolled = subscriber.Poll(TOPIC)

	if (valuePolled == nil) {
		t.Error()
	}

	valuePolled = subscriber.Poll(TOPIC)

	if (valuePolled != nil) {
		t.Error()
	}

}