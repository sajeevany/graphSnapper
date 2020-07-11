package credentials

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	as "github.com/sajeevany/graph-snapper/internal/db/aerospike"
	"github.com/sajeevany/graph-snapper/internal/db/aerospike/record"
	"github.com/sirupsen/logrus"
	"net/http"
)

const PostCredentialsEndpoint = "/{accountID}"

//@Summary Add credentials to an account
//@Description Non-authenticated endpoint that adds grafana and confluence-server users to an account. Assumes entries are pre-validated
//@Produce json
//@Param account body AddCredentialsV1 true "Add credentials"
//@Success 200 {object} AddedCredentialsV1
//@Fail 404 {object} gin.H
//@Fail 500 {object} gin.H
//@Router /credentials [put]
//@Tags credentials
func PutCredentialsV1(logger *logrus.Logger, aeroClient *as.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Bind add credentials object
		var addReq AddCredentialsV1
		if bErr := ctx.BindJSON(&addReq); bErr != nil {
			msg := fmt.Sprintf("Unable to bind request body to AddCredentialsV1 object %v", bErr)
			logger.Errorf(msg)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		//Validate account. Returns account key since it validates if the record exists
		vErr, returnCode, actKey, actKeyExists := validateRequest(logger, aeroClient, addReq)
		if vErr != nil {
			if actKeyExists {
				logger.WithFields(addReq.GetFields()).Errorf("Input credentials are invalid <%v>", vErr)
				ctx.JSON(returnCode, gin.H{
					"humanReadableError": "Input credentials variables are invalid. Host, user, password, apikey must be non empty. Port must be within 0 and 65535",
					"error":              vErr.Error(),
				})
				return
			} else {
				logger.WithFields(addReq.GetFields()).Errorf("Account id doesn't exist <%v>. err <%v>", addReq.AccountID, vErr)
				ctx.JSON(returnCode, gin.H{
					"humanReadableError": fmt.Sprintf("No account exists with ID %v", addReq.AccountID),
					"error":              vErr.Error(),
				})
				return
			}
		}

		//Check if any users have been provided. If not, skip and return 200
		if addReq.HasNoUsers() {
			logger.Debug("No credentials were provided for the account id <%v>. returning bad request", addReq.AccountID)
			ctx.Status(http.StatusBadRequest)
			return
		}

		rec, aErr := setAccountUsers(logger, aeroClient, addReq, actKey)
		if aErr != nil {
			hMsg := "Internal error when adding users to Aerospike data store"
			logger.WithFields(addReq.GetFields()).Error(hMsg, aErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"humanReadableError": hMsg,
				"error":              aErr.Error()})
			return
		}
		ctx.JSON(http.StatusOK, rec.ToRecordViewV1())
	}
}

//Checks if input is in acceptable and a record exists with the specified key. Returns a non-zero return code if an error is present. Returns no error and a statusOk(200).
func validateRequest(logger *logrus.Logger, aeroClient *as.ASClient, addReq AddCredentialsV1) (error, int, *aerospike.Key, bool) {

	//Validate the account info. Checks if record exists with the ID
	returnCode, aErr, actKey, actKeyExists := validateAcctID(logger, aeroClient, addReq.AccountID)
	if aErr != nil || !actKeyExists {
		//if an error occurred or the key doesn't exist return with the http code
		return aErr, returnCode, nil, actKeyExists
	}

	//Validate grafana users
	for _, gUser := range addReq.GrafanaReadUsers {
		if !gUser.IsValid() {
			logger.WithFields(gUser.GetFields()).Errorf("Grafana user has invalid attributes")
			return fmt.Errorf("grafana user <%#v> is invalid", gUser), http.StatusBadRequest, nil, actKeyExists
		}
	}

	//Validate confluence users
	for _, csUser := range addReq.ConfluenceServerUsers {
		if !csUser.IsValid() {
			logger.WithFields(csUser.GetFields()).Errorf("Confluence user has invalid attributes")
			return fmt.Errorf("confluence user <%#v> is invalid", csUser), http.StatusBadRequest, nil, actKeyExists
		}
	}

	logger.Debugf("Validate request passed for account id <%v>", addReq.AccountID)
	return nil, http.StatusOK, actKey, actKeyExists
}

//Returns error if invalid. int value is the http return code to use
func validateAcctID(logger *logrus.Logger, aeroClient *as.ASClient, id string) (int, error, *aerospike.Key, bool) {

	if id == "" {
		return http.StatusBadRequest, fmt.Errorf("account ID is empty and must be defined"), nil, false
	}

	//check if account exists
	actExists, actKey, rErr := aeroClient.GetReader().KeyExists(id)
	if rErr != nil {
		logger.Errorf("Error when reading from aerospike to check if key exists <%v>", rErr)
		return http.StatusInternalServerError, rErr, actKey, actExists
	}
	if !actExists {
		msg := fmt.Sprintf("Key <%v> doesn't exist", id)
		logger.Debug(msg)
		return http.StatusNotFound, fmt.Errorf(msg), actKey, actExists
	}

	logger.Debugf("Valid account id provided <%v>", id)
	return http.StatusOK, nil, actKey, actExists
}

//setAccountUsers - adds specified users to the record at the specified account. Assumes that the record at the provided key has already been checked for existence
func setAccountUsers(logger *logrus.Logger, client *as.ASClient, req AddCredentialsV1, actKey *aerospike.Key) (record.Record, error) {

	logger.Debugf("Starting overwrite users to account with id <%v> operation", actKey.String())
	//Get the current record
	record, err := client.GetReader().ReadRecord(actKey)
	if err != nil {
		logger.Errorf("Failed to read record using key <%v>. err <%v>", actKey.String(), err)
		return nil, err
	}

	//Update the local record copy and overwrite it in the db
	logger.Debugf("Record has been read for account with id <%v>. ", actKey.String())
	record.SetUserCredentialsV1(logger, req.GrafanaReadUsers, req.ConfluenceServerUsers)
	if wErr := client.GetWriter().WriteRecordWithASKey(actKey, record); wErr != nil{
		logger.Errorf("Error when writing record to db. err <%v>", wErr)
		return record, wErr
	}

	logger.Debugf("Record written for setAccountUsers pk <%v>", actKey.String())
	return record, nil
}
