package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"inventory-service/helpers/exception"

	"inventory-service/model"
	"inventory-service/protocgen/inventory/v1/global/meta"

	"inventory-service/helpers/utils/converter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	invalidParameter = "invalid %s parameter"
)

var (
	orderRegex     = regexp.MustCompile("(\\w+):(\\w+)")
	OrderOperators = map[string]string{
		"desc": "desc",
		"asc":  "asc",
	}
	filterRegex    = regexp.MustCompile(`(\w+):([^|]+):(\w+)`)
	FilterOperator = map[string]string{
		"eq":   "=",
		"lt":   "<",
		"gt":   ">",
		"lte":  "<=",
		"gte":  ">=",
		"in":   "in",
		"like": "like",
		"is":   "is",
		"not":  "not in",
	}
)

type GRPCHandler struct {
}

func (h *GRPCHandler) StreamContextError(ctx context.Context) error {
	switch err := ctx.Err(); {
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, "request is canceled")
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
func (h *GRPCHandler) GetGRPCCode(code codes.Code) int64 {
	return int64(code)
}

func (h *GRPCHandler) ResponseOK(messages string) *meta.Meta {
	return &meta.Meta{
		StatusCode:  h.GetGRPCCode(codes.OK),
		MessageCode: codes.OK.String(),
		Message:     messages,
	}
}

func (h *GRPCHandler) ResponseOKPagination(messages string) *meta.Meta {
	return &meta.Meta{
		StatusCode:  h.GetGRPCCode(codes.OK),
		MessageCode: codes.OK.String(),
		Message:     messages,
		Pagination:  &meta.PaginationResponse{},
	}
}

func (h *GRPCHandler) ResponseError(err error) error {
	return status.Error(codes.Internal, err.Error())
}

func (h *GRPCHandler) ResponseErrorException(errException *exception.Exception) error {
	return errException.ReturnGRPCError()
}

func (h *GRPCHandler) ResponseErrorCode(err error, code codes.Code) error {
	return status.Error(code, err.Error())
}

func GetOrderValue(value string) (string, error) {
	if op, ok := OrderOperators[value]; ok {
		return op, nil
	}
	return "", fmt.Errorf(invalidParameter, value)
}

func (h *GRPCHandler) ParsePageLimitParam(page, pageSize string) (model.PaginationParam, error) {
	var p model.PaginationParam
	var err error
	if page == "" || page == "0" {
		p.Offset = -1
	} else {
		p.Offset, err = strconv.Atoi(page)
		if err != nil {
			return model.PaginationParam{}, err
		}
	}
	if pageSize == "" || pageSize == "0" {
		p.Limit = -1
	} else {
		p.Limit, err = strconv.Atoi(pageSize)
		if err != nil {
			return model.PaginationParam{}, err
		}
	}
	return p, nil
}

func (h *GRPCHandler) ParseOrderParam(order string) (model.OrderParam, error) {
	var p model.OrderParam
	if order != "" {
		listOrder := strings.Split(order, ",")
		for _, o := range listOrder {
			if !orderRegex.MatchString(o) {
				continue
			}
			condition := strings.Split(o, ":")
			if len(condition) != 2 {
				return model.OrderParam{}, fmt.Errorf(invalidParameter, "order")
			}
			value, err := GetOrderValue(condition[1])
			if err != nil {
				return model.OrderParam{}, err
			}
			p.OrderBy = condition[0]
			p.Order = value
		}
	}
	return p, nil
}

func (h *GRPCHandler) ParseFilterParams(ctx context.Context, f string) (model.FilterParams, context.Context, error) {
	var p model.FilterParams

	if f != "" {
		listFilter := strings.Split(f, "|")
		for _, v := range listFilter {
			if !filterRegex.MatchString(v) {
				continue
			}
			filter := strings.Split(v, ":")
			if len(filter) != 3 {
				return model.FilterParams{}, ctx, fmt.Errorf(invalidParameter, filter)
			}
			operator, err := h.getFilterOperator(filter[2])
			if err != nil {
				return model.FilterParams{}, ctx, err
			}

			p = append(p, &model.FilterParam{
				Field:    filter[0],
				Value:    filter[1],
				Operator: operator,
			})
		}
	}

	return p, ctx, nil
}

func (h *GRPCHandler) ParseKeywordParam(keyword string) model.KeywordParam {
	return model.KeywordParam{
		Value: keyword,
	}
}

func (h *GRPCHandler) getFilterOperator(operator string) (string, error) {
	if op, ok := FilterOperator[operator]; ok {
		return op, nil
	}
	return "", fmt.Errorf(invalidParameter, operator)
}

func (h *GRPCHandler) ParseListParams(ctx context.Context, page, pageSize int32, order, filter, keyword string) (
	model.PaginationParam, model.OrderParam, model.FilterParams, model.KeywordParam, context.Context, error,
) {
	pagination, err := h.ParsePageLimitParam(converter.ToString(page), converter.ToString(pageSize))
	if err != nil {
		return model.PaginationParam{}, model.OrderParam{}, model.FilterParams{}, model.KeywordParam{}, ctx, err
	}
	orders, err := h.ParseOrderParam(order)
	if err != nil {
		return model.PaginationParam{}, model.OrderParam{}, model.FilterParams{}, model.KeywordParam{}, ctx, err
	}
	filters, ctx, err := h.ParseFilterParams(ctx, filter)
	if err != nil {
		return model.PaginationParam{}, model.OrderParam{}, model.FilterParams{}, model.KeywordParam{}, ctx, err
	}

	return pagination, orders, filters, h.ParseKeywordParam(keyword), ctx, nil
}

// Transform function with integer type conversion support
func (h *GRPCHandler) Transform(src interface{}, dest interface{}) error {

	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() == reflect.Ptr && srcVal.Elem().Kind() == reflect.Struct {
		srcVal = srcVal.Elem()
	} else {
		return errors.New("source must be a pointer to a struct")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		fieldName := srcVal.Type().Field(i).Name
		srcField := srcVal.Field(i)
		destField := reflect.Value{}

		if destVal.Kind() == reflect.Ptr {
			// If destVal is a pointer but is nil, initialize it with a new value of the underlying type
			if destVal.IsNil() {
				destVal.Set(reflect.New(destVal.Type().Elem()))
			}
			destField = reflect.Indirect(destVal).FieldByName(fieldName)
		} else {
			destField = destVal.FieldByName(fieldName)
		}
		// Check if the destination field exists and is settable
		// fmt.Println(fieldName, srcField.Kind(), destField.Kind())
		if destField.IsValid() && destField.CanSet() {
			// If it's a slice, handle array of structs or array of basic types

			if srcField.Kind() == reflect.Slice && destField.Kind() == reflect.Slice {
				srcElemType := srcField.Type().Elem()
				destElemType := destField.Type().Elem()
				// Check if the slice contains structs or basic types
				if srcElemType.Kind() == reflect.Struct || (srcElemType.Kind() == reflect.Ptr && srcElemType.Elem().Kind() == reflect.Struct) {
					// Slice of structs or pointers to structs - use TransformSlice
					if destElemType.Kind() == reflect.Uint8 {
						raws, err := reflectValueToRawMessage(srcVal.Field(i), fieldName)
						if err != nil {
							return err
						}
						destField.Set(reflect.ValueOf(raws))
					} else {
						if err := h.TransformSlice(srcField, destField); err != nil {
							return err
						}
					}
				} else if srcElemType == destElemType {
					// Directly copy slice if types match (for basic types)
					destField.Set(srcField)
				}
			} else if srcField.Kind() == reflect.String && destField.Type() == reflect.TypeOf(json.RawMessage{}) {

				str := srcField.Interface().(string)

				// Try to unescape if needed
				cleanStr := str
				if strings.Contains(str, `\"`) {
					if unquoted, err := strconv.Unquote(`"` + str + `"`); err == nil {
						cleanStr = unquoted
					}
				}

				if str == "" {
					destField.Set(reflect.Zero(destField.Type())) // sets to nil json.RawMessage
				} else {
					destField.Set(reflect.ValueOf(json.RawMessage(cleanStr)))
				}
			} else if srcField.Type() == destField.Type() {
				// Directly set the field if types match
				destField.Set(srcField)
			} else if h.isConvertibleInteger(srcField, destField) {

				destField.Set(h.convertInteger(srcField, destField.Type()))

			} else if h.isConvertibleTimeToString(srcField, destField) {
				convertedValue, err := h.convertTimeToString(srcField)
				if err != nil {
					return err
				}
				destField.Set(convertedValue)
			} else if h.isConvertibleStringToTime(srcField, destField) {
				convertedValue, err := h.convertStringToTime(srcField, destField.Type())
				if err != nil {
					return err
				}
				destField.Set(convertedValue)
			} else if h.isConvertibleString(srcField, destField) {
				convertedValue, err := h.convertString(srcField, destField.Type())
				if err != nil {
					return err
				}
				destField.Set(convertedValue)
			} else if h.isConvertibleFloat(srcField, destField) {

				destField.Set(h.convertFloat(srcField, destField.Type()))

			} else if srcField.Kind() == reflect.Struct && destField.Kind() == reflect.Ptr {
				destElem := reflect.New(destField.Type().Elem()).Interface()
				srcElem := reflect.New(srcField.Type())
				srcElem.Elem().Set(srcField)
				// srcElem.Elem().Set(originalValue)
				// interfaceValue := srcElem.Interface()
				if err := h.Transform(srcElem.Interface(), destElem); err != nil {
					return err
				}
				destField.Set(reflect.ValueOf(destElem))
			} else if srcField.Kind() == reflect.Ptr && !srcField.IsNil() {
				// elemType := destField.Type().Elem()

				newDestElem := reflect.New(destField.Type())
				if err := h.Transform(srcField.Interface(), newDestElem.Interface()); err != nil {
					return err
				}
				destField.Set(newDestElem.Elem())
			} else if srcField.Type() == reflect.TypeOf(json.RawMessage{}) && destField.Kind() == reflect.String {

				raw := srcField.Interface().(json.RawMessage)
				destField.SetString(string(raw)) // NO Marshal!

			}
		}
	}
	return nil
}

// TransformSlice function remains the same as previously defined

// Check if the source and destination fields are integers and can be converted

// TransformSlice transforms a slice of structs or pointers from the source type to the destination type
func (h *GRPCHandler) TransformSlice(src reflect.Value, dest reflect.Value) error {
	elemType := dest.Type().Elem() // Get the type of elements in the destination slice
	// Create a new slice for the destination with the same capacity as the source
	destSlice := reflect.MakeSlice(dest.Type(), 0, src.Len())

	// Check if the element type is a basic type (not a struct or pointer to struct)
	isBasicType := elemType.Kind() != reflect.Struct && elemType.Kind() != reflect.Ptr

	if isBasicType {
		// Directly copy elements if the slice contains basic types
		for i := 0; i < src.Len(); i++ {
			destSlice = reflect.Append(destSlice, src.Index(i))
		}
		dest.Set(destSlice)
		return nil
	}

	// If the element type is a struct or a pointer, perform recursive transformation
	isElemPointer := elemType.Kind() == reflect.Ptr

	for i := 0; i < src.Len(); i++ {
		srcElem := src.Index(i)
		var destElem reflect.Value

		if isElemPointer {
			// If the destination element is a pointer, create a new instance of the element type
			destElem = reflect.New(elemType.Elem()) // Allocate new pointer to element
			if err := h.Transform(srcElem.Interface(), destElem.Interface()); err != nil {
				return err
			}
		} else {
			// If the destination element is not a pointer, create a new value and pass its address
			destElem = reflect.New(elemType).Elem()
			if err := h.Transform(srcElem.Interface(), destElem.Addr().Interface()); err != nil {
				return err
			}
		}

		//update

		// Append the transformed element to the destination slice
		destSlice = reflect.Append(destSlice, destElem)
	}

	// Set the transformed slice to the destination field
	dest.Set(destSlice)
	return nil
}

func (h *GRPCHandler) isConvertibleInteger(srcField, destField reflect.Value) bool {
	srcType := srcField.Type()
	destType := destField.Type()

	srcKind := srcType.Kind()
	if srcKind == reflect.Ptr && !srcField.IsNil() {
		srcType = srcType.Elem()
		srcKind = srcType.Kind()
	}

	destKind := destType.Kind()
	if destKind == reflect.Ptr {
		destType = destType.Elem()
		destKind = destType.Kind()
	}

	return (srcKind == reflect.Int || srcKind == reflect.Int8 || srcKind == reflect.Int16 || srcKind == reflect.Int32 || srcKind == reflect.Int64) &&
		(destKind == reflect.Int || destKind == reflect.Int8 || destKind == reflect.Int16 || destKind == reflect.Int32 || destKind == reflect.Int64)
}

// Convert the integer from src to the target destination type
func (h *GRPCHandler) convertInteger(srcField reflect.Value, destType reflect.Type) reflect.Value {
	// Dereference the source if it is a pointer
	if srcField.Kind() == reflect.Ptr {
		if srcField.IsNil() {
			// Return a nil pointer for the destination type if source is nil
			if destType.Kind() == reflect.Ptr {
				return reflect.Zero(destType) // Equivalent to returning nil
			}
			return reflect.Zero(destType) // Zero value for non-pointer types
		}
		// Dereference the pointer
		srcField = srcField.Elem()
	}

	// Return nil if the source value is invalid or zero
	if !srcField.IsValid() || srcField.IsZero() {
		if destType.Kind() == reflect.Ptr {
			return reflect.Zero(destType) // Return a nil pointer
		}
		return reflect.Zero(destType) // Return zero value for non-pointer types
	}
	switch destType.Kind() {
	case reflect.Int:
		return reflect.ValueOf(int(srcField.Int()))
	case reflect.Int8:
		return reflect.ValueOf(int8(srcField.Int()))
	case reflect.Int16:
		return reflect.ValueOf(int16(srcField.Int()))
	case reflect.Int32:
		return reflect.ValueOf(int32(srcField.Int()))
	case reflect.Int64:
		return reflect.ValueOf(srcField.Int())
	case reflect.Ptr, reflect.Uintptr:
		elemType := destType.Elem()
		switch elemType.Kind() {
		case reflect.Int:
			val := int(srcField.Int())
			return reflect.ValueOf(&val)
		case reflect.Int8:
			val := int8(srcField.Int())
			return reflect.ValueOf(&val)
		case reflect.Int16:
			val := int16(srcField.Int())
			return reflect.ValueOf(&val)
		case reflect.Int32:
			val := int32(srcField.Int())
			return reflect.ValueOf(&val)
		case reflect.Int64:
			val := srcField.Int()
			return reflect.ValueOf(&val)
		default:
			panic("unsupported pointer to integer type conversion")
		}
	default:
		panic("unsupported integer type conversion")
	}
}

func (h *GRPCHandler) isConvertibleTimeToString(srcField, destField reflect.Value) bool {
	srcType := srcField.Type()
	destType := destField.Type()

	return (srcType == reflect.TypeOf(time.Time{}) || srcType == reflect.PointerTo(reflect.TypeOf(time.Time{}))) &&
		destType.Kind() == reflect.String
}

func (h *GRPCHandler) convertTimeToString(srcField reflect.Value) (reflect.Value, error) {
	var timeValue time.Time

	if srcField.Type() == reflect.TypeOf(time.Time{}) {
		// If src is time.Time
		timeValue = srcField.Interface().(time.Time)
	} else if srcField.Type() == reflect.PointerTo(reflect.TypeOf(time.Time{})) {
		// If src is *time.Time
		if srcField.IsNil() {
			// Handle nil pointer case
			return reflect.ValueOf(""), nil
		}
		timeValue = srcField.Elem().Interface().(time.Time)
	} else {
		return reflect.Value{}, fmt.Errorf("unsupported source type for time to string conversion: %v", srcField.Type())
	}

	// Format the time value to a string (RFC3339 or customize as needed)
	formattedTime := timeValue.Format(time.RFC3339)
	return reflect.ValueOf(formattedTime), nil
}

func (h *GRPCHandler) isConvertibleStringToTime(srcField, destField reflect.Value) bool {
	srcType := srcField.Type()
	destType := destField.Type()

	return srcType.Kind() == reflect.String &&
		(destType == reflect.TypeOf(time.Time{}) || destType == reflect.PointerTo(reflect.TypeOf(time.Time{})))
}

func (h *GRPCHandler) convertStringToTime(srcField reflect.Value, destType reflect.Type) (reflect.Value, error) {
	if srcField.Kind() != reflect.String {
		return reflect.Value{}, fmt.Errorf("unsupported source type for string to time conversion: %v", srcField.Type())
	}

	srcString := srcField.String()
	if srcString == "" {
		// Handle empty string case
		if destType == reflect.TypeOf(time.Time{}) {
			// Return zero value for time.Time
			return reflect.ValueOf(time.Time{}), nil
		} else if destType == reflect.PointerTo(reflect.TypeOf(time.Time{})) {
			// Return nil for *time.Time
			return reflect.Zero(destType), nil
		}
		return reflect.Value{}, fmt.Errorf("unsupported destination type for empty string conversion: %v", destType)
	}

	// Parse the string into a time.Time value
	parsedTime, err := time.Parse(time.RFC3339, srcString)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("failed to parse string to time: %w", err)
	}

	if destType == reflect.TypeOf(time.Time{}) {
		// Destination is time.Time
		return reflect.ValueOf(parsedTime), nil
	} else if destType == reflect.PointerTo(reflect.TypeOf(time.Time{})) {
		// Destination is *time.Time
		return reflect.ValueOf(&parsedTime), nil
	}

	return reflect.Value{}, fmt.Errorf("unsupported destination type for string to time conversion: %v", destType)
}

func (h *GRPCHandler) isConvertibleString(srcField, destField reflect.Value) bool {
	srcType := srcField.Type()
	destType := destField.Type()

	return (srcType.Kind() == reflect.String && destType.Kind() == reflect.String) ||
		(srcType.Kind() == reflect.String && destType == reflect.PointerTo(reflect.TypeOf(""))) ||
		(srcType == reflect.PointerTo(reflect.TypeOf("")) && destType.Kind() == reflect.String) ||
		(srcType == reflect.PointerTo(reflect.TypeOf("")) && destType == reflect.PointerTo(reflect.TypeOf("")))
}

func (h *GRPCHandler) isConvertibleFloat(srcField, destField reflect.Value) bool {
	srcType := srcField.Type()
	destType := destField.Type()

	srcKind := srcType.Kind()
	if srcKind == reflect.Ptr && !srcField.IsNil() {
		srcType = srcType.Elem()
		srcKind = srcType.Kind()
	}

	destKind := destType.Kind()
	if destKind == reflect.Ptr {
		destType = destType.Elem()
		destKind = destType.Kind()
	}

	return (srcKind == reflect.Float32 || srcKind == reflect.Float64) &&
		(destKind == reflect.Float32 || destKind == reflect.Float64)
}

func (h *GRPCHandler) convertFloat(srcField reflect.Value, destType reflect.Type) reflect.Value {
	// Dereference the source if it is a pointer
	if srcField.Kind() == reflect.Ptr {
		if srcField.IsNil() {
			// Return a nil pointer for the destination type if source is nil
			if destType.Kind() == reflect.Ptr {
				return reflect.Zero(destType) // Equivalent to returning nil
			}
			return reflect.Zero(destType) // Zero value for non-pointer types
		}
		// Dereference the pointer
		srcField = srcField.Elem()
	}

	// Return nil if the source value is invalid or zero
	if !srcField.IsValid() || srcField.IsZero() {
		if destType.Kind() == reflect.Ptr {
			return reflect.Zero(destType) // Return a nil pointer
		}
		return reflect.Zero(destType) // Return zero value for non-pointer types
	}

	switch destType.Kind() {
	case reflect.Float32:
		return reflect.ValueOf(float32(srcField.Float()))
	case reflect.Float64:
		return reflect.ValueOf(srcField.Float())
	case reflect.Ptr:
		elemType := destType.Elem()
		switch elemType.Kind() {
		case reflect.Float32:
			val := float32(srcField.Float())
			return reflect.ValueOf(&val)
		case reflect.Float64:
			val := srcField.Float()
			return reflect.ValueOf(&val)
		default:
			panic("unsupported pointer to float type conversion")
		}
	default:
		panic("unsupported float type conversion")
	}
}

func (h *GRPCHandler) convertString(srcField reflect.Value, destType reflect.Type) (reflect.Value, error) {
	if srcField.Kind() == reflect.Ptr && srcField.IsNil() {
		// Handle nil pointer case
		return reflect.Zero(destType), nil
	}

	switch {
	case srcField.Kind() == reflect.String && destType.Kind() == reflect.String:
		// Directly copy string to string
		return srcField, nil

	case srcField.Kind() == reflect.String && destType == reflect.PointerTo(reflect.TypeOf("")):
		// Convert string to *string
		strValue := srcField.Interface().(string)
		return reflect.ValueOf(&strValue), nil

	case srcField.Kind() == reflect.Ptr && srcField.Type() == reflect.PointerTo(reflect.TypeOf("")) && destType.Kind() == reflect.String:
		// Convert *string to string
		return reflect.ValueOf(srcField.Elem().Interface().(string)), nil

	case srcField.Kind() == reflect.Ptr && srcField.Type() == reflect.PointerTo(reflect.TypeOf("")) && destType == reflect.PointerTo(reflect.TypeOf("")):
		// Copy *string to *string
		return srcField, nil

	default:
		return reflect.Zero(destType), fmt.Errorf("unsupported string conversion from %v to %v", srcField.Type(), destType)
	}
}

func reflectValueToRawMessage(v reflect.Value, name string) (json.RawMessage, error) {
	// Check if the reflect.Value is valid and of kind Slice or Array
	if !v.IsValid() {
		return nil, fmt.Errorf("invalid reflect.Value")
	}

	// Ensure it's of kind Slice or Array
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("expected slice or array, got %s", v.Kind())
	}

	// Convert reflect.Value to []byte
	bytes, err := json.Marshal(v.Interface())
	if err != nil {
		return nil, fmt.Errorf("error marshaling value %s: %w", name, err)
	}

	// Return as json.RawMessage
	return json.RawMessage(bytes), nil
}

func ConvertStringToRawMessage(v reflect.Value) (json.RawMessage, error) {
	if v.Kind() != reflect.String {
		return nil, fmt.Errorf("expected a string type but got %s", v.Kind())
	}

	// Extract the string value
	jsonString := v.Interface()
	// Unmarshal the JSON string into json.RawMessage
	var raw json.RawMessage
	raw, err := json.Marshal(jsonString)
	if err != nil {
		return nil, fmt.Errorf("error converting string to json.RawMessage: %v", err)
	}

	return raw, nil
}
