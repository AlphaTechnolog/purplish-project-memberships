package core

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alphatechnolog/purplish-memberships/database"
	"github.com/gin-gonic/gin"
)

func getMemberships(d *sql.DB, c *gin.Context) error {
	memberships, err := database.GetMemberships(d)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"memberships": memberships})

	return nil
}

func getMembership(d *sql.DB, c *gin.Context) error {
	membershipID := c.Param("MembershipID")
	if membershipID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify Membership ID"})
		return nil
	}

	membership, err := database.GetMembership(d, membershipID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"membership": membership})

	return nil
}

func getCompanyMembership(d *sql.DB, c *gin.Context) error {
    companyID := c.Param("CompanyID")
    if companyID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Specify company id"})
        return nil
    }

    companyMembership, err := database.GetCompanyMembership(d, companyID)
    if err != nil {
        return err
    }

    c.JSON(http.StatusOK, gin.H{"company_membership": companyMembership})

    return nil
}

// TODO: This will require we to implement authentication to check user permissions and user token.
func createMembership(d *sql.DB, c *gin.Context) error {
	bodyContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	var createPayload database.CreateMembershipPayload
	if err = json.Unmarshal(bodyContents, &createPayload); err != nil {
		return err
	}

	if err = database.CreateMembership(d, createPayload); err != nil {
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})

	return nil
}

func removeMembership(d *sql.DB, c *gin.Context) error {
	membershipID := c.Param("MembershipID")
	if membershipID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify Membership ID"})
		return nil
	}

	if err := database.RemoveMembership(d, membershipID); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	return nil
}

func CreateMembershipsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getMemberships)))
	r.GET("/:MembershipID", WrapError(WithDB(d, getMembership)))
	r.GET("/company-membership/:CompanyID", WrapError(WithDB(d, getCompanyMembership)))
	r.POST("/", WrapError(WithDB(d, createMembership)))
	r.DELETE("/:MembershipID", WrapError(WithDB(d, removeMembership)))
}
