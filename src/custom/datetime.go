package custom

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"time"
)

type Datetime time.Time

func (d *Datetime) MarshalBSON() ([]byte, error) {
	return d.MarshalText()
}

func (d *Datetime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(*d))
}

func (d *Datetime) UnmarshalBSON(bytes []byte) error {
	if bytes == nil {
		return nil
	}

	t, _, ok := bsoncore.ReadTime(bytes)
	if !ok {
		return errors.New("not enough bytes to unmarshal Datetime")
	}

	*d = Datetime(t)
	return nil
}

func (d *Datetime) MarshalJSON() ([]byte, error) {
	return []byte(d.UTC().String()), nil
}

func (d *Datetime) MarshalText() ([]byte, error) {
	return []byte(d.UTC().String()), nil
}

func (d *Datetime) UnmarshalText(text []byte) error {
	loc, err := time.LoadLocation(UtcLocation)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(DateTimeWithNumericAndUtcZone, string(text), loc)
	if err != nil {
		return err
	}

	t = t.In(time.Now().Location())

	*d = Datetime(t)
	return nil
}

func (d *Datetime) Scan(src any) error {
	loc, err := time.LoadLocation(UtcLocation)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(DateTimeWithNumericAndUtcZone, src.(string), loc)
	if err != nil {
		return err
	}

	t = t.In(time.Now().Location())

	*d = Datetime(t)
	return nil
}

func (d *Datetime) String() string {
	return time.Time(*d).Format(time.RFC3339)
}

func (d *Datetime) TimeValue() *time.Time {
	if d == nil {
		return nil
	}

	var t = time.Time(*d)
	return &t
}

func (d *Datetime) Display() string {
	return d.TimeValue().Format(DateTimeDisplay)
}

func (d *Datetime) UTC() *Datetime {
	if d == nil {
		return nil
	}

	var t = Datetime(d.TimeValue().UTC())
	return &t
}
