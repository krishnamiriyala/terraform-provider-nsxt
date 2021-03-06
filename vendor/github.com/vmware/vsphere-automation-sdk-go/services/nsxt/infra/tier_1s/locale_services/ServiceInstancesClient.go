/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceInstances
 * Used by client-side stubs.
 */

package locale_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceInstancesClient interface {

    // Delete Tier1 policy service instance
    //
    // @param tier1IdParam Tier-0 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Tier1 Service instance id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string) error

    // Read Tier1 service instance
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceInstance
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string) (model.PolicyServiceInstance, error)

    // Read all service instance objects under a tier-1
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyServiceInstanceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, localeServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyServiceInstanceListResult, error)

    // Create Tier1 Service Instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Tier1 Service instance id (required)
    // @param policyServiceInstanceParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, policyServiceInstanceParam model.PolicyServiceInstance) error

    // Create Tier1 service instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Tier1 Service instance id (required)
    // @param policyServiceInstanceParam (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceInstance
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, policyServiceInstanceParam model.PolicyServiceInstance) (model.PolicyServiceInstance, error)
}
