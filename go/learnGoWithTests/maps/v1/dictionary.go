package v1

type dictionary map[string]string

func (d dictionary) Search(key string) string {
	return d[key]
}
