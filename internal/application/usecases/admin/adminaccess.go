package admin

import "GoProject1/internal/adapter/persistence"

// UC_Accesadmin просто прокидывает проверку прав админа из persistence.
func UC_Accesadmin(userID int) bool {
	return persistence.A_adminaccess(userID)
}
