package domain

type NluReq struct {
	Text       string `json:"text"`
	TextOrigin string `json:"textOrigin"`
}

type NluResp struct {
	Text   string      `json:"text"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

type RasaResp struct {
	Entities         []Entity         `json:"entities"`
	Intent           Intent           `json:"intent"`
	IntentRanking    []IntentRanking  `json:"intent_ranking"`
	ResponseSelector ResponseSelector `json:"response_selector"`
	Text             string           `json:"text"`
	TextOrigin       string           `json:"textOrigin"`
}

type Entity struct {
	ConfidenceEntity float64  `json:"confidence_entity"`
	End              int64    `json:"end"`
	Entity           string   `json:"entity"`
	Extractor        string   `json:"extractor"`
	Processors       []string `json:"processors"`
	Start            int64    `json:"start"`
	Value            string   `json:"value"`
	ValueOrigin      string   `json:"valueOrigin"`
}

type Intent struct {
	Confidence float32 `json:"confidence"`
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
}
type IntentRanking struct {
	Confidence int64  `json:"confidence"`
	ID         int64  `json:"id"`
	Name       string `json:"name"`
}
type Default struct {
	Ranking  []interface{} `json:"ranking"`
	Response Response      `json:"response"`
}
type ResponseSelector struct {
	AllRetrievalIntents []interface{} `json:"all_retrieval_intents"`
	Default             Default       `json:"default"`
}
type Response struct {
	Confidence        int64       `json:"confidence"`
	ID                interface{} `json:"id"`
	IntentResponseKey interface{} `json:"intent_response_key"`
	ResponseTemplates interface{} `json:"response_templates"`
	Responses         interface{} `json:"responses"`
	TemplateName      string      `json:"template_name"`
	UtterAction       string      `json:"utter_action"`
}

func (resp *NluResp) SetResult(result interface{}) {
	//bytes, _ := json.Marshal(msg)
	//resp.Result = string(bytes)

	resp.Result = result
}
