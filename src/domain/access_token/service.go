package access_token


import(
	"github.com/ArturoAguilar1/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}