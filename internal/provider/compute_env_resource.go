package provider

import (
	"context"
	"fmt"
	"terraform-provider-nftower/internal/provider/client"
	"terraform-provider-nftower/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type computeEnvResourceType struct{}

func (t computeEnvResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "nftower compute env",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"workspace_id": {
				Type:     types.StringType,
				Optional: true,
			},
			"credentials_id": {
				Type:     types.StringType,
				Optional: true,
			},
			"config": {
				Required: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"region": {
						Type:     types.StringType,
						Required: true,
					},
					"compute_queue": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"compute_job_role": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"head_queue": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"head_job_role": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"cli_path": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"work_dir": {
						Type:        types.StringType,
						Required:    true,
						Description: "Working directory for the compute environment. This must be an absolute path.",
					},
					"pre_run_script": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"post_run_script": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"head_job_cpus": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
					"head_job_memory_mb": {
						Type:     types.Int64Type,
						Computed: true,
						Optional: true,
					},
					"forge": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"type": {
								Type:     types.StringType,
								Required: true,
								Validators: []tfsdk.AttributeValidator{
									enumValidator{
										allowedValues: []string{"SPOT", "EC2"},
									},
								},
							},
							"min_cpus": {
								Type:     types.Int64Type,
								Computed: true,
								Optional: true,
							},
							"max_cpus": {
								Type:     types.Int64Type,
								Computed: true,
								Optional: true,
							},
							"gpu_enabled": {
								Type:     types.BoolType,
								Computed: true,
								Optional: true,
							},
							"ebs_autoscale": {
								Type:     types.BoolType,
								Computed: true,
								Optional: true,
							},
							"instance_types": {
								Type: types.ListType{
									ElemType: types.StringType,
								},
								Computed: true,
								Optional: true,
							},
							"alloc_strategy": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
								Validators: []tfsdk.AttributeValidator{
									enumValidator{
										allowedValues: []string{"BEST_FIT", "BEST_FIT_PROGRESSIVE", "SPOT_CAPACITY_OPTIMIZED"},
									},
								},
							},
							"image_id": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"vpc_id": {
								Type:     types.StringType,
								Required: true,
							},
							"subnets": {
								Type: types.ListType{
									ElemType: types.StringType,
								},
								Required: true,
							},
							"security_groups": {
								Type: types.ListType{
									ElemType: types.StringType,
								},
								Computed: true,
								Optional: true,
							},
							"fsx_mount": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"fsx_name": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"fsx_size": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"dispose_on_deletion": {
								Type:     types.BoolType,
								Computed: true,
								Optional: true,
							},
							"ec2_key_pair": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"allow_buckets": {
								Type: types.ListType{
									ElemType: types.StringType,
								},
								Computed: true,
								Optional: true,
							},
							"ebs_block_size": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"fusion_enabled": {
								Type:     types.BoolType,
								Computed: true,
								Optional: true,
							},
							"bid_percentage": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"efs_create": {
								Type:     types.BoolType,
								Computed: true,
								Optional: true,
							},
							"efs_id": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
							"efs_mount": {
								Type:     types.StringType,
								Computed: true,
								Optional: true,
							},
						})},
					"discriminator": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
				}),
			},
		},
	}, nil
}

func (t computeEnvResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return computeEnvResource{
		provider: provider,
	}, diags
}

type computeEnvResourceData struct {
	Id            types.String                 `tfsdk:"id"`
	Name          types.String                 `tfsdk:"name"`
	WorkspaceID   types.String                 `tfsdk:"workspace_id"`
	CredentialsID types.String                 `tfsdk:"credentials_id"`
	Config        computeEnvResourceDataConfig `tfsdk:"config"`
}

type computeEnvResourceDataConfig struct {
	Region         types.String `json:"region" tfsdk:"region"`
	ComputeQueue   types.String `json:"computeQueue" tfsdk:"compute_queue" `
	ComputeJobRole types.String `json:"computeJobRole" tfsdk:"compute_job_role"`
	HeadQueue      types.String `json:"headQueue" tfsdk:"head_queue"`
	HeadJobRole    types.String `json:"headJobRole" tfsdk:"head_job_role"`
	CliPath        types.String `json:"cliPath" tfsdk:"cli_path"`
	// Volumes         []types.String `json:"volumes" tfsdk:"volumes"`
	WorkDir         types.String `json:"workDir" tfsdk:"work_dir"`
	PreRunScript    types.String `json:"preRunScript" tfsdk:"pre_run_script"`
	PostRunScript   types.String `json:"postRunScript" tfsdk:"post_run_script"`
	HeadJobCpus     types.String `json:"headJobCpus" tfsdk:"head_job_cpus"`
	HeadJobMemoryMb types.Int64  `json:"headJobMemoryMb" tfsdk:"head_job_memory_mb"`
	// Environment     []Environment     `json:"environment" tfsdk:"environment"`
	Forge computeEnvResourceDataForge `json:"forge" tfsdk:"forge"`
	// ForgedResources map[types.String]types.String `json:"forgedResources" tfsdk:"forged_resources"`
	Discriminator types.String `json:"discriminator" tfsdk:"discriminator"`
}

type computeEnvResourceDataForge struct {
	Type              types.String   `json:"type" tfsdk:"type"`
	MinCpus           types.Int64    `json:"minCpus" tfsdk:"min_cpus"`
	MaxCpus           types.Int64    `json:"maxCpus" tfsdk:"max_cpus"`
	GpuEnabled        types.Bool     `json:"gpuEnabled" tfsdk:"gpu_enabled"`
	EbsAutoScale      types.Bool     `json:"ebsAutoScale" tfsdk:"ebs_autoscale"`
	InstanceTypes     []types.String `json:"instanceTypes" tfsdk:"instance_types"`
	AllocStrategy     types.String   `json:"allocStrategy" tfsdk:"alloc_strategy"`
	ImageID           types.String   `json:"imageId" tfsdk:"image_id"`
	VpcID             types.String   `json:"vpcId" tfsdk:"vpc_id"`
	Subnets           []types.String `json:"subnets" tfsdk:"subnets"`
	SecurityGroups    []types.String `json:"securityGroups" tfsdk:"security_groups"`
	FsxMount          types.String   `json:"fsxMount" tfsdk:"fsx_mount"`
	FsxName           types.String   `json:"fsxName" tfsdk:"fsx_name"`
	FsxSize           types.String   `json:"fsxSize" tfsdk:"fsx_size"`
	DisposeOnDeletion types.Bool     `json:"disposeOnDeletion" tfsdk:"dispose_on_deletion"`
	Ec2KeyPair        types.String   `json:"ec2KeyPair" tfsdk:"ec2_key_pair"`
	AllowBuckets      []types.String `json:"allowBuckets" tfsdk:"allow_buckets"`
	EbsBlockSize      types.String   `json:"ebsBlockSize"  tfsdk:"ebs_block_size"`
	FusionEnabled     types.Bool     `json:"fusionEnabled" tfsdk:"fusion_enabled"`
	BidPercentage     types.String   `json:"bidPercentage" tfsdk:"bid_percentage"`
	EfsCreate         types.Bool     `json:"efsCreate" tfsdk:"efs_create"`
	EfsID             types.String   `json:"efsId" tfsdk:"efs_id"`
	EfsMount          types.String   `json:"efsMount" tfsdk:"efs_mount"`
}

type computeEnvResource struct {
	provider provider
}

func (r computeEnvResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data computeEnvResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	fmt.Printf("config data: %+v", data)

	payload := client.CreateComputeEnvPayload{
		ComputeEnv: client.CreateComputeEnvComputeEnv{
			Name:          data.Name.Value,
			Platform:      "aws-batch",
			CredentialsID: data.CredentialsID.Value,

			Config: client.CreateComputeEnvConfig{
				Region:          data.Config.Region.Value,
				WorkDir:         data.Config.WorkDir.Value,
				HeadJobMemoryMb: data.Config.HeadJobMemoryMb.Value,
				HeadJobRole:     data.Config.HeadJobRole.Value,
				ComputeJobRole:  data.Config.ComputeJobRole.Value,
				HeadJobCpus:     data.Config.HeadJobCpus.Value,
				CliPath:         "/home/ec2-user/miniconda/bin/aws",

				Forge: client.CreateComputeEnvForge{
					Type:              data.Config.Forge.Type.Value,
					MinCpus:           data.Config.Forge.MinCpus.Value,
					MaxCpus:           data.Config.Forge.MaxCpus.Value,
					GpuEnabled:        data.Config.Forge.GpuEnabled.Value,
					EbsAutoScale:      data.Config.Forge.EbsAutoScale.Value,
					AllowBuckets:      stringTypeArrayToStringArray(data.Config.Forge.AllowBuckets),
					DisposeOnDeletion: data.Config.Forge.DisposeOnDeletion.Value,
					InstanceTypes:     stringTypeArrayToStringArray(data.Config.Forge.InstanceTypes),
					AllocStrategy:     data.Config.Forge.AllocStrategy.Value,
					VpcID:             data.Config.Forge.VpcID.Value,
					Subnets:           stringTypeArrayToStringArray(data.Config.Forge.Subnets),
					FusionEnabled:     data.Config.Forge.FusionEnabled.Value,
					EfsCreate:         data.Config.Forge.EfsCreate.Value,
					FsxMode:           "None",
					EfsMode:           "None",
				},
				ConfigMode: "Batch Forge",
			},
		},
	}
	newComputeEnv, err := r.provider.client.CreateComputeEnv(payload, data.WorkspaceID.Value)

	if err != nil {
		resp.Diagnostics.AddError("Error creating new compute env", err.Error())
		return
	}

	createResource(&data, newComputeEnv)

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
}

func (r computeEnvResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data computeEnvResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	computeEnv, err := r.provider.client.ReadComputeEnv(data.Id.Value, data.WorkspaceID.Value)
	if err != nil {
		// resp.Diagnostics.AddError("Error reading compute env", err.Error())
		resp.State.Set(ctx, nil)
		return
	}
	fmt.Printf("computeEnv: %+v", computeEnv)

	err = createResource(&data, computeEnv)
	fmt.Printf("data: %+v", data)
	if err != nil {
		var path = tftypes.NewAttributePath().WithAttributeName("id")
		resp.State.SetAttribute(ctx, path, "")
		resp.Diagnostics.AddError("error converting api response to resource data", err.Error())
		return
	}

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
}

func (r computeEnvResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data computeEnvResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var currentState computeEnvResourceData

	diags = req.State.Get(ctx, &currentState)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	payload := client.CreateComputeEnvPayload{
		ComputeEnv: client.CreateComputeEnvComputeEnv{
			Name:          data.Name.Value,
			Platform:      "aws-batch",
			CredentialsID: data.CredentialsID.Value,

			Config: client.CreateComputeEnvConfig{
				HeadJobCpus:     data.Config.HeadJobCpus.Value,
				Region:          data.Config.Region.Value,
				WorkDir:         data.Config.WorkDir.Value,
				HeadJobMemoryMb: data.Config.HeadJobMemoryMb.Value,
				ComputeJobRole:  data.Config.ComputeJobRole.Value,
				HeadJobRole:     data.Config.HeadJobRole.Value,
				CliPath:         "/home/ec2-user/miniconda/bin/aws",
				Forge: client.CreateComputeEnvForge{
					Type:              data.Config.Forge.Type.Value,
					MinCpus:           data.Config.Forge.MinCpus.Value,
					MaxCpus:           data.Config.Forge.MaxCpus.Value,
					GpuEnabled:        data.Config.Forge.GpuEnabled.Value,
					EbsAutoScale:      data.Config.Forge.EbsAutoScale.Value,
					AllowBuckets:      stringTypeArrayToStringArray(data.Config.Forge.AllowBuckets),
					DisposeOnDeletion: data.Config.Forge.DisposeOnDeletion.Value,
					InstanceTypes:     stringTypeArrayToStringArray(data.Config.Forge.InstanceTypes),
					AllocStrategy:     data.Config.Forge.AllocStrategy.Value,
					VpcID:             data.Config.Forge.VpcID.Value,
					Subnets:           stringTypeArrayToStringArray(data.Config.Forge.Subnets),
					FusionEnabled:     data.Config.Forge.FusionEnabled.Value,
					EfsCreate:         data.Config.Forge.EfsCreate.Value,
					FsxMode:           "None",
					EfsMode:           "None",
				},
				ConfigMode: "Batch Forge",
			},
		},
	}
	fmt.Printf("payload: %+v", payload)

	newComputeEnv, err := r.provider.client.UpdateComputeEnv(currentState.Id.Value, payload, data.WorkspaceID.Value)

	if err != nil {
		resp.Diagnostics.AddError("Error creating new compute env", err.Error())
		return
	}

	fmt.Printf("newComputeEnv: %+v", *newComputeEnv)

	err = createResource(&data, newComputeEnv)
	fmt.Printf("data: %+v", data)
	if err != nil {
		resp.Diagnostics.AddError("error converting api response to resource data", err.Error())
		return
	}

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
}

func (r computeEnvResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data computeEnvResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.provider.client.DeleteComputeEnv(data.Id.Value, data.WorkspaceID.Value)
	if err != nil {
		resp.Diagnostics.AddError("Error reading compute env", err.Error())
		return
	}

	if err != nil {
		resp.Diagnostics.AddError("error converting api response to resource data", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r computeEnvResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}

func stringTypeArrayToStringArray(in []types.String) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.Value
	}
	fmt.Println(out)
	return out
}

func stringArrayToStringTypeArray(in []string) []types.String {
	out := make([]types.String, len(in))
	for i, v := range in {
		out[i] = types.String{Value: v}
	}
	return out
}

func createResource(data *computeEnvResourceData, computeEnv *models.ComputeEnv) error {
	data.Id = types.String{Value: computeEnv.ID}
	data.Name = types.String{Value: computeEnv.Name}

	data.Config = computeEnvResourceDataConfig{
		Region:          types.String{Value: computeEnv.Config.Region},
		ComputeQueue:    types.String{Value: computeEnv.Config.ComputeQueue},
		ComputeJobRole:  types.String{Value: computeEnv.Config.ComputeJobRole},
		HeadQueue:       types.String{Value: computeEnv.Config.HeadQueue},
		HeadJobRole:     types.String{Value: computeEnv.Config.HeadJobRole},
		CliPath:         types.String{Value: computeEnv.Config.CliPath},
		WorkDir:         types.String{Value: computeEnv.Config.WorkDir},
		PreRunScript:    types.String{Value: computeEnv.Config.PreRunScript},
		PostRunScript:   types.String{Value: computeEnv.Config.PostRunScript},
		HeadJobCpus:     types.String{Value: computeEnv.Config.HeadJobCpus},
		HeadJobMemoryMb: types.Int64{Value: computeEnv.Config.HeadJobMemoryMb},
		// Environment:     newComputeEnv.Config.Environment,
		Forge: computeEnvResourceDataForge{
			Type:              types.String{Value: computeEnv.Config.Forge.Type},
			MinCpus:           types.Int64{Value: computeEnv.Config.Forge.MinCpus},
			MaxCpus:           types.Int64{Value: computeEnv.Config.Forge.MaxCpus},
			GpuEnabled:        types.Bool{Value: computeEnv.Config.Forge.GpuEnabled},
			EbsAutoScale:      types.Bool{Value: computeEnv.Config.Forge.EbsAutoScale},
			InstanceTypes:     stringArrayToStringTypeArray(computeEnv.Config.Forge.InstanceTypes),
			AllocStrategy:     types.String{Value: computeEnv.Config.Forge.AllocStrategy},
			VpcID:             types.String{Value: computeEnv.Config.Forge.VpcID},
			Subnets:           stringArrayToStringTypeArray(computeEnv.Config.Forge.Subnets),
			FusionEnabled:     types.Bool{Value: computeEnv.Config.Forge.FusionEnabled},
			SecurityGroups:    stringArrayToStringTypeArray(computeEnv.Config.Forge.SecurityGroups),
			FsxMount:          types.String{Value: computeEnv.Config.Forge.FsxMount},
			FsxName:           types.String{Value: computeEnv.Config.Forge.FsxName},
			FsxSize:           types.String{Value: computeEnv.Config.Forge.FsxSize},
			DisposeOnDeletion: types.Bool{Value: computeEnv.Config.Forge.DisposeOnDeletion},
			EfsCreate:         types.Bool{Value: computeEnv.Config.Forge.EfsCreate},
			Ec2KeyPair:        types.String{Value: computeEnv.Config.Forge.Ec2KeyPair},
			AllowBuckets:      stringArrayToStringTypeArray(computeEnv.Config.Forge.AllowBuckets),
			EbsBlockSize:      types.String{Value: computeEnv.Config.Forge.EbsBlockSize},
			BidPercentage:     types.String{Value: computeEnv.Config.Forge.BidPercentage},
			EfsID:             types.String{Value: computeEnv.Config.Forge.EfsID},
			EfsMount:          types.String{Value: computeEnv.Config.Forge.EfsMount},
		},
		// ForgedResources: newComputeEnv.Config.ForgedResources,
		Discriminator: types.String{Value: computeEnv.Config.Discriminator},
	}

	return nil
}
