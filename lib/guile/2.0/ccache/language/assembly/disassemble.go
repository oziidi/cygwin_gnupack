GOOF----LE-4-2.0å0      ] ¢ 4    h[      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  assembly¤	g  disassemble¤		 ¤	
g  filenameS¤	f  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm¤	g  importsS¤	g  ice-9¤	g  format¤	 ¤	 ¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	g  system¤	g  vm¤	g  instruction¤	 ¤	 ¤	g  program¤	 ¤	 ¤	g  base¤	g  pmatch¤	 ¤	  ¤	! ¤	"! ¤	#g  compile¤	$# ¤	%$ ¤	& "% ¤	'g  exportsS¤	( ¤	)g  set-current-module¤	*) ¤	+) ¤	,f  Disassembly of ~A:

¤	-g  disassemble-load-program¤	.g  	decompile¤	/g  fromS¤	0g  value¤	1g  toS¤	2g  load-program¤	3g  assq-ref¤	4g  objects¤	5g  	free-vars¤	6g  meta¤	7g  blocs¤	8g  sources¤	9g  newline¤	:g  for-each¤	;f  Embedded program ~A:

¤	<g  reverse!¤	=g  byte-length¤	>g  nop¤	?g  
print-info¤	@g  code-annotation¤	Ag  and=>¤	Bg  assq¤	Cg  source->string¤	Dg  gensym¤	Ef   ¤	Fg  disassemble-free-vars¤	Gg  disassemble-meta¤	Hg  program?¤	Ig  display¤	Jf  )----------------------------------------
¤	Kg  vector->list¤	Lg  error¤	Mf  bad load-program form¤	Nf  Free variables:

¤	Og  fold¤	Pg  make-syntax-transformer¤	QP ¤	RP ¤	Sg  unless¤	Tg  macro¤	Ug  $sc-dispatch¤	VU ¤	WU ¤	Xg  _¤	Yg  any¤	ZXY¤	[g  syntax->datum¤	\[ ¤	][ ¤	^g  datum->syntax¤	_^ ¤	`^ ¤	ag  if¤	bg  not¤	cg  begin¤	dg  syntax-violation¤	ed ¤	fd ¤	gf  -source expression failed to match any pattern¤	hg  name¤	ih ¤	jg  *uninteresting-props*¤	kg  filter¤	lg  memq¤	mf  Properties:

¤	nf  ~a:~a:~a¤	og  source:file¤	pf  (unknown file)¤	qg  source:line-for-user¤	rg  source:column¤	sg  
make-int16¤	tg  assembly-unpack¤	ug  list¤	vg  vector¤	wf  ~a element~:p¤	xg  br¤	yg  br-if¤	zg  br-if-eq¤	{g  	br-if-not¤	|g  br-if-not-eq¤	}g  br-if-not-null¤	~g  
br-if-null¤	f  -> ~A¤ g  br-if-nargs-ne¤ g  br-if-nargs-lt¤ g  br-if-nargs-gt¤ g  bind-optionals/shuffle-or-br¤ g  	last-pair¤ g  
object-ref¤ f  ~s¤ g  	local-ref¤ g  local-boxed-ref¤ g  	local-set¤ g  local-boxed-set¤ g  binding:start¤ g  binding:end¤ f  `~a'~@[ (arg)~]¤ g  binding:name¤ g  binding:index¤ g  list-ref¤ g  assert-nargs-ee/locals¤ g  assert-nargs-ge/locals¤ f  ~a arg~:p, ~a local~:p¤ g  free-ref¤ g  free-boxed-ref¤ g  free-boxed-set¤ f  (closure variable)¤ g  toplevel-ref¤ g  toplevel-set¤ f  `~s'¤ g  	variable?¤ g  mv-call¤ f  MV -> ~A¤ g  prompt¤ f  H -> ~A¤  g  assembly->object¤ ¡f  ,~4@S    ~32S~@[;; ~1{~@?~}~]~@[~61t at ~a~]
¤C 5hÈ#  û   ]4	
&'(5 4+ >  "  G   ,-./01       h   °   ] L 6¨       g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	!		
	!	%		!	1		!	 		
   C     h0   ß   ]4 >  "  G  4 O >   6<       ×       g  x
		)  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	
											 	 		)  g  nameg  disassemble CR23456789:;- h    æ   ]4 >  "  G   6Þ       g  sym+asm
		   g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	1			2			2			2	4		2			3	(		3	6	 	3	 			    C<=>?@ABCDEFGHIJ      h(   Ø   ]4 5$  4>  "  G   6C Ð       g  x
		'  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	S			T			T			U			U	"		U		%	V	 		'   CKLM       hP  4  ] $ 9  & $ $ ÷$ ß$  45"  $  45"  $  45"  $  45"  $  45"  
	" (  -4>   "  G  4	
45>  " éG  " â45" I$  é&  x(  45"ÿÿ|44	ÿ54$  45"  5>  "  G  45"ÿÿ#44	ÿ54$  45"  5>  "  G  45"ÿþÊ44	ÿ54$  45"  5>  "  G  45"ÿþk$  ^&  F454 >  "  G  45"ÿþ"ÿþU"ÿþN
"ÿýë	$  4	>  "  G  "   
$  4
>  "  G  "   $  	456C 6 6 6 6 6      ,      g  asm
	J g  env	J g  vx		B g  vy		B g  vx		(2 g  vy		(2 g  vy		6* g  vy		F" g  objs	 °" g  	free-vars		 °" g  meta	
 °" g  blocs	 °" g  srcs	 °" g  pos	 ¾Å g  code	 ¾Å g  programs	 ¾Å g  asm	 ôÅ g  len	 ýÂ g  end	¿ g  vx	ú g  vy	ú g  vx	]¸ g  sym	r±  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	$
		%		R	'		S	'		Y	'	)	[	'		f	(		g	(		m	(	-	o	(		z	)		{	)	 	)	) 	)	 	*	 	*	 	*	) 	*	 ¢	+	 £	+	 ©	+	) «	+	 °	'	 ¾	,	 Ä	-		 Å	/	 Õ	0	 Ú	4	 å	0	 ô	6	 ô	6	 ÷	7	 ý	6		8		6	
	9	,	@	3	@	6	@	.B	@	C	B	J	E	[	G	c	G	#d	G	-u	G	z	B		H		H		H	.	H		B	£	E	´	G	¼	G	#½	G	-Î	G	Ó	B	Þ	H	å	H	è	H	.ô	H	û	B		E		G		G	#	G	--	G	2	B	=	H	D	H	G	H	.S	H	S	9	l	;	p	;	#r	;	r	;	u	<	{	<	"	<		=		=		=	0¡	>	®	=	µ	9	Å	,	É	,	.Ó	,	Ö	J	Ú	J	Û	K	ö	L	÷	M		Q		W		W	 	R	&	Y	*	Y	.	Y	2	Y	6	Y	:	Y	>	Y	B	Y	F	Y	J	Y	 n	J	  g  nameg  disassemble-load-program C-RINO?    h    Î   ]4 >  "  G  C    Æ       g  free-var
		 g  i		  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	]			^	
		_	
 			   C       h    ñ   ]4>  "  G  
 6   é       g  	free-vars
		  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	[
		\			\			\			]	 		  g  nameg  disassemble-free-vars CFR4RSTWZ]`abc h    Å   - 1 3    C     ½       g  test
			 g  body			  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	c
		d	 			
   C        h   ¹   ]	4 5L 4?6±       g  args
		 g  v			  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm		c
 		   Cfg        h(   ·   ]	4 5$   O @ 6 ¯       g  y
		' g  tmp		'  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	c
 		'   C5SRijRklj    h   È   ]4 5C  À       g  x
		  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	i			j			j	$		j			j	 		   CIm:?    h   °   ] 6     ¨       g  x
		  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	n			n	 		   C9   hH   )  ]	4 5(  C4>  "  G  4>  "  G  6       !      g  meta
		B g  props		B  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	h
		i		
	k			i			i			l			m			m		"	m		+	n		B	o	 		B  g  nameg  disassemble-meta CGRnopqr    h8   
  ]	4 5$  "  4 54 56             g  src
		1 g  t		!  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	q
		r			r			r			r	.	"	s	
	)	s	%	1	r	 			1  g  nameg  source->string CCR h   å   ]  C    Ý       g  byte1
		 g  byte2		  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	u
		v			v	 			  g  nameg  
make-int16 CsRtuvwsxyz{|}~3A  h   ½   ]  C      µ       g  obj
		
  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
 ¤		 ¤	"		 ¤	 		
   C 2       h°    ].45$  "  $  4? C$  "  K$  "  =$  "  /	$  "  !
$  "  $  "  $  45 C$  "  $  "  $  45 C$  4455 C$  $  £ CC$  "  !$  "  $  "  $  s$  k"  W	$  N	
"  			"ÿÿç4
5 $  *4	5 $  4
54
5 C"ÿÿ¿"ÿÿ»C45	"ÿÿC $  "  !$  "	 	ý C#$  "  $$  "  %$  & C'$  "  ($  ?$  7£	"  )	 C4*	5$  	$  		 C"ÿÿÙ"ÿÿÕC+$  ,45 C-$  .45 C/40516             g  end-addr
	© g  code	© g  objs		© g  nargs		© g  blocs		© g  labels		© g  code			© g  inst		© g  args		© g  bindings		U¬ g  b	
_ª g  v		5f  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
	x
		y				y			z			y			{			y			|		.	~		/	~		8	~		B	|	  	  	   	& ¢ 	 ¥ 	 ¯	|	 Í 	 Î 	 Õ 	& Ù 	 Ü 	 æ	|	 è 	 é 	 î 	+ õ 	& ÷ 	 ú 		|	
 	 	 	- 	 	!	|	Q 	U 	X 	\ 	_ 	_ 	h 	n 	n 	o 	!s 	0u 	!x 	| 	} 	" 	/ 	" 	 	 	 	 	3 	0¡ 	¬ 	­ 	´ 	/¶ 	¼ 	Ç	|	× 	Ú 	Ý 	à 	*ã 	%æ 	ð	|	 	 		|	/ 	4 	&5 	5 	= 	B 	C 	D 	N 	Q 	%U 	W 	Z 	] 	q	|	s 	t 	{ 	)~ 	 		|	 ¡	 ¡	 ¡	( ¡	 ¡	 £	© £	 j	©	  g  nameg  code-annotation C@R¡  h   
  ] 6             g  addr
		 g  info		 g  extra			 g  src			  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm
 §
	 ¨		 ¨	 			  g  nameg  
print-info C?RC ó       g  m
		,  g  filenamef  e/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/disassemble.scm		
2	
/	$
\	[
F	f	I	f
	h
õ	q
ø	u
"	x
#Å §
 	#Ç
   C6 