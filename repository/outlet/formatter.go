package outlet

type OutletFormatter struct {
	ID 				int `json:"id"`
	Name 			string `json:"name"`
	Description 	string `json:"description"`
}

// func FormatterOutlet (outlet Outlet) OutletFormatter {
// 	formatter := OutletFormatter{
// 		ID:				outlet.ID,
// 		Name:			outlet.Name,
// 		Description:	outlet.Description,
// 	}

// 	return formatter
// }

// func FormatOutlets(outlets []Outlet) []OutletFormatter {
// 	outletsFormatter := []OutletFormatter{}

// 	if len(outlets) == 0 {
// 		return []OutletFormatter{}
// 	}

// 	for _, outlet := range outlets {
// 		FormatterOutlet(outlet)
// 	}

// 	return outletsFormatter
// }
