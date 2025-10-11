package exceptions

type RatelimitError struct {

}


func (fe *RatelimitError) Error() string {
	return ""
}

func NewRatelimitError() *RatelimitError {
	return &RatelimitError{
	}
}