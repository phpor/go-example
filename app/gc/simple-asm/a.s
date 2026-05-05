
simple-asm:     file format elf64-x86-64


Disassembly of section .text:

00000000004526d0 <main.(*BigStruct).reset2>:
	b.name = ""
	b.addr = ""
	b.age = 0
}

func (b *BigStruct)reset2() {
  4526d0:	64 48 8b 0c 25 f8 ff 	mov    %fs:0xfffffffffffffff8,%rcx
  4526d7:	ff ff 
  4526d9:	48 3b 61 10          	cmp    0x10(%rcx),%rsp
  4526dd:	76 58                	jbe    452737 <main.(*BigStruct).reset2+0x67>
  4526df:	48 83 ec 18          	sub    $0x18,%rsp
  4526e3:	48 89 6c 24 10       	mov    %rbp,0x10(%rsp)
  4526e8:	48 8d 6c 24 10       	lea    0x10(%rsp),%rbp
	*b = BigStruct{}
  4526ed:	48 8b 44 24 20       	mov    0x20(%rsp),%rax
  4526f2:	84 00                	test   %al,(%rax)
  4526f4:	83 3d a5 d9 08 00 00 	cmpl   $0x0,0x8d9a5(%rip)        # 4e00a0 <runtime.writeBarrier>
  4526fb:	74 02                	je     4526ff <main.(*BigStruct).reset2+0x2f>
  4526fd:	eb 21                	jmp    452720 <main.(*BigStruct).reset2+0x50>
  4526ff:	48 c7 00 00 00 00 00 	movq   $0x0,(%rax)
  452706:	0f 57 c0             	xorps  %xmm0,%xmm0
  452709:	0f 11 40 08          	movups %xmm0,0x8(%rax)
  45270d:	0f 57 c0             	xorps  %xmm0,%xmm0
  452710:	0f 11 40 18          	movups %xmm0,0x18(%rax)
  452714:	eb 00                	jmp    452716 <main.(*BigStruct).reset2+0x46>
  452716:	48 8b 6c 24 10       	mov    0x10(%rsp),%rbp
  45271b:	48 83 c4 18          	add    $0x18,%rsp
  45271f:	c3                   	retq   
	*b = BigStruct{}
  452720:	48 8d 0d 59 5d 01 00 	lea    0x15d59(%rip),%rcx        # 468480 <type.*+0x15480>
  452727:	48 89 0c 24          	mov    %rcx,(%rsp)
  45272b:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  452730:	e8 7b b6 fb ff       	callq  40ddb0 <runtime.typedmemclr>
  452735:	eb df                	jmp    452716 <main.(*BigStruct).reset2+0x46>
func (b *BigStruct)reset2() {
  452737:	e8 b4 79 ff ff       	callq  44a0f0 <runtime.morestack_noctxt>
  45273c:	eb 92                	jmp    4526d0 <main.(*BigStruct).reset2>
