package wrapper

import (
	"errors"

	errorkit "github.com/Epritka/gokit/errors"
	"github.com/go-pg/pg"
)

func PgWrap(err error) error {
	if errors.Is(err, pg.ErrNoRows) {
		return errorkit.NotFoundError()
	} else if errors.Is(err, pg.ErrMultiRows) {
		return errorkit.NotFoundError()
	}

	return nil
}
