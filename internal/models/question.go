package models

type QuestionData struct {
	Title        string   `json:"title"`
	Header       string   `json:"header"`
	Content      string   `json:"content"`
	Questions    []string `json:"questions"`
	Answer       string   `json:"answer"`
	Timestamp    string   `json:"timestamp"`
	QuestionLink string   `json:"question_link"`
	Comments     string   `json:"comments"`
}

type FileInfo struct {
	URL    string
	Name   string
	Number int
}

type JSONResponse struct {
	PageProps struct {
		Questions []struct {
			Choices           map[string]string `json:"choices"`
			ID                string            `json:"id"`
			ExamID            int               `json:"exam_id"`
			QuestionText      string            `json:"question_text"`
			Answer            string            `json:"answer"`
			AnswerET          string            `json:"answer_ET"`
			Topic             string            `json:"topic"`
			IsMC              bool              `json:"isMC"`
			AnswerDescription string            `json:"answer_description"`
			Discussion        []struct {
				Content     string `json:"content"`
				UpvoteCount string `json:"upvote_count"`
				Poster      string `json:"poster"`
				Timestamp   string `json:"timestamp"`
			} `json:"discussion"`
			AnswerImages   []string `json:"answer_images"`
			QuestionImages []string `json:"question_images"`
			URL            string   `json:"url"`
			Timestamp      string   `json:"timestamp"`
		} `json:"questions"`
	} `json:"pageProps"`
}
