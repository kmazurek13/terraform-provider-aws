// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
)

func unsuccessfulItemErrorV2(apiObject *awstypes.UnsuccessfulItemError) error {
	if apiObject == nil {
		return nil
	}

	return &smithy.GenericAPIError{
		Code:    aws.ToString(apiObject.Code),
		Message: aws.ToString(apiObject.Message),
	}
}

func unsuccessfulItemsErrorV2(apiObjects []awstypes.UnsuccessfulItem) error {
	var errs []error

	for _, apiObject := range apiObjects {
		if err := unsuccessfulItemErrorV2(apiObject.Error); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", aws.ToString(apiObject.ResourceId), err))
		}
	}

	return errors.Join(errs...)
}
