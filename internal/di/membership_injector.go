package di

import (
	"database/sql"

	"github.com/alphatechnolog/purplish-memberships/delivery/http"
	"github.com/alphatechnolog/purplish-memberships/infrastructure/database"
	"github.com/alphatechnolog/purplish-memberships/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MembershipInjector struct {
	db *sql.DB
}

func NewMembershipInjector(db *sql.DB) ModuleInjector {
	return &MembershipInjector{db: db}
}

func (mi *MembershipInjector) Inject(routerGroup *gin.RouterGroup) {
	sqliteRepo := database.NewSQLiteRepository(mi.db)
	membershipUseCase := usecase.NewMembershipUsecase(sqliteRepo)
	membershipHandler := http.NewMembershipHandler(membershipUseCase)

	routerGroup.GET("/", membershipHandler.GetMemberships)
	routerGroup.GET("/:id/", membershipHandler.GetMembership)
	routerGroup.GET("/company-membership/:id", membershipHandler.GetCompanyMembership)
	routerGroup.POST("/", http.APIGatewayScopeCheck("c:memberships"), membershipHandler.CreateMembership)
	routerGroup.PUT("/:id/", http.APIGatewayScopeCheck("u:memberships"), membershipHandler.UpdateMembership)
	routerGroup.DELETE("/:id/", http.APIGatewayScopeCheck("d:memberships"), membershipHandler.DeleteMembership)
}
