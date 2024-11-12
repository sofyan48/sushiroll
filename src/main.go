// Package main
package main

import (
	"github.com/sofyan48/sushiroll/src/cmd"
)

// @title           sushiroll API
// @version         1.0
// @description     sushiroll API server for interact the app.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Sofyan Saputra
// @contact.url    http://kiriminaja.id
// @contact.email  sofyan.saputra@kiriminaja.co.id

// @license.name  Kiriminaja License 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					add your toke with this format | bearer (access_token) |
func main() {
	cmd.Start()
}
