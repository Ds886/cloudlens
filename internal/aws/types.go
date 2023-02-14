package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

type EC2Resp struct {
	Instance         ec2.Instance
	InstanceId       string
	InstanceType     string
	AvailabilityZone string
	InstanceState    string
	PublicDNS        string
	PublicIPv4       string
	MonitoringState  string
	LaunchTime       string
}

type S3Object struct {
	Name, ObjectType, LastModified, Size, StorageClass string
}

type BucketInfo struct {
	EncryptionConfiguration *s3.ServerSideEncryptionConfiguration
	LifeCycleRules          []*s3.LifecycleRule
}

type IAMUSerResp struct {
	UserId       string
	UserName     string
	ARN          string
	CreationTime string
}

type IAMUSerGroupResp struct {
	GroupId      string
	GroupName    string
	ARN          string
	CreationTime string
}

type IAMUSerPolicyResponse struct {
	PolicyArn  string
	PolicyName string
}

type EBSResp struct {
	VolumeId         string
	Size             string
	VolumeType       string
	State            string
	AvailabilityZone string
	Snapshot         string
	CreationTime     string
}

type IAMUSerGroupPolicyResponse struct {
	PolicyArn  string
	PolicyName string
}

type IamRoleResp struct {
	RoleId       string
	RoleName     string
	ARN          string
	CreationTime string
}

type IamRolePolicyResponse struct {
	PolicyArn  string
	PolicyName string
}

type SQSResp struct {
	Name              string
	URL               string
	Type              string
	Created           string
	MessagesAvailable string
	Encryption        string
	MaxMessageSize    string
}

type Snapshot struct {
	SnapshotId string
	OwnerId    string
	VolumeId   string
	VolumeSize string
	StartTime  string
	State      string
}

type ImageResp struct {
	ImageId       string
	OwnerId       string
	ImageLocation string
	Name          string
	ImageType     string
}

type VpcResp struct {
	VpcId           string
	OwnerId         string
	CidrBlock       string
	InstanceTenancy string
	State           string
}

type LambdaResp struct {
	FunctionName string
	Description  string
	Role         string
	FunctionArn  string
	CodeSize     string
	LastModified string
}
