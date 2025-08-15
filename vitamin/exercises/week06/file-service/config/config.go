package config

const MaxFileSize = 5 * 1024 * 1024

var AllowedExtensions = map[string]struct{}{
	".jpg":  {},
	".png":  {},
	".js":   {},
	".css":  {},
	".html": {},
}
