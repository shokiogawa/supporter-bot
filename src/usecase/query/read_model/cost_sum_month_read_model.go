package read_model

type CostSumReamModel struct {
	OutCome int
	Date    string
}

type CostSumListReadModel struct {
	LineUserId  string
	CostSumList []CostSumReamModel
}
