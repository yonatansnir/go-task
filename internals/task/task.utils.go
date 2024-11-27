package task

func updateTask(id string, reqBody *TaskRequestBody) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			if reqBody.Description != "" {
				tasks[i].Description = reqBody.Description
			}
			if reqBody.Title != "" {
				tasks[i].Title = reqBody.Title
			}
			return
		}
	}
}