package models

type IndexRequest struct {
    Items []*Item `json:"items"`
}

type ToastResponse struct {
    ClassName string
    Message string
}

type Item struct {
    Id int
    Name string
    Quantity int
}
