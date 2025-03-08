package dto

type ListReqData struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	// Offset int64  `json:"offset"`
	// Name   string `json:"name"`
	Order int32 `json:"order"`
}
