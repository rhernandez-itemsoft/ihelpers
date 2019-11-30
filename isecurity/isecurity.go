package isecurity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strings"

	"log"
	"time"

	"github.com/rhernandez-itemsoft/ihelpers/ijwt"
	"github.com/rhernandez-itemsoft/ihelpers/iresponse"
	isecuritystt "github.com/rhernandez-itemsoft/ihelpers/isecurity/structs"

	"github.com/kataras/iris/v12"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

var _iresponse *iresponse.Definition
var _ijwt ijwt.Definition

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto

}

//New Crea una nueva instancia de  Definition
func New(ctx iris.Context) *Definition {
	_iresponse = iresponse.New(ctx)
	_ijwt = ijwt.New(ctx)
	_ijwt.LoadKeys()
	return &Definition{
		Ctx: ctx,
	}
}

//NewToken - Handler encargado de validar el login con username & password
func (def *Definition) NewToken(data interface{}) (isecuritystt.Token, error) {
	var token isecuritystt.Token
	var err error
	if _ijwt.Conf.PrivateKey != nil {
		/*
			//create a rsa 256 signer
			signer := jwt.New(jwt.GetSigningMethod("RS256"))

			//set claims
			claims := make(jwt.MapClaims)
			//claims["iss"] = "admin"
			claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
			claims["Data"] = data
			signer.Claims = claims
		*/
		signer := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"exp":  time.Now().Add((time.Hour * 24) * 365).Unix(),
			"Data": data,
			//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		tokenString, err := signer.SignedString(_ijwt.Conf.PrivateKey)
		if err == nil {
			token = isecuritystt.Token{
				Token: tokenString,
			}
		}
	}

	//retorna IRESPONSE
	return token, err
}

//EncriptSha512 encripta la contraseña en sha512
func (def *Definition) EncriptSha512(data string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(data))) // return [64]byte
	//sha256.Sum256([]byte(password))
}

//EncriptSha256 encripta la contraseña en sha256
func (def *Definition) EncriptSha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Println(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}
	return plaintext
}

//JWTMiddleware permite asegurar nuestra api con JWT
//regresa:
//Claims del JWT,
//Error code
//Message
func (def *Definition) JWTMiddleware(ctx iris.Context) {
	//def.Ctx = ctx

	//fmt.Println("1.- ")

	//revisa que se reciba un token
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		_iresponse.JSONResponse(iris.StatusUnauthorized, nil, "missing_token")
		return
	}
	//fmt.Println("2.- ")
	//revisa que el token tenga formato correcto
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		_iresponse.JSONResponse(iris.StatusUnauthorized, nil, "error_token")
		return
	}
	//fmt.Println("3.- ")
	// valida el token
	tokenString = tokenArr[1]
	token, err := jwt.ParseWithClaims(tokenString, &isecuritystt.IClaim{}, func(token *jwt.Token) (interface{}, error) {
		return _ijwt.Conf.PublicKey, nil
		//return []byte("AllYourBase"), nil  //si no hubiera encriptación RSA, sería algo así
	})
	//fmt.Println("4.- ")
	if err != nil || !token.Valid {
		_iresponse.JSONResponse(iris.StatusUnauthorized, nil, "error_token")
		return
	}
	//No podemos generar una respuesta JSON porque se generaría doble response
	//y obtendríamos error en el cliente
	//El token es válido y retorna una respuesta válida
	//_iresponse.JSON(iris.StatusOK, token.Claims, "success")

	ctx.Next()
	return
}
