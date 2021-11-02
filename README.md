# Nostress Errors Package

Usage example:

```golang
import "github.com/nostressdev/nerrors"

err := nerrors.Internal.New("my errors")

err := nerrors.BadRequest.Wrap(originalErr, "wrapped error:")

switch GetType(err) {
case nerrors.BadRequest:
	...
case nerrors.Internal:
	...
}

errors.Is(originalErr, nerrors.GetError(err)) == True
```