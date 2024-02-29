package serializers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
)

type UuidSerializer struct{}

func init() {
	schema.RegisterSerializer("uuid", UuidSerializer{})
}

func (UuidSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	if dbValue == nil {
		return nil
	}
	uid, ok := dbValue.(string)
	if !ok {
		return fmt.Errorf("expected hex string as the database value: %T", dbValue)
	}
	parseUid, err := uuid.Parse(uid)
	if err != nil {
		return fmt.Errorf("expected hex string as the database value: %T", dbValue)
	}
	field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(parseUid))
	return nil
}

func (s UuidSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	if fieldValue == nil || (field.FieldType.Kind() == reflect.Pointer && reflect.ValueOf(fieldValue).IsNil()) {
		return nil, nil
	}
	uid, ok := fieldValue.(uuid.UUID)
	if !ok {
		return nil, fmt.Errorf("expected hex string as the database value: %T", fieldValue)
	}
	uuidStr := strings.ReplaceAll(uid.String(), "-", "")
	//field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(uuidStr))
	return uuidStr, nil
}
