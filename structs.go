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
	Categories []UpworkCategoriesCategory `json:"categories"`
}

type UpworkCategoriesCategory struct {
	Title  string                          `json:"title"`
	Topics []UpworkCategoriesCategoryTopic `json:"topics"`
}

type UpworkCategoriesCategoryTopic struct {
	Title string `json:"title"`
}

type UpworkJobs struct {
	Jobs   []Job            `json:"jobs"`
	Paging UpworkJobsPaging `json:"paging"`
}

type UpworkJobsPaging struct {
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
	Total  int64 `json:"total"`
}

type UpworkSkills struct {
	Skills []string `json:"skills"`
}

type Category struct {
	Id       int64  `db:"id"`
	Title    string `db:"title"`
	SubTitle string `db:"sub_title"`
	Status   string `db:"status"`
}

type Job struct {
	Id                 int64    `db:"id" json:"omit"`
	Budget             int64    `db:"budget" json:"budget"`
	Category           string   `db:"category" json:"category2"`
	ClientCountry      string   `db:"client_country" json:"client>country"`
	ClientFeedback     float64  `db:"client_feedback" json:"client>feedback"`
	ClientJobsPosted   int64    `db:"client_jobs_posted" json:"client>jobs_posted"`
	ClientPastHires    int64    `db:"client_past_hires" json:"client>past_hires"`
	ClientReviewsCount int64    `db:"client_reviews_count" json:"client>reviews_count"`
	DateCreated        string   `db:"date_created" json:"date_created"`
	Duration           string   `db:"duration" json:"duration"`
	JobStatus          string   `db:"job_status" json:"job_status"`
	JobType            string   `db:"job_type" json:"job_type"`
	Skills             []string `db:"skills" json:"skills"`
	Snippet            string   `db:"snippet" json:"snippet"`
	SubCategory        string   `db:"sub_category" json:"subcategory2"`
	Title              string   `db:"title" json:"title"`
	Url                string   `db:"url" json:"url"`
	Workload           string   `db:"workload" json:"workload"`
}

type Skill struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	Status string `db:"status"`
}
