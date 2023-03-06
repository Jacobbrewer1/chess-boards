package custom

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"time"
)

type Date time.Time

func (d *Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(*d))
}

func (d *Date) UnmarshalBSON(bytes []byte) error {
	if bytes == nil {
		return nil
	}

	t, _, ok := bsoncore.ReadTime(bytes)
	if !ok {
		return errors.New("not enough bytes to unmarshal Date")
	}

	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(DateLayout, string(text))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d *Date) Scan(src any) error {
	t, err := time.Parse(DateTimeWithNumericAndUtcZone, src.(string))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d *Date) String() string {
	return time.Time(*d).Format(DateLayout)
}

func (d *Date) TimeValue() *time.Time {
	if d == nil {
		return nil
	}

	var t = time.Time(*d)
	return &t
}

func (d *Date) Display() string {
	return d.TimeValue().Format(DateDisplay)
}
