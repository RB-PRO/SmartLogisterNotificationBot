package direct

import (
	"os"

	"github.com/dogenzaka/tsv"
)

type AnswerTSV struct {
	Cost         int    `tsv:"Cost"`         // Расход с НДС, ₽
	CampaignName string `tsv:"CampaignName"` // Название рекламной компании
}

func UnwrapTSV(FileName string) ([]AnswerTSV, error) {
	file, ErrOpen := os.Open(FileName)
	if ErrOpen != nil {
		return nil, ErrOpen
	}
	defer file.Close()

	var datas []AnswerTSV
	data := AnswerTSV{}
	parser, ErrNewParser := tsv.NewParser(file, &data)
	if ErrNewParser != nil {
		return nil, ErrNewParser
	}

	for {
		eof, ErrNext := parser.Next()
		if eof {
			break
		}
		if ErrNext != nil {
			return nil, ErrNext
		}

		datas = append(datas, data)
	}
	return datas, nil
}
