# github.com/phpor/go-example/app/asm
"".main STEXT size=58 args=0x0 locals=0x8
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	TEXT	"".main(SB), ABIInternal, $8-0
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	JLS	51
	0x000f 00015 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	SUBQ	$8, SP
	0x0013 00019 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	MOVQ	BP, (SP)
	0x0017 00023 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	LEAQ	(SP), BP
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:15)	PCDATA	$0, $0
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:15)	PCDATA	$1, $0
	0x001b 00027 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:15)	CALL	"".f1(SB)
	0x0020 00032 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:16)	CALL	"".f2(SB)
	0x0025 00037 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:17)	CALL	"".f3(SB)
	0x002a 00042 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:18)	MOVQ	(SP), BP
	0x002e 00046 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:18)	ADDQ	$8, SP
	0x0032 00050 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:18)	RET
	0x0033 00051 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:18)	NOP
	0x0033 00051 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	PCDATA	$1, $-1
	0x0033 00051 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	PCDATA	$0, $-1
	0x0033 00051 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x0038 00056 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:14)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 24 48  dH..%....H;a.v$H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 e8 00 00 00 00  ...H.,$H.,$.....
	0x0020 e8 00 00 00 00 e8 00 00 00 00 48 8b 2c 24 48 83  ..........H.,$H.
	0x0030 c4 08 c3 e8 00 00 00 00 eb c6                    ..........
	rel 5+4 t=16 TLS+0
	rel 28+4 t=8 "".f1+0
	rel 33+4 t=8 "".f2+0
	rel 38+4 t=8 "".f3+0
	rel 52+4 t=8 runtime.morestack_noctxt+0
"".f1 STEXT size=318 args=0x0 locals=0x88
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	TEXT	"".f1(SB), ABIInternal, $136-0
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	LEAQ	-8(SP), AX
	0x000e 00014 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	CMPQ	AX, 16(CX)
	0x0012 00018 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	JLS	308
	0x0018 00024 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	SUBQ	$136, SP
	0x001f 00031 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	MOVQ	BP, 128(SP)
	0x0027 00039 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	LEAQ	128(SP), BP
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	FUNCDATA	$1, gclocals·94d172d993ad121e3746e9b06f1f5579(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	FUNCDATA	$2, gclocals·33b901baab2acec3083d16f1ab81c65a(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	FUNCDATA	$3, "".f1.stkobj(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	PCDATA	$0, $0
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	PCDATA	$1, $0
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	XORPS	X0, X0
	0x0032 00050 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	MOVUPS	X0, "".p+8(SB)
	0x0039 00057 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	PCDATA	$0, $-2
	0x0039 00057 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	PCDATA	$1, $-2
	0x0039 00057 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	CMPL	runtime.writeBarrier(SB), $0
	0x0040 00064 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	JEQ	71
	0x0042 00066 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	JMP	289
	0x0047 00071 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	MOVQ	$0, "".p(SB)
	0x0052 00082 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	JMP	84
	0x0054 00084 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x0054 00084 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $1
	0x0054 00084 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	XORPS	X0, X0
	0x0057 00087 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVUPS	X0, ""..autotmp_0+88(SP)
	0x005c 00092 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x005c 00092 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $0
	0x005c 00092 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	LEAQ	""..autotmp_0+88(SP), AX
	0x0061 00097 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x0061 00097 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $2
	0x0061 00097 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, ""..autotmp_2+64(SP)
	0x0066 00102 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x0066 00102 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	LEAQ	type."".Person(SB), AX
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, (SP)
	0x0071 00113 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x0071 00113 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	LEAQ	"".p(SB), AX
	0x0078 00120 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x0078 00120 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, 8(SP)
	0x007d 00125 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	CALL	runtime.convT2E(SB)
	0x0082 00130 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x0082 00130 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	24(SP), AX
	0x0087 00135 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	16(SP), CX
	0x008c 00140 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	CX, ""..autotmp_3+72(SP)
	0x0091 00145 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, ""..autotmp_3+80(SP)
	0x0096 00150 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $2
	0x0096 00150 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	""..autotmp_2+64(SP), DX
	0x009b 00155 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	TESTB	AL, (DX)
	0x009d 00157 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	CX, (DX)
	0x00a0 00160 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $3
	0x00a0 00160 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	LEAQ	8(DX), DI
	0x00a4 00164 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $-2
	0x00a4 00164 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $-2
	0x00a4 00164 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	CMPL	runtime.writeBarrier(SB), $0
	0x00ab 00171 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	JEQ	175
	0x00ad 00173 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	JMP	282
	0x00af 00175 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, 8(DX)
	0x00b3 00179 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	JMP	181
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $0
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	""..autotmp_2+64(SP), AX
	0x00ba 00186 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	TESTB	AL, (AX)
	0x00bc 00188 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	JMP	190
	0x00be 00190 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x00be 00190 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $3
	0x00be 00190 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, ""..autotmp_1+104(SP)
	0x00c3 00195 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	$1, ""..autotmp_1+112(SP)
	0x00cc 00204 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	$1, ""..autotmp_1+120(SP)
	0x00d5 00213 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x00d5 00213 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	LEAQ	go.string."%#v"(SB), AX
	0x00dc 00220 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x00dc 00220 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, (SP)
	0x00e0 00224 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	$3, 8(SP)
	0x00e9 00233 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $1
	0x00e9 00233 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $0
	0x00e9 00233 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	""..autotmp_1+104(SP), AX
	0x00ee 00238 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $0
	0x00ee 00238 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	AX, 16(SP)
	0x00f3 00243 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	$1, 24(SP)
	0x00fc 00252 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	MOVQ	$1, 32(SP)
	0x0105 00261 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	CALL	fmt.Printf(SB)
	0x010a 00266 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:22)	MOVQ	128(SP), BP
	0x0112 00274 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:22)	ADDQ	$136, SP
	0x0119 00281 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:22)	RET
	0x011a 00282 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$0, $-2
	0x011a 00282 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	PCDATA	$1, $-2
	0x011a 00282 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	CALL	runtime.gcWriteBarrier(SB)
	0x011f 00287 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:21)	JMP	181
	0x0121 00289 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	LEAQ	"".p(SB), DI
	0x0128 00296 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	XORL	AX, AX
	0x012a 00298 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	CALL	runtime.gcWriteBarrier(SB)
	0x012f 00303 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	JMP	84
	0x0134 00308 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:20)	NOP
	0x0134 00308 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	PCDATA	$1, $-1
	0x0134 00308 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	PCDATA	$0, $-1
	0x0134 00308 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	CALL	runtime.morestack_noctxt(SB)
	0x0139 00313 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:19)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 f8 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 1c 01 00 00 48 81 ec 88 00 00 00 48  A.......H......H
	0x0020 89 ac 24 80 00 00 00 48 8d ac 24 80 00 00 00 0f  ..$....H..$.....
	0x0030 57 c0 0f 11 05 00 00 00 00 83 3d 00 00 00 00 00  W.........=.....
	0x0040 74 05 e9 da 00 00 00 48 c7 05 00 00 00 00 00 00  t......H........
	0x0050 00 00 eb 00 0f 57 c0 0f 11 44 24 58 48 8d 44 24  .....W...D$XH.D$
	0x0060 58 48 89 44 24 40 48 8d 05 00 00 00 00 48 89 04  XH.D$@H......H..
	0x0070 24 48 8d 05 00 00 00 00 48 89 44 24 08 e8 00 00  $H......H.D$....
	0x0080 00 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24  ..H.D$.H.L$.H.L$
	0x0090 48 48 89 44 24 50 48 8b 54 24 40 84 02 48 89 0a  HH.D$PH.T$@..H..
	0x00a0 48 8d 7a 08 83 3d 00 00 00 00 00 74 02 eb 6b 48  H.z..=.....t..kH
	0x00b0 89 42 08 eb 00 48 8b 44 24 40 84 00 eb 00 48 89  .B...H.D$@....H.
	0x00c0 44 24 68 48 c7 44 24 70 01 00 00 00 48 c7 44 24  D$hH.D$p....H.D$
	0x00d0 78 01 00 00 00 48 8d 05 00 00 00 00 48 89 04 24  x....H......H..$
	0x00e0 48 c7 44 24 08 03 00 00 00 48 8b 44 24 68 48 89  H.D$.....H.D$hH.
	0x00f0 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44 24  D$.H.D$.....H.D$
	0x0100 20 01 00 00 00 e8 00 00 00 00 48 8b ac 24 80 00   .........H..$..
	0x0110 00 00 48 81 c4 88 00 00 00 c3 e8 00 00 00 00 eb  ..H.............
	0x0120 94 48 8d 3d 00 00 00 00 31 c0 e8 00 00 00 00 e9  .H.=....1.......
	0x0130 20 ff ff ff e8 00 00 00 00 e9 c2 fe ff ff         .............
	rel 5+4 t=16 TLS+0
	rel 53+4 t=15 "".p+8
	rel 59+4 t=15 runtime.writeBarrier+-1
	rel 74+4 t=15 "".p+-4
	rel 105+4 t=15 type."".Person+0
	rel 116+4 t=15 "".p+0
	rel 126+4 t=8 runtime.convT2E+0
	rel 166+4 t=15 runtime.writeBarrier+-1
	rel 216+4 t=15 go.string."%#v"+0
	rel 262+4 t=8 fmt.Printf+0
	rel 283+4 t=8 runtime.gcWriteBarrier+0
	rel 292+4 t=15 "".p+0
	rel 299+4 t=8 runtime.gcWriteBarrier+0
	rel 309+4 t=8 runtime.morestack_noctxt+0
"".f2 STEXT size=347 args=0x0 locals=0xa0
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	TEXT	"".f2(SB), ABIInternal, $160-0
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	LEAQ	-32(SP), AX
	0x000e 00014 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	CMPQ	AX, 16(CX)
	0x0012 00018 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	JLS	337
	0x0018 00024 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	SUBQ	$160, SP
	0x001f 00031 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	MOVQ	BP, 152(SP)
	0x0027 00039 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	LEAQ	152(SP), BP
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	FUNCDATA	$1, gclocals·4b62fca614ba1b83331bcc57182ac46b(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	FUNCDATA	$2, gclocals·33b901baab2acec3083d16f1ab81c65a(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	FUNCDATA	$3, "".f2.stkobj(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:24)	PCDATA	$0, $0
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:24)	PCDATA	$1, $0
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:24)	XORPS	X0, X0
	0x0032 00050 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:24)	MOVUPS	X0, "".b+104(SP)
	0x0037 00055 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:24)	MOVQ	$0, "".b+120(SP)
	0x0040 00064 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	XORPS	X0, X0
	0x0043 00067 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	MOVUPS	X0, "".p+8(SB)
	0x004a 00074 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	PCDATA	$0, $-2
	0x004a 00074 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	PCDATA	$1, $-2
	0x004a 00074 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	CMPL	runtime.writeBarrier(SB), $0
	0x0051 00081 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	JEQ	88
	0x0053 00083 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	JMP	318
	0x0058 00088 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	MOVQ	$0, "".p(SB)
	0x0063 00099 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	JMP	101
	0x0065 00101 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x0065 00101 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $1
	0x0065 00101 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	XORPS	X0, X0
	0x0068 00104 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVUPS	X0, ""..autotmp_1+88(SP)
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $0
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	LEAQ	""..autotmp_1+88(SP), AX
	0x0072 00114 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x0072 00114 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $2
	0x0072 00114 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, ""..autotmp_3+64(SP)
	0x0077 00119 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x0077 00119 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	LEAQ	type."".Person(SB), AX
	0x007e 00126 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x007e 00126 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, (SP)
	0x0082 00130 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x0082 00130 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	LEAQ	"".p(SB), AX
	0x0089 00137 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x0089 00137 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, 8(SP)
	0x008e 00142 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	CALL	runtime.convT2E(SB)
	0x0093 00147 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x0093 00147 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	24(SP), AX
	0x0098 00152 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	16(SP), CX
	0x009d 00157 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	CX, ""..autotmp_4+72(SP)
	0x00a2 00162 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, ""..autotmp_4+80(SP)
	0x00a7 00167 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $2
	0x00a7 00167 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	""..autotmp_3+64(SP), DX
	0x00ac 00172 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	TESTB	AL, (DX)
	0x00ae 00174 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	CX, (DX)
	0x00b1 00177 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $3
	0x00b1 00177 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	LEAQ	8(DX), DI
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $-2
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $-2
	0x00b5 00181 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	CMPL	runtime.writeBarrier(SB), $0
	0x00bc 00188 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	JEQ	192
	0x00be 00190 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	JMP	311
	0x00c0 00192 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, 8(DX)
	0x00c4 00196 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	JMP	198
	0x00c6 00198 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x00c6 00198 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $0
	0x00c6 00198 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	""..autotmp_3+64(SP), AX
	0x00cb 00203 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	TESTB	AL, (AX)
	0x00cd 00205 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	JMP	207
	0x00cf 00207 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x00cf 00207 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $3
	0x00cf 00207 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, ""..autotmp_2+128(SP)
	0x00d7 00215 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	$1, ""..autotmp_2+136(SP)
	0x00e3 00227 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	$1, ""..autotmp_2+144(SP)
	0x00ef 00239 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x00ef 00239 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	LEAQ	go.string."%#v"(SB), AX
	0x00f6 00246 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x00f6 00246 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, (SP)
	0x00fa 00250 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	$3, 8(SP)
	0x0103 00259 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $1
	0x0103 00259 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $0
	0x0103 00259 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	""..autotmp_2+128(SP), AX
	0x010b 00267 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $0
	0x010b 00267 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	AX, 16(SP)
	0x0110 00272 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	$1, 24(SP)
	0x0119 00281 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	MOVQ	$1, 32(SP)
	0x0122 00290 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	CALL	fmt.Printf(SB)
	0x0127 00295 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:27)	MOVQ	152(SP), BP
	0x012f 00303 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:27)	ADDQ	$160, SP
	0x0136 00310 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:27)	RET
	0x0137 00311 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$0, $-2
	0x0137 00311 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	PCDATA	$1, $-2
	0x0137 00311 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	CALL	runtime.gcWriteBarrier(SB)
	0x013c 00316 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:26)	JMP	198
	0x013e 00318 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	LEAQ	"".p(SB), DI
	0x0145 00325 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	XORL	AX, AX
	0x0147 00327 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	CALL	runtime.gcWriteBarrier(SB)
	0x014c 00332 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	JMP	101
	0x0151 00337 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:25)	NOP
	0x0151 00337 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	PCDATA	$1, $-1
	0x0151 00337 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	PCDATA	$0, $-1
	0x0151 00337 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	CALL	runtime.morestack_noctxt(SB)
	0x0156 00342 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:23)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 e0 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 39 01 00 00 48 81 ec a0 00 00 00 48  A...9...H......H
	0x0020 89 ac 24 98 00 00 00 48 8d ac 24 98 00 00 00 0f  ..$....H..$.....
	0x0030 57 c0 0f 11 44 24 68 48 c7 44 24 78 00 00 00 00  W...D$hH.D$x....
	0x0040 0f 57 c0 0f 11 05 00 00 00 00 83 3d 00 00 00 00  .W.........=....
	0x0050 00 74 05 e9 e6 00 00 00 48 c7 05 00 00 00 00 00  .t......H.......
	0x0060 00 00 00 eb 00 0f 57 c0 0f 11 44 24 58 48 8d 44  ......W...D$XH.D
	0x0070 24 58 48 89 44 24 40 48 8d 05 00 00 00 00 48 89  $XH.D$@H......H.
	0x0080 04 24 48 8d 05 00 00 00 00 48 89 44 24 08 e8 00  .$H......H.D$...
	0x0090 00 00 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c  ...H.D$.H.L$.H.L
	0x00a0 24 48 48 89 44 24 50 48 8b 54 24 40 84 02 48 89  $HH.D$PH.T$@..H.
	0x00b0 0a 48 8d 7a 08 83 3d 00 00 00 00 00 74 02 eb 77  .H.z..=.....t..w
	0x00c0 48 89 42 08 eb 00 48 8b 44 24 40 84 00 eb 00 48  H.B...H.D$@....H
	0x00d0 89 84 24 80 00 00 00 48 c7 84 24 88 00 00 00 01  ..$....H..$.....
	0x00e0 00 00 00 48 c7 84 24 90 00 00 00 01 00 00 00 48  ...H..$........H
	0x00f0 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24 08 03  ......H..$H.D$..
	0x0100 00 00 00 48 8b 84 24 80 00 00 00 48 89 44 24 10  ...H..$....H.D$.
	0x0110 48 c7 44 24 18 01 00 00 00 48 c7 44 24 20 01 00  H.D$.....H.D$ ..
	0x0120 00 00 e8 00 00 00 00 48 8b ac 24 98 00 00 00 48  .......H..$....H
	0x0130 81 c4 a0 00 00 00 c3 e8 00 00 00 00 eb 88 48 8d  ..............H.
	0x0140 3d 00 00 00 00 31 c0 e8 00 00 00 00 e9 14 ff ff  =....1..........
	0x0150 ff e8 00 00 00 00 e9 a5 fe ff ff                 ...........
	rel 5+4 t=16 TLS+0
	rel 70+4 t=15 "".p+8
	rel 76+4 t=15 runtime.writeBarrier+-1
	rel 91+4 t=15 "".p+-4
	rel 122+4 t=15 type."".Person+0
	rel 133+4 t=15 "".p+0
	rel 143+4 t=8 runtime.convT2E+0
	rel 183+4 t=15 runtime.writeBarrier+-1
	rel 242+4 t=15 go.string."%#v"+0
	rel 291+4 t=8 fmt.Printf+0
	rel 312+4 t=8 runtime.gcWriteBarrier+0
	rel 321+4 t=15 "".p+0
	rel 328+4 t=8 runtime.gcWriteBarrier+0
	rel 338+4 t=8 runtime.morestack_noctxt+0
"".f3 STEXT size=381 args=0x0 locals=0xa8
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	TEXT	"".f3(SB), ABIInternal, $168-0
	0x0000 00000 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	LEAQ	-40(SP), AX
	0x000e 00014 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	CMPQ	AX, 16(CX)
	0x0012 00018 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	JLS	371
	0x0018 00024 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	SUBQ	$168, SP
	0x001f 00031 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	MOVQ	BP, 160(SP)
	0x0027 00039 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	LEAQ	160(SP), BP
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	FUNCDATA	$0, gclocals·3e27b3aa6b89137cce48b3379a2a6610(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	FUNCDATA	$1, gclocals·9cf841ba61ca75101c532af3ea6af54c(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	FUNCDATA	$2, gclocals·fd8cf83e1c48dd0b84c741be7b1e4c9c(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	FUNCDATA	$3, "".f3.stkobj(SB)
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	PCDATA	$0, $0
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	PCDATA	$1, $1
	0x002f 00047 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	XORPS	X0, X0
	0x0032 00050 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	MOVUPS	X0, ""..autotmp_2+136(SP)
	0x003a 00058 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	MOVQ	$0, ""..autotmp_2+152(SP)
	0x0046 00070 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	PCDATA	$0, $1
	0x0046 00070 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	LEAQ	""..autotmp_2+136(SP), CX
	0x004e 00078 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:30)	MOVQ	CX, "".b+64(SP)
	0x0053 00083 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	PCDATA	$0, $0
	0x0053 00083 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	TESTB	AL, (CX)
	0x0055 00085 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	""..autotmp_2+144(SP), CX
	0x005d 00093 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	PCDATA	$0, $2
	0x005d 00093 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	""..autotmp_2+136(SP), AX
	0x0065 00101 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	PCDATA	$1, $0
	0x0065 00101 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	""..autotmp_2+152(SP), DX
	0x006d 00109 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	CX, "".p+8(SB)
	0x0074 00116 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	DX, "".p+16(SB)
	0x007b 00123 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	PCDATA	$0, $-2
	0x007b 00123 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	PCDATA	$1, $-2
	0x007b 00123 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	CMPL	runtime.writeBarrier(SB), $0
	0x0082 00130 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	JEQ	137
	0x0084 00132 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	JMP	354
	0x0089 00137 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	MOVQ	AX, "".p(SB)
	0x0090 00144 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	JMP	146
	0x0092 00146 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x0092 00146 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $2
	0x0092 00146 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	XORPS	X0, X0
	0x0095 00149 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVUPS	X0, ""..autotmp_1+96(SP)
	0x009a 00154 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x009a 00154 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $0
	0x009a 00154 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	LEAQ	""..autotmp_1+96(SP), AX
	0x009f 00159 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x009f 00159 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $3
	0x009f 00159 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00a4 00164 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x00a4 00164 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	LEAQ	type."".Person(SB), AX
	0x00ab 00171 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x00ab 00171 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, (SP)
	0x00af 00175 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x00af 00175 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	LEAQ	"".p(SB), AX
	0x00b6 00182 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x00b6 00182 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, 8(SP)
	0x00bb 00187 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	CALL	runtime.convT2E(SB)
	0x00c0 00192 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x00c0 00192 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	24(SP), AX
	0x00c5 00197 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	16(SP), CX
	0x00ca 00202 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	CX, ""..autotmp_5+80(SP)
	0x00cf 00207 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, ""..autotmp_5+88(SP)
	0x00d4 00212 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $3
	0x00d4 00212 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	""..autotmp_4+72(SP), DX
	0x00d9 00217 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	TESTB	AL, (DX)
	0x00db 00219 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	CX, (DX)
	0x00de 00222 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $4
	0x00de 00222 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	LEAQ	8(DX), DI
	0x00e2 00226 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $-2
	0x00e2 00226 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $-2
	0x00e2 00226 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	CMPL	runtime.writeBarrier(SB), $0
	0x00e9 00233 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	JEQ	237
	0x00eb 00235 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	JMP	347
	0x00ed 00237 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, 8(DX)
	0x00f1 00241 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	JMP	243
	0x00f3 00243 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x00f3 00243 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $0
	0x00f3 00243 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	""..autotmp_4+72(SP), AX
	0x00f8 00248 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	TESTB	AL, (AX)
	0x00fa 00250 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	JMP	252
	0x00fc 00252 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x00fc 00252 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $4
	0x00fc 00252 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, ""..autotmp_3+112(SP)
	0x0101 00257 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	$1, ""..autotmp_3+120(SP)
	0x010a 00266 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	$1, ""..autotmp_3+128(SP)
	0x0116 00278 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x0116 00278 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	LEAQ	go.string."%#v"(SB), AX
	0x011d 00285 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x011d 00285 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, (SP)
	0x0121 00289 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	$3, 8(SP)
	0x012a 00298 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $2
	0x012a 00298 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $0
	0x012a 00298 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	""..autotmp_3+112(SP), AX
	0x012f 00303 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $0
	0x012f 00303 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	AX, 16(SP)
	0x0134 00308 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	$1, 24(SP)
	0x013d 00317 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	MOVQ	$1, 32(SP)
	0x0146 00326 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	CALL	fmt.Printf(SB)
	0x014b 00331 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:33)	MOVQ	160(SP), BP
	0x0153 00339 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:33)	ADDQ	$168, SP
	0x015a 00346 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:33)	RET
	0x015b 00347 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$0, $-2
	0x015b 00347 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	PCDATA	$1, $-2
	0x015b 00347 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	CALL	runtime.gcWriteBarrier(SB)
	0x0160 00352 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:32)	JMP	243
	0x0162 00354 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	LEAQ	"".p(SB), DI
	0x0169 00361 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	CALL	runtime.gcWriteBarrier(SB)
	0x016e 00366 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	JMP	146
	0x0173 00371 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:31)	NOP
	0x0173 00371 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	PCDATA	$1, $-1
	0x0173 00371 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	PCDATA	$0, $-1
	0x0173 00371 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	CALL	runtime.morestack_noctxt(SB)
	0x0178 00376 (/Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go:29)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 d8 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 5b 01 00 00 48 81 ec a8 00 00 00 48  A...[...H......H
	0x0020 89 ac 24 a0 00 00 00 48 8d ac 24 a0 00 00 00 0f  ..$....H..$.....
	0x0030 57 c0 0f 11 84 24 88 00 00 00 48 c7 84 24 98 00  W....$....H..$..
	0x0040 00 00 00 00 00 00 48 8d 8c 24 88 00 00 00 48 89  ......H..$....H.
	0x0050 4c 24 40 84 01 48 8b 8c 24 90 00 00 00 48 8b 84  L$@..H..$....H..
	0x0060 24 88 00 00 00 48 8b 94 24 98 00 00 00 48 89 0d  $....H..$....H..
	0x0070 00 00 00 00 48 89 15 00 00 00 00 83 3d 00 00 00  ....H.......=...
	0x0080 00 00 74 05 e9 d9 00 00 00 48 89 05 00 00 00 00  ..t......H......
	0x0090 eb 00 0f 57 c0 0f 11 44 24 60 48 8d 44 24 60 48  ...W...D$`H.D$`H
	0x00a0 89 44 24 48 48 8d 05 00 00 00 00 48 89 04 24 48  .D$HH......H..$H
	0x00b0 8d 05 00 00 00 00 48 89 44 24 08 e8 00 00 00 00  ......H.D$......
	0x00c0 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24 50 48  H.D$.H.L$.H.L$PH
	0x00d0 89 44 24 58 48 8b 54 24 48 84 02 48 89 0a 48 8d  .D$XH.T$H..H..H.
	0x00e0 7a 08 83 3d 00 00 00 00 00 74 02 eb 6e 48 89 42  z..=.....t..nH.B
	0x00f0 08 eb 00 48 8b 44 24 48 84 00 eb 00 48 89 44 24  ...H.D$H....H.D$
	0x0100 70 48 c7 44 24 78 01 00 00 00 48 c7 84 24 80 00  pH.D$x....H..$..
	0x0110 00 00 01 00 00 00 48 8d 05 00 00 00 00 48 89 04  ......H......H..
	0x0120 24 48 c7 44 24 08 03 00 00 00 48 8b 44 24 70 48  $H.D$.....H.D$pH
	0x0130 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44  .D$.H.D$.....H.D
	0x0140 24 20 01 00 00 00 e8 00 00 00 00 48 8b ac 24 a0  $ .........H..$.
	0x0150 00 00 00 48 81 c4 a8 00 00 00 c3 e8 00 00 00 00  ...H............
	0x0160 eb 91 48 8d 3d 00 00 00 00 e8 00 00 00 00 e9 1f  ..H.=...........
	0x0170 ff ff ff e8 00 00 00 00 e9 83 fe ff ff           .............
	rel 5+4 t=16 TLS+0
	rel 112+4 t=15 "".p+8
	rel 119+4 t=15 "".p+16
	rel 125+4 t=15 runtime.writeBarrier+-1
	rel 140+4 t=15 "".p+0
	rel 167+4 t=15 type."".Person+0
	rel 178+4 t=15 "".p+0
	rel 188+4 t=8 runtime.convT2E+0
	rel 228+4 t=15 runtime.writeBarrier+-1
	rel 281+4 t=15 go.string."%#v"+0
	rel 327+4 t=8 fmt.Printf+0
	rel 348+4 t=8 runtime.gcWriteBarrier+0
	rel 357+4 t=15 "".p+0
	rel 362+4 t=8 runtime.gcWriteBarrier+0
	rel 372+4 t=8 runtime.morestack_noctxt+0
type..hash."".BaseInfo STEXT dupok size=158 args=0x18 locals=0x38
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash."".BaseInfo(SB), DUPOK|ABIInternal, $56-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	148
	0x0013 00019 (<autogenerated>:1)	SUBQ	$56, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 48(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	48(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$1, $0
	0x0021 00033 (<autogenerated>:1)	MOVQ	$0, "".~r2+80(SP)
	0x002a 00042 (<autogenerated>:1)	PCDATA	$0, $1
	0x002a 00042 (<autogenerated>:1)	MOVQ	"".p+64(SP), AX
	0x002f 00047 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+40(SP)
	0x0034 00052 (<autogenerated>:1)	PCDATA	$0, $0
	0x0034 00052 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x0038 00056 (<autogenerated>:1)	MOVQ	"".h+72(SP), AX
	0x003d 00061 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0042 00066 (<autogenerated>:1)	CALL	runtime.strhash(SB)
	0x0047 00071 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x004c 00076 (<autogenerated>:1)	MOVQ	AX, "".h+72(SP)
	0x0051 00081 (<autogenerated>:1)	PCDATA	$0, $1
	0x0051 00081 (<autogenerated>:1)	PCDATA	$1, $1
	0x0051 00081 (<autogenerated>:1)	MOVQ	"".p+64(SP), AX
	0x0056 00086 (<autogenerated>:1)	ADDQ	$16, AX
	0x005a 00090 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+32(SP)
	0x005f 00095 (<autogenerated>:1)	PCDATA	$0, $0
	0x005f 00095 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x0063 00099 (<autogenerated>:1)	MOVQ	"".h+72(SP), AX
	0x0068 00104 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x006d 00109 (<autogenerated>:1)	MOVQ	$8, 16(SP)
	0x0076 00118 (<autogenerated>:1)	CALL	runtime.memhash(SB)
	0x007b 00123 (<autogenerated>:1)	MOVQ	24(SP), AX
	0x0080 00128 (<autogenerated>:1)	MOVQ	AX, "".h+72(SP)
	0x0085 00133 (<autogenerated>:1)	MOVQ	AX, "".~r2+80(SP)
	0x008a 00138 (<autogenerated>:1)	MOVQ	48(SP), BP
	0x008f 00143 (<autogenerated>:1)	ADDQ	$56, SP
	0x0093 00147 (<autogenerated>:1)	RET
	0x0094 00148 (<autogenerated>:1)	NOP
	0x0094 00148 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0094 00148 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0094 00148 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0099 00153 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 81  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 38 48 89 6c 24 30 48 8d 6c 24  ...H..8H.l$0H.l$
	0x0020 30 48 c7 44 24 50 00 00 00 00 48 8b 44 24 40 48  0H.D$P....H.D$@H
	0x0030 89 44 24 28 48 89 04 24 48 8b 44 24 48 48 89 44  .D$(H..$H.D$HH.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24  $......H.D$.H.D$
	0x0050 48 48 8b 44 24 40 48 83 c0 10 48 89 44 24 20 48  HH.D$@H...H.D$ H
	0x0060 89 04 24 48 8b 44 24 48 48 89 44 24 08 48 c7 44  ..$H.D$HH.D$.H.D
	0x0070 24 10 08 00 00 00 e8 00 00 00 00 48 8b 44 24 18  $..........H.D$.
	0x0080 48 89 44 24 48 48 89 44 24 50 48 8b 6c 24 30 48  H.D$HH.D$PH.l$0H
	0x0090 83 c4 38 c3 e8 00 00 00 00 e9 62 ff ff ff        ..8.......b...
	rel 5+4 t=16 TLS+0
	rel 67+4 t=8 runtime.strhash+0
	rel 119+4 t=8 runtime.memhash+0
	rel 149+4 t=8 runtime.morestack_noctxt+0
type..eq."".BaseInfo STEXT dupok size=198 args=0x18 locals=0x50
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq."".BaseInfo(SB), DUPOK|ABIInternal, $80-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	188
	0x0013 00019 (<autogenerated>:1)	SUBQ	$80, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 72(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	72(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·7e7fcb5c7cd183fbe200fb26b1d44a90(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·91d432edea9e3c468c5aec7a805d99d2(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$2, gclocals·6e8d7ea4abad763909b26991048ee1fe(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$1, $0
	0x0021 00033 (<autogenerated>:1)	MOVB	$0, "".~r2+104(SP)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $1
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+88(SP), AX
	0x002b 00043 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x002f 00047 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$1, $1
	0x0032 00050 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+56(SP)
	0x0037 00055 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_4+64(SP)
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $1
	0x003c 00060 (<autogenerated>:1)	MOVQ	"".q+96(SP), AX
	0x0041 00065 (<autogenerated>:1)	PCDATA	$0, $2
	0x0041 00065 (<autogenerated>:1)	MOVQ	(AX), CX
	0x0044 00068 (<autogenerated>:1)	PCDATA	$0, $3
	0x0044 00068 (<autogenerated>:1)	MOVQ	8(AX), AX
	0x0048 00072 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_5+40(SP)
	0x004d 00077 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_5+48(SP)
	0x0052 00082 (<autogenerated>:1)	CMPQ	""..autotmp_4+64(SP), AX
	0x0057 00087 (<autogenerated>:1)	SETEQ	AL
	0x005a 00090 (<autogenerated>:1)	JEQ	94
	0x005c 00092 (<autogenerated>:1)	JMP	186
	0x005e 00094 (<autogenerated>:1)	PCDATA	$0, $2
	0x005e 00094 (<autogenerated>:1)	MOVQ	""..autotmp_4+56(SP), AX
	0x0063 00099 (<autogenerated>:1)	PCDATA	$0, $3
	0x0063 00099 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x0067 00103 (<autogenerated>:1)	PCDATA	$0, $0
	0x0067 00103 (<autogenerated>:1)	MOVQ	CX, 8(SP)
	0x006c 00108 (<autogenerated>:1)	PCDATA	$1, $0
	0x006c 00108 (<autogenerated>:1)	MOVQ	""..autotmp_4+64(SP), AX
	0x0071 00113 (<autogenerated>:1)	MOVQ	AX, 16(SP)
	0x0076 00118 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x007b 00123 (<autogenerated>:1)	MOVBLZX	24(SP), AX
	0x0080 00128 (<autogenerated>:1)	JMP	130
	0x0082 00130 (<autogenerated>:1)	MOVB	AL, ""..autotmp_3+39(SP)
	0x0086 00134 (<autogenerated>:1)	TESTB	AL, AL
	0x0088 00136 (<autogenerated>:1)	JNE	140
	0x008a 00138 (<autogenerated>:1)	JMP	184
	0x008c 00140 (<autogenerated>:1)	PCDATA	$0, $1
	0x008c 00140 (<autogenerated>:1)	PCDATA	$1, $2
	0x008c 00140 (<autogenerated>:1)	MOVQ	"".p+88(SP), AX
	0x0091 00145 (<autogenerated>:1)	PCDATA	$0, $2
	0x0091 00145 (<autogenerated>:1)	PCDATA	$1, $3
	0x0091 00145 (<autogenerated>:1)	MOVQ	"".q+96(SP), CX
	0x0096 00150 (<autogenerated>:1)	PCDATA	$0, $1
	0x0096 00150 (<autogenerated>:1)	MOVQ	16(CX), CX
	0x009a 00154 (<autogenerated>:1)	PCDATA	$0, $0
	0x009a 00154 (<autogenerated>:1)	CMPQ	16(AX), CX
	0x009e 00158 (<autogenerated>:1)	SETEQ	""..autotmp_3+39(SP)
	0x00a3 00163 (<autogenerated>:1)	JMP	165
	0x00a5 00165 (<autogenerated>:1)	MOVBLZX	""..autotmp_3+39(SP), AX
	0x00aa 00170 (<autogenerated>:1)	MOVB	AL, "".~r2+104(SP)
	0x00ae 00174 (<autogenerated>:1)	MOVQ	72(SP), BP
	0x00b3 00179 (<autogenerated>:1)	ADDQ	$80, SP
	0x00b7 00183 (<autogenerated>:1)	RET
	0x00b8 00184 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00b8 00184 (<autogenerated>:1)	PCDATA	$1, $-2
	0x00b8 00184 (<autogenerated>:1)	JMP	165
	0x00ba 00186 (<autogenerated>:1)	JMP	130
	0x00bc 00188 (<autogenerated>:1)	NOP
	0x00bc 00188 (<autogenerated>:1)	PCDATA	$1, $-1
	0x00bc 00188 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00bc 00188 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00c1 00193 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a9  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 50 48 89 6c 24 48 48 8d 6c 24  ...H..PH.l$HH.l$
	0x0020 48 c6 44 24 68 00 48 8b 44 24 58 48 8b 48 08 48  H.D$h.H.D$XH.H.H
	0x0030 8b 00 48 89 44 24 38 48 89 4c 24 40 48 8b 44 24  ..H.D$8H.L$@H.D$
	0x0040 60 48 8b 08 48 8b 40 08 48 89 4c 24 28 48 89 44  `H..H.@.H.L$(H.D
	0x0050 24 30 48 39 44 24 40 0f 94 c0 74 02 eb 5c 48 8b  $0H9D$@...t..\H.
	0x0060 44 24 38 48 89 04 24 48 89 4c 24 08 48 8b 44 24  D$8H..$H.L$.H.D$
	0x0070 40 48 89 44 24 10 e8 00 00 00 00 0f b6 44 24 18  @H.D$........D$.
	0x0080 eb 00 88 44 24 27 84 c0 75 02 eb 2c 48 8b 44 24  ...D$'..u..,H.D$
	0x0090 58 48 8b 4c 24 60 48 8b 49 10 48 39 48 10 0f 94  XH.L$`H.I.H9H...
	0x00a0 44 24 27 eb 00 0f b6 44 24 27 88 44 24 68 48 8b  D$'....D$'.D$hH.
	0x00b0 6c 24 48 48 83 c4 50 c3 eb eb eb c6 e8 00 00 00  l$HH..P.........
	0x00c0 00 e9 3a ff ff ff                                ..:...
	rel 5+4 t=16 TLS+0
	rel 119+4 t=8 runtime.memequal+0
	rel 189+4 t=8 runtime.morestack_noctxt+0
type..hash."".Person STEXT dupok size=99 args=0x18 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash."".Person(SB), DUPOK|ABIInternal, $40-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	92
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (<autogenerated>:1)	PCDATA	$0, $0
	0x001d 00029 (<autogenerated>:1)	PCDATA	$1, $0
	0x001d 00029 (<autogenerated>:1)	MOVQ	$0, "".~r2+64(SP)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $1
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $1
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+48(SP), AX
	0x002b 00043 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+24(SP)
	0x0030 00048 (<autogenerated>:1)	PCDATA	$0, $0
	0x0030 00048 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x0034 00052 (<autogenerated>:1)	MOVQ	"".h+56(SP), AX
	0x0039 00057 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x003e 00062 (<autogenerated>:1)	CALL	type..hash."".BaseInfo(SB)
	0x0043 00067 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x0048 00072 (<autogenerated>:1)	MOVQ	AX, "".h+56(SP)
	0x004d 00077 (<autogenerated>:1)	MOVQ	AX, "".~r2+64(SP)
	0x0052 00082 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x0057 00087 (<autogenerated>:1)	ADDQ	$40, SP
	0x005b 00091 (<autogenerated>:1)	RET
	0x005c 00092 (<autogenerated>:1)	NOP
	0x005c 00092 (<autogenerated>:1)	PCDATA	$1, $-1
	0x005c 00092 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005c 00092 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0061 00097 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 4d 48  dH..%....H;a.vMH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 44  ..(H.l$ H.l$ H.D
	0x0020 24 40 00 00 00 00 48 8b 44 24 30 48 89 44 24 18  $@....H.D$0H.D$.
	0x0030 48 89 04 24 48 8b 44 24 38 48 89 44 24 08 e8 00  H..$H.D$8H.D$...
	0x0040 00 00 00 48 8b 44 24 10 48 89 44 24 38 48 89 44  ...H.D$.H.D$8H.D
	0x0050 24 40 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00  $@H.l$ H..(.....
	0x0060 00 eb 9d                                         ...
	rel 5+4 t=16 TLS+0
	rel 63+4 t=8 type..hash."".BaseInfo+0
	rel 93+4 t=8 runtime.morestack_noctxt+0
type..eq."".Person STEXT dupok size=189 args=0x18 locals=0x48
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq."".Person(SB), DUPOK|ABIInternal, $72-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	179
	0x0013 00019 (<autogenerated>:1)	SUBQ	$72, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 64(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	64(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·7e7fcb5c7cd183fbe200fb26b1d44a90(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·91d432edea9e3c468c5aec7a805d99d2(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$1, $0
	0x0021 00033 (<autogenerated>:1)	MOVB	$0, "".~r2+96(SP)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $1
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+80(SP), AX
	0x002b 00043 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x002f 00047 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$1, $1
	0x0032 00050 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+48(SP)
	0x0037 00055 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_3+56(SP)
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $1
	0x003c 00060 (<autogenerated>:1)	MOVQ	"".q+88(SP), AX
	0x0041 00065 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x0045 00069 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0048 00072 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+32(SP)
	0x004d 00077 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_4+40(SP)
	0x0052 00082 (<autogenerated>:1)	CMPQ	""..autotmp_3+56(SP), CX
	0x0057 00087 (<autogenerated>:1)	SETEQ	CL
	0x005a 00090 (<autogenerated>:1)	JEQ	94
	0x005c 00092 (<autogenerated>:1)	JMP	175
	0x005e 00094 (<autogenerated>:1)	PCDATA	$0, $2
	0x005e 00094 (<autogenerated>:1)	MOVQ	""..autotmp_3+48(SP), CX
	0x0063 00099 (<autogenerated>:1)	PCDATA	$0, $1
	0x0063 00099 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x0067 00103 (<autogenerated>:1)	PCDATA	$0, $0
	0x0067 00103 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x006c 00108 (<autogenerated>:1)	PCDATA	$1, $0
	0x006c 00108 (<autogenerated>:1)	MOVQ	""..autotmp_3+56(SP), AX
	0x0071 00113 (<autogenerated>:1)	MOVQ	AX, 16(SP)
	0x0076 00118 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x007b 00123 (<autogenerated>:1)	MOVBLZX	24(SP), AX
	0x0080 00128 (<autogenerated>:1)	JMP	130
	0x0082 00130 (<autogenerated>:1)	TESTB	AL, AL
	0x0084 00132 (<autogenerated>:1)	JNE	136
	0x0086 00134 (<autogenerated>:1)	JMP	173
	0x0088 00136 (<autogenerated>:1)	PCDATA	$0, $1
	0x0088 00136 (<autogenerated>:1)	PCDATA	$1, $2
	0x0088 00136 (<autogenerated>:1)	MOVQ	"".p+80(SP), AX
	0x008d 00141 (<autogenerated>:1)	PCDATA	$0, $2
	0x008d 00141 (<autogenerated>:1)	PCDATA	$1, $3
	0x008d 00141 (<autogenerated>:1)	MOVQ	"".q+88(SP), CX
	0x0092 00146 (<autogenerated>:1)	PCDATA	$0, $1
	0x0092 00146 (<autogenerated>:1)	MOVQ	16(CX), CX
	0x0096 00150 (<autogenerated>:1)	PCDATA	$0, $0
	0x0096 00150 (<autogenerated>:1)	CMPQ	16(AX), CX
	0x009a 00154 (<autogenerated>:1)	SETEQ	AL
	0x009d 00157 (<autogenerated>:1)	JMP	159
	0x009f 00159 (<autogenerated>:1)	MOVB	AL, "".~r2+96(SP)
	0x00a3 00163 (<autogenerated>:1)	MOVQ	64(SP), BP
	0x00a8 00168 (<autogenerated>:1)	ADDQ	$72, SP
	0x00ac 00172 (<autogenerated>:1)	RET
	0x00ad 00173 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00ad 00173 (<autogenerated>:1)	PCDATA	$1, $-2
	0x00ad 00173 (<autogenerated>:1)	JMP	159
	0x00af 00175 (<autogenerated>:1)	PCDATA	$0, $0
	0x00af 00175 (<autogenerated>:1)	PCDATA	$1, $0
	0x00af 00175 (<autogenerated>:1)	MOVL	CX, AX
	0x00b1 00177 (<autogenerated>:1)	JMP	130
	0x00b3 00179 (<autogenerated>:1)	NOP
	0x00b3 00179 (<autogenerated>:1)	PCDATA	$1, $-1
	0x00b3 00179 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00b3 00179 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00b8 00184 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a0  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 48 48 89 6c 24 40 48 8d 6c 24  ...H..HH.l$@H.l$
	0x0020 40 c6 44 24 60 00 48 8b 44 24 50 48 8b 48 08 48  @.D$`.H.D$PH.H.H
	0x0030 8b 00 48 89 44 24 30 48 89 4c 24 38 48 8b 44 24  ..H.D$0H.L$8H.D$
	0x0040 58 48 8b 48 08 48 8b 00 48 89 44 24 20 48 89 4c  XH.H.H..H.D$ H.L
	0x0050 24 28 48 39 4c 24 38 0f 94 c1 74 02 eb 51 48 8b  $(H9L$8...t..QH.
	0x0060 4c 24 30 48 89 0c 24 48 89 44 24 08 48 8b 44 24  L$0H..$H.D$.H.D$
	0x0070 38 48 89 44 24 10 e8 00 00 00 00 0f b6 44 24 18  8H.D$........D$.
	0x0080 eb 00 84 c0 75 02 eb 25 48 8b 44 24 50 48 8b 4c  ....u..%H.D$PH.L
	0x0090 24 58 48 8b 49 10 48 39 48 10 0f 94 c0 eb 00 88  $XH.I.H9H.......
	0x00a0 44 24 60 48 8b 6c 24 40 48 83 c4 48 c3 eb f0 89  D$`H.l$@H..H....
	0x00b0 c8 eb cf e8 00 00 00 00 e9 43 ff ff ff           .........C...
	rel 5+4 t=16 TLS+0
	rel 119+4 t=8 runtime.memequal+0
	rel 180+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFINFO dupok size=0
	0x0000 2d 4e 20 2d 6c                                   -N -l
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+58
	rel 27+4 t=29 gofile../Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 1f 00                             .......
go.string."%#v" SRODATA dupok size=3
	0x0000 25 23 76                                         %#v
go.loc."".f1 SDWARFLOC size=0
go.info."".f1 SDWARFINFO size=31
	0x0000 03 22 22 2e 66 31 00 00 00 00 00 00 00 00 00 00  ."".f1..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 00     ...............
	rel 7+8 t=1 "".f1+0
	rel 15+8 t=1 "".f1+318
	rel 25+4 t=29 gofile../Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go+0
go.range."".f1 SDWARFRANGE size=0
go.isstmt."".f1 SDWARFMISC size=0
	0x0000 04 18 04 17 03 03 01 22 02 03 01 26 02 05 01 83  ......."...&....
	0x0010 01 02 15 01 1a 02 0a 00                          ........
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
go.loc."".f2 SDWARFLOC size=0
go.info."".f2 SDWARFINFO size=42
	0x0000 03 22 22 2e 66 32 00 00 00 00 00 00 00 00 00 00  ."".f2..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0a 62  ...............b
	0x0020 00 18 00 00 00 00 02 91 40 00                    ........@.
	rel 7+8 t=1 "".f2+0
	rel 15+8 t=1 "".f2+347
	rel 25+4 t=29 gofile../Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go+0
	rel 34+4 t=28 go.info."".BaseInfo+0
go.range."".f2 SDWARFRANGE size=0
go.isstmt."".f2 SDWARFMISC size=0
	0x0000 04 18 04 17 03 03 01 0e 02 03 01 15 02 0b 01 02  ................
	0x0010 02 03 01 26 02 05 01 8f 01 02 15 01 07 02 07 01  ...&............
	0x0020 0c 02 0a 00                                      ....
go.loc."".f3 SDWARFLOC size=0
go.info."".f3 SDWARFINFO size=43
	0x0000 03 22 22 2e 66 33 00 00 00 00 00 00 00 00 00 00  ."".f3..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0a 62  ...............b
	0x0020 00 1e 00 00 00 00 03 91 90 7f 00                 ...........
	rel 7+8 t=1 "".f3+0
	rel 15+8 t=1 "".f3+381
	rel 25+4 t=29 gofile../Users/junjie2/data1/github.com/phpor/go-example/app/asm/main.go+0
	rel 34+4 t=28 go.info.*"".BaseInfo+0
go.range."".f3 SDWARFRANGE size=0
go.isstmt."".f3 SDWARFMISC size=0
	0x0000 04 18 04 17 03 03 01 21 02 02 01 3d 02 03 01 26  .......!...=...&
	0x0010 02 05 01 86 01 02 15 01 18 02 0a 00              ............
go.loc.type..hash."".BaseInfo SDWARFLOC dupok size=0
go.info.type..hash."".BaseInfo SDWARFINFO dupok size=85
	0x0000 03 74 79 70 65 2e 2e 68 61 73 68 2e 22 22 2e 42  .type..hash."".B
	0x0010 61 73 65 49 6e 66 6f 00 00 00 00 00 00 00 00 00  aseInfo.........
	0x0020 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f  ................
	0x0030 70 00 00 01 00 00 00 00 01 9c 0f 68 00 00 01 00  p..........h....
	0x0040 00 00 00 02 91 08 0f 7e 72 32 00 01 01 00 00 00  .......~r2......
	0x0050 00 02 91 10 00                                   .....
	rel 24+8 t=1 type..hash."".BaseInfo+0
	rel 32+8 t=1 type..hash."".BaseInfo+158
	rel 42+4 t=29 gofile..<autogenerated>+0
	rel 52+4 t=28 go.info.*"".BaseInfo+0
	rel 63+4 t=28 go.info.uintptr+0
	rel 77+4 t=28 go.info.uintptr+0
go.range.type..hash."".BaseInfo SDWARFRANGE dupok size=0
go.isstmt.type..hash."".BaseInfo SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 09 01 18 02 05 01 2f 02 05 01 19  .........../....
	0x0010 02 0a 00                                         ...
go.loc.type..eq."".BaseInfo SDWARFLOC dupok size=0
go.info.type..eq."".BaseInfo SDWARFINFO dupok size=83
	0x0000 03 74 79 70 65 2e 2e 65 71 2e 22 22 2e 42 61 73  .type..eq."".Bas
	0x0010 65 49 6e 66 6f 00 00 00 00 00 00 00 00 00 00 00  eInfo...........
	0x0020 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f 70 00  ..............p.
	0x0030 00 01 00 00 00 00 01 9c 0f 71 00 00 01 00 00 00  .........q......
	0x0040 00 02 91 08 0f 7e 72 32 00 01 01 00 00 00 00 02  .....~r2........
	0x0050 91 10 00                                         ...
	rel 22+8 t=1 type..eq."".BaseInfo+0
	rel 30+8 t=1 type..eq."".BaseInfo+198
	rel 40+4 t=29 gofile..<autogenerated>+0
	rel 50+4 t=28 go.info.*"".BaseInfo+0
	rel 61+4 t=28 go.info.*"".BaseInfo+0
	rel 75+4 t=28 go.info.bool+0
go.range.type..eq."".BaseInfo SDWARFRANGE dupok size=0
go.isstmt.type..eq."".BaseInfo SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 50 02 05 01 2a 02 05 01 12  .......P...*....
	0x0010 02 0a 00                                         ...
type..hashfunc."".BaseInfo SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash."".BaseInfo+0
type..eqfunc."".BaseInfo SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".BaseInfo+0
type..alg."".BaseInfo SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc."".BaseInfo+0
	rel 8+8 t=1 type..eqfunc."".BaseInfo+0
type..namedata.*main.BaseInfo. SRODATA dupok size=17
	0x0000 01 00 0e 2a 6d 61 69 6e 2e 42 61 73 65 49 6e 66  ...*main.BaseInf
	0x0010 6f                                               o
type.*"".BaseInfo SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2e 76 aa db 00 08 08 36 00 00 00 00 00 00 00 00  .v.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.BaseInfo.+0
	rel 48+8 t=1 type."".BaseInfo+0
type..importpath."". SRODATA dupok size=7
	0x0000 00 00 04 6d 61 69 6e                             ...main
type..namedata.name- SRODATA dupok size=7
	0x0000 00 00 04 6e 61 6d 65                             ...name
type..namedata.age- SRODATA dupok size=6
	0x0000 00 00 03 61 67 65                                ...age
type."".BaseInfo SRODATA size=144
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9d 86 0e 96 07 08 08 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 24+8 t=1 type..alg."".BaseInfo+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.BaseInfo.+0
	rel 44+4 t=5 type.*"".BaseInfo+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".BaseInfo+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.name-+0
	rel 104+8 t=1 type.string+0
	rel 120+8 t=1 type..namedata.age-+0
	rel 128+8 t=1 type.int+0
go.string..gostring.114.846015892bbcbd04658e6b9c1b05d4c6f13b84b69e861ac13302486d79b57703 SRODATA dupok size=114
	0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
	0x0010 70 61 74 68 09 67 69 74 68 75 62 2e 63 6f 6d 2f  path.github.com/
	0x0020 70 68 70 6f 72 2f 67 6f 2d 65 78 61 6d 70 6c 65  phpor/go-example
	0x0030 2f 61 70 70 2f 61 73 6d 0a 6d 6f 64 09 67 69 74  /app/asm.mod.git
	0x0040 68 75 62 2e 63 6f 6d 2f 70 68 70 6f 72 2f 67 6f  hub.com/phpor/go
	0x0050 2d 65 78 61 6d 70 6c 65 09 28 64 65 76 65 6c 29  -example.(devel)
	0x0060 09 0a f9 32 43 31 86 18 20 72 00 82 42 10 41 16  ...2C1.. r..B.A.
	0x0070 d8 f2                                            ..
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
"".p SBSS size=24
runtime.modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 72 00 00 00 00 00 00 00  ........r.......
	rel 0+8 t=1 go.string..gostring.114.846015892bbcbd04658e6b9c1b05d4c6f13b84b69e861ac13302486d79b57703+0
go.loc.type..hash."".Person SDWARFLOC dupok size=0
go.info.type..hash."".Person SDWARFINFO dupok size=83
	0x0000 03 74 79 70 65 2e 2e 68 61 73 68 2e 22 22 2e 50  .type..hash."".P
	0x0010 65 72 73 6f 6e 00 00 00 00 00 00 00 00 00 00 00  erson...........
	0x0020 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f 70 00  ..............p.
	0x0030 00 01 00 00 00 00 01 9c 0f 68 00 00 01 00 00 00  .........h......
	0x0040 00 02 91 08 0f 7e 72 32 00 01 01 00 00 00 00 02  .....~r2........
	0x0050 91 10 00                                         ...
	rel 22+8 t=1 type..hash."".Person+0
	rel 30+8 t=1 type..hash."".Person+99
	rel 40+4 t=29 gofile..<autogenerated>+0
	rel 50+4 t=28 go.info.*"".Person+0
	rel 61+4 t=28 go.info.uintptr+0
	rel 75+4 t=28 go.info.uintptr+0
go.range.type..hash."".Person SDWARFRANGE dupok size=0
go.isstmt.type..hash."".Person SDWARFMISC dupok size=0
	0x0000 04 0f 04 0e 03 09 01 18 02 05 01 19 02 07 00     ...............
go.loc.type..eq."".Person SDWARFLOC dupok size=0
go.info.type..eq."".Person SDWARFINFO dupok size=81
	0x0000 03 74 79 70 65 2e 2e 65 71 2e 22 22 2e 50 65 72  .type..eq."".Per
	0x0010 73 6f 6e 00 00 00 00 00 00 00 00 00 00 00 00 00  son.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 0f 70 00 00 01  ............p...
	0x0030 00 00 00 00 01 9c 0f 71 00 00 01 00 00 00 00 02  .......q........
	0x0040 91 08 0f 7e 72 32 00 01 01 00 00 00 00 02 91 10  ...~r2..........
	0x0050 00                                               .
	rel 20+8 t=1 type..eq."".Person+0
	rel 28+8 t=1 type..eq."".Person+189
	rel 38+4 t=29 gofile..<autogenerated>+0
	rel 48+4 t=28 go.info.*"".Person+0
	rel 59+4 t=28 go.info.*"".Person+0
	rel 73+4 t=28 go.info.bool+0
go.range.type..eq."".Person SDWARFRANGE dupok size=0
go.isstmt.type..eq."".Person SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 50 02 05 01 07 02 02 01 1b  .......P........
	0x0010 02 04 01 0a 02 02 01 04 02 0a 00                 ...........
type..hashfunc."".Person SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash."".Person+0
type..eqfunc."".Person SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".Person+0
type..alg."".Person SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc."".Person+0
	rel 8+8 t=1 type..eqfunc."".Person+0
type..namedata.*main.Person. SRODATA dupok size=15
	0x0000 01 00 0c 2a 6d 61 69 6e 2e 50 65 72 73 6f 6e     ...*main.Person
type.*"".Person SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 91 18 70 14 00 08 08 36 00 00 00 00 00 00 00 00  ..p....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Person.+0
	rel 48+8 t=1 type."".Person+0
type..namedata.Base. SRODATA dupok size=7
	0x0000 01 00 04 42 61 73 65                             ...Base
type."".Person SRODATA size=120
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fd cf 91 06 07 08 08 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg."".Person+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Person.+0
	rel 44+4 t=5 type.*"".Person+0
	rel 56+8 t=1 type."".Person+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.Base.+0
	rel 104+8 t=1 type."".BaseInfo+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·94d172d993ad121e3746e9b06f1f5579 SRODATA dupok size=12
	0x0000 04 00 00 00 08 00 00 00 00 10 01 20              ........... 
gclocals·33b901baab2acec3083d16f1ab81c65a SRODATA dupok size=12
	0x0000 04 00 00 00 07 00 00 00 00 01 05 45              ...........E
"".f1.stkobj SRODATA dupok size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals·4b62fca614ba1b83331bcc57182ac46b SRODATA dupok size=16
	0x0000 04 00 00 00 0b 00 00 00 00 00 10 00 01 00 00 01  ................
"".f2.stkobj SRODATA dupok size=24
	0x0000 01 00 00 00 00 00 00 00 c0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals·3e27b3aa6b89137cce48b3379a2a6610 SRODATA dupok size=8
	0x0000 05 00 00 00 00 00 00 00                          ........
gclocals·9cf841ba61ca75101c532af3ea6af54c SRODATA dupok size=18
	0x0000 05 00 00 00 0c 00 00 00 00 00 00 02 20 00 02 00  ............ ...
	0x0010 40 00                                            @.
gclocals·fd8cf83e1c48dd0b84c741be7b1e4c9c SRODATA dupok size=13
	0x0000 05 00 00 00 07 00 00 00 00 02 01 05 45           ............E
"".f3.stkobj SRODATA dupok size=40
	0x0000 02 00 00 00 00 00 00 00 c0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00 e8 ff ff ff ff ff ff ff  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
	rel 32+8 t=1 type."".BaseInfo+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·7e7fcb5c7cd183fbe200fb26b1d44a90 SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 03 03 02 00              ............
gclocals·91d432edea9e3c468c5aec7a805d99d2 SRODATA dupok size=12
	0x0000 04 00 00 00 04 00 00 00 00 04 00 00              ............
gclocals·6e8d7ea4abad763909b26991048ee1fe SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 00 01 03 02              ............
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
note: module requires Go 1.14
