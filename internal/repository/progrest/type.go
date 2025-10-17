package progress

type ProgressEntity struct {
	CustomerId  string
	MilestoneId uint64
	PassCount   int32
	Progress    int32
}
