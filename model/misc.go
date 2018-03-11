package model

func FetchUserCount() int32 {
	query := "SELECT COUNT(id) from users"
	var count int32
	if err := DBInstance.QueryRow(query).Scan(&count); err != nil {
		return -1
	}
	return count
}

func FetchCellCount() int64 {
	query := "SELECT COUNT(id) from cells"
	var count int64
	if err := DBInstance.QueryRow(query).Scan(&count); err != nil {
		return -1
	}
	return count
}
