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

	for i := range *props {
		prop := (*props)[i]
		compProperty := Property{
			Property: &prop,
		}
		buffer := (2 - len(exclude))
		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(PropertyGetExcludeValues, exclude) {
			go func() {
				values, err := db.GetPropertyValues(h, prop.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- values
			}()
		}

		if !misc.InSlice(PropertyGetExcludeExtras, exclude) {
			go func() {
				extras, err := db.GetPropertyValueExtras(h, prop.ID)
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

		for i := range *propertyValues {
			value := (*propertyValues)[i]
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

func GetPropertiesByTraits[H db.Handle](h H, collectionID string, traits map[string]string, exclude ...PropertyGetExclude) (*[]Property, error) {
	const op errors.Op = "GetPropertiesByAssetID"
	var properties []Property

	traitKeys := []string{}
	for key := range traits {
		traitKeys = append(traitKeys, key)
	}

	props, err := db.GetPropertiesWhereNameIn(h, collectionID, traitKeys...)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for i := range *props {
		prop := (*props)[i]
		compProperty := Property{
			Property: &prop,
		}

		valueName, ok := traits[prop.Name]
		if !ok {
			return nil, errors.E(op, fmt.Errorf("trait %s not found", prop.Name))
		}

		buffer := (2 - len(exclude))
		rChan := make(chan interface{}, buffer)

		if !misc.InSlice(PropertyGetExcludeValues, exclude) {
			go func() {
				values, err := db.GetPropertyValueByName(h, prop.ID, valueName)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}

				rChan <- values
			}()
		}

		if !misc.InSlice(PropertyGetExcludeExtras, exclude) {
			go func() {
				extras, err := db.GetPropertyValueExtrasByName(h, prop.ID, valueName)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- extras
			}()
		}

		var propertyValues []db.PropertyValue
		var propertyValueExtras []db.PropertyValueExtras
		var errs []error
		for i := 0; i < buffer; i++ {
			data := <-rChan
			switch result := data.(type) {
			case *db.PropertyValue:
				if result != nil {
					propertyValues = append(propertyValues, *result)
				}
			case *[]db.PropertyValueExtras:
				propertyValueExtras = *result
			case error:
				errs = append(errs, result)
			}
		}
		close(rChan)

		if len(errs) > 0 {
			msg := ""
			for _, err := range errs {
				msg += err.Error() + "\n"
			}

			return nil, errors.E(op, fmt.Errorf(msg))
		}

		for i := range propertyValues {
			value := propertyValues[i]
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

func GetCumulativeProperties[H db.Handle](h H, collectionID string, exclude ...PropertyGetExclude) (*[]Property, error) {

	ids, err := GetCollectionAssetIDs(h, collectionID)
	if err != nil {
		return nil, err
	}

	if ids == nil || (ids != nil && !(len(*ids) > 0)) {
		return nil, errors.E("no assets found in collection")
	}

	traits, err := db.GetAssetPropertiesInAsMap(h, *ids)
	if err != nil {
		return nil, err
	}

	if len(traits) == 0 {
		return nil, errors.E("no traits found")
	}

	declaredProperties, err := GetProperties(h, collectionID)
	if err != nil {
		return nil, err
	}
	// has declared properties
	hdp := declaredProperties != nil && len(*declaredProperties) > 0

	interumProperties := []Property{}

	for key, values := range traits {

		prop := Property{
			Property: &db.Property{
				Name: key,
			},
			Values: []PropertyValue{},
		}

		for _, value := range values {
			if hdp {
				for _, declaredProperty := range *declaredProperties {
					if declaredProperty.Name == key {
						for _, declaredValue := range declaredProperty.Values {
							if declaredValue.Name == value {
								prop.Values = append(prop.Values, declaredValue)
							}
						}
					}
				}
			} else {
				prop.Values = append(prop.Values, PropertyValue{
					PropertyValue: &db.PropertyValue{
						Name: value,
					},
				})
			}

		}

		interumProperties = append(interumProperties, prop)
	}

	return &interumProperties, nil
}
