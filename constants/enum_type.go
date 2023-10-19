package constants

type GenderType string
type MembershipType string
type StatusTransactions string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
)

const (
	Silver   MembershipType = "siler"
	Gold     MembershipType = "gold"
	Platinum MembershipType = "platinum"
)

const (
	Failed  StatusTransactions = "failed"
	Pending StatusTransactions = "gold"
	Success StatusTransactions = "success"
)
