GOOF----LE-4-2.0Z*      ] \ 4  h¯      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  assembly¤	g  decompile-bytecode¤		 ¤	
g  filenameS¤	f  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm¤	g  importsS¤	g  system¤	g  vm¤	g  instruction¤	 ¤	 ¤	g  base¤	g  pmatch¤	 ¤	 ¤	g  srfi¤	g  srfi-4¤	 ¤	 ¤	g  rnrs¤	g  bytevectors¤	 ¤	 ¤	 ¤	 ¤	 g  objcode¤	!  ¤	"g  selectS¤	#g  
byte-order¤	$# ¤	%!"$ ¤	&% ¤	'g  exportsS¤	( ¤	)g  set-current-module¤	*) ¤	+) ¤	,g  u8vector-length¤	-g  decode-load-program¤	.g  error¤	/f  tried to decode too many bytes¤	0f  -bad bytecode: only decoded ~a out of ~a bytes¤	1g  memq¤	2g  br¤	3g  br-if¤	4g  	br-if-not¤	5g  br-if-eq¤	6g  br-if-not-eq¤	7g  
br-if-null¤	8g  br-if-not-null¤	92345678 ¤	:g  br-instruction?¤	;g  br-if-nargs-ne¤	<g  br-if-nargs-lt¤	=g  br-if-nargs-gt¤	>g  br-if-nargs-lt/non-kw¤	?;<=> ¤	@g  br-nargs-instruction?¤	Ag  
bytes->s24¤	Bf  -error decoding program -- read too many bytes¤	Cg  load-program¤	Dg  map¤	Eg  reverse¤	Fg  reverse!¤	Gg  decode-bytecode¤	Hg  prompt¤	Ig  assv-ref¤	Jg  gensym¤	Kf  :L¤	Lg  mv-call¤	Mg  bind-optionals/shuffle-or-br¤	Ng  and=>¤	Og  opcode->instruction¤	Pg  instruction-length¤	Qg  
load-array¤	Rg  load-wide-string¤	Sg  make-bytevector¤	Tg  make-string¤	Ug  bytevector-u8-set!¤	V	U ¤	W	U ¤	Xg  string-set!¤	Yg  integer->char¤	Zg  utf32->string¤	[g  native-endianness¤C 5    hH"  à   ]4	
&'(5 4+ >  "  G   ,-./       h@     ]ML $  	LM¶"  ML $  "  45  $  	MN"    C         g  b
	)	=  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
				 			 			 	!		!			 		#	"		'	"	"	)	"		)	 		1	#		4	#		6	#	 		=
  g  nameg  pop C.0  h@   j  ]
4 5H4 O 5J$  DJ6       b      g  x
		9 g  env		9 g  opts			9 g  i		
	9 g  size		
	9 g  ret			9  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	
				
				%			%		%	&	
	)	&		/	'	
	3	(		9	(	
 		9	  g  nameg  decompile-bytecode CR19   h   Ò   ] 6      Ê       g  x
		
  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	*
	
	+	 		
  g  nameg  br-instruction? C:R1?   h   Ø   ] 6      Ð       g  x
		
  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	,
	
	-	 		
  g  nameg  br-nargs-instruction? C@R   h8   B  ] 		        
$  C       C    :      g  a
		4 g  b		4 g  c			4 g  x			4  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	/
		0			0			0			0			1			1			1		$	1		2	3		3	3	 		4	  g  nameg  
bytes->s24 CAR.BCDh   Ç   ]  C      ¿       g  x
		
  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	N	"		N	4		N	<			N	. 		
   CE-FG./ h@   #  ]ML $  	4L5 "  ML $  "  45  $  	MN"    C         g  b
	)	=  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	D			E			E			E	 		F			E		#	G		'	G	"	)	G		)	E		1	H		4	H		6	H	 		=
  g  nameg  sub-pop CHAIJKLM@:        h¸  1  , 3 %4 5 4 5 4 5 4 5 4 5 4 5 4 5 4 5 							
H
H" <J	$  6J	$  244J55	

$  "  4 545C4	
	 O 5" " (" Ý" $  ó&  Ö$  Ä$  ¨$  $  p(  UJ454J5$  "  45JK "ÿþÚ"ÿþÏ"ÿþÄ"ÿþ¹"ÿþ®"ÿþ£"ÿþ"ÿþ$  É&  ³$  ¨$  $  ~$  i(  UJ454J5$  "  45JK "ÿýÕ"ÿþ?"ÿþ;"ÿþ7"ÿþ3"ÿþ/"ÿþ+"ÿþ'$ <& &$ $ $  ñ$  Ü$  Ç$  ²$  $    $  s  "!"(  _J4!5#4J#5$$$  $"  45$#$JK$$$# "ÿü¢"ÿý"ÿý"ÿý "ÿüü"ÿüø"ÿüô"ÿüð"ÿüì"ÿüè"ÿüä"ÿüà"ÿüÜ$  ã$  Î$  ¹$  ¤$  $  z(  f45$  WJ454J5$  "  45JK "ÿû¤"ÿü"ÿü"ÿûþ"ÿûú"ÿûö"ÿûò"ÿûî"ÿûê$  µ$   $  $  v(  b45$  SJ454J5$  "  45JK "ÿúÜ"ÿû:"ÿû6"ÿû2"ÿû."ÿû*"ÿû&"ÿú½ )      g  pop
	· g  a	· g  b		· g  c		· g  d		 · g  e		'· g  f		.· g  g		5· g  h		<· g  len			R· g  metalen	
	h· g  labels		k· g  i		n· g  out		t° g  exp	 Ð° g  vx	 ïÜ g  vy	 ïÜ g  vx	Æ g  vy	Æ g  vx	» g  vy	» g  vx	*° g  vy	*° g  vx	;¥ g  vy	;¥ g  where	W g  t	b g  l	w g  vx	ô· g  vy	ô· g  vx	¯ g  vy	¯ g  vx	« g  vy	« g  vx	/§ g  vy	/§ g  vx	@£ g  vy	@£ g  where	\ g  t	g g  l	| g  vx	Èþ g  vy	Èþ g  vx	áö g  vy	áö g  vx	òò g  vy	òò g  vx	î g  vy	î g  vx	ê g  vy	ê g  vx	%æ g  vy	%æ g  vx	6â g  vy	6â g  vx	GÞ g  vy	GÞ g  vx	XÚ g  vy	 XÚ g  vx	!iÖ g  vy	"iÖ g  where	#Æ g  t	$Ã g  l	$¯À g  vx	ì g  vy	ì g  vx	 è g  vy	 è g  vx	1ä g  vy	1ä g  vx	Bà g  vy	Bà g  vx	SÜ g  vy	SÜ g  vx	dØ g  vy	dØ g  where	Ä g  t	Á g  l	­¾ g  vx	ý¬ g  vy	ý¬ g  vx	¨ g  vy	¨ g  vx	¤ g  vy	¤ g  vx	0  g  vy	0  g  where	U g  t	` g  l	u  Zg  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	6
		7			7			7			7			7	 		7			7	*	 	7		#	8		'	7		*	8		.	7		1	8	 	5	7		8	8	*	<	7		E	9		F	9		K	9		L	9		Q	9	(	R	9		R	7		[	:		\	:		a	:	!	b	:		g	:	,	h	:		h	7		k	;		k	7		t	J		y	K		}	K	 	L	 	L	 	M	 	K	 	N	 	N	 	O	" 	N	 £	Q	! ¨	Q	 ®	Q	4 µ	R	 ¼	N	 Á	T	 Ð	T	 Ö	U	G	g	L	>	W	>	W	>	Z	?	b	?	q	A	u	A	w	A	w	A	~	B		B		g		g		g		i	¥	i	ª	i	°	i	µ	i	»	i	À	i	Æ	i	Ë	i	Ñ	i	Ö	i	Ü	i	á	i	ç	i	ç	U	L	e	Q	>	\	>	\	>	_	?	g	?	v	A	z	A	|	A	|	A		B		B		e		e		e	£	U	u	^		>		>		>		?		?	©	A	­	A	¯	A	¯	A	¶	B	»	B	É	^	Ì	^	Ò	^	Ö	U	o	X	9y	U		>		>		>		?		?	§	A	«	A	­	A	­	A	´	B	¹	B	Ç	Y	Ê	Y	Ð	Y	Ô	U	;	V	1E	U	J	>	U	>	U	>	X	?	`	?	o	A	s	A	u	A	u	A	|	B		B		W		W		W		U	°	J	±	J	·	J	 	·  g  nameg  decode-load-program C-RNOC-PQRSTWXY      h   å   ] 456       Ý       g  str
		 g  pos		 g  value			  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	}				~	2		~	 			   CZ[F  h0  s  ]14 5&  L 645
$  Í&  "  $  "  &  "  $  	"  
4L 5 4L 5 4L 5 		45"  K$   &  445 5"   C44L 5 >  "  G  "ÿÿµ
"ÿÿ®"  !
$  64L 5 "ÿÿß45 "ÿÿË    k      g  opcode
	, g  inst		, g  make-sequence		D ó g  sequence-set!		d ó g  a		k  g  b		r  g  c		y  g  len	  ó g  seq	  ó g  i	 ¡ ì g  n	 ÷ g  out	 ÷  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	m				n				n			p			o			q			s		"	s		&	o		.	w		<	w		D	v		N	{		\	{		d	v		g		%	k			n		/	r			u		9	y		  	   	+  	 	v	  	 	v	 ¡ 	 ¦ 	 ª 	 ° 	. ´ 	  µ 	$ º 	7 À 	$ É 	 Ë 	 Ò 	. Û 	 æ 	 ì 	 ì 	 ÷ 	 ú 	 ÿ 	 	 		 	& 	  	 	 	$ 	;, 	 7	,   C  h   Ü   ]4 5  O 6      Ô       g  pop
		  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm
	k
		l				l	 		  g  nameg  decode-bytecode CGRC       Ø       g  m
		,  g  filenamef  l/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/assembly/decompile-bytecode.scm		
e	
]	*
[	,
å	/
\	6
"?	k
 	"A
   C6 