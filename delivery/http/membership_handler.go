package http

import (
	"net/http"

	"github.com/alphatechnolog/purplish-memberships/internal/domain"
	"github.com/alphatechnolog/purplish-memberships/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MembershipHandler struct {
	membershipUsecase *usecase.MembershipUsecase
}

func NewMembershipHandler(membershipUsecase *usecase.MembershipUsecase) *MembershipHandler {
	return &MembershipHandler{membershipUsecase}
}

func (h *MembershipHandler) GetMemberships(c *gin.Context) {
	memberships, err := h.membershipUsecase.GetMemberships()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"memberships": memberships})
}

func (h *MembershipHandler) GetMembership(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	membership, err := h.membershipUsecase.GetMembership(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"membership": membership})
}

func (h *MembershipHandler) GetCompanyMembership(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	membership, err := h.membershipUsecase.GetCompanyMembership(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company_membership": membership})
}

func (h *MembershipHandler) CreateMembership(c *gin.Context) {
	var membership domain.Membership
	if err := c.ShouldBindJSON(&membership); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.membershipUsecase.CreateMembership(&membership); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create membership"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"membership": membership})
}

func (h *MembershipHandler) UpdateMembership(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var membership domain.Membership
	if err := c.ShouldBindJSON(&membership); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	membership.ID = id

	if err := h.membershipUsecase.UpdateMembership(&membership); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update membership"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"membership": membership})
}

func (h *MembershipHandler) DeleteMembership(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.membershipUsecase.DeleteMembership(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete membership"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
