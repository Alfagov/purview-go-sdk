package utils

import "github.com/Alfagov/pwCatalog/src/models"

func ValidateGuid(guid string) error {
	if guid == "" {
		return models.EmptyGuidError
	}
	return nil
}

func ValidateGuids(guids []string) error {
	if guids == nil {
		return models.EmptyGuidError
	}
	return nil
}

func ValidateClassificationName(name string) error {
	if name == "" {
		return models.ClassificationsError
	}
	return nil
}

func ValidateAttributes(attributes map[string]string) error {
	if attributes == nil {
		return models.NoAttributesProvided
	}
	return nil
}

func ValidateTypeName(typeName string) error {
	if typeName == "" {
		return models.EmptyTypeNameError
	}
	return nil
}

func ValidateData(data interface{}) error {
	if data == nil {
		return models.NoDataProvidedError
	}
	return nil
}

func ValidateQueryAttributes(attributes *models.QueryAttributes) error {
	if attributes == nil {
		return models.NoAttributesProvided
	}

	if attributes.TypeName == "" {
		return models.EmptyTypeNameError
	}

	if attributes.Attributes == nil {
		return models.NoAttributesProvided
	}
	return nil
}
