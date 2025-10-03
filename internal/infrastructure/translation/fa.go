package translation

var Persian = map[string]interface{}{
	"name": "نام",
	"errors": map[string]interface{}{
		"required": "فیلد {0} اجباری است.",
		"generic":  "مشکلی پیش آمده است. لطفا دوباره تلاش کنید.",
		"numeric":  "`{0}` باید عدد باشد.",
	},
	"success": map[string]interface{}{
		"hello": "سلام {0}",
		"add":   "با موفقیت اضافه شد.",
	},
}
