// Package register registers all services
package register

import (

	// register services.
	_ "go.viam.com/rdk/services/baseremotecontrol"
	_ "go.viam.com/rdk/services/framesystem"
	_ "go.viam.com/rdk/services/navigation"
	_ "go.viam.com/rdk/services/objectmanipulation"
	_ "go.viam.com/rdk/services/web"
)
