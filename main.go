package main

import (
	_ "bluebell/docs"
	"bluebell/router"
	"log"
)

//	@title			Bluebell
//	@version		0.4.7
//	@description	This is a sample bbs server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:5912
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	defer closure()
	var err error
	if err = initialize(); err != nil {
		log.Fatalln("infrastructure init failed: ", err)
	}
	if err = router.SetupRouter(); err != nil {
		log.Fatalln("server start failed: ", err)
	}
	return
}
