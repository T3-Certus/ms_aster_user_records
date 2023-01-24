package utils_methods

import (
	utils_interfaces "github.com/ssssshel/ms_aster_user_data_go/src/utils/interfaces"
)

func GenericSuccessResponse(status uint16) *utils_interfaces.GenericSuccessResponse {
	// var serverMessage string
	// switch status {
	// case 200:
	// 	serverMessage = ""
	// }
	response := &utils_interfaces.GenericSuccessResponse{

		MoreDetails: "prueba",
	}

	return response
}
