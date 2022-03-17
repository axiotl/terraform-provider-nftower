package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"terraform-provider-nftower/internal/provider/models"
	"time"
)

type CreateComputeEnvPayload struct {
	ComputeEnv CreateComputeEnvComputeEnv `json:"computeEnv,omitempty"`
}
type CreateComputeEnvForge struct {
	FsxMode           string      `json:"fsxMode,omitempty"`
	EfsMode           string      `json:"efsMode,omitempty"`
	Type              string      `json:"type,omitempty"`
	MinCpus           int64       `json:"minCpus"`
	MaxCpus           int64       `json:"maxCpus,omitempty"`
	GpuEnabled        bool        `json:"gpuEnabled,omitempty"`
	EbsAutoScale      bool        `json:"ebsAutoScale,omitempty"`
	AllowBuckets      []string    `json:"allowBuckets,omitempty"`
	DisposeOnDeletion bool        `json:"disposeOnDeletion,omitempty"`
	InstanceTypes     []string    `json:"instanceTypes,omitempty"`
	AllocStrategy     string      `json:"allocStrategy,omitempty"`
	Ec2KeyPair        string      `json:"ec2KeyPair,omitempty"`
	VpcID             string      `json:"vpcId,omitempty"`
	ImageID           interface{} `json:"imageId,omitempty"`
	Subnets           []string    `json:"subnets,omitempty"`
	SecurityGroups    []string    `json:"securityGroups,omitempty"`
	EbsBlockSize      int64       `json:"ebsBlockSize,omitempty"`
	FusionEnabled     bool        `json:"fusionEnabled,omitempty"`
	EfsCreate         bool        `json:"efsCreate,omitempty"`
	ContainerRegIds   string      `json:"containerRegIds,omitempty"`
}
type CreateComputeEnvConfig struct {
	Credentials     string                `json:"credentials,omitempty"`
	Region          string                `json:"region,omitempty"`
	WorkDir         string                `json:"workDir,omitempty"`
	ComputeJobRole  string                `json:"computeJobRole,omitempty"`
	HeadJobRole     string                `json:"headJobRole,omitempty"`
	HeadJobCpus     string                `json:"headJobCpus,omitempty"`
	HeadJobMemoryMb int64                 `json:"headJobMemoryMb,omitempty"`
	ConfigMode      string                `json:"configMode,omitempty"`
	PreRunScript    string                `json:"preRunScript,omitempty"`
	PostRunScript   string                `json:"postRunScript,omitempty"`
	CliPath         string                `json:"cliPath,omitempty"`
	Environment     string                `json:"environment,omitempty"`
	Forge           CreateComputeEnvForge `json:"forge,omitempty"`
}
type CreateComputeEnvComputeEnv struct {
	Name          string                 `json:"name,omitempty"`
	Platform      string                 `json:"platform,omitempty"`
	Config        CreateComputeEnvConfig `json:"config,omitempty"`
	CredentialsID string                 `json:"credentialsId,omitempty"`
	DateCreated   time.Time              `json:"dateCreated,omitempty"`
}

type CreateComputeEnvResponse struct {
	ComputeEnvID string `json:"computeEnvId"`
}

/*

{
    "computeEnv": {
        "name": "default_copy_2",
        "platform": "aws-batch",
        "config": {
            "credentials": null,
            "region": "us-east-1",
            "workDir": "s3://convergence-default-dev",
            "computeJobRole": null,
            "headJobRole": null,
            "headJobCpus": null,
            "headJobMemoryMb": 8192,
            "configMode": "Batch Forge",
            "preRunScript": null,
            "postRunScript": null,
            "cliPath": "/home/ec2-user/miniconda/bin/aws",
            "environment": null,
            "forge": {
                "fsxMode": "None",
                "efsMode": "None",
                "type": "EC2",
                "minCpus": 0,
                "maxCpus": 512,
                "gpuEnabled": false,
                "ebsAutoScale": true,
                "allowBuckets": [
                    "s3://convergence-default-data",
                    "s3://convergence-default-run",
                    "s3://convergence-default-shared",
                    "s3://convergence-tf-tower",
                    "s3://convergence-default-nf-tower",
                    "s3://convergence-default-dev"
                ],
                "disposeOnDeletion": true,
                "instanceTypes": [],
                "allocStrategy": "BEST_FIT_PROGRESSIVE",
                "ec2KeyPair": "convergence_dev_box",
                "vpcId": "vpc-c3eed5b9",
                "imageId": null,
                "subnets": [
                    "subnet-0114aa4a598a2b8f2"
                ],
                "securityGroups": [],
                "ebsBlockSize": null,
                "fusionEnabled": true,
                "efsCreate": false,
                "containerRegIds": null
            }
        },
        "credentialsId": "7Y9dwU2JKHwuASqOdDLJ77",
        "dateCreated": "2021-12-30T01:30:48.086Z"
    }
}
*/

func (client *Client) CreateComputeEnv(env CreateComputeEnvPayload, workspaceId string) (*models.ComputeEnv, error) {
	body, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Create Compute Env Payload: %+v\n", string(body))
	endpoint := fmt.Sprintf("%s/compute-envs?workspaceId=%s", client.HostURL, workspaceId)
	fmt.Println("endpoint: ", endpoint)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	res, err := client.DoRequest(req)

	if err != nil {
		return nil, err
	}

	var responseData CreateComputeEnvResponse
	json.Unmarshal(res, &responseData)

	newComputeEnv, err := waitForAvailableState(client, responseData.ComputeEnvID, workspaceId)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(&newComputeEnv)
	fmt.Println("jsonData: ", string(jsonData))

	if err != nil {
		return nil, err
	}

	return newComputeEnv, nil
}

func (client *Client) ReadComputeEnv(computeEnvId string, workspaceId string) (*models.ComputeEnv, error) {
	endpoint := fmt.Sprintf("%s/compute-envs/%s?workspaceId=%s", client.HostURL, computeEnvId, workspaceId)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.DoRequest(req)

	if err != nil {
		return nil, err
	}

	var data models.ComputeEnvResponse
	json.Unmarshal(res, &data)
	fmt.Printf("%+v\n", data)
	return &data.ComputeEnv, nil
}

// UpdateComputeEnv
func (client *Client) UpdateComputeEnv(computeEnvId string, env CreateComputeEnvPayload, workspaceId string) (*models.ComputeEnv, error) {
	fmt.Printf("id: %s\n", computeEnvId)
	// in order to update compute env, we need to delete it first
	fmt.Println("deleting compute env")
	err := client.DeleteComputeEnv(computeEnvId, workspaceId)
	if err != nil {
		fmt.Printf("error deleting compute env: %s\n", err.Error())
		return nil, err
	}
	fmt.Println("creating compute env")
	newComputeEnvId, err := client.CreateComputeEnv(env, workspaceId)
	if err != nil {
		fmt.Printf("error creating compute env: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("reading compute env")
	newComputeEnv, err := client.ReadComputeEnv(newComputeEnvId.ID, workspaceId)
	if err != nil {
		fmt.Printf("error reading compute env: %s\n", err.Error())
		return nil, err
	}

	return newComputeEnv, nil
}

// delete compute env
func (client *Client) DeleteComputeEnv(computeEnvId string, workspaceId string) error {
	fmt.Printf("id: %s\n", computeEnvId)
	endpoint := fmt.Sprintf("%s/compute-envs/%s?workspaceId=%s", client.HostURL, computeEnvId, workspaceId)
	fmt.Println("endpoint: ", endpoint)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	_, err = client.DoRequest(req)

	if err != nil {
		return err
	}

	return nil
}

func waitForAvailableState(client *Client, computeEnvId string, workspaceId string) (*models.ComputeEnv, error) {
	numTries := 6
	for i := 0; i < numTries; i++ {
		fmt.Println("Waiting for compute env to be available...")
		e, err := client.ReadComputeEnv(computeEnvId, workspaceId)
		if err != nil {
			return nil, err
		}
		fmt.Println("status :", e.Status)
		if e.Status == "AVAILABLE" {
			fmt.Println("Compute env is available")
			return e, nil
		}

		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("compute env not available")
}
