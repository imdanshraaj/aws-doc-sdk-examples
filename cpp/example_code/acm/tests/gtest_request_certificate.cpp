// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
/*
 * Test types are indicated by the test label ending.
 *
 * _1_ Requires credentials, permissions, and AWS resources.
 * _2_ Requires credentials and permissions.
 * _3_ Does not require credentials.
 *
 */

#include <gtest/gtest.h>
#include "acm_samples.h"
#include "acm_gtests.h"

namespace AwsDocTest {
    // NOLINTNEXTLINE(readability-named-parameter)
    TEST_F(ACM_GTests, request_certificate_3_) {
        MockHTTP mockHttp;
        bool result = mockHttp.addResponseWithBody("mock_input/RequestCertificate.json");
        ASSERT_TRUE(result) << preconditionError() << std::endl;

        result = AwsDoc::ACM::requestCertificate("www.example.com", "123456",
                                                 *s_clientConfig);
        ASSERT_TRUE(result);
    }
} // namespace AwsDocTest