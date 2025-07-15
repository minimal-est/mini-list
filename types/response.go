package types

type PostPreview struct {
	Sequence   int64  `json:"sequence"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	FolderName string `json:"folderName"`
	HitCount   int64  `json:"hitCount"`
	CreatedAt  string `json:"createdAt"`
}

type Data struct {
	Content []PostPreview `json:"content"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       Data   `json:"data"`
	Timestamp  string `json:"timestamp"`
}
