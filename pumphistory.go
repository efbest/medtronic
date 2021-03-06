package medtronic

import (
	"log"
	"time"
)

// History returns the history records since the specified time.
// Note that the results may include records with a zero timestamp or
// an earlier timestamp than the cutoff (in the case of DailyTotal records).
func (pump *Pump) History(since time.Time) History {
	newer := pump.Family() >= 23
	count := pump.HistoryPageCount()
	if pump.Error() != nil {
		return nil
	}
	var results History
	for page := 0; page < count && pump.Error() == nil; page++ {
		data := pump.HistoryPage(page)
		records, err := DecodeHistory(data, newer)
		if err != nil {
			pump.SetError(err)
		}
		i := findOlder(records, since)
		results = append(results, records[:i]...)
		if i < len(records) {
			break
		}
	}
	return results
}

// findOlder finds the first record that is older than the given time and returns its index,
// or len(records) if all the records occur more recently.
func findOlder(records History, cutoff time.Time) int {
	for i, r := range records {
		// Don't use DailyTotal timestamps to decide when to stop,
		// because they appear out of order (at the end of the day).
		switch r.Type() {
		case DailyTotal:
		case DailyTotal522:
		case DailyTotal523:
		default:
			t := time.Time(r.Time)
			if !t.IsZero() && t.Before(cutoff) {
				log.Printf("stopping pump history scan at %s", t.Format(UserTimeLayout))
				return i
			}
		}
	}
	return len(records)
}
