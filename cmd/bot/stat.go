package main

type statistics struct {
	UsersCount             int    `json:"users_count"`
	ActiveUsersCount       int    `json:"active_users_count"`
	HeavyUsersCount        int    `json:"heavy_users_count"`
	ModelsCount            int    `json:"models_count"`
	ModelsToQueryCount     int    `json:"models_to_query_count"`
	OnlineModelsCount      int    `json:"online_models_count"`
	QueriesDurationSeconds int    `json:"queries_duration"`
	ErrorRate              [2]int `json:"error_rate"`
	Rss                    int64  `json:"rss"`
	MaxRss                 int64  `json:"max_rss"`
}
