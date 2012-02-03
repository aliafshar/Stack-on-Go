package stackongo

import (
	"testing"
)

func TestEvents(t *testing.T) {
	dummy_server := returnDummyResponseForPathAndParams("/2.0/events", map[string]string{"key": "app123", "access_token": "abc"}, dummyEventsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	events, err := session.Events(map[string]string{"page": "1"}, map[string]string{"key": "app123", "access_token": "abc"})

	if err != nil {
		t.Error(err.String())
	}

	if len(events) != 3 {
		t.Error("Number of items wrong.")
	}

	if events[0].Event_type != "comment_posted" {
		t.Error("Event type invalid.")
	}

	if events[0].Event_id != 11462515 {
		t.Error("Event id invalid.")
	}

	if events[0].Creation_date != 1328226264 {
		t.Error("Date invalid.")
	}

}

//Test Data

var dummyEventsResponse string = `
{
  "items": [
    {
      "event_type": "comment_posted",
      "event_id": 11462515,
      "creation_date": 1328226264
    },
    {
      "event_type": "answer_posted",
      "event_id": 9121759,
      "creation_date": 1328226257
    },
    {
      "event_type": "question_posted",
      "event_id": 9121758,
      "creation_date": 1328226255
    }
  ]
}
`