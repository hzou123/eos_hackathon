package entity

type AssetEntity struct {
	Assets []struct {
		Token string  `json:"token"`
		Quantity string `json:"quantity"`
	}
}