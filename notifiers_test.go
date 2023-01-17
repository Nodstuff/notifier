package notifier

import "testing"

func TestNotifiers_Add_GetNotifier(t *testing.T) {
	testCases := []struct {
		name     string
		notifier *Notifier
	}{
		{
			"should contain notifier-1",
			NewNotifier("notifier-1"),
		},
		{
			"should contain notifier-2",
			NewNotifier("notifier-2"),
		},
		{
			"should contain notifier-3",
			NewNotifier("notifier-3"),
		},
	}
	notifiers := Notifiers{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			notifiers.AddNotifier(tc.notifier)
			nn := notifiers.GetNotifier(tc.notifier.identifier)
			if nn == nil {
				t.Fatalf("notifier should exist but %s didn't", tc.notifier.identifier)
			}
		})
	}
}

func TestNotifiers_RemoveNotifier(t *testing.T) {
	testCases := []struct {
		name     string
		notifier *Notifier
	}{
		{
			"should not contain notifier-1",
			NewNotifier("notifier-1"),
		},
		{
			"should not contain notifier-2",
			NewNotifier("notifier-2"),
		},
		{
			"should not contain notifier-3",
			NewNotifier("notifier-3"),
		},
	}
	notifiers := Notifiers{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			notifiers.AddNotifier(tc.notifier)
			nn := notifiers.GetNotifier(tc.notifier.identifier)
			if nn == nil {
				t.Fatalf("notifier should exist but %s didn't", tc.notifier.identifier)
			}
			notifiers.RemoveNotifier(tc.notifier.identifier)
			nn = notifiers.GetNotifier(tc.notifier.identifier)
			if nn != nil {
				t.Fatalf("notifier should exist but %s didn't", tc.notifier.identifier)
			}
		})
	}
}
