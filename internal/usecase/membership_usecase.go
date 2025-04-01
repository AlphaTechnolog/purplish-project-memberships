package usecase

import (
	"fmt"

	"github.com/alphatechnolog/purplish-memberships/internal/domain"
	"github.com/alphatechnolog/purplish-memberships/internal/repository"
	"github.com/google/uuid"
)

type MembershipUsecase struct {
	sqldbRepo repository.SQLDBRepository
}

func NewMembershipUsecase(sqldbRepo repository.SQLDBRepository) *MembershipUsecase {
	return &MembershipUsecase{
		sqldbRepo,
	}
}

func (uc *MembershipUsecase) GetMemberships() ([]domain.Membership, error) {
	query := "SELECT id, name, description, scopes FROM memberships;"
	memberships := []domain.Membership{}

	rows, err := uc.sqldbRepo.Query(query)
	if err != nil {
		return memberships, fmt.Errorf("unable to query memberships: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var membership domain.Membership
		err = rows.Scan(&membership.ID, &membership.Name, &membership.Description, &membership.Scopes)
		if err != nil {
			return memberships, fmt.Errorf("cannot scan queryset: %w", err)
		}
		memberships = append(memberships, membership)
	}

	return memberships, nil
}

func (uc *MembershipUsecase) GetMembership(id string) (*domain.Membership, error) {
	query := "SELECT id, name, description, scopes FROM memberships WHERE id = ?"
	row := uc.sqldbRepo.QueryRow(query, id)

	membership := &domain.Membership{}
	err := row.Scan(&membership.ID, &membership.Name, &membership.Description, &membership.Scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to scan membership: %w", err)
	}

	return membership, nil
}

func (uc *MembershipUsecase) GetCompanyMembership(id string) (domain.CompanyMembership, error) {
	query := `
	SELECT cm.company_id, cm.membership_id, m.name, m.description, m.scopes
	FROM company_memberships cm
	INNER JOIN memberships m
	ON m.id = cm.membership_id
	WHERE cm.company_id = ?
	LIMIT 1;
	`

	var companyMembership domain.CompanyMembership
	row := uc.sqldbRepo.QueryRow(query, id)
	err := row.Scan(&companyMembership.CompanyID, &companyMembership.MembershipID, &companyMembership.Name, &companyMembership.Description, &companyMembership.Scopes)
	if err != nil {
		return companyMembership, fmt.Errorf("failed to scan company membership: %w", err)
	}

	return companyMembership, nil
}

func (uc *MembershipUsecase) CreateMembership(membership *domain.Membership) error {
	query := "INSERT INTO memberships (id, name, description, scopes) VALUES (?, ?, ?, ?)"
	membership.ID = uuid.New().String()

	_, err := uc.sqldbRepo.Execute(query, membership.ID, membership.Name, membership.Description, membership.Scopes)
	if err != nil {
		return fmt.Errorf("failed to create membership: %w", err)
	}

	return nil
}

func (uc *MembershipUsecase) UpdateMembership(membership *domain.Membership) error {
	query := "UPDATE memberships SET name = ?, description = ?, scopes = ? WHERE id = ?"

	_, err := uc.sqldbRepo.Execute(query, membership.Name, membership.Description, membership.Scopes, membership.ID)
	if err != nil {
		return fmt.Errorf("failed to update membership: %w", err)
	}

	return nil
}

func (uc *MembershipUsecase) DeleteMembership(id string) error {
	query := "DELETE FROM memberships WHERE id = ?"
	_, err := uc.sqldbRepo.Execute(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove membership by id '%s': %w", id, err)
	}

	return nil
}
