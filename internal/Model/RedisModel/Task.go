package RedisModel

type Task struct {
	TaskType    string      `json:"task_type"`
	TaskMessage interface{} `json:"task_message"`
}

