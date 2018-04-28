package slp

import (
	"fmt"
	"math"
)

func Maxargs(s Stm) int {
	switch s.Kind {
	case CompoundStm:
		c := s.U.(Compound)
		return Maxargs(c.Stm1) + Maxargs(c.Stm2)
	case AssignStm:
		a := s.U.(Assign)
		return maxargsInExp(a.Exp)
	case PrintStm:
		p := s.U.(Print)
		return 1 + maxargsInExpList(p.ExpList)
	default:
		return 0
	}
}

func maxargsInExp(e Exp) int {
	switch e.Kind {
	case OpExp:
		o := e.U.(Op)
		return maxargsInExp(o.Left) + maxargsInExp(o.Right)
	case EseqExp:
		q := e.U.(Eseq)
		return Maxargs(q.Stm) + maxargsInExp(q.Exp)
	default:
		return 0
	}
}

func maxargsInExpList(list ExpList) int {
	switch list.Kind {
	case PairExpList:
		p := list.U.(Pair)
		return maxargsInExp(p.Head) + maxargsInExpList(p.Tail)
	case LastExpList:
		l := list.U.(LastExp)
		return maxargsInExp(l)
	default:
		return 0
	}
}

func Interp(s Stm) {
	interpStm(s, nil)
}

func interpStm(s Stm, t *table) *table {
	switch s.Kind {
	case CompoundStm:
		c := s.U.(Compound)
		t2 := interpStm(c.Stm1, t)
		return interpStm(c.Stm2, t2)
	case AssignStm:
		a := s.U.(Assign)
		r := interpExp(a.Exp, t)
		t2 := newTable(a.Id, r.i, r.t)
		return t2
	case PrintStm:
		p := s.U.(Print)
		for {
			switch p.ExpList.Kind {
			case PairExpList:
				pair := p.ExpList.U.(Pair)
				r := interpExp(pair.Head, t)
				fmt.Print(r.i)
				fmt.Print(" ")
				t = r.t
				p.ExpList = pair.Tail
			case LastExpList:
				last := p.ExpList.U.(LastExp)
				r := interpExp(last, t)
				fmt.Print(r.i)
				fmt.Print(" ")
				t = r.t
				goto OUT
			default:
				goto OUT
			}
		}
OUT: 
		fmt.Print("\n")
		return t
	default:
		return t
	}
}

func interpExp(e Exp, t *table) intAndTable {
	switch e.Kind {
	case IdExp: // Id表达式
		id := e.U.(Id)
		return intAndTable{
			i: lookup(t, id),
			t: t,
		}
	case NumExp: // 数字表达式
		num := e.U.(Num)
		return intAndTable{
			i: num,
			t: t,
		}
	case OpExp: // 运算表达式
		op := e.U.(Op)
		r1 := interpExp(op.Left, t)
		r2 := interpExp(op.Right, r1.t)
		switch op.Oper {
		case Plus: // +
			return intAndTable{
				i: r1.i + r2.i,
				t: r2.t,
			}
		case Minus: // -
			return intAndTable{
				i: r1.i - r2.i,
				t: r2.t,
			}
		case Times: // *
			return intAndTable{
				i: r1.i * r2.i,
				t: r2.t,
			}
		case Div: // /
			return intAndTable{
				i: r1.i / r2.i,
				t: r2.t,
			}
		}
	case EseqExp: // 逗号表达式
		q := e.U.(Eseq)
		t2 := interpStm(q.Stm, t)
		return interpExp(q.Exp, t2)
	default:
	}
	return intAndTable{}
}

func newTable(id string, value int, tail *table) *table {
	return &table{
		id:    id,
		value: value,
		tail:  tail,
	}
}

func update(t *table, id string, value int) *table {
	return newTable(id, value, t)
}

func lookup(t *table, id string) int {
	for p := t; p != nil; p = t.tail {
		if p.id == id {
			return p.value
		}
	}
	return math.MinInt32
}
