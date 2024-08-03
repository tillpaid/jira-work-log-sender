package page_dump_work_logs

func GetFooter(delimiter string) []string {
	helpRow := "Action keys: R-Back to work logs table | [Q/Space/Return/Esc]-Exit"

	return []string{delimiter, helpRow, delimiter}
}
