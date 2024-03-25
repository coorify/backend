package option

type SignatureOption struct {
	Enable bool   `default:"false"`
	Pri    string `default:""`
	Pub    string `default:""`
}
