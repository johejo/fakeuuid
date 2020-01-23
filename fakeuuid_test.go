package fakeuuid

import "testing"

func TestMock(t *testing.T) {
	got := New().String()
	if got == "" {
		t.Fatalf("should not empty: got=%s", got)
	}
	t.Log(got)

	const fake = "fake"
	unMock := Mock(fake)
	got = New().String()
	if got != fake {
		t.Fatalf("failed to mock: got=%s, fake=%s", got, fake)
	}
	t.Log(got)

	unMock()
	got = New().String()
	if got == fake {
		t.Fatalf("failed to unMock: got=%s, fake=%s", got, fake)
	}
	t.Log(got)
}
