package db

func getTable(obj interface{}) string {
	switch obj.(type) {
	default:
		return ""
	}
}
