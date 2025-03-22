package database

import (
    "database/sql"
    "fmt"
    "strings"

    "github.com/google/uuid"
)

type Membership struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Description *string `json:"description"`
    Scopes string `json:"-"`
}

type CreateMembershipPayload struct {
    Name string `json:"name"`
    Description *string `json:"description"`
    Scopes string `json:"scopes"`
}

type CompanyMembership struct {
    CompanyID string `json:"company_id"`
    MembershipID string `json:"membership_id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Scopes string `json:"scopes"`
}

func GetMemberships(d *sql.DB) ([]Membership, error) {
    var memberships []Membership

    sql := "SELECT m.id, m.name, m.description, m.scopes FROM memberships m";
    rows, err := d.Query(sql)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    for rows.Next() {
        var membership Membership
        err = rows.Scan(&membership.ID, &membership.Name, &membership.Description, &membership.Scopes)
        if err != nil {
            return nil, err
        }

        // FIXME: Stop hardcoding this.
        if strings.ToLower(membership.Name) != "admin" {
            memberships = append(memberships, membership)
        }
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return memberships, nil
}

func GetMembership(d *sql.DB, ID string) (Membership, error) {
    var membership Membership

    sql := "SELECT m.id, m.name, m.description, m.scopes FROM memberships m WHERE m.id = ? LIMIT 1;"
    row := d.QueryRow(sql, ID)
    err := row.Scan(&membership.ID, &membership.Name, &membership.Description, &membership.Scopes)

    if err != nil {
        return Membership{}, err
    }

    return membership, nil
}

func GetCompanyMembership(d *sql.DB, companyID string) (CompanyMembership, error) {
    var companyMembership CompanyMembership

    sql := `
        SELECT cm.company_id, cm.membership_id, m.name, m.description, m.scopes
        FROM company_memberships cm
        INNER JOIN memberships m
        ON m.id = cm.membership_id
        WHERE cm.company_id = ?
        LIMIT 1;
    `

    row := d.QueryRow(sql, companyID)
    err := row.Scan(
        &companyMembership.CompanyID,
        &companyMembership.MembershipID,
        &companyMembership.Name,
        &companyMembership.Description,
        &companyMembership.Scopes,
    )

    if err != nil {
        return CompanyMembership{}, err
    }

    return companyMembership, nil
}

func CreateMembership(d *sql.DB, createPayload CreateMembershipPayload) error {
    sql := `
        INSERT INTO memberships (id, name, description, scopes)
        VALUES
            (?, ?, ?, ?);
    `

    _, err := d.Exec(sql, uuid.New().String(), createPayload.Name, createPayload.Description, createPayload.Scopes)
    if err != nil {
        return fmt.Errorf("Unable to create new membership: %w", err)
    }

    return nil
}

func RemoveMembership(d *sql.DB, membershipID string) error {
    sql := `
        DELETE FROM memberships WHERE id = ?
    `

    if _, err := d.Exec(sql, membershipID); err != nil {
        return fmt.Errorf("Unable to remove non-matching membership by id = '%s': %w", membershipID, err)
    }

    return nil
}