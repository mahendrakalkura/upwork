package main

type Settings struct {
	Raven  SettingsRaven  `toml:"raven"`
	SQLX   SettingsSQLX   `toml:"sqlx"`
	Upwork SettingsUpwork `toml:"upwork"`
}

type SettingsRaven struct {
	Dsn string `toml:"dsn"`
}

type SettingsSQLX struct {
	Database string `toml:"database"`
	Hostname string `toml:"hostname"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
}

type SettingsUpwork struct {
	AccessSecret   string `toml:"access_secret"`
	AccessToken    string `toml:"access_token"`
	ConsumerKey    string `toml:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret"`
	Debug          string `toml:"debug"`
}

type UpworkCategories struct {
	Items []UpworkCategoriesCategory `json:"categories"`
}

type UpworkCategoriesCategory struct {
	Title  string                          `json:"title"`
	Topics []UpworkCategoriesCategoryTopic `json:"topics"`
}

type UpworkCategoriesCategoryTopic struct {
	Title string `json:"title"`
}

type UpworkSkills struct {
	Items []string `json:"skills"`
}

type Category struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	SubTitle string `db:"sub_title"`
	Status   string `db:"status"`
}

type Skill struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Status string `db:"status"`
}
