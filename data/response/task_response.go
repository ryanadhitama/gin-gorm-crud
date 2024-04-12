package response

type TaskResponse struct {
	Id   		int    `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
}