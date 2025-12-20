package admin

import "GoProject1/internal/adapter/persistence"

func UC_Accesadmin(userID int) bool {
	if persistence.A_adminaccess(userID) {
		return true // права админа подтверждены
	} else {
		return false // права админа не подтверждены
	}
}
