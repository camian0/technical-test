package model

type ShardInfo struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []struct {
		Index     string  `json:"_index"`
		Type      string  `json:"_type"`
		ID        string  `json:"_id"`
		Score     float64 `json:"_score"`
		Timestamp string  `json:"@timestamp"`
		Source    struct {
			Timestamp               string   `json:"@timestamp"`
			ContentTransferEncoding string   `json:"Content-Transfer-Encoding"`
			ContentType             string   `json:"Content-Type"`
			Date                    string   `json:"Date"`
			From                    string   `json:"From"`
			MessageID               string   `json:"Message-ID"`
			MimeVersion             string   `json:"Mime-Version"`
			Subject                 string   `json:"Subject"`
			To                      []string `json:"To"`
			XFileName               string   `json:"X-FileName"`
			XFolder                 string   `json:"X-Folder"`
			XFrom                   string   `json:"X-From"`
			XOrigin                 string   `json:"X-Origin"`
			XTo                     string   `json:"X-To"`
			XBcc                    string   `json:"X-bcc"`
			XCc                     string   `json:"X-cc"`
			Content                 string   `json:"content"`
		} `json:"_source"`
	} `json:"hits"`
}

type Response struct {
	Took     int       `json:"took"`
	TimedOut bool      `json:"timed_out"`
	Shards   ShardInfo `json:"_shards"`
	Hits     Hits      `json:"hits"`
}
