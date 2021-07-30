package adret

import (
	"fmt"
)

// Use to config the information about remote ubac listener to avoid specifying
// in CLi all the time
func ConfigRemote(hostname string, port string) {
	urlUbac := "http://" + hostname + ":" + port + "/"
	export := "export REMOTE_UBAC_URL='" + urlUbac + "'"
	//os.Setenv("REMOTE_UBAC_URL", urlUbac)
	fmt.Println(export)
}
