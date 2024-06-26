// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func decodeOutputb64(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}

	return string(decoded)
}

func main() {
	// Create new EC2 service
	ec2Svc := ec2.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	})))

	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <instance-id>\n", os.Args[0])
		os.Exit(1)
	}

	instanceId := os.Args[1]

	// Call EC2 GetConsoleOutput API on the given instance according
	//   https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.GetConsoleOutput

	input := ec2.GetConsoleOutputInput{InstanceId: aws.String(instanceId)}
	json, err := ec2Svc.GetConsoleOutput(&input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(decodeOutputb64(*json.Output))
	}
}
