# Amazon ECR code examples for the SDK for Java 2.x

## Overview

Shows how to use the AWS SDK for Java 2.x to work with Amazon Elastic Container Registry (Amazon ECR).

<!--custom.overview.start-->
<!--custom.overview.end-->

_Amazon ECR is a fully managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images._

## ⚠ Important

* Running this code might result in charges to your AWS account. For more details, see [AWS Pricing](https://aws.amazon.com/pricing/) and [Free Tier](https://aws.amazon.com/free/).
* Running the tests might result in charges to your AWS account.
* We recommend that you grant your code least privilege. At most, grant only the minimum permissions required to perform the task. For more information, see [Grant least privilege](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#grant-least-privilege).
* This code is not tested in every AWS Region. For more information, see [AWS Regional Services](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services).

<!--custom.important.start-->
<!--custom.important.end-->

## Code examples

### Prerequisites

For prerequisites, see the [README](../../README.md#Prerequisites) in the `javav2` folder.


<!--custom.prerequisites.start-->
<!--custom.prerequisites.end-->

### Get started

- [Hello Amazon ECR](src/main/java/com/example/ecr/HelloECR.java#L6) (`listImages`)


### Basics

Code examples that show you how to perform the essential operations within a service.

- [Learn the basics](src/main/java/com/example/ecr/scenario/ECRScenario.java)


### Single actions

Code excerpts that show you how to call individual service functions.

- [CreateRepository](src/main/java/com/example/ecr/scenario/ECRActions.java#L54)
- [DeleteRepository](src/main/java/com/example/ecr/scenario/ECRActions.java#L101)
- [DescribeImages](src/main/java/com/example/ecr/scenario/ECRActions.java#L244)
- [DescribeRepositories](src/main/java/com/example/ecr/scenario/ECRActions.java#L284)
- [GetAuthorizationToken](src/main/java/com/example/ecr/scenario/ECRActions.java#L329)
- [GetRepositoryPolicy](src/main/java/com/example/ecr/scenario/ECRActions.java#L361)
- [PushImageCmd](src/main/java/com/example/ecr/scenario/ECRActions.java#L451)
- [SetRepositoryPolicy](src/main/java/com/example/ecr/scenario/ECRActions.java#L396)
- [StartLifecyclePolicyPreview](src/main/java/com/example/ecr/scenario/ECRActions.java#L244)


<!--custom.examples.start-->
<!--custom.examples.end-->

## Run the examples

### Instructions


<!--custom.instructions.start-->
<!--custom.instructions.end-->

#### Hello Amazon ECR

This example shows you how to get started using Amazon ECR.


#### Learn the basics

This example shows you how to do the following:

- Create an Amazon ECR repository.
- Set repository policies.
- Retrieve repository URIs.
- Get Amazon ECR authorization tokens.
- Set lifecycle policies for Amazon ECR repositories.
- Push a Docker image to an Amazon ECR repository.
- Verify the existence of an image in an Amazon ECR repository.
- List Amazon ECR repositories for your account and get details about them.
- Delete Amazon ECR repositories.

<!--custom.basic_prereqs.ecr_Scenario_RepositoryManagement.start-->
<!--custom.basic_prereqs.ecr_Scenario_RepositoryManagement.end-->


<!--custom.basics.ecr_Scenario_RepositoryManagement.start-->
<!--custom.basics.ecr_Scenario_RepositoryManagement.end-->


### Tests

⚠ Running tests might result in charges to your AWS account.


To find instructions for running these tests, see the [README](../../README.md#Tests)
in the `javav2` folder.



<!--custom.tests.start-->
<!--custom.tests.end-->

## Additional resources

- [Amazon ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/what-is-ecr.html)
- [Amazon ECR API Reference](https://docs.aws.amazon.com/AmazonECR/latest/APIReference/Welcome.html)
- [SDK for Java 2.x Amazon ECR reference](https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/services/ecr/package-summary.html)

<!--custom.resources.start-->
<!--custom.resources.end-->

---

Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0