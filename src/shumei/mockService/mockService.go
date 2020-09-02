package mockService

import (
	"shumei/config"
	"shumei/log"
	"shumei/mockService/prediction"
)

type RecordService struct {
	Conf   *config.Config
	Logger *log.Log
}

func (this *RecordService) Predict(request *prediction.PredictRequest) (*prediction.PredictResult_, error) {
	var score int32
	riskLevel := "PASS"
	detail := "{\"code\":1100,\"message\":\"成功\"}"
	score = 0

	preRes := &prediction.PredictResult_{
		Score:     &score,
		RiskLevel: &riskLevel,
		Detail:    &detail}

	return preRes, nil
}
