package http

import (
	"encoding/json"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/sub_domain"
	"io"
	"net/http"
	"os"
)

type OnboardingClient struct {
}

func ConstructOnboardingClient() OnboardingClient {
	return OnboardingClient{}
}

func (OnboardingClient) CheckIfUserExists(userIdentifier string) (sub_domain.UserStatusResponse, error) {
	url := os.Getenv("MS_ONBOARDING_URL")
	uri := os.Getenv("MS_ONBOARDING_USER_STATUS_URI")

	resp, err := http.Get(url + uri + "/" + userIdentifier)
	if err != nil {
		// TODO
		return sub_domain.UserStatusResponse{}, err
	}

	switch resp.StatusCode {
	case 200:
		body, IOErr := io.ReadAll(resp.Body)

		if IOErr != nil {
			// TODO
			return sub_domain.UserStatusResponse{}, err
		}

		var userStatus sub_domain.UserStatusResponse
		err := json.Unmarshal(body, &userStatus)

		if err != nil {
			// TODO
			return sub_domain.UserStatusResponse{}, err
		}

		return userStatus, err
	default:
		return sub_domain.UserStatusResponse{}, err
	}
}
