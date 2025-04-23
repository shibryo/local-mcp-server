package domain

// Weather はドメインの天気情報を表現する構造体です
type Weather struct {
	Location    Location
	Condition   string // 天気状態
	Temperature Temperature
	Description string // 天気概況文
}

type Location struct {
	Area       string // 地方名
	Prefecture string // 都道府県
	District   string // 地域
	City       string // 市区町村
}

type Temperature struct {
	Current string // 現在の気温（摂氏）
	Max     string // 最高気温（摂氏）
	Min     string // 最低気温（摂氏）
}
