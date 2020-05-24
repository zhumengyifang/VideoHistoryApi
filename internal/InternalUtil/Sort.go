package InternalUtil

import (
	"gindemo/internal/Model/RedisModel"
	"sort"
)

func SortBySubmitDate(infos []*RedisModel.HistoryInfoParameter) []*RedisModel.HistoryInfoParameter {
	if infos == nil {
		return nil
	}

	sort.Sort(RedisModel.HistoryInfoParameters(infos))
	return infos
}
