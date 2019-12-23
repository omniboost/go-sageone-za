package aktiva

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time)
}

func (d DateTime) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("20060102150405", value)
	return
}

func (d DateTime) String() string {
	return d.Time.Format("20060102150405")
}
