package solver

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	curExpression, err := readToNewLine(r)
	fmt.Println("curExpression:", curExpression)
	if err != nil {
		return 0, err
	}
	if len(curExpression) == 0 {
		return 0, errors.New("no expression to read")
	}
	answer, err := p.Solver.Resolve(ctx, curExpression)
	return answer, err
}

func readToNewLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	fmt.Println(line, string(line))
	return string(line), nil
}

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}
