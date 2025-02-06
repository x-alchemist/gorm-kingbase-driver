package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Datetime time.Time

func (date *Datetime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Datetime(nullTime.Time)
	return
}

func (date Datetime) Value() (driver.Value, error) {
	t := time.Time(date)
	return t, nil
}

// GormDataType gorm common data type
func (date Datetime) GormDataType() string {
	return "datetime"
}

func (date Datetime) GobEncode() ([]byte, error) {

	encode, err := time.Time(date).GobEncode()
	if err != nil {
		return nil, err
	}
	return encode, err
}

func (date *Datetime) GobDecode(b []byte) error {

	return (*time.Time)(date).GobDecode(b)
}

func (date Datetime) MarshalJSON() ([]byte, error) {

	tTime := time.Time(date)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (date *Datetime) UnmarshalJSON(b []byte) error {
	t := (*time.Time)(date)
	if t.IsZero() {
		return nil
	}
	return t.UnmarshalJSON(b)
}
