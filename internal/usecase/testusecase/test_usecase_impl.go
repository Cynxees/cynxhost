package testusecase

import (
	"cynxhost/internal/app"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
)

type TestUseCaseImpl struct {
	repos        *app.Repos
	dependencies *app.Dependencies
}

func New(repos *app.Repos, dependencies *app.Dependencies) *TestUseCaseImpl {
	return &TestUseCaseImpl{
		repos:        repos,
		dependencies: dependencies,
	}
}

func (usecase *TestUseCaseImpl) CreateDNS(resp *response.APIResponse) {

	porkbunManager := usecase.repos.PorkbunManager

	response, err := porkbunManager.CreateDNS("node1", "47.130.69.79")
	if err != nil {
		resp.Code = responsecode.CodePorkBunError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = response
}

func (usecase *TestUseCaseImpl) RetrieveDNS(resp *response.APIResponse) {
	porkbunManager := usecase.repos.PorkbunManager

	response, err := porkbunManager.RetrieveDNSByTypeSubdomain("A", "node1")
	if err != nil {
		resp.Code = responsecode.CodePorkBunError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = response
}

func (usecase *TestUseCaseImpl) UpdateDNS(resp *response.APIResponse) {
	porkbunManager := usecase.repos.PorkbunManager

	err := porkbunManager.UpdateDNS("A", "node1", "47.130.69.80")
	if err != nil {
		resp.Code = responsecode.CodePorkBunError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
}
