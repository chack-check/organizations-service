package subscriptions

import "github.com/chack-check/organizations-service/domain/organizations/models"

type SubscriptionsAdapter struct{}

func (adapter SubscriptionsAdapter) GetUserOrganizationConditions(userId int) (*models.OrganizationConditions, error) {
	conditions := models.NewOrganizationConditions(
		5, 5, 5,
	)
	return &conditions, nil
}
