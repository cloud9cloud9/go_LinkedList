package main

import (
	"testing"

	"linkedlist/linked"

	"github.com/stretchr/testify/require"
)

func TestLinkedAddFirst(t *testing.T) {
	list := &linked.LinkedList[int]{}
	list.AddFirst(1)
	list.AddFirst(2)

	if list.Head.Value != 2 {
		t.Errorf("expected 2, got %v", list.Head.Value)
	}
	if list.Head.Next.Value != 1 {
		t.Errorf("expected 1, got %v", list.Head.Next.Value)
	}
	require.Equal(t, "{2,1}", list.String())
}

func TestLinkedAddLast(t *testing.T) {
	list := &linked.LinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)

	if list.Head.Value != 1 {
		t.Errorf("expected 1, got %v", list.Head.Value)
	}
	if list.Head.Next.Value != 2 {
		t.Errorf("expected 2, got %v", list.Head.Next.Value)
	}

	require.Equal(t, "{1,2}", list.String())
}

func TestRemoveFromList(t *testing.T) {
	list := &linked.LinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)
	list.Remove(1)
	if list.Head.Value != 2 {
		t.Errorf("expected 2, got %v", list.Head.Value)
	}

	require.Equal(t, "{2}", list.String())
}
