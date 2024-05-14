// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontacts_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// SSMContacts resources depend on a replication set existing and
// only one replication set resource can be active at once, so we must have serialised tests
func TestAccSSMContacts_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Contact Resource Tests": {
			"basic":             testAccContact_basic,
			"disappears":        testAccContact_disappears,
			"updateAlias":       testAccContact_updateAlias,
			"updateDisplayName": testAccContact_updateDisplayName,
			"updateTags":        testAccContact_updateTags,
			"updateType":        testAccContact_updateType,
		},
		"Contact Data Source Tests": {
			"basic": testAccContactDataSource_basic,
		},
		"Contact Channel Resource Tests": {
			"basic":           testAccContactChannel_basic,
			"contactId":       testAccContactChannel_contactID,
			"deliveryAddress": testAccContactChannel_deliveryAddress,
			"disappears":      testAccContactChannel_disappears,
			names.AttrName:    testAccContactChannel_name,
			names.AttrType:    testAccContactChannel_type,
		},
		"Contact Channel Data Source Tests": {
			"basic": testAccContactChannelDataSource_basic,
		},
		"Plan Resource Tests": {
			"basic":                   testAccPlan_basic,
			"disappears":              testAccPlan_disappears,
			"updateChannelTargetInfo": testAccPlan_updateChannelTargetInfo,
			"updateContactId":         testAccPlan_updateContactId,
			"updateContactTargetInfo": testAccPlan_updateContactTargetInfo,
			"updateDurationInMinutes": testAccPlan_updateDurationInMinutes,
			"updateStages":            testAccPlan_updateStages,
			"updateTargets":           testAccPlan_updateTargets,
		},
		"Plan Data Source Tests": {
			"basic":             testAccPlanDataSource_basic,
			"channelTargetInfo": testAccPlanDataSource_channelTargetInfo,
		},
		"RotationResource": {
			"basic":        testAccRotation_basic,
			"disappears":   testAccRotation_disappears,
			"update":       testAccRotation_updateRequiredFields,
			"startTime":    testAccRotation_startTime,
			"contactIds":   testAccRotation_contactIds,
			"recurrence":   testAccRotation_recurrence,
			names.AttrTags: testAccRotation_tags,
		},
		"RotationDataSource": {
			"basic":           testAccRotationDataSource_basic,
			"dailySettings":   testAccRotationDataSource_dailySettings,
			"monthlySettings": testAccRotationDataSource_monthlySettings,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
