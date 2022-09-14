package models

// WhereData: struct to old the params given to the get users endpoints
type WhereData struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// BuildWHereQuery: function to build the query according to the where clauses
func BuildWHereQuery(wds []WhereData) (query string, args []interface{}) {

	for i, w := range wds {
		if i > 0 {
			query += " and "
		}
		query += w.fieldToWhere()
		args = append(args, "%"+w.Value+"%")
	}
	return query, args
}

// fieldToWhere: convenience function to make the where data to a string in the query.
func (wd WhereData) fieldToWhere() string {

	//This is a trade off, would you prefer to match(=) or like the strings. I think the like is way more convenient for user searches.
	//return wd.Field + " = ?"

	return wd.Field + " LIKE ?"
}
