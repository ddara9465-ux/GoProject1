package masters

import "GoProject1/internal/adapter/persistence"

func UC_CreateMaster(name, specialization string) {
	persistence.A_CreateMaster(name, specialization)
}

func UC_GetMasters() []map[string]interface{} {
	return persistence.A_GetMasters()
}

func UC_DeleteMaster(masterID int) {
	persistence.A_DeleteMaster(masterID)
}
