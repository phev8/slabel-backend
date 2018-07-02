package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ValidateToken checks if API key correct
func ValidateToken(token string) bool {
	const APIkey = "dicjvE543!-.,sdf54///vdfsdf45"
	if token != APIkey {
		return false
	}
	return true
}

// KeyRequired is a middleware function to check if API key present and valid
func KeyRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		req := context.Request
		// Get token string from url or header field
		var token string
		tokens, ok := req.Header["Api"]

		if ok && len(tokens) >= 1 {
			token = tokens[0]
		} else if len(req.FormValue("token")) > 0 {
			token = req.FormValue("token")
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"msg": "No API key found."})
			context.Abort()
			return
		}

		// Parse and validate token
		ok = ValidateToken(token)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{"msg": "Wrong API key."})
			context.Abort()
			return
		}

		context.Next()
	}
}

// TestAPI for test purposes
func TestAPI(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// CreateLabelSetHandl for creating a new labelset
func CreateLabelSetHandl(context *gin.Context) {
	params := LabelSet{}

	if err := context.BindJSON(&params); err != nil {
		//log.Println(err.(validator.ValidationErrors));
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	labelset, err := CreateNewLabelSet(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"labelset": labelset})
}

// UpdateLabelSetHandl for updating a labelset
func UpdateLabelSetHandl(context *gin.Context) {
	params := LabelSet{}

	if err := context.BindJSON(&params); err != nil {
		//log.Println(err.(validator.ValidationErrors));
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	labelset, err := UpdateLabelSet(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"labelset": labelset})
}

// GetLabelSetsHandl to retrieve all Labelsets
func GetLabelSetsHandl(context *gin.Context) {
	labelsets, err := GetLabelSets()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"labelsets": labelsets})
}

// DeleteLabelSetHandl is used to Delete Labelset and all of its labels
func DeleteLabelSetHandl(context *gin.Context) {
	labelsetID, ok := context.GetQuery("id")
	parsedLabelsetID, err := strconv.ParseUint(labelsetID, 10, 64)
	if !ok || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "You should provide a correct label set id."})
		return
	}

	if err := DeleteLabelSet(uint(parsedLabelsetID)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "label set deleted successfully"})
}

// CreateLabelTemplateHandl for add new label template (to labelset and parent)
func CreateLabelTemplateHandl(context *gin.Context) {
	params := LabelTemplate{}

	if err := context.BindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	labelTemplate, err := CreateNewLabelTemplate(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"label_template": labelTemplate})
}

// UpdateLabelTemplateHandl for updating a label template
func UpdateLabelTemplateHandl(context *gin.Context) {
	params := LabelTemplate{}

	if err := context.BindJSON(&params); err != nil {
		//log.Println(err.(validator.ValidationErrors));
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	label, err := UpdateLabelTemplate(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"label_template": label})
}

// DeleteLabelTemplateHandl to Remove label and its sublabels from labelset
func DeleteLabelTemplateHandl(context *gin.Context) {
	ltID, ok := context.GetQuery("id")
	parsedLabelTemplateID, err := strconv.ParseUint(ltID, 10, 64)
	if !ok || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "You should provide a correct label set id."})
		return
	}

	if err := DeleteLabelTemplate(uint(parsedLabelTemplateID)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "subtree deleted"})
}

// GetSingleLabelSetHandl to get a populated labelset
func GetSingleLabelSetHandl(context *gin.Context) {
	lsID, ok := context.GetQuery("id")
	parsedLsID, err := strconv.ParseUint(lsID, 10, 64)
	if !ok || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "You should provide a correct customer id."})
		return
	}

	ls, err := GetSingleLabelSet(uint(parsedLsID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"labelset": ls})
}

// CreateSessionHandl for creating a new session
func CreateSessionHandl(context *gin.Context) {
	params := Session{}

	if err := context.BindJSON(&params); err != nil {
		//log.Println(err.(validator.ValidationErrors));
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	session, err := CreateNewSession(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"session": session})
}

// UpdateSessionHandl for updating a session
func UpdateSessionHandl(context *gin.Context) {
	params := Session{}

	if err := context.BindJSON(&params); err != nil {
		//log.Println(err.(validator.ValidationErrors));
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	session, err := UpdateSession(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"session": session})
}

// GetSessionsHandl to retrieve all Sessions
func GetSessionsHandl(context *gin.Context) {
	sessions, err := GetSessions()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"sessions": sessions})
}

// DeleteSessionHandl to Delete session and all of its labels
func DeleteSessionHandl(context *gin.Context) {
	sID, ok := context.GetQuery("id")
	parsedSID, err := strconv.ParseUint(sID, 10, 64)
	if !ok || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "You should provide a correct session id."})
		return
	}

	if err := DeleteSession(uint(parsedSID)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "session deleted successfully"})
}

// GetSingleSessionHandl to get a populated session instance by id
func GetSingleSessionHandl(context *gin.Context) {
	sID, ok := context.GetQuery("id")
	parsedSID, err := strconv.ParseUint(sID, 10, 64)
	if !ok || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "You should provide a correct session id."})
		return
	}

	s, err := GetSingleSession(uint(parsedSID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"session": s})
}
