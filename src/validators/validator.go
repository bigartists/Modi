package validators

import (
	"errors"
	"github.com/bigartists/Modi/src/handler"
	"github.com/bigartists/Modi/src/reason"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"log/slog"
	"reflect"
	"strings"
	"unicode"
)

// MyValidator my validator
type MyValidator struct {
	Validate *validator.Validate
}

var GlobalValidator *MyValidator

func init() {
	val := createDefaultValidator()
	GlobalValidator = &MyValidator{Validate: val}
}

// FormErrorField indicates the current form error content. which field is error and error message.
type FormErrorField struct {
	ErrorField string `json:"error_field"`
	ErrorMsg   string `json:"error_msg"`
}

func createDefaultValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("notblank", NotBlank)
	_ = validate.RegisterValidation("sanitizer", Sanitizer)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) (res string) {

		if jsonTag := fld.Tag.Get("json"); len(jsonTag) > 0 {
			if jsonTag == "-" {
				return ""
			}
			return jsonTag
		}
		if formTag := fld.Tag.Get("form"); len(formTag) > 0 {
			return formTag
		}
		return fld.Name
	})
	return validate
}

func NotBlank(fl validator.FieldLevel) (res bool) {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		trimSpace := strings.TrimSpace(field.String())
		res := len(trimSpace) > 0
		if !res {
			field.SetString(trimSpace)
		}
		return true
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func Sanitizer(fl validator.FieldLevel) (res bool) {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		filter := bluemonday.UGCPolicy()
		content := strings.Replace(filter.Sanitize(field.String()), "&amp;", "&", -1)
		field.SetString(content)
		return true
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

// Check /
func (m *MyValidator) Check(value interface{}) (errFields []*FormErrorField, err error) {
	defer func() {
		if len(errFields) == 0 {
			return
		}
		for _, field := range errFields {
			if len(field.ErrorField) == 0 {
				continue
			}
			firstRune := []rune(field.ErrorMsg)[0]
			if !unicode.IsLetter(firstRune) || !unicode.Is(unicode.Latin, firstRune) {
				continue
			}
			upperFirstRune := unicode.ToUpper(firstRune)
			field.ErrorMsg = string(upperFirstRune) + field.ErrorMsg[1:]
			if !strings.HasSuffix(field.ErrorMsg, ".") {
				field.ErrorMsg += "."
			}
		}
	}()
	err = m.Validate.Struct(value)
	if err != nil {
		var valErrors validator.ValidationErrors
		if !errors.As(err, &valErrors) {

			slog.Error("validate check exception", "error", err)
			return nil, errors.New("validate check exception")
		}

		for _, fieldError := range valErrors {
			errField := &FormErrorField{
				ErrorField: fieldError.Field(),
				ErrorMsg:   "invalid value",
			}

			// get original tag name from value for set err field key.
			structNamespace := fieldError.StructNamespace()
			_, fieldName, found := strings.Cut(structNamespace, ".")
			if found {
				originalTag := getObjectTagByFieldName(value, fieldName)
				if len(originalTag) > 0 {
					errField.ErrorField = originalTag
				}
			}
			errFields = append(errFields, errField)
		}
		if len(errFields) > 0 {
			errMsg := ""
			if len(errFields) == 1 {
				errMsg = errFields[0].ErrorMsg
			}
			return errFields, handler.NewCustomError().BadRequest(reason.RequestFormatError).WithMsg(errMsg)
		}
	}

	if v, ok := value.(Checker); ok {
		errFields, err = v.Check()
		if err == nil {
			return nil, nil
		}
		errMsg := ""
		for _, errField := range errFields {
			errMsg = errField.ErrorMsg
		}
		return errFields, handler.NewCustomError().BadRequest(reason.RequestFormatError).WithMsg(errMsg)
	}
	return nil, nil
}

// Checker .
type Checker interface {
	Check() (errField []*FormErrorField, err error)
}

func getObjectTagByFieldName(obj interface{}, fieldName string) (tag string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("getObjectTagByFieldName panic", "error", err)
		}
	}()

	objT := reflect.TypeOf(obj)
	objT = objT.Elem()

	structField, exists := objT.FieldByName(fieldName)
	if !exists {
		return ""
	}
	tag = structField.Tag.Get("json")
	if len(tag) == 0 {
		return structField.Tag.Get("form")
	}
	return tag
}
