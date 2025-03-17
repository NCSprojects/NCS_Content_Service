package dto

type UpdateRnkDTO struct {
	ID  int `json:"id" binding:"required"`  
	Rnk int `json:"rnk" binding:"required"`
}