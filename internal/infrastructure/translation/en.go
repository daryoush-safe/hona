package translation

var English = map[string]interface{}{
	"name": "name",
	"errors": map[string]interface{}{
		"required": "The {0} is required.",
		"generic":  "An error occurred, please try again.",
		"numeric":  "The {0} should be a numeric value.",
	},
	"success": map[string]interface{}{
		"hello": "hello {0}",
		"add":   "Added successfully.",
	},
}
