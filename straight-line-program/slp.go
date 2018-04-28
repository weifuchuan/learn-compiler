package slp

// Any type
type Any = interface{}

// 二元运算符
type Binop int

const (
	Plus  Binop = iota // +
	Minus              // -
	Times              // *
	Div                // /
)

// 本质
type Kind int

const (
	CompoundStm Kind = iota // 复合语句
	AssignStm               // 赋值语句
	PrintStm                // 打印语句

	IdExp   // Id表达式
	NumExp  // 数字表达式
	OpExp   // 运算表达式
	EseqExp // 逗号表达式

	PairExpList // 表达式列表前面
	LastExpList // 表达式列表最后一位
)

// 语句
type Stm = *stm

// 表达式
type Exp = *exp

// 表达式列表
type ExpList = *expList

type stm struct {
	Kind Kind
	U    Any
}

type Compound struct {
	Stm1 Stm
	Stm2 Stm
}

type Assign struct {
	Id  string
	Exp Exp
}

type Print struct {
	ExpList ExpList
}

type exp struct {
	Kind Kind
	U    Any
}

type Id = string

type Num = int

type Op struct {
	Left, Right Exp
	Oper        Binop
}

type Eseq struct {
	Stm Stm
	Exp Exp
}

type expList struct {
	Kind Kind
	U    Any
}

type Pair struct {
	Head Exp
	Tail ExpList
}

type LastExp = Exp

type table struct {
	id    string
	value int
	tail  *table
}

type intAndTable struct {
	i int
	t *table
}

var StmForUse = &stm{
	Kind: CompoundStm,
	U: Compound{
		Stm1: &stm{
			Kind: AssignStm,
			U: Assign{
				Id: "a",
				Exp: &exp{
					Kind: OpExp,
					U: Op{
						Left: &exp{
							Kind: NumExp,
							U:    5,
						},
						Oper: Plus,
						Right: &exp{
							Kind: NumExp,
							U:    3,
						},
					},
				},
			},
		},
		Stm2: &stm{
			Kind: CompoundStm,
			U: Compound{
				Stm1: &stm{
					Kind: AssignStm,
					U: Assign{
						Id: "b",
						Exp: &exp{
							Kind: EseqExp,
							U: Eseq{
								Stm: &stm{
									Kind: PrintStm,
									U: Print{
										ExpList: &expList{
											Kind: PairExpList,
											U: Pair{
												Head: &exp{
													Kind: IdExp,
													U:    "a",
												},
												Tail: &expList{
													Kind: LastExpList,
													U: &exp{
														Kind: OpExp,
														U: Op{
															Left: &exp{
																Kind: IdExp,
																U:    "a",
															},
															Oper: Minus,
															Right: &exp{
																Kind: NumExp,
																U:    1,
															},
														},
													},
												},
											},
										},
									},
								},
								Exp: &exp{
									Kind: OpExp,
									U: Op{
										Left: &exp{
											Kind: NumExp,
											U:    10,
										},
										Oper: Times,
										Right: &exp{
											Kind: IdExp,
											U:    "a",
										},
									},
								},
							},
						},
					},
				},
				Stm2: &stm{
					Kind: PrintStm,
					U: Print{
						ExpList: &expList{
							Kind: LastExpList,
							U: &exp{
								Kind: IdExp,
								U:    "b",
							},
						},
					},
				},
			},
		},
	},
}
