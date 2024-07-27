package common

type Stack struct {
	StackName     string
	DefaultBranch string
	Branches      []Branch
}

type Branch struct {
	BranchName string
	Priority   int
}

type ByPriority []Branch

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
