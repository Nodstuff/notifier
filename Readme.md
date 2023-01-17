# Notifier

This is a simple project to allow basic signalling in Go using only the standard library with zero external dependencies.

The intention for this library is to allow processes to await an async operation that they have no control or awareness of.
The channels returned from `Subscribe` will allow the caller to block until a signal is sent via the `Notify` method.

The `Notifier` itself is meant to be injected into processes that may have an async element which can then `Notify` when a certain condition is met.

Example:

```go
licenceValidationNotifier := NewNotifier("licence-notifier")

go doSomethingAfterSignal(licenceValidationNotifier)

svc := SomeAsyncService(licenceValidationNotifier)

svc.DoAsyncThing()
....
	
	
func (s *SomeAsyncService)DoAsyncThing() {
	res := externalResource.WaitForSomething()
	...
	// Signal to subscribers that the async operation is complete
	s.notifier.Notify()
}


func doSomethingAfterSignal(notifier *Notifier){
	// block until signal
	<-notifier.Subscribe("doSomethingAfterSignal-subscriber")
}
```