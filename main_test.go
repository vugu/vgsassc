package main_test

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

func TestVgsassc(t *testing.T) {

	tmpDir := t.TempDir()

	vgsasscExe := filepath.Join(tmpDir, "vgsassc")

	b, err := exec.Command("go", "build", "-o", vgsasscExe, ".").CombinedOutput()
	if err != nil {
		t.Logf("go build output:\n%s", b)
		t.Fatal(err)
	}

	t.Run("001-simple", func(t *testing.T) {
		outPath := filepath.Join(tmpDir, "001-out.css")
		mustRun(t, exec.Command(vgsasscExe, "-o", outPath, "testdata/001-simple.scss"))
		bstr := string(mustRead(t, outPath))
		restrmatch(t, "some comment here", bstr)
		//t.Logf("out: %s", b)
	})
	t.Run("001a-simple-min", func(t *testing.T) {
		outPath := filepath.Join(tmpDir, "001-out.min.css")
		mustRun(t, exec.Command(vgsasscExe, "-m", "-o", outPath, "testdata/001-simple.scss"))
		bstr := string(mustRead(t, outPath))
		restrmatch(t, `^\.test1\{color:green\}\.test1 \.test2\{color:red\}\s*$`, bstr)
	})
	t.Run("002-use", func(t *testing.T) {
		outPath := filepath.Join(tmpDir, "002-out.min.css")
		mustRun(t, exec.Command(vgsasscExe, "-m", "-o", outPath, "-I", "testdata", "testdata/002-use.scss"))
		bstr := string(mustRead(t, outPath))
		restrmatch(t, `^\.include-a\{color:purple\}`, bstr)
	})

}

func restrmatch(t *testing.T, re string, shouldMatch string) {
	t.Helper()
	r, err := regexp.Compile(re)
	if err != nil {
		t.Fatal(err)
	}
	if !r.MatchString(shouldMatch) {
		t.Errorf("regexp match failure for: %s\nshould have matched: %s", re, shouldMatch)
	}
}

func mustRead(t *testing.T, p string) []byte {
	t.Helper()
	b, err := ioutil.ReadFile(p)
	if err != nil {
		t.Fatalf("error while reading %q: %v", p, err)
	}
	return b
}

func mustRun(t *testing.T, cmd *exec.Cmd) {
	t.Helper()
	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("output:\n%s", b)
		t.Fatal(err)
	}
}

func mustRunShow(t *testing.T, cmd *exec.Cmd) {
	t.Helper()
	b, err := cmd.CombinedOutput()
	t.Logf("output:\n%s", b)
	if err != nil {
		t.Fatal(err)
	}
}
