package slp

import (
	"testing"
)

func TestMaxargs(t *testing.T) {
	type args struct {
		s Stm
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success 1",
			args: args{
				s: &stm{
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
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maxargs(tt.args.s); got != tt.want {
				t.Errorf("Maxargs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterp(t *testing.T) {
	type args struct {
		s Stm
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success 1",
			args: args{
				s: &stm{
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
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interp(tt.args.s)
		})
	}
}
