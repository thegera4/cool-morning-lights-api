package utils

// Validates the data of a store that is going to be updated.
func StoreUpdateIsValid(store map[string]interface{}) bool {
	for key, value := range store {
		switch v := value.(type) {
			case string:
				if v == "" { return false }
			case int:
				if v <= 0 { return false }
			case []interface{}:
				if len(v) == 0 { return false }
		}

		// Validations for name
		if key == "name" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for address
		if key == "address" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for zipCode
		if key == "zipCode" {
			if _, ok := value.(float64); !ok { // value is not a number
				return false
			}
			if zipCode, ok := value.(float64); ok { // value is a number but is less than or equal to 0
				if zipCode <= 0 {
					return false
				}
			} 
		}

		// Validations for city
		if key == "city" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for state
		if key == "state" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for phone
		if key == "phone" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for email
		if key == "email" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for openTime
		if key == "openTime" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for closeTime
		if key == "closeTime" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for workingDays
		if key == "workingDays" {
			if _, ok := value.([]interface{}); !ok { // value is not an slice of strings
				return false
			}
		}

		// Validations for active
		if key == "active" {
			if _, ok := value.(bool); !ok {
				return false
			}
		}
	}

	return true
}