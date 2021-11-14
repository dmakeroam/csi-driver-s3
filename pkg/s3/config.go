package s3

// Config holds values to configure the driver
type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	BucketPrefix    string
	Region          string
	Endpoint        string
	Mounter         string
}
