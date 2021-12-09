package http

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/sub_domain"
	"io"
	"log"
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
		return sub_domain.UserStatusResponse{}, err
	}

	switch resp.StatusCode {
	case 200:
		defer resp.Body.Close()
		body, IOErr := io.ReadAll(resp.Body)
		if IOErr != nil {
			return sub_domain.UserStatusResponse{}, err
		}

		log.Print(body)

		return sub_domain.UserStatusResponse{}, err
	default:
		return sub_domain.UserStatusResponse{}, err
	}
}
