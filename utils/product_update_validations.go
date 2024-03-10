package utils

// Validates the data of a product that is going to be updated.
func ProductUpdateIsValid(product map[string]interface{}) bool {
	for key, value := range product {
		switch v := value.(type) {
		case string:
			if v == "" { return false }
		case int:
			if v <= 0 && key != "price" { return false }
		case []interface{}:
			if len(v) == 0 { return false }
		}

		// Validations for name
		if key == "name" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for description
		if key == "description" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for price
		if key == "price" {
			if _, ok := value.(float64); !ok { // value is not a number
				return false
			}
			if price, ok := value.(float64); ok { // value is a number but is less than or equal to 0
				if price <= 0 {
					return false
				}
			} 
		}

		// Validations for stock
		if key == "stock" {
			if _, ok := value.(float64); !ok { // value is not a number
				return false
			}
			if stock, ok := value.(float64); ok { // value is a number but is less than 0
				if stock < 0 {
					return false
				}
			}
		}

		// Validations for store
		if key == "store" {
			if _, ok := value.(string); !ok { // value is not a string
				return false
			}
		}

		// Validations for pictures
		if key == "pictures" {
			if _, ok := value.([]string); !ok { // value is not a []string
				return false
			}
		}

		// Validations for categories
		if key == "categories" {
			if _, ok := value.([]string); !ok { // value is not a []string
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