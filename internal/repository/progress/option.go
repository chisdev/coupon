package progress

type ProgrestOption struct {
	Progress     int32
	PassCount    int32
	MilestoneIds []uint64
	CustomerIds  []string
}

type Option interface {
	Apply(*ProgrestOption)
}

type funcOption func(*ProgrestOption)

func (f funcOption) Apply(o *ProgrestOption) {
	f(o)
}

func WithProgress(progress int32) Option {
	return funcOption(func(po *ProgrestOption) {
		po.Progress = progress
	})
}

func WithPassCount(passCount int32) Option {
	return funcOption(func(po *ProgrestOption) {
		po.PassCount = passCount
	})
}

func WithCustomerIDs(ids []string) Option {
	return funcOption(func(po *ProgrestOption) {
		po.CustomerIds = ids
	})
}

func WithMilestoneIds(ids []uint64) Option {
	return funcOption(func(po *ProgrestOption) {
		po.MilestoneIds = ids
	})
}
