﻿// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved. 
// SPDX-License-Identifier: Apache-2.0

namespace BedrockRuntimeTests
{
    public class ActionTest_InvokeModelWithResponseStream
    {
        [Theory, Trait("Category", "Integration")]
        [InlineData(typeof(AmazonTitanText.InvokeModelWithResponseStream))]
        [InlineData(typeof(Mistral.InvokeModelWithResponseStream))]
        [InlineData(typeof(MetaLlama3.InvokeModelWithResponseStream))]
        [InlineData(typeof(CohereCommand.InvokeModelWithResponseStream))]
        [InlineData(typeof(CohereCommandR.InvokeModelWithResponseStream))]
        [InlineData(typeof(AnthropicClaude.InvokeModelWithResponseStream))]

        public void InvokeModelWithResponseStreamDoesNotThrow(Type type)
        {
            Thread.Sleep(5000);
            var entryPoint = type.Assembly.EntryPoint!;
            var exception = Record.Exception(() => entryPoint.Invoke(null, [Array.Empty<string>()]));
            Assert.Null(exception);
        }
    }
}