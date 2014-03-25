package iso8601

import (
	"encoding/json"
	"testing"
	"time"
)

func TestISO8601Time(t *testing.T) {
	now := Time(time.Now().UTC())

	data, err := json.Marshal(now)
	if err != nil {
		t.Fatal(err)
	}

	_, err = time.Parse(`"`+Format+`"`, string(data))
	if err != nil {
		t.Fatal(err)
	}

	var now2 Time
	err = json.Unmarshal(data, &now2)
	if err != nil {
		t.Fatal(err)
	}

	// ISO8601 truncates the time to nearest second
	trunc := time.Time(now).Truncate(time.Second)
	if !trunc.Equal(time.Time(now2)) {
		t.Fatalf("Time %s does not equal expected %s", now2, trunc)
	}

	if trunc.String() != now2.String() {
		t.Fatalf("String format for %s does not equal expected %s", now2, trunc)
	}
}
