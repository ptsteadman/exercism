package erratum

import "fmt"

// openWithRetry attempts to open resource o. If an error of type
// TransientError is returned, it continually retries until it can
// successfully return the resource or another type of error is returned
func openWithRetry(o ResourceOpener) (Resource, error) {
	var resource Resource
	var err error
	for err == nil {
		resource, err = o()
		if err == nil {
			return resource, nil
		}
		if _, ok := err.(TransientError); ok {
			// if the error is a TransientError, try again
			err = nil
		}
	}
	return nil, err
}

// Use opens a resource, Frobs a string to it, and ensures that the
// resource is appropriately closed and defrobbed if necessary
func Use(o ResourceOpener, input string) (err error) {
	resource, err := openWithRetry(o)
	if err != nil {
		return err
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e.inner
			case error:
				err = e
			default:
				err = fmt.Errorf("%s", r)
			}
		}
	}()
	resource.Frob(input)
	return err
}
