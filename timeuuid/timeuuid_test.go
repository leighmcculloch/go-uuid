package timeuuid

import (
	"testing"
	"time"

	uuid "github.com/leighmcculloch/go-uuid"
)

func TestNow(t *testing.T) {
	t.Log(Now())
}

func TestNowUnique(t *testing.T) {
	var seen map[uuid.UUID]bool

	for i := 0; i < 1000; i++ {
		id := Now()
		if _, ok := seen[id]; ok {
			t.Fatal("Seen", id, "already, expected unique")
		}
	}
}

func TestNowOrder(t *testing.T) {
	var lastID uuid.UUID
	for i := 0; i < 100000; i++ {
		id := Now()
		if id.String() < lastID.String() {
			t.Fatal(id, "sorts earlier than last", lastID, "want to sort later")
		}
		lastID = id
	}
}

func TestNewOrder(t *testing.T) {
	testPeriodYears := 20
	timeStart := time.Now().AddDate(-testPeriodYears/2, 0, 0)
	timeEnd := timeStart.AddDate(testPeriodYears, -1, 0)
	var lastID uuid.UUID
	for time := timeStart; time.Before(timeEnd); time = time.AddDate(0, 0, 1) {
		id := New(time)
		if id.Timestamp() < lastID.Timestamp() {
			t.Fatal(id, "sorts earlier than last", lastID, "want to sort later")
		}
		lastID = id
	}
}
