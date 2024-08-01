package service

import "github.com/tillpaid/paysera-log-time-golang/internal/model"

func GetWorkLogs() []model.WorkLog {
	return []model.WorkLog{
		{Number: 1, OriginalTime: "1h40m", ModifiedTime: "1h40m", IssueNumber: "TIME-505", Description: "MEETING- Daily- Call with Pashkevich about the new ticket"},
		{Number: 2, OriginalTime: "20m", ModifiedTime: "25m", IssueNumber: "COMP-800", Description: "OTHER"},
		{Number: 3, OriginalTime: "40m", ModifiedTime: "50m", IssueNumber: "COMP-766", Description: "REVIEW"},
		{Number: 4, OriginalTime: "40m", ModifiedTime: "50m", IssueNumber: "COMP-781", Description: "REVIEW"},
		{Number: 5, OriginalTime: "20m", ModifiedTime: "25m", IssueNumber: "COMP-774", Description: "CODING- Code review fixes"},
		{Number: 6, OriginalTime: "2h40m", ModifiedTime: "3h50m", IssueNumber: "COMP-782", Description: "CODING- Retested all the changes- New code changes after final tests, found issue with checking data and some another long text"},
		{Number: 7, OriginalTime: "1h40m", ModifiedTime: "1h40m", IssueNumber: "TIME-505", Description: "MEETING- Daily- Call with Pashkevich about the new ticket"},
		{Number: 8, OriginalTime: "20m", ModifiedTime: "25m", IssueNumber: "COMP-800", Description: "OTHER"},
		{Number: 9, OriginalTime: "40m", ModifiedTime: "50m", IssueNumber: "COMP-766", Description: "REVIEW"},
		{Number: 10, OriginalTime: "40m", ModifiedTime: "50m", IssueNumber: "COMP-781", Description: "REVIEW"},
		{Number: 11, OriginalTime: "20m", ModifiedTime: "25m", IssueNumber: "COMP-774", Description: "CODING- Code review fixes"},
		{Number: 12, OriginalTime: "2h40m", ModifiedTime: "3h50m", IssueNumber: "COMP-782", Description: "CODING- Retested all the changes- New code changes after final tests, found issue with checking data and some another long text"},
	}
}
