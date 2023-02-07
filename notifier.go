package notifier

import "sync"

// Notifiers is a collection of pointers to Notifier structs, they are stored using their identifier as their key.
type Notifiers map[string]*Notifier

// AddNotifier provides the ability to add a Notifier to the Notifiers collection, it is added using its identifier.
func (n Notifiers) AddNotifier(notifier *Notifier) {
	n[notifier.identifier] = notifier
}

// GetNotifier provides the ability to get a Notifier from the Notifiers collection using its identifier.
func (n Notifiers) GetNotifier(identifier string) *Notifier {
	return n[identifier]
}

// RemoveNotifier provides the ability to remove a Notifier from the Notifiers collection using its identifier.
func (n Notifiers) RemoveNotifier(identifier string) {
	delete(n, identifier)
}

// Notifier allows for the broadcasting of a signal to subscribers to allow them to react to async actions
// they may not be in control of or be aware of. It can be specifically identified using its identifier field.
type Notifier struct {
	sync.RWMutex
	identifier string
	subs       map[string]chan struct{}
}

// NewNotifier Returns a pointer to an instance of Notifier created with the passed identifier
func NewNotifier(identifier string) *Notifier {
	return &Notifier{
		identifier: identifier,
		subs:       map[string]chan struct{}{},
	}
}

// Subscribe provides the ability for the caller to subscribe to the specific notifier and uses the passed identifier to identify
// the caller. It returns a channel for the caller to block on to await the signal.
func (n *Notifier) Subscribe(identifier string) <-chan struct{} {
	n.Lock()
	defer n.Unlock()
	if _, ok := n.subs[identifier]; ok {
		return n.subs[identifier]
	}
	n.subs[identifier] = make(chan struct{}, 3)
	return n.subs[identifier]
}

// Unsubscribe provides the caller with the ability to remove itself from the notifier should they need to or need to clean up.
func (n *Notifier) Unsubscribe(identifier string) {
	n.Lock()
	defer n.Unlock()
	delete(n.subs, identifier)
}

// Notify provides the ability to fire a signal to each of the subscribers to signal them to unblock themselves or take some other action
// based on their needs once they receive the signal.
func (n *Notifier) Notify() {
	n.RLock()
	defer n.RUnlock()
	for k := range n.subs {
		n.subs[k] <- struct{}{}
	}
}
