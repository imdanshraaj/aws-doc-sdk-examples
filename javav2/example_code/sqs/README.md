# Amazon SQS code examples for the SDK for Java 2.x

## Overview

Shows how to use the AWS SDK for Java 2.x to work with Amazon Simple Queue Service (Amazon SQS).

<!--custom.overview.start-->
<!--custom.overview.end-->

_Amazon SQS is a fully managed message queuing service that makes it easy to decouple and scale microservices, distributed systems, and serverless applications._

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

- [Hello Amazon SQS](src/main/java/com/example/sqs/HelloSQS.java#L6) (`ListQueues`)


### Single actions

Code excerpts that show you how to call individual service functions.

- [CreateQueue](src/main/java/com/example/sqs/SQSExample.java#L6)
- [DeleteMessage](src/main/java/com/example/sqs/SQSExample.java#L190)
- [DeleteQueue](src/main/java/com/example/sqs/DeleteQueue.java#L6)
- [GetQueueUrl](src/main/java/com/example/sqs/SQSExample.java#L66)
- [ListQueues](src/main/java/com/example/sqs/SQSExample.java#L82)
- [ReceiveMessage](src/main/java/com/example/sqs/SQSExample.java#L151)
- [SendMessage](src/main/java/com/example/sqs/SendMessages.java#L7)
- [SendMessageBatch](src/main/java/com/example/sqs/SQSExample.java#L132)
- [SetQueueAttributes](src/main/java/com/example/sqs/SSEncryptionExample.java#L64)

### Scenarios

Code examples that show you how to accomplish a specific task by calling multiple
functions within the same service.

- [Create and publish to a FIFO topic](../sns/src/main/java/com/example/sns/PriceUpdateExample.java)
- [Manage large messages using S3](src/main/java/com/example/sqs/SqsExtendedClientExample.java)
- [Process S3 event notifications](../s3/src/main/java/com/example/s3/ProcessS3EventNotification.java)
- [Publish messages to queues](../../usecases/topics_and_queues/src/main/java/com/example/sns/SNSWorkflow.java)
- [Send and receive batches of messages](src/main/java/com/example/sqs/SendRecvBatch.java)
- [Use the Amazon SQS Java Messaging Library to work with the JMS interface](../sqs-jms/src/main/java/com/example/sqs/jms/stdqueue/TextMessageSender.java)
- [Work with queue tags](src/main/java/com/example/sqs/TagExamples.java)


<!--custom.examples.start-->
<!--custom.examples.end-->

## Run the examples

### Instructions


<!--custom.instructions.start-->
<!--custom.instructions.end-->

#### Hello Amazon SQS

This example shows you how to get started using Amazon SQS.



#### Create and publish to a FIFO topic

This example shows you how to create and publish to a FIFO Amazon SNS topic.


<!--custom.scenario_prereqs.sns_PublishFifoTopic.start-->
<!--custom.scenario_prereqs.sns_PublishFifoTopic.end-->


<!--custom.scenarios.sns_PublishFifoTopic.start-->
<!--custom.scenarios.sns_PublishFifoTopic.end-->

#### Manage large messages using S3

This example shows you how to use the Amazon SQS Extended Client Library to work with large Amazon SQS messages.


<!--custom.scenario_prereqs.sqs_Scenario_SqsExtendedClient.start-->
<!--custom.scenario_prereqs.sqs_Scenario_SqsExtendedClient.end-->


<!--custom.scenarios.sqs_Scenario_SqsExtendedClient.start-->
<!--custom.scenarios.sqs_Scenario_SqsExtendedClient.end-->

#### Process S3 event notifications

This example shows you how to work with S3 event notifications in an object-oriented way.


<!--custom.scenario_prereqs.s3_Scenario_ProcessS3EventNotification.start-->
<!--custom.scenario_prereqs.s3_Scenario_ProcessS3EventNotification.end-->


<!--custom.scenarios.s3_Scenario_ProcessS3EventNotification.start-->
<!--custom.scenarios.s3_Scenario_ProcessS3EventNotification.end-->

#### Publish messages to queues

This example shows you how to do the following:

- Create topic (FIFO or non-FIFO).
- Subscribe several queues to the topic with an option to apply a filter.
- Publish messages to the topic.
- Poll the queues for messages received.

<!--custom.scenario_prereqs.sqs_Scenario_TopicsAndQueues.start-->
<!--custom.scenario_prereqs.sqs_Scenario_TopicsAndQueues.end-->


<!--custom.scenarios.sqs_Scenario_TopicsAndQueues.start-->
<!--custom.scenarios.sqs_Scenario_TopicsAndQueues.end-->

#### Send and receive batches of messages

This example shows you how to do the following:

- Create an Amazon SQS queue.
- Send batches of messages to the queue.
- Receive batches of messages from the queue.
- Delete batches of messages from the queue.

<!--custom.scenario_prereqs.sqs_Scenario_SendReceiveBatch.start-->
<!--custom.scenario_prereqs.sqs_Scenario_SendReceiveBatch.end-->


<!--custom.scenarios.sqs_Scenario_SendReceiveBatch.start-->
<!--custom.scenarios.sqs_Scenario_SendReceiveBatch.end-->

#### Use the Amazon SQS Java Messaging Library to work with the JMS interface

This example shows you how to use the Amazon SQS Java Messaging Library to work with the JMS interface.


<!--custom.scenario_prereqs.sqs_Scenario_UseJMS.start-->
<!--custom.scenario_prereqs.sqs_Scenario_UseJMS.end-->


<!--custom.scenarios.sqs_Scenario_UseJMS.start-->
<!--custom.scenarios.sqs_Scenario_UseJMS.end-->

#### Work with queue tags

This example shows you how to perform tagging operation with Amazon SQS.


<!--custom.scenario_prereqs.sqs_Scenario_WorkWithTags.start-->
<!--custom.scenario_prereqs.sqs_Scenario_WorkWithTags.end-->


<!--custom.scenarios.sqs_Scenario_WorkWithTags.start-->
<!--custom.scenarios.sqs_Scenario_WorkWithTags.end-->

### Tests

⚠ Running tests might result in charges to your AWS account.


To find instructions for running these tests, see the [README](../../README.md#Tests)
in the `javav2` folder.



<!--custom.tests.start-->
<!--custom.tests.end-->

## Additional resources

- [Amazon SQS Developer Guide](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html)
- [Amazon SQS API Reference](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/Welcome.html)
- [SDK for Java 2.x Amazon SQS reference](https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/services/sqs/package-summary.html)

<!--custom.resources.start-->
<!--custom.resources.end-->

---

Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
