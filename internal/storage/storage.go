package storage

func Save(c, url string) error {
	link := Link{Code: c, URL: url}
	result := DB.Create(&link)
	return result.Error
}

func Get(c string) (Link, error) {
	var link Link
	result := DB.Where("code = ?", c).First(&link)
	return link, result.Error
}
