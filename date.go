package sage

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("20060102"))
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
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

	d.Time, err = time.Parse("20060102", value)
	return
}

func (d Date) String() string {
	return d.Time.Format("20060102")
}
