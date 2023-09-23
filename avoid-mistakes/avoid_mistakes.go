package avoid_mistakes

import "errors"

func correctWayOfIfElse(s string) (bool, error) {

	/*
		When the "if" block returns we should omit the else block
	*/
	if len(s) <= 3 {
		return true, errors.New("string must have more than 3 characters")
	}

	if s == "" {
		return false, errors.New("empty string")
	}

	return true, nil
}
