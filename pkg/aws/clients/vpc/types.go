package vpc

// VPC defines the desired state of an AWS VPC
type VPC struct {
	ID                          uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	AmazonProvidedIpv6CidrBlock bool   `gorm:""`
	CidrBlock                   string `gorm:"not null"`
	EnableDnsHostNames          bool   `gorm:""`
	EnableDnsSupport            bool   `gorm:""`
	InstanceTenancy             string `gorm:""`
	Ipv6CidrBlock               string `gorm:""`
	Ipv6Pool                    string `gorm:""`
	Region                      string `gorm:""`
}
