# Amazon Bedrock Runtime code examples for the SDK for .NET

## Overview

Shows how to use the AWS SDK for .NET to work with Amazon Bedrock Runtime.

<!--custom.overview.start-->
<!--custom.overview.end-->

_Amazon Bedrock Runtime is a fully managed service that makes it easy to use foundation models from third-party providers and Amazon._

## ⚠ Important

* Running this code might result in charges to your AWS account. For more details, see [AWS Pricing](https://aws.amazon.com/pricing/) and [Free Tier](https://aws.amazon.com/free/).
* Running the tests might result in charges to your AWS account.
* We recommend that you grant your code least privilege. At most, grant only the minimum permissions required to perform the task. For more information, see [Grant least privilege](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#grant-least-privilege).
* This code is not tested in every AWS Region. For more information, see [AWS Regional Services](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services).

<!--custom.important.start-->
<!--custom.important.end-->

## Code examples

### Prerequisites

For prerequisites, see the [README](../README.md#Prerequisites) in the `dotnetv3` folder.


<!--custom.prerequisites.start-->
<!--custom.prerequisites.end-->
### Scenarios

Code examples that show you how to accomplish a specific task by calling multiple
functions within the same service.

- [Tool use with the Converse API](Scenarios/ConverseToolScenario/ConverseToolScenario.cs)

### Amazon Nova

- [Converse](Models/AmazonNova/AmazonNovaText/Converse/Converse.cs#L4)
- [ConverseStream](Models/AmazonNova/AmazonNovaText/ConverseStream/ConverseStream.cs#L4)
- [Scenario: Tool use with the Converse API](Scenarios/ConverseToolScenario/ConverseToolScenario.cs#L4)

### Amazon Nova Canvas

- [InvokeModel](Models/AmazonNova/AmazonNovaCanvas/InvokeModel/InvokeModel.cs#L4)

### Amazon Titan Text

- [InvokeModelWithResponseStream](Models/AmazonTitanText/InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)

### Anthropic Claude

- [InvokeModelWithResponseStream](Models/AnthropicClaude/InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)

### Cohere Command

- [InvokeModel: Command R and R+](Models/CohereCommand/Command_R_InvokeModel/InvokeModel.cs#L4)
- [InvokeModel: Command and Command Light](Models/CohereCommand/Command_InvokeModel/InvokeModel.cs#L4)
- [InvokeModelWithResponseStream: Command R and R+](Models/CohereCommand/Command_R_InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)
- [InvokeModelWithResponseStream: Command and Command Light](Models/CohereCommand/Command_InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)

### Meta Llama

- [InvokeModel](Models/MetaLlama/Llama3_InvokeModel/InvokeModel.cs#L4)
- [InvokeModelWithResponseStream](Models/MetaLlama/Llama3_InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)

### Mistral AI

- [InvokeModel](Models/Mistral/InvokeModel/InvokeModel.cs#L4)
- [InvokeModelWithResponseStream](Models/Mistral/InvokeModelWithResponseStream/InvokeModelWithResponseStream.cs#L4)


<!--custom.examples.start-->
<!--custom.examples.end-->

## Run the examples

### Instructions

For general instructions to run the examples, see the
[README](../README.md#building-and-running-the-code-examples) in the `dotnetv3` folder.

Some projects might include a settings.json file. Before compiling the project,
you can change these values to match your own account and resources. Alternatively,
add a settings.local.json file with your local settings, which will be loaded automatically
when the application runs.

After the example compiles, you can run it from the command line. To do so, navigate to
the folder that contains the .csproj file and run the following command:

```
dotnet run
```

Alternatively, you can run the example from within your IDE.


<!--custom.instructions.start-->
<!--custom.instructions.end-->



#### Tool use with the Converse API

This example shows you how to build a typical interaction between an application, a generative AI model, and connected tools or APIs to mediate interactions between the AI and the outside world. It uses the example of connecting an external weather API to the AI model so it can provide real-time weather information based on user input.


<!--custom.scenario_prereqs.bedrock-runtime_Scenario_ToolUse.start-->
<!--custom.scenario_prereqs.bedrock-runtime_Scenario_ToolUse.end-->


<!--custom.scenarios.bedrock-runtime_Scenario_ToolUse.start-->
<!--custom.scenarios.bedrock-runtime_Scenario_ToolUse.end-->

### Tests

⚠ Running tests might result in charges to your AWS account.


To find instructions for running these tests, see the [README](../README.md#Tests)
in the `dotnetv3` folder.



<!--custom.tests.start-->
<!--custom.tests.end-->

## Additional resources

- [Amazon Bedrock Runtime User Guide](https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-bedrock.html)
- [Amazon Bedrock Runtime API Reference](https://docs.aws.amazon.com/bedrock/latest/APIReference/welcome.html)
- [SDK for .NET Amazon Bedrock Runtime reference](https://docs.aws.amazon.com/sdkfornet/v3/apidocs/items/Bedrock-runtime/NBedrock-runtime.html)

<!--custom.resources.start-->
<!--custom.resources.end-->

---

Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
