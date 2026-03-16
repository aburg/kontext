package util

func CreateKontextFilename(kind string) string {
	filename := "." + kind + ".kontext"
	return filename
}
