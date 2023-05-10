package main

import (
	"fmt"
)

type Document struct {
	title string
}

type Machine interface {
	Print(d Document)
	Scan(d Document)
	Fax(d Document)
}

func RunMachine(m Machine, d Document) {
	m.Print(d)
	m.Scan(d)
	m.Fax(d)
}

type MultiFunctionalPrinter struct{}

func (m MultiFunctionalPrinter) Print(d Document) {
	fmt.Printf("Multi Functional Printer is printing document: %s\n", d.title)
}

func (m MultiFunctionalPrinter) Scan(d Document) {
	fmt.Printf("Multi Functional Printer is scanning document: %s\n", d.title)
}

func (m MultiFunctionalPrinter) Fax(d Document) {
	fmt.Printf("Multi Functional Printer is faxing document: %s\n", d.title)
}

type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Printf("Old Fashioned Printer is printing document: %s\n", d.title)
}

func (o OldFashionedPrinter) Scan(d Document) {
	fmt.Println("Old Fashioned Printer does not support scanning")
}

func (o OldFashionedPrinter) Fax(d Document) {
	fmt.Println("Old Fashioned Printer does not support faxing")
}

type Printer interface {
	Print(d Document)
}

func RunPrinter(p Printer, d Document) {
	p.Print(d)
}

type Scanner interface {
	Scan(d Document)
}

func RunScanner(s Scanner, d Document) {
	s.Scan(d)
}

type Faxer interface {
	Fax(d Document)
}

func RunFaxer(f Faxer, d Document) {
	f.Fax(d)
}

type AnotherOldFashionedPrinter struct{}

func (a AnotherOldFashionedPrinter) Print(d Document) {
	fmt.Printf("Another Old Fashioned Printer is printing document: %s\n", d.title)
}

type PrinterScanner interface {
	Printer
	Scanner
}

func RunPrinterScanner(p PrinterScanner, d Document) {
	p.Print(d)
	p.Scan(d)
}

type PhotoCopier struct{}

func (p PhotoCopier) Print(d Document) {
	fmt.Printf("Photo Copier is printing document: %s\n", d.title)
}

func (p PhotoCopier) Scan(d Document) {
	fmt.Printf("Photo Copier is scanning document: %s\n", d.title)
}

type DecoratedPhotoCopier struct {
	PrinterScanner
}

func (dpc DecoratedPhotoCopier) Print(d Document) {
	fmt.Print("Decorated printer = ")
	dpc.PrinterScanner.Print(d)
}

func (dpc DecoratedPhotoCopier) Scan(d Document) {
	fmt.Print("Decorated scanner = ")
	dpc.PrinterScanner.Scan(d)
}

func main() {
	document := Document{"To be or not to be"}

	mfp := MultiFunctionalPrinter{}
	RunMachine(mfp, document)

	// the old fashioned printer can print but cannot scan and fax
	// it needs to implement Scan and Fax methods in order to be a Machine
	ofp := OldFashionedPrinter{}
	RunMachine(ofp, document)

	//therefore segregate the Machine interface into
	//Printer, Scan and Fax interfaxes
	//the another old fashioned printer can only print and only implements the Print interface
	aofm := AnotherOldFashionedPrinter{}
	RunPrinter(aofm, document)

	//photo copier can print and scan
	//photo copier implements PrinterScanner interface
	//which is an aggregation of Printer and Scanner interface
	pc := PhotoCopier{}
	RunPrinterScanner(pc, document)

	//Decorate a value that implement the PrinterScanner interface
	dpc := DecoratedPhotoCopier{pc}
	RunPrinterScanner(dpc, document)
}
