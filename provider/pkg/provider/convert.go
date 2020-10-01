// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types map[string]AzureAPIType
}

func (k *SdkShapeConverter) sdkPropertyToRequest(prop *AzureAPIProperty, value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		// For union types, iterate through types and find the first one that matches the shape.
		for _, t := range prop.OneOf {
			typeName := strings.TrimPrefix(t, "#/types/")
			typ := k.Types[typeName]
			request := k.SdkPropertiesToRequestBody(typ.Properties, value)
			if request != nil {
				return request
			}
		}

		if prop.Ref == "" {
			// Return untyped dictionaries as-is.
			return value
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ := k.Types[typeName]
		return k.SdkPropertiesToRequestBody(typ.Properties, value)
	case []interface{}:
		var result []interface{}
		for _, item := range value {
			result = append(result, k.sdkPropertyToRequest(prop.Items, item))
		}
		return result
	}
	return value
}

// SdkPropertiesToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
func (k *SdkShapeConverter) SdkPropertiesToRequestBody(props map[string]AzureAPIProperty,
	values map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, has := values[sdkName]; has {
			if prop.Const != nil && value != prop.Const {
				return nil
			}

			if prop.Container == "" {
				result[name] = k.sdkPropertyToRequest(&p, value)
			} else {
				// "Unflatten" the property into a container property.
				container := map[string]interface{}{}
				if v, ok := result[prop.Container]; ok {
					if v, ok := v.(map[string]interface{}); ok {
						container = v
					}
				}
				container[name] = k.sdkPropertyToRequest(&p, value)
				result[prop.Container] = container
			}
		}
	}
	return result
}

func (k *SdkShapeConverter) bodyPropertyToSdk(prop *AzureAPIProperty, value interface{}) interface{} {
	if prop == nil {
		return value
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		// For union types, iterate through types and find the first one that matches the shape.
		for _, t := range prop.OneOf {
			typeName := strings.TrimPrefix(t, "#/types/")
			typ := k.Types[typeName]
			response := k.BodyPropertiesToSDK(typ.Properties, value.(map[string]interface{}))
			if response != nil {
				return response
			}
		}

		if prop.Ref == "" {
			// Return untyped dictionaries as-is.
			return value
		}

		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		typ, ok := k.Types[typeName]
		if !ok {
			return value.(map[string]interface{})
		}

		return k.BodyPropertiesToSDK(typ.Properties, value.(map[string]interface{}))
	case reflect.Slice, reflect.Array:
		var result []interface{}
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			if prop.Items != nil {
				result = append(result, k.bodyPropertyToSdk(prop.Items, s.Index(i).Interface()))
			} else {
				result = append(result, s.Index(i).Interface())
			}
		}
		return result
	}
	return value
}

// BodyPropertiesToSDK converts a JSON request- or response body to the SDK shape.
func (k *SdkShapeConverter) BodyPropertiesToSDK(props map[string]AzureAPIProperty,
	response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		values := response
		if prop.Container != "" {
			if v, has := values[prop.Container]; has {
				if v, ok := v.(map[string]interface{}); ok {
					values = v
				}
			}
		}

		if value, has := values[name]; has {
			if prop.Const != nil && value != prop.Const {
				return nil
			}

			if value != nil {
				result[sdkName] = k.bodyPropertyToSdk(&p, value)
			}
		}
	}
	return result
}

// ResponseToSdkInputs calculates a map of input values that would produce the given resource path and
// response. This is useful when we need to import an existing resource based on its current properties.
func (k *SdkShapeConverter) ResponseToSdkInputs(parameters []AzureAPIParameter,
	pathValues map[string]string, response map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, param := range parameters {
		switch {
		case strings.EqualFold(param.Name, "subscriptionId"):
			// Ignore
		case param.Location == "path":
			name := param.Name
			result[name] = pathValues[name]
		case param.Location == body:
			bodyProps := k.BodyPropertiesToSDK(param.Body.Properties, response)
			for k, v := range bodyProps {
				switch {
				case k == "id":
					// Some resources have a top-level `id` property which is (probably incorrectly) marked as
					// non-readonly. `id` is a special property to Pulumi and will always cause diffs if set in
					// the result of a Read operation and block import. So, don't copy it to inputs.
					continue
				default:
					// Attempt to exclude insignificant properties from the inputs. A resource response would
					// contain a lot of default values, e.g. empty arrays when no values were specified, empty
					// strings, or false booleans. The decision to remove them is somewhat arbitrary but it
					// seems to make the practical import experience smoother.
					result[k] = removeDefaultValues(v)
				}
			}
		}
	}
	return result
}

// parseResourceID extracts templated values from the given resource ID based on the names of those templated
// values in an HTTP path. The structure of id and path must match: same segment count and segment names.
func parseResourceID(id, path string) (map[string]string, error) {
	pathParts := strings.Split(path, "/")
	idParts := strings.Split(id, "/")
	if len(pathParts) != len(idParts) {
		return nil, errors.Errorf("length of id doesn't match the path: '%s' vs. '%s'", id, path)
	}

	result := map[string]string{}
	for i, s := range pathParts {
		value := idParts[i]
		switch {
		case strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}"):
			name := s[1 : len(s)-1]
			result[name] = value
		case !strings.EqualFold(s, value):
			return nil, errors.Errorf("failed parsing id: %s <> %s", s, value)
		}
	}
	return result, nil
}

// removeDefaultValues returns nil if the given value is a default for its type (e.g. `false`, or an
// empty string). It also applies this recursively for values in arrays and maps.
func removeDefaultValues(value interface{}) interface{} {
	switch value := value.(type) {
	case map[string]interface{}:
		result := map[string]interface{}{}
		for k, v := range value {
			resultValue := removeDefaultValues(v)
			if resultValue != nil {
				result[k] = resultValue
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, v := range value {
			resultValue := removeDefaultValues(v)
			if resultValue != nil {
				result = append(result, resultValue)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case bool:
		if !value {
			return nil
		}
	case string:
		if len(value) == 0 {
			return nil
		}
	}
	return value
}
