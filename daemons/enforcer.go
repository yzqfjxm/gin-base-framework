package daemons

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
)

//RoleMenusEnforcerDaemon enforcer daemon will modify policy and reload
/*func RoleMenusEnforcerDaemon(pType string, deletePolicyActions []common.PolicyAction, roleId int) {
var (
	addRmMenus []*models.RoleMenu
	globalErr  error
)
en := common.Enforcer
db := common.DB
if len(deletePolicyActions) > 0 {
	for _, deP := range deletePolicyActions {
		en.RemovePolicy(deP.PType, deP.Address, deP.Method)
	}
}

db.Where("role_id = ?", roleId).Find(&addRmMenus)
//  todo need modify
*/ /*for _, rm := range addRmMenus {
	var reApi models.ResourceApi
	if e := db.Where("id = ?", rm.ResourceApiID).First(&reApi).Error; e == nil {
		en.AddPolicy(pType, reApi.Address, reApi.Method)

	} else {
		globalErr = e
		log.Errorf("daemons casbin add policy role menus failed: %s", e)
	}
}*/ /*
	if globalErr == nil {
		en.SavePolicy()
		en.LoadPolicy()
	} else {
		log.Errorf("daemons casbin add policy role menus failed: %s", globalErr)
	}

}*/

func RoleSystemApiEnforcerDaemon(policies []common.PolicyAction) (err error) {
	en := common.Enforcer
	for _, rule := range policies {
		en.AddPolicy(rule.PType, rule.Address, rule.Method)

	}
	err = en.SavePolicy()
	if err == nil {
		en.LoadPolicy()

	}
	return
}

//UserOrRoleEnforcerDaemon enforcer daemon will modify policy and reload
func UserOrRoleEnforcerDaemon(groupPolicyActions []common.GroupPolicyAction) {
	en := common.Enforcer
	if len(groupPolicyActions) > 0 {
		for _, gpa := range groupPolicyActions {
			switch gpa.Action {
			case "delete":
				en.RemoveGroupingPolicy(gpa.UserOrRole, gpa.Role)
			case "add":
				en.AddGroupingPolicy(gpa.UserOrRole, gpa.Role)

			}
		}
		en.SavePolicy()
		en.LoadPolicy()
	}
}
