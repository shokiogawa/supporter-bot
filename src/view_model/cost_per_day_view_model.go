package view_model

type CostPerDayListViewModel struct {
	CostPerDay []CostPerDayViewModel `json:"cost_per_day"`
}

type CostPerDayViewModel struct {
	Title   string `json:"title"`
	OutCome int    `json:"outCome"`
}
