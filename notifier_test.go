package notifier

import (
	"testing"
)

func TestNotifier_Subscribe_Notify(t *testing.T) {
	testCases := []struct {
		name   string
		events int
		want   int
	}{
		{"Test 3 signal received", 3, 3},
		{"Test 1 signal received", 1, 1},
		{"Test 99 signal received", 99, 99},
		{"Test 12 signal received", 12, 12},
		{"Test 7 signal received", 7, 7},
		{"Test 0 signal received", 0, 0},
	}
	n := NewNotifier("test-subscribe")
	c := n.Subscribe("test-sub1")
	c2 := n.Subscribe("test-sub2")
	c3 := n.Subscribe("test-sub3")
	for _, tc := range testCases {
		got1 := 0
		got2 := 0
		got3 := 0
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.want; i++ {
				go n.Notify()
				select {
				case <-c:
					got1++
				}

				select {
				case <-c2:
					got2++
				}

				select {
				case <-c3:
					got3++
				}
			}

			if got1 != tc.want {
				t.Errorf("unexpected count, got %d, wanted %d", got1, tc.want)
			}

			if got2 != tc.want {
				t.Errorf("unexpected count, got %d, wanted %d", got2, tc.want)
			}

			if got3 != tc.want {
				t.Errorf("unexpected count, got %d, wanted %d", got3, tc.want)
			}
		})
	}
}

func TestNotifier_Unsubscribe(t *testing.T) {
	testCases := []struct {
		name   string
		events int
		want   int
	}{
		{"Test 99 signal received", 99, 99},
		{"Test 0 signal received", 10, 0},
	}
	n := NewNotifier("test-unsubscribe")
	c := n.Subscribe("test-unsub")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := 0
			for i := 0; i < tc.want; i++ {
				n.Notify()
				select {
				case <-c:
					got++
				}
			}
			// assert
			if got != tc.want {
				t.Errorf("unexpected count, got %d, wanted %d", got, tc.want)
			}
		})
		n.Unsubscribe("test-unsub")
	}
}
