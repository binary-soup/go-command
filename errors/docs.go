// Package errors is a replacement/improvement on the standard errors package.
//
// Use the Chain function to add extra context before returning an error:
//
//	err := SomeFunction()
//	if err != nil {
//		return errors.Chain(err, "error running some function")
//	}
//
// Use the Errors type to track multiple errors at once:
//
//	func Foo() error {
//		errs := errors.Errors{}
//
//		err := CheckSomething()
//		if err != nil {
//			errs.Add(err)
//		}
//
//		err = CheckSomethingElse()
//		if err != nil {
//			errs.Add(err)
//		}
//
//		return errs.Collapse(",")
//	}
package errors
