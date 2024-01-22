package catalog

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestExternalLocationCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					Owner:          "x",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "administrators",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		owner = "administrators"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationReadOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					ReadOnly:       true,
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					ReadOnly:       true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
					ReadOnly:    true,
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		read_only = true
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationWithAPAndEncryptionDetails(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					AccessPoint:    "some_access_point",
					EncryptionDetails: &catalog.EncryptionDetails{
						SseEncryptionDetails: &catalog.SseEncryptionDetails{
							Algorithm:    "AWS_SSE_KMS",
							AwsKmsKeyArn: "some_key_arn",
						},
					},
					Comment: "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					AccessPoint:    "some_access_point",
					EncryptionDetails: &catalog.EncryptionDetails{
						SseEncryptionDetails: &catalog.SseEncryptionDetails{
							Algorithm:    "AWS_SSE_KMS",
							AwsKmsKeyArn: "some_key_arn",
						},
					},
					Comment: "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		access_point = "some_access_point"
	    encryption_details {
          sse_encryption_details {
            algorithm     = "AWS_SSE_KMS"
            aws_kms_key_arn = "some_key_arn"
		  }
        }
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationOnlyOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "abc",
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationOwnerAndOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "abc",
					Comment:        "def",
					Owner:          "administrators",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateExternalLocationRollbackError(t *testing.T) {
	serverErrMessage := "Something unexpected happened"
	rollbackErrMessage := "Internal error happened"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "administrators",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.Apply(t)
	errOccurred := fmt.Sprintf("%s. Owner rollback also failed: %s", serverErrMessage, rollbackErrMessage)
	qa.AssertErrorStartsWith(t, err, errOccurred)
}

func TestUpdateExternalLocationForce(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Force:          true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		force_update = true
		`,
	}.ApplyNoError(t)
}
