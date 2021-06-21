package godesignpatterns

import "testing"

func TestPrinter(t *testing.T) {
	testData := []struct {
		printer  Printer
		document Document
		want     string
	}{
		{
			PrinterMachine{},
			"This is a good printer",
			"Printed: This is a good printer",
		},
		{
			MultiFunctionalMachine{PrinterMachine{}, ScannerMachine{}, CopierMachine{}},
			"This is a good printer",
			"Printed: This is a good printer",
		},
	}

	for _, v := range testData {
		if r := v.printer.Print(v.document); r != v.want {
			t.Errorf("Printer does not print right, got: %s, expected: %s\n", r, v.want)
		}
	}
}

func TestScanner(t *testing.T) {
	testData := []struct {
		scanner  Scanner
		document Document
		want     string
	}{
		{
			ScannerMachine{},
			"This is a good scanner",
			"Scanned: This is a good scanner",
		},
		{
			MultiFunctionalMachine{PrinterMachine{}, ScannerMachine{}, CopierMachine{}},
			"This is a good scanner",
			"Scanned: This is a good scanner",
		},
	}

	for _, v := range testData {
		if r := v.scanner.Scan(v.document); r != v.want {
			t.Errorf("Scanner does not scan right, got: %s, expected: %s\n", r, v.want)
		}
	}
}

func TestCopier(t *testing.T) {
	testData := []struct {
		copier   Copier
		document Document
		want     string
	}{
		{
			CopierMachine{},
			"This is a good copier",
			"Copied: This is a good copier",
		},
		{
			MultiFunctionalMachine{PrinterMachine{}, ScannerMachine{}, CopierMachine{}},
			"This is a good copier",
			"Copied: This is a good copier",
		},
	}

	for _, v := range testData {
		if r := v.copier.Copy(v.document); r != v.want {
			t.Errorf("Copier does not copy right, got: %s, expected: %s\n", r, v.want)
		}
	}
}
