package main

import (
	"context"
	"fmt"
	"time"
)

type AOut struct{}
type BOut struct{}
type COut struct{}
type CIn struct {
	A AOut
	B BOut
}

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

type Input struct {
	A string
	B string
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()
	go func() {
		select {
		case <-ctx.Done():
			return
		case inputC := <-p.inC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
}

func getResultA(ctx context.Context, s string) (struct{}, error) {
	time.Sleep(60 * time.Millisecond)
	fmt.Println("in getResultA():", ctx, s)
	return struct{}{}, nil
}

func getResultB(ctx context.Context, s string) (struct{}, error) {
	fmt.Println("in getResultB():", ctx, s)
	return struct{}{}, nil
}

func getResultC(ctx context.Context, cin CIn) (struct{}, error) {
	//time.Sleep(30 * time.Millisecond)
	fmt.Println("in getResultC():", ctx, cin)
	return struct{}{}, nil
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0
	for count < 2 {
		select {
		case a := <-p.outA:
			inputC.A = a
			count++
		case b := <-p.outB:
			inputC.B = b
			count++
		case err := <-p.errs:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2),
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

func main512() {
	cOut, err := GatherAndProcess(context.Background(), Input{A: "aa", B: "bb"})
	fmt.Println(cOut, err)
}
