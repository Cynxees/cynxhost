package helper

import (
	"bytes"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response/responsecode"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"text/template"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidateRequest(r *http.Request, dst interface{}, v *validator.Validate) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return errors.New("invalid request payload: " + err.Error())
	}

	if err := v.Struct(dst); err != nil {
		return errors.New("validation failed: " + err.Error())
	}

	return nil
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to decode request: "+err.Error(), http.StatusBadRequest)
	}
}

func GetResponseCodeName(code responsecode.ResponseCode) string {
	if name, exists := responsecode.ResponseCodeNames[code]; exists {
		return name
	}
	return "Unknown Code"
}

// replacePlaceholders replaces {{}} placeholders in a script with real values.
func ReplacePlaceholders(script string, variables map[string]string) (string, error) {
	tmpl, err := template.New("script").Parse(script)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, variables); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return output.String(), nil
}

func GetClientIP(r *http.Request) string {
	// If the request is behind a reverse proxy, the IP address might be forwarded in the X-Forwarded-For header.
	// First, check for the X-Forwarded-For header.
	ips := r.Header.Get("X-Forwarded-For")
	if ips != "" {
		// The X-Forwarded-For header contains a comma-separated list of IPs
		// The first IP in the list is the original client IP.
		return strings.Split(ips, ",")[0]
	}

	// Otherwise, fallback to the remote address.
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func FormatServerTemplateVariableKey(key string) string {
	return "VARIABLE_" + strings.ToUpper(key)
}

func FormatScriptVariables(scriptVariables entity.Variables, variableInput []request.ServerTemplateScriptVariable) (map[string]string, error) {
	result := make(map[string]string)
	reqVariableMap := make(map[string]string)
	missingVariables := []string{}

	// Map input variables by name for easy lookup
	for _, reqVar := range variableInput {
		fmt.Println("Mapping: ", reqVar.Name, " - ", reqVar.Value)
		reqVariableMap[reqVar.Name] = reqVar.Value
	}

	fmt.Println("Request variables: ", reqVariableMap)
	fmt.Println("Script variables: ", scriptVariables)
	// Process each required variable
	for _, variable := range scriptVariables {
		fmt.Println("Variable name: ", variable.Name)

		// Check if the variable name exists in the input
		reqValue, found := reqVariableMap[variable.Name]
		if !found {
			// Add to missing variables if not found
			missingVariables = append(missingVariables, variable.Name)
			continue
		}

		fmt.Println("Request value: ", reqValue)

		// Check if the value matches any content in the variable
		valueSet := false
		for _, content := range variable.Content {
			if content.Name == reqValue {
				fmt.Println("Found value: ", content.Value)

				// Add each key-value pair from content.Value into the result
				for key, val := range content.Value {
					result[FormatServerTemplateVariableKey(key)] = fmt.Sprintf("%v", val) // Convert value to string
				}

				valueSet = true
				break
			}
		}

		// If no matching content was found, treat it as missing
		if !valueSet {
			missingVariables = append(missingVariables, variable.Name)
		}
	}

	// Return an error if any variables are missing
	if len(missingVariables) > 0 {
		return nil, fmt.Errorf("missing or invalid variables: %v", missingVariables)
	}

	return result, nil
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func StructToMapStringArray(data interface{}) ([]map[string]string, error) {
	var result []map[string]string

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}
