package context

import (
	"github.com/loconsolutions/smalldocs/config"
	"labix.org/v2/mgo"
)

// Configuration object
var Config *config.Config

// Database session
var DBSession *mgo.Session
