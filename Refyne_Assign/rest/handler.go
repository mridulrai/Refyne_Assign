package rest

import (
	"net/http"
	"strconv"
	"strings"

	"../storage"
	"github.com/labstack/echo/v4"
)

// index is a handler function
func index(c echo.Context) error {
	return c.String(http.StatusOK, "Account service")
}

// health is a handler function
func health(c echo.Context) error {
	var errResp ErrorResponseData
	err := storage.Health()
	if err != nil {
		errResp.Data.Code = "database_connection_error"
		errResp.Data.Description = err.Error()
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}
	return c.String(http.StatusOK, "Healthy")
}

// createUser is a handler function for creating a new account
func createUser(c echo.Context) error {
	var errResp ErrorResponseData
	var resp UserResponseData

	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		errResp.Data.Code = "request_binding_error"
		errResp.Data.Description = "Unable to bind request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	user := req.mapToModel()
	err := storage.CreateUser(&user)

	if err != nil {
		errResp.Data.Code = "create_account_error"
		errResp.Data.Description = "Unable to create account"
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	resp.mapFromModel(user)
	return c.JSON(http.StatusCreated, resp)
}

//addCars is a Handler Function to add Cars
func addCars(c echo.Context) error {
	var errResp ErrorResponseData
	var resp UserResponseData

	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		errResp.Data.Code = "request_binding_error"
		errResp.Data.Description = "Unable to bind request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	account := req.mapToModel()
	err := storage.CreateUser(&account)

	if err != nil {
		errResp.Data.Code = "create_account_error"
		errResp.Data.Description = "Unable to create account"
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	resp.mapFromModel(account)
	return c.JSON(http.StatusCreated, resp)
}

// getAccount is a handler function for fetching account based on id
func getAccount(c echo.Context) error {
	var errResp ErrorResponseData
	var resp UserResponseData

	id := strings.TrimSpace(c.Param("id"))
	if len(id) == 0 {
		errResp.Data.Code = "invalid_param_error"
		errResp.Data.Description = "Value for account id not set in request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	account, err := storage.GetAccount(id)

	if err != nil {
		errResp.Data.Code = "get_account_error"
		errResp.Data.Description = "Unable to fetch account details"
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	if account == nil {
		errResp.Data.Code = "no_account_found"
		errResp.Data.Description = "No account with id " + id + " exists"
		errResp.Data.Status = strconv.Itoa(http.StatusNotFound)
		return c.JSON(http.StatusNotFound, errResp)
	}
	resp.mapFromModel(account)

	return c.JSON(http.StatusOK, resp)
}

// listAccounts is a handler for listing accounts in paginated format
func listAccounts(c echo.Context) error {
	var errResp ErrorResponseData
	var resp AccountListResponseData

	fromDateTime, err := strconv.Atoi(c.Param("dromDateTime"))

	if (err != nil) || (fromDateTime <= 0) {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value   in query parameter fromDateTime"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}
	toDateTime, err := strconv.Atoi(c.Param("toDateTime"))

	if (err != nil) || (toDateTime <= 0) || toDateTime < fromDateTime {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value "
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	totalItems, accounts, err := storage.GetAccountList(fromDateTime, toDateTime)

	if err != nil {
		errResp.Data.Code = "error"
		errResp.Data.Description = "Unable to fetch list "
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	for _, account := range accounts {
		var respAccount UserResponseData
		respAccount.mapFromModel(account)
		resp.Data = append(resp.Data, respAccount.Data)
	}

	pageSize := 10
	resp.Meta.TotalPages = (totalItems / pageSize) + 1

	return c.JSON(http.StatusOK, resp)
}

// listAccounts is a handler for listing accounts in paginated format
func calculatePrice(c echo.Context) error {
	var errResp ErrorResponseData
	var resp AccountListResponseData

	fromDateTime, err := strconv.Atoi(c.Param("fromDateTime"))

	if (err != nil) || (fromDateTime <= 0) {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value   in query parameter fromDateTime"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}
	toDateTime, err := strconv.Atoi(c.Param("toDateTime"))

	if (err != nil) || (toDateTime <= 0) || toDateTime < fromDateTime {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value "
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	totalItems, accounts, err := storage.GetAccountList(fromDateTime, toDateTime)

	if err != nil {
		errResp.Data.Code = "error"
		errResp.Data.Description = "Unable to fetch list "
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	for _, account := range accounts {
		var respAccount UserResponseData
		respAccount.mapFromModel(account)
		resp.Data = append(resp.Data, respAccount.Data)
	}

	pageSize := 10
	resp.Meta.TotalPages = (totalItems / pageSize) + 1

	return c.JSON(http.StatusOK, resp)
}

// listUserBookings is a handler for listing bookings of a user in paginated format
func listUserBookings(c echo.Context) error {
	var errResp ErrorResponseData
	var resp AccountListResponseData

	id := strings.TrimSpace(c.Param("id"))
	if len(id) == 0 {
		errResp.Data.Code = "invalid_param_error"
		errResp.Data.Description = "Value for account id not set in request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	//change here
	totalItems, bookings, err := storage.GetBookingList(fromDateTime, toDateTime)

	if err != nil {
		errResp.Data.Code = "error"
		errResp.Data.Description = "Unable to fetch list "
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	for _, account := range bookings {
		var respAccount UserResponseData
		respAccount.mapFromModel(account)
		resp.Data = append(resp.Data, respAccount.Data)
	}

	pageSize := 10
	resp.Meta.TotalPages = (totalItems / pageSize) + 1

	return c.JSON(http.StatusOK, resp)
}
func listCarBookings(c echo.Context) error {
	var errResp ErrorResponseData
	var resp AccountListResponseData

	id := strings.TrimSpace(c.Param("id"))
	if len(id) == 0 {
		errResp.Data.Code = "invalid_param_error"
		errResp.Data.Description = "Value for account id not set in request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	//change here
	totalItems, accounts, err := storage.GetAccountList(fromDateTime, toDateTime)

	if err != nil {
		errResp.Data.Code = "error"
		errResp.Data.Description = "Unable to fetch list "
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	for _, account := range accounts {
		var respAccount UserResponseData
		respAccount.mapFromModel(account)
		resp.Data = append(resp.Data, respAccount.Data)
	}

	pageSize := 10
	resp.Meta.TotalPages = (totalItems / pageSize) + 1

	return c.JSON(http.StatusOK, resp)
}

//BookCar hnadler function
func bookCar(c echo.Context) error {
	var errResp ErrorResponseData
	var resp AccountListResponseData

	fromDateTime, err := strconv.Atoi(c.Param("fromDateTime"))

	if (err != nil) || (fromDateTime <= 0) {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value   in query parameter fromDateTime"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}
	toDateTime, err := strconv.Atoi(c.Param("toDateTime"))

	if (err != nil) || (toDateTime <= 0) || toDateTime < fromDateTime {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value "
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}
	user_id, err := strconv.Atoi(c.Param("userId"))

	if (err != nil) || (user_id <= 0) {
		errResp.Data.Code = "invalid_parameter_error"
		errResp.Data.Description = "Invalid value   in query parameter fromDateTime"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	//Change here

	totalItems, accounts, err := storage.GetAccountList(fromDateTime, toDateTime)

	if err != nil {
		errResp.Data.Code = "error"
		errResp.Data.Description = "Unable to fetch list "
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	for _, account := range accounts {
		var respAccount UserResponseData
		respAccount.mapFromModel(account)
		resp.Data = append(resp.Data, respAccount.Data)
	}

	pageSize := 10
	resp.Meta.TotalPages = (totalItems / pageSize) + 1

	return c.JSON(http.StatusOK, resp)
}

// deleteAccount is a handler function for deleting an account based on id
func deleteAccount(c echo.Context) error {
	var errResp ErrorResponseData

	id := strings.TrimSpace(c.Param("id"))
	if len(id) == 0 {
		errResp.Data.Code = "invalid_param_error"
		errResp.Data.Description = "Value for account id not set in request"
		errResp.Data.Status = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	noRecords, err := storage.DeleteAccount(id)

	if err != nil {
		errResp.Data.Code = "delete_account_error"
		errResp.Data.Description = "Unable to delete account details"
		errResp.Data.Status = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	if noRecords == 0 {
		errResp.Data.Code = "no_account_found"
		errResp.Data.Description = "No account with id " + id + " exists"
		errResp.Data.Status = strconv.Itoa(http.StatusNotFound)
		return c.JSON(http.StatusNotFound, errResp)
	}

	return c.JSON(http.StatusNoContent, nil)
}
