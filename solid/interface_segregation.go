package solid

import "fmt"

type Document string

type Printer interface {
	Print(d Document) string
}

type Scanner interface {
	Scan(d Document) string
}

type Copier interface {
	Copy(d Document) string
}

type PrinterMachine struct{}

func (p PrinterMachine) Print(d Document) string {
	return fmt.Sprintf("Printed: %s", d)
}

type ScannerMachine struct{}

func (s ScannerMachine) Scan(d Document) string {
	return fmt.Sprintf("Scanned: %s", d)
}

type CopierMachine struct{}

func (c CopierMachine) Copy(d Document) string {
	return fmt.Sprintf("Copied: %s", d)
}

type MultiFunctionalMachine struct {
	Printer
	Scanner
	Copier
}
