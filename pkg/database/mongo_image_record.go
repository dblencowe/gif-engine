package database

type MongoImageRecord struct {
	Filepath string `bson:"url"`
	Tags     []string
}

func (r *MongoImageRecord) Url() string {
	return r.Filepath
}