package models

// model = ใน database
type Book struct {
	ID     uint64 `json:"id" gorm:"primary_key"`
	Title  string `json:"title" `
	Author string `json:"author" `
}

// schema มีไว้เพื่อให้ผู้ใช้กรอกข้อมูลได้ถูกต้อง
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

//updateBook
type UpdateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
