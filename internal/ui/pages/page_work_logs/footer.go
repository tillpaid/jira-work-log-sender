package page_work_logs

func GetFooter(delimiter string) []string {
	helpRow := "Action keys: R-Reload | L-Dump work logs | [Q/Space/Return/Esc]-Exit"

	return []string{delimiter, helpRow, delimiter}
}
