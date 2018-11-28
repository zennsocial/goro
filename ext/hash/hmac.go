package hash

import (
	"crypto/hmac"
	"encoding/hex"
	"fmt"

	"github.com/MagicalTux/goro/core"
	"github.com/MagicalTux/goro/core/phpv"
)

//> func string hash_hmac ( string $algo , string $data , string $key [, bool $raw_output = FALSE ] )
func fncHashHmac(ctx phpv.Context, args []*phpv.ZVal) (*phpv.ZVal, error) {
	var algo phpv.ZString
	var data phpv.ZString
	var key phpv.ZString
	var raw *phpv.ZBool

	_, err := core.Expand(ctx, args, &algo, &data, &key, &raw)
	if err != nil {
		return nil, err
	}

	algN, ok := algos[algo.ToLower()]
	if !ok {
		return nil, fmt.Errorf("Unknown hashing algorithm: %s", algo)
	}

	a := hmac.New(algN, []byte(key))
	_, err = a.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	r := a.Sum(nil)

	if raw != nil && *raw {
		// return as raw
		return phpv.ZString(r).ZVal(), nil
	}

	// convert to hex
	return phpv.ZString(hex.EncodeToString(r)).ZVal(), nil
}
