package models

//TemplateData hold data used to transfer data to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	FlashMsg  string
	Warning   string
	Error     string
}
