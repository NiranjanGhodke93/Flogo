package OracleConn

import (
    
    "database/sql"
    _ "github.com/mattn/go-oci8"
	"github.com/TIBCOSoftware/flogo-lib/logger"	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"

	
)

// log is the default package logger which we'll use to log
	var log = logger.GetLogger("activity-TestConnection")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	
	sqlurl := context.GetInput("SqlURL").(string)
	log.Debugf("sqlurl");
	
	
    db, err := sql.Open("oci8",sqlurl)
	log.Debugf("Connection Successfull");
    if err != nil {
         log.Debugf("Connection Refused");
        return
    }
    defer db.Close()
    
    
    if err = db.Ping(); err != nil {
     
	  log.Debugf("Error connecting to the database: [%s]",err);
        return
    }
	
	context.SetOutput("output","")
	return true, nil
}

