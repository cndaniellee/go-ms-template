package es

type Properties map[string]map[string]any

type IndexMapping struct {
	Settings struct {
		NumberOfShards   int `json:"number_of_shards"`
		NumberOfReplicas int `json:"number_of_replicas"`
	} `json:"settings"`
	Mappings struct {
		Properties `json:"properties"`
	} `json:"mappings"`
}

const Analyzer = "ik_max_word"
