package erratum

func openWithRetry(o ResourceOpener) (Resource, error) {
	resource, err := o()
	if err == nil {
		return resource, nil
	}
	if _, ok := err.(TransientError); ok {
		return openWithRetry(o)
	}
	return nil, err
}

func Use(o ResourceOpener, input string) (err error) {
	resource, err := openWithRetry(o)
	if err != nil {
		return err
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
			if frobError, ok := r.(FrobError); ok {
				resource.Defrob(frobError.defrobTag)
				err = frobError.inner
			}
		}
	}()
	resource.Frob(input)
	return err
}
