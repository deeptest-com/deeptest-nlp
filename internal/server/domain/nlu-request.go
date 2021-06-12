package domain

type NluReq struct {
	Text string `json:"text"`
}

type NluResp struct {
	Text   string      `json:"text"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type RasaResp struct {
	Entities []struct {
		ConfidenceEntity float64  `json:"confidence_entity"`
		End              int64    `json:"end"`
		Entity           string   `json:"entity"`
		Extractor        string   `json:"extractor"`
		Processors       []string `json:"processors"`
		Start            int64    `json:"start"`
		Value            string   `json:"value"`
	} `json:"entities"`
	Intent struct {
		Confidence float32 `json:"confidence"`
		ID         int64   `json:"id"`
		Name       string  `json:"name"`
	} `json:"intent"`
	IntentRanking []struct {
		Confidence int64  `json:"confidence"`
		ID         int64  `json:"id"`
		Name       string `json:"name"`
	} `json:"intent_ranking"`
	ResponseSelector struct {
		AllRetrievalIntents []interface{} `json:"all_retrieval_intents"`
		Default             struct {
			Ranking  []interface{} `json:"ranking"`
			Response struct {
				Confidence        int64       `json:"confidence"`
				ID                interface{} `json:"id"`
				IntentResponseKey interface{} `json:"intent_response_key"`
				ResponseTemplates interface{} `json:"response_templates"`
				Responses         interface{} `json:"responses"`
				TemplateName      string      `json:"template_name"`
				UtterAction       string      `json:"utter_action"`
			} `json:"response"`
		} `json:"default"`
	} `json:"response_selector"`
	Text string `json:"text"`
}
