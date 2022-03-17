package models

type ComputeEnvResponse struct {
	ComputeEnv ComputeEnv `json:"computeEnv" tfsdk:"computeEnv"`
}

type Environment struct {
	Name    string `json:"name" tfsdk:"name"`
	Value   string `json:"value" tfsdk:"value"`
	Head    bool   `json:"head" tfsdk:"head"`
	Compute bool   `json:"compute" tfsdk:"compute"`
}
type Forge struct {
	Type              string   `json:"type" tfsdk:"type"`
	MinCpus           int64    `json:"minCpus" tfsdk:"min_cpus"`
	MaxCpus           int64    `json:"maxCpus" tfsdk:"max_cpus"`
	GpuEnabled        bool     `json:"gpuEnabled" tfsdk:"gpu_enabled"`
	EbsAutoScale      bool     `json:"ebsAutoScale" tfsdk:"ebs_autoscale"`
	InstanceTypes     []string `json:"instanceTypes" tfsdk:"instance_types"`
	AllocStrategy     string   `json:"allocStrategy" tfsdk:"alloc_strategy"`
	ImageID           string   `json:"imageId" tfsdk:"image_id"`
	VpcID             string   `json:"vpcId" tfsdk:"vpc_id"`
	Subnets           []string `json:"subnets" tfsdk:"subnets"`
	SecurityGroups    []string `json:"securityGroups" tfsdk:"security_groups"`
	FsxMount          string   `json:"fsxMount" tfsdk:"fsx_mount"`
	FsxName           string   `json:"fsxName" tfsdk:"fsx_name"`
	FsxSize           string   `json:"fsxSize" tfsdk:"fsx_size"`
	DisposeOnDeletion bool     `json:"disposeOnDeletion" tfsdk:"dispose_on_deletion"`
	Ec2KeyPair        string   `json:"ec2KeyPair" tfsdk:"ec2_key_pair"`
	AllowBuckets      []string `json:"allowBuckets" tfsdk:"allow_buckets"`
	EbsBlockSize      string   `json:"ebsBlockSize"  tfsdk:"ebs_block_size"`
	FusionEnabled     bool     `json:"fusionEnabled" tfsdk:"fusion_enabled"`
	BidPercentage     string   `json:"bidPercentage" tfsdk:"bid_percentage"`
	EfsCreate         bool     `json:"efsCreate" tfsdk:"efs_create"`
	EfsID             string   `json:"efsId" tfsdk:"efs_id"`
	EfsMount          string   `json:"efsMount" tfsdk:"efs_mount"`
}

type Config struct {
	Region          string            `json:"region" tfsdk:"region"`
	ComputeQueue    string            `json:"computeQueue" tfsdk:"compute_queue" `
	ComputeJobRole  string            `json:"computeJobRole" tfsdk:"compute_job_role"`
	HeadQueue       string            `json:"headQueue" tfsdk:"head_queue"`
	HeadJobRole     string            `json:"headJobRole" tfsdk:"head_job_role"`
	CliPath         string            `json:"cliPath" tfsdk:"cli_path"`
	Volumes         []string          `json:"volumes" tfsdk:"volumes"`
	WorkDir         string            `json:"workDir" tfsdk:"work_dir"`
	PreRunScript    string            `json:"preRunScript" tfsdk:"pre_run_script"`
	PostRunScript   string            `json:"postRunScript" tfsdk:"post_run_script"`
	HeadJobCpus     string            `json:"headJobCpus" tfsdk:"head_job_cpus"`
	HeadJobMemoryMb int64             `json:"headJobMemoryMb" tfsdk:"head_job_memory_mb"`
	Environment     []Environment     `json:"environment" tfsdk:"environment"`
	Forge           Forge             `json:"forge" tfsdk:"forge"`
	ForgedResources map[string]string `json:"forgedResources" tfsdk:"forged_resources"`
	Discriminator   string            `json:"discriminator" tfsdk:"discriminator"`
}

type ComputeEnv struct {
	ID            string `json:"id" tfsdk:"id"`
	Name          string `json:"name" tfsdk:"name"`
	Description   string `json:"description" tfsdk:"description"`
	Platform      string `json:"platform" tfsdk:"platform"`
	Config        Config `json:"config" tfsdk:"config"`
	DateCreated   string `json:"dateCreated" tfsdk:"date_created"`
	LastUpdated   string `json:"lastUpdated" tfsdk:"last_updated"`
	LastUsed      string `json:"lastUsed" tfsdk:"last_used"`
	Deleted       bool   `json:"deleted" tfsdk:"deleted"`
	Status        string `json:"status" tfsdk:"status"`
	Message       string `json:"message" tfsdk:"message"`
	Primary       bool   `json:"primary" tfsdk:"primary"`
	CredentialsID string `json:"credentialsId" tfsdk:"credentials_id"`
	OrgID         string `json:"orgId" tfsdk:"org_id"`
	WorkspaceID   string `json:"workspaceId" tfsdk:"workspace_id"`
}
