package compound

import (
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Property struct {
	*db.Property
	Values []PropertyValue `json:"values,omitempty"`
}

type PropertyValue struct {
	*db.PropertyValue
	Extras map[string]string `json:"extras,omitempty"`
}

type PropertyGetExclude string

const (
	PropertyGetExcludeValues PropertyGetExclude = "values"
	PropertyGetExcludeExtras PropertyGetExclude = "extras"
)

func GetProperties[H db.Handle](h H, id string, exclude ...PropertyGetExclude) (*[]Property, error) {
	const op errors.Op = "GetProperties"
	var properties []Property

	props, err := db.GetProperties(h, id)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for _, prop := range *props {
		compProperty := Property{
			Property: &prop,
		}
		buffer := (2 - len(exclude))
		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(PropertyGetExcludeValues, exclude) {
			go func() {
				values, err := db.GetPropertyValues(h, id)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- values
			}()
		}

		if !misc.InSlice(PropertyGetExcludeExtras, exclude) {
			go func() {
				extras, err := db.GetPropertyValueExtras(h, id)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- extras
			}()
		}

		var propertyValues *[]db.PropertyValue
		var propertyValueExtras []db.PropertyValueExtras
		var errs []error
		for i := 0; i < buffer; i++ {
			data := <-rChan
			switch result := data.(type) {
			case *[]db.PropertyValue:
				propertyValues = result
			case *[]db.PropertyValueExtras:
				propertyValueExtras = *result
			case error:
				errs = append(errs, result)
			}
		}

		if len(errs) > 0 {
			msg := ""
			for _, err := range errs {
				msg += err.Error() + "\n"
			}

			return nil, errors.E(op, fmt.Errorf(msg))
		}

		for _, value := range *propertyValues {
			compPropertyValue := PropertyValue{
				PropertyValue: &value,
				Extras:        make(map[string]string),
			}

			for _, extra := range propertyValueExtras {
				if extra.Name == value.Name {
					compPropertyValue.Extras[extra.Key] = extra.Value
				}
			}

			compProperty.Values = append(compProperty.Values, compPropertyValue)
		}

		properties = append(properties, compProperty)
	}

	return &properties, nil
}
