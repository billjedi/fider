package actions_test

import (
	"testing"

	"github.com/getfider/fider/app/actions"
	"github.com/getfider/fider/app/models"

	. "github.com/onsi/gomega"
)

func TestCreateTenant_ShouldHaveVerificationKey(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{}
	action.Initialize()

	Expect(action.Model.VerificationKey).NotTo(Equal(""))
}

func TestCreateTenant_EmptyToken(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Token: ""}}
	result := action.Validate(services)
	ExpectFailed(result, "token", "tenantName", "subdomain")
}

func TestCreateTenant_EmptyTenantName(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Token: jonSnowToken, TenantName: ""}}
	result := action.Validate(services)
	ExpectFailed(result, "tenantName", "subdomain")
}

func TestCreateTenant_EmptyEmail(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Name: "Jon Snow", Email: ""}}
	result := action.Validate(services)
	ExpectFailed(result, "email", "tenantName", "subdomain")
}

func TestCreateTenant_InvalidEmail(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Name: "Jon Snow", Email: "jonsnow"}}
	result := action.Validate(services)
	ExpectFailed(result, "email", "tenantName", "subdomain")
}

func TestCreateTenant_EmptyName(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Name: "", Email: "jon.snow@got.com"}}
	result := action.Validate(services)
	ExpectFailed(result, "name", "tenantName", "subdomain")
}

func TestCreateTenant_EmptySubdomain(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Token: jonSnowToken, TenantName: "My Company"}}
	result := action.Validate(services)
	ExpectFailed(result, "subdomain")
}

func TestCreateTenant_UpperCaseSubdomain(t *testing.T) {
	RegisterTestingT(t)

	action := actions.CreateTenant{Model: &models.CreateTenant{Token: jonSnowToken, TenantName: "My Company", Subdomain: "MyCompany"}}
	result := action.Validate(services)
	ExpectSuccess(result)
	Expect(action.Model.Subdomain).To(Equal("mycompany"))
}

func TestUpdateTenantSettings_EmptyTitle(t *testing.T) {
	RegisterTestingT(t)

	action := actions.UpdateTenantSettings{Model: &models.UpdateTenantSettings{}}
	result := action.Validate(services)
	ExpectFailed(result, "title")
}

func TestUpdateTenantSettings_LargeTitle(t *testing.T) {
	RegisterTestingT(t)

	action := actions.UpdateTenantSettings{Model: &models.UpdateTenantSettings{Title: "123456789012345678901234567890123456789012345678901234567890123"}}
	result := action.Validate(services)
	ExpectFailed(result, "title")
}

func TestUpdateTenantSettings_LargeInvitation(t *testing.T) {
	RegisterTestingT(t)

	action := actions.UpdateTenantSettings{Model: &models.UpdateTenantSettings{Title: "Ok", Invitation: "123456789012345678901234567890123456789012345678901234567890123"}}
	result := action.Validate(services)
	ExpectFailed(result, "invitation")
}
