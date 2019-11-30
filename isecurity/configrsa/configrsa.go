package isec_conf_rsa

import (
	"github.com/rhernandez-itemsoft/ihelpers/icommon"
	isecuritystt "github.com/rhernandez-itemsoft/ihelpers/isecurity/structs"
)

//DefaultConfig regresa la configuración por default
func DefaultConfig() isecuritystt.Config {
	return isecuritystt.Config{
		//PrivKeyPath - Llave pública
		PrivKeyPath: icommon.AppPath() + "resources/keys/private_key.rsa",

		//PubKeyPath - Llave privada
		PubKeyPath: icommon.AppPath() + "resources/keys/public_key.rsa.pub",
	}
}
