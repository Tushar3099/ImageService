package types

type Error struct {
	StoreId string `json:"store_id"`
	Message string `json:"error"`
}
