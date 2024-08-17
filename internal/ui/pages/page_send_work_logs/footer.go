package page_send_work_logs

func GetFooter(delimiter string) []string {
	helpRow := "Action keys: R-Reload | L-Send work logs (double press) | [Q/Space/Return/Esc]-Exit"

	return []string{delimiter, helpRow, delimiter}
}
