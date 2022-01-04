package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type computeEnvDataSourceType struct{}

func (t computeEnvDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Required: true,
			},
			"name": {
				Type:     types.StringType,
				Computed: true,
				Optional: true,
			},
			"config": {
				Computed: true,
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"region": {
						Type:     types.StringType,
						Computed: true,
					},
					"compute_queue": {
						Type:     types.StringType,
						Computed: true,
					},
					"ComputeJobRole": {
						Type:     types.StringType,
						Computed: true,
					},
					"HeadQueue": {
						Type:     types.StringType,
						Computed: true,
					},
					"HeadJobRole": {
						Type:     types.StringType,
						Computed: true,
					},
					"CliPath": {
						Type:     types.StringType,
						Computed: true,
					},
					"WorkDir": {
						Type:     types.StringType,
						Computed: true,
					},
					"PreRunScript": {
						Type:     types.StringType,
						Computed: true,
					},
					"PostRunScript": {
						Type:     types.StringType,
						Computed: true,
					},
					"HeadJobCpus": {
						Type:     types.StringType,
						Computed: true,
					},
					"HeadJobMemoryMb": {
						Type:     types.StringType,
						Computed: true,
					},
					"Forge": {
						Type:     types.StringType,
						Computed: true,
					},
					"Discriminator": {
						Type:     types.StringType,
						Computed: true,
					},
				}),
			},
		},
	}, nil
}

func (t computeEnvDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return computeEnvDataSource{
		provider: provider,
	}, diags
}

type computeEnvDataSourceData struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	// Config computeEnvConfigDataSourceData `tfsdk:"config"`
}

type computeEnvConfigDataSourceData struct {
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
	HeadJobMemoryMb types.String `json:"headJobMemoryMb" tfsdk:"head_job_memory_mb"`
	// Environment     []Environment     `json:"environment" tfsdk:"environment"`
	// Forge           Forge             `json:"forge" tfsdk:"forge"`
	// ForgedResources map[types.String]types.String `json:"forgedResources" tfsdk:"forged_resources"`
	Discriminator types.String `json:"discriminator" tfsdk:"discriminator"`
}

type computeEnvDataSource struct {
	provider provider
}

func (d computeEnvDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data computeEnvDataSourceData
	var data2 interface{}
	diags := req.Config.Get(ctx, &data)
	req.Config.Get(ctx, &data2)
	resp.Diagnostics.Append(diags...)

	log.Printf("got here")
	fmt.Printf("data: %+v", data)
	fmt.Printf("data2: %+v", data2)

	computeEnv, err := d.provider.client.ReadComputeEnv(data.Id.Value, "197562422694202")
	if err != nil {
		resp.Diagnostics.AddError("Failed to read compute env", err.Error())
		return
	}
	fmt.Printf("Compute asdfenv: %+v", computeEnv)

	log.Printf("got heasdfre")

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.ReadExample(...)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.String{Value: computeEnv.ID}
	data.Name = types.String{Value: computeEnv.Name}
	// data.Config = computeEnvConfigDataSourceData{
	// 	Region:         types.String{Value: computeEnv.Config.Region},
	// 	ComputeQueue:   types.String{Value: computeEnv.Config.ComputeQueue},
	// 	ComputeJobRole: types.String{Value: computeEnv.Config.ComputeJobRole},
	// 	HeadQueue:      types.String{Value: computeEnv.Config.HeadQueue},
	// 	HeadJobRole:    types.String{Value: computeEnv.Config.HeadJobRole},
	// 	CliPath:        types.String{Value: computeEnv.Config.CliPath},
	// 	// Volumes:         []types.String{types.String{Value: computeEnv.Config.Volumes}},
	// 	WorkDir:         types.String{Value: computeEnv.Config.WorkDir},
	// 	PreRunScript:    types.String{Value: computeEnv.Config.PreRunScript},
	// 	PostRunScript:   types.String{Value: computeEnv.Config.PostRunScript},
	// 	HeadJobCpus:     types.string{Value: computeEnv.Config.HeadJobCpus},
	// 	HeadJobMemoryMb: types.string{Value: computeEnv.Config.HeadJobMemoryMb},
	// 	// Environment:     computeEnv.Environment,
	// 	// Forge:           computeEnv.Forge,
	// 	// ForgedResources: computeEnv.ForgedResources,
	// 	Discriminator: types.String{Value: computeEnv.Config.Discriminator},
	// }

	fmt.Printf("compute env data: %+v", data)
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
