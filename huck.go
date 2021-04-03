package huck

var kernel *HuckKernel = NewHuckKernel()

type HuckKernel struct {
	stat map[string]DataStatistical
}

func NewHuckKernel() *HuckKernel {
	return &HuckKernel{
		stat: make(map[string]DataStatistical),
	}
}
