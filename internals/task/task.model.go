package task

type TaskItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}