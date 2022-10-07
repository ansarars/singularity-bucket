// (c) Copyright 2022 Hewlett Packard Enterprise Development LP

package main

import (
	"fmt"
	"github.com/ansarars/singularity-s3-bucket/client"
	"github.com/ansarars/singularity-s3-bucket/utils"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/hpe-storage/common-host-libs/logger"
	"github.com/spf13/cobra"
	"github.com/sylabs/singularity/pkg/cmdline"
	pluginapi "github.com/sylabs/singularity/pkg/plugin"
	clicallback "github.com/sylabs/singularity/pkg/plugin/callback/cli"
)

const (
	CREATE          string = "create"
	LIST            string = "list"
	RESOURCE_BUCKET string = "bucket"
	RESOURCE_OBJECT string = "object"
	BUCKET_NAME     string = "bucket-name"
)

const BucketUsage = `bucket <create|list> <name>


Options
- create
    Specifies bucket creation operation.
- list
    Specifies bucket list operation.
- name
    Specifies bucket name. Required for create operation only.
`

const ObjectUsage = `object <list> <bucket-name>


Options
- list
    Specifies bucket list operation.
- bucket-name
    Specifies bucket name. Required for object list operation.
`

var Plugin = pluginapi.Plugin{
	Manifest: pluginapi.Manifest{
		Name:        "singularity-s3-bucket-plugin",
		Author:      "Arshad Alam Ansari",
		Version:     "v1.0",
		Description: "singularity plugin for s3 bucket operations",
	},
	Callbacks: []pluginapi.Callback{
		(clicallback.Command)(callbackBucketCmd),
		(clicallback.Command)(callbackObjectCmd),
	},
}

func callbackBucketCmd(manager *cmdline.CommandManager) {
	manager.RegisterCmd(&cobra.Command{
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Use:                   BucketUsage,
		Short:                 "bucket",
		Long:                  "Allows life-cycle management of a bucket",
		Example:               "singularity bucket list",
		Run:                   run,
		TraverseChildren:      true,
	})
}

func callbackObjectCmd(manager *cmdline.CommandManager) {
	manager.RegisterCmd(&cobra.Command{
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Use:                   ObjectUsage,
		Short:                 "object",
		Long:                  "Allows life-cycle management of a objects",
		Example:               "singularity object list name=bucketName",
		Run:                   run,
		TraverseChildren:      true,
	})
}

func run(cmd *cobra.Command, args []string) {
	resourceType := cmd.Short
	operationType := args[0]
	if err := log.InitLogging("singularity.log", nil, false); err != nil {
		log.Errorf("Error: InitLogging: %v\n", err)
		fmt.Printf("Error: %v\n", err)
		return
	}
	argsMap, err := utils.MakeCommand(args[1:])
	if err != nil {
		log.Errorln(err)
		fmt.Printf("Error: Response is %v\n", err)
		return
	}
	log.Infof("Processing command %v %v\n", resourceType, args)
	defer log.Infof("Processed command %v %v\n", resourceType, args)
	cl := client.NewClient("default", "us-west-2")
	session, err := cl.CreateSession()
	if err != nil {
		log.Errorln(err)
		fmt.Printf("Error: Response is %v", err)
		return
	}
	s3Client := s3.New(session)
	if resourceType == RESOURCE_BUCKET {
		if operationType == CREATE {
			err = cl.CreateBucket(s3Client, argsMap["name"])
			if err != nil {
				log.Errorln(err)
				fmt.Printf("bucket creation failed with error: %v", err)
				return
			}
		} else if operationType == LIST {
			buckets, err := cl.ListBuckets(s3Client)
			if err != nil {
				log.Errorln(err)
				fmt.Printf("Couldn't list buckets: %v", err)
				return
			}
			for _, bucket := range buckets.Buckets {
				fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
			}
		} else {
			log.Errorf("Error: invalid operations %v", operationType)
			fmt.Printf("Error: invalid operations %v", operationType)
			return
		}
	} else if resourceType == RESOURCE_OBJECT {
		if operationType == LIST {
			bucketObjects, err := cl.ListObjects(s3Client, argsMap[BUCKET_NAME], "")
			if err != nil {
				fmt.Printf("Couldn't retrieve bucket items: %v", err)
				return
			}

			for _, item := range bucketObjects.Contents {
				fmt.Printf("Name: %s, Last Modified: %s\n", *item.Key, *item.LastModified)
			}
			return
		} else {
			fmt.Printf("Error: invalid operations %v", operationType)
			log.Errorf("Error: invalid operations %v", operationType)
			return
		}
	} else {
		fmt.Printf("Error: invalid resource type %v", resourceType)
		return
	}

	log.Infof("Command '%v %v' successfully performed\n", resourceType, args)
}
