package database

type MongoImageRecord struct {
	Filepath string
	Tags     []string
}

func (r *MongoImageRecord) Url() string {
	return r.Filepath
}