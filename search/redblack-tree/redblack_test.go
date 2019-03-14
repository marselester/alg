package redblack

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	tree := abcTree()

	tt := []struct {
		name string
		key  string
		n    *node
		want *node
	}{
		{
			name: "blank tree",
			key:  "fizz",
			n:    nil,
			want: nil,
		},
		{
			name: "S to S",
			key:  "S",
			n:    tree.root,
			want: tree.root,
		},
		{
			name: "S to E",
			key:  "E",
			n:    tree.root,
			want: tree.root.left,
		},
		{
			name: "S to A",
			key:  "A",
			n:    tree.root,
			want: tree.root.left.left,
		},
		{
			name: "S to C",
			key:  "C",
			n:    tree.root,
			want: tree.root.left.left.right,
		},
		{
			name: "E to C",
			key:  "C",
			n:    tree.root.left,
			want: tree.root.left.left.right,
		},
		{
			name: "S to R",
			key:  "R",
			n:    tree.root,
			want: tree.root.left.right,
		},
		{
			name: "S to H",
			key:  "H",
			n:    tree.root,
			want: tree.root.left.right.left,
		},
		{
			name: "S to M",
			key:  "M",
			n:    tree.root,
			want: tree.root.left.right.left.right,
		},
		{
			name: "S to X",
			key:  "X",
			n:    tree.root,
			want: tree.root.right,
		},
		{
			name: "X to X",
			key:  "X",
			n:    tree.root.right,
			want: tree.root.right,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.key, tc.n)
			if got != tc.want {
				t.Errorf("search(%q, %+v) got %+v, want %+v", tc.key, tc.n, got, tc.want)
			}
		})
	}
}

func TestTree_Get(t *testing.T) {
	tree := abcTree()

	tt := []struct {
		key  string
		want []byte
	}{
		{"S", []byte("sea")},
		{"R", []byte("rock")},
		{"unknown", nil},
	}
	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			got := tree.Get(tc.key)
			if !bytes.Equal(got, tc.want) {
				t.Errorf("Get(%q) got %q, want %q", tc.key, got, tc.want)
			}
		})
	}

	blank := Tree{}
	key := "missing"
	got := blank.Get(key)
	if got != nil {
		t.Errorf("Get(%q) got %+v from blank tree, want nil", key, got)
	}
}

func TestTree_Set(t *testing.T) {
	tree := &Tree{}

	key := "name"
	value := []byte("Bob")
	tree.Set(key, value)
	if tree.root == nil {
		t.Fatalf("Set(%q, %q) root is nil", key, value)
	}
	if tree.root.key != key {
		t.Errorf("Set(%q, %q) wrong root key, got %q", key, value, tree.root.key)
	}
	if !bytes.Equal(tree.root.value, value) {
		t.Errorf("Set(%q, %q) wrong root value, got %q", key, value, tree.root.value)
	}

	key = "planet"
	value = []byte("Earth")
	tree.Set(key, value)
	if tree.root == nil {
		t.Fatalf("Set(%q, %q) root is nil", key, value)
	}
	if tree.root.key != key {
		t.Errorf("Set(%q, %q) wrong root key, got %q", key, value, tree.root.key)
	}
	if !bytes.Equal(tree.root.value, value) {
		t.Errorf("Set(%q, %q) wrong root value, got %q", key, value, tree.root.value)
	}

	key = "name"
	value = []byte("Bob")
	if tree.root.left == nil {
		t.Fatalf("Set(%q, %q) left child is nil", key, value)
	}
	if tree.root.left.key != key {
		t.Errorf("Set(%q, %q) wrong left child key, got %q", key, value, tree.root.left.key)
	}
	if !bytes.Equal(tree.root.left.value, value) {
		t.Errorf("Set(%q, %q) wrong left child value, got %q", key, value, tree.root.left.value)
	}
}

func TestTree_Set_inserts(t *testing.T) {
	tree := &Tree{}

	tree.Set("s", nil)
	want := []string{"s(,)"}
	got := shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(s) got %q, want %q", got, want)
	}

	tree.Set("e", nil)
	want = []string{"E(,)", "s(E,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(e) got %q, want %q", got, want)
	}

	tree.Set("a", nil)
	want = []string{"a(,)", "e(a,s)", "s(,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(a) got %q, want %q", got, want)
	}

	tree.Set("r", nil)
	want = []string{"a(,)", "e(a,s)", "R(,)", "s(R,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(r) got %q, want %q", got, want)
	}

	tree.Set("c", nil)
	want = []string{"A(,)", "c(A,)", "e(c,s)", "R(,)", "s(R,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(c) got %q, want %q", got, want)
	}

	tree.Set("h", nil)
	want = []string{"A(,)", "c(A,)", "E(c,h)", "h(,)", "r(E,s)", "s(,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(h) got %q, want %q", got, want)
	}

	tree.Set("x", nil)
	want = []string{"A(,)", "c(A,)", "E(c,h)", "h(,)", "r(E,x)", "S(,)", "x(S,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(x) got %q, want %q", got, want)
	}

	tree.Set("m", nil)
	want = []string{"A(,)", "c(A,)", "E(c,m)", "H(,)", "m(H,)", "r(E,x)", "S(,)", "x(S,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(m) got %q, want %q", got, want)
	}

	tree.Set("p", nil)
	want = []string{"A(,)", "c(A,)", "e(c,h)", "h(,)", "m(e,r)", "p(,)", "r(p,x)", "S(,)", "x(S,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(p) got %q, want %q", got, want)
	}

	tree.Set("l", nil)
	want = []string{"A(,)", "c(A,)", "e(c,l)", "H(,)", "l(H,)", "m(e,r)", "p(,)", "r(p,x)", "S(,)", "x(S,)"}
	got = shape(nil, tree.root)
	if !equal(got, want) {
		t.Errorf("Set(l) got %q, want %q", got, want)
	}
}

// shape returns all keys in order. Keys with red links are upper case,
// black links are lower case.
func shape(kk []string, n *node) []string {
	if n == nil {
		return kk
	}

	coloredKey := func(n *node) string {
		switch {
		case n == nil:
			return ""
		case n.isRed():
			return strings.ToUpper(n.key)
		default:
			return strings.ToLower(n.key)
		}
	}
	kk = shape(kk, n.left)
	kk = append(kk, fmt.Sprintf("%s(%s,%s)", coloredKey(n), coloredKey(n.left), coloredKey(n.right)))
	kk = shape(kk, n.right)
	return kk
}

func TestRotateLeft(t *testing.T) {
	// Before: B has a right child X.
	h := rightLeaning()
	// After: X has a left child B.
	want := leftLeaning()

	got := rotateLeft(h)
	gotkeys := keys(nil, got)
	wantkeys := keys(nil, want)
	if !equal(gotkeys, wantkeys) {
		t.Errorf("rotateLeft() got %v, want %v", gotkeys, wantkeys)
	}

	if got == nil || got.left.color != red {
		t.Errorf("rotateLeft() left link must be red")
	}
	if got == nil || got.right.color != black {
		t.Errorf("rotateLeft() right link must be black")
	}
}

func TestRotateRight(t *testing.T) {
	// Before: X has a left child B.
	h := leftLeaning()
	// After: B has a right child X.
	want := rightLeaning()

	got := rotateRight(h)
	gotkeys := keys(nil, got)
	wantkeys := keys(nil, want)
	if !equal(gotkeys, wantkeys) {
		t.Errorf("rotateRight() got %v, want %v", gotkeys, wantkeys)
	}

	if got == nil || got.left.color != black {
		t.Errorf("rotateRight() left link must be black")
	}
	if got == nil || got.right.color != red {
		t.Errorf("rotateRight() right link must be red")
	}
}

func TestFlipColors(t *testing.T) {
	h := &node{
		key:   "E",
		color: black,
		left: &node{
			key:   "A",
			color: red,
		},
		right: &node{
			key:   "S",
			color: red,
		},
	}

	flipColors(h)
	if h.color != red {
		t.Errorf("flipColors() parent color must be red")
	}
	if h.left.color != black {
		t.Errorf("flipColors() left link must be black")
	}
	if h.right.color != black {
		t.Errorf("flipColors() right link must be black")
	}
}

func TestKeys(t *testing.T) {
	tree := abcTree()
	kk := keys(nil, tree.root)
	want := []string{"A", "C", "E", "H", "M", "R", "S", "X"}
	if !equal(kk, want) {
		t.Errorf("keys(nil, root) got %v, want %v", kk, want)
	}
}

func TestTree_Keys(t *testing.T) {
	tree := &Tree{}
	kk := tree.Keys()
	if kk != nil {
		t.Errorf("Keys() got %v, want nil", kk)
	}

	tree = abcTree()
	kk = tree.Keys()
	want := []string{"A", "C", "E", "H", "M", "R", "S", "X"}
	if !equal(kk, want) {
		t.Errorf("Keys() got %v, want %v", kk, want)
	}
}

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func abcTree() *Tree {
	r := node{
		key:   "S",
		value: []byte("sea"),
		left: &node{
			key:   "E",
			value: []byte("ear"),
			left: &node{
				key:   "A",
				value: []byte("apple"),
				right: &node{
					key:   "C",
					value: []byte("cat"),
				},
			},
			right: &node{
				key:   "R",
				value: []byte("rock"),
				left: &node{
					key:   "H",
					value: []byte("house"),
					right: &node{
						key:   "M",
						value: []byte("mouse"),
					},
				},
			},
		},
		right: &node{
			key: "X",
		},
	}
	return &Tree{root: &r}
}

func rightLeaning() *node {
	return &node{
		key:   "B",
		value: []byte("b"),
		color: black,
		// left child A is less than B.
		left: &node{
			key:   "A",
			value: []byte("a"),
		},
		// right child X is greater than B (right-leaning red link).
		right: &node{
			key:   "X",
			value: []byte("x"),
			color: red,
			// left child C is between B and X.
			left: &node{
				key:   "C",
				value: []byte("c"),
			},
			// right child Z is greater than X.
			right: &node{
				key:   "Z",
				value: []byte("z"),
			},
		},
	}
}

func leftLeaning() *node {
	return &node{
		key:   "X",
		value: []byte("x"),
		// left child B is less than X.
		left: &node{
			key:   "B",
			value: []byte("b"),
			// left child A is less than B.
			left: &node{
				key:   "A",
				value: []byte("a"),
			},
			// right child C is between B and X.
			right: &node{
				key:   "C",
				value: []byte("c"),
			},
		},
		// right child Z is greater than X.
		right: &node{
			key:   "Z",
			value: []byte("z"),
		},
	}
}
