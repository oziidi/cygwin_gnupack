GOOF----LE-4-2.0±7      ] l 4  h³      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  sxml¤	g  ssax¤	g  input-parse¤		 ¤	
g  filenameS¤	f  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm¤	g  importsS¤	g  ice-9¤	g  rdelim¤	 ¤	 ¤	 ¤	g  exportsS¤	g  peek-next-char¤	g  assert-curr-char¤	g  
skip-until¤	g  
skip-while¤	g  
next-token¤	g  next-token-of¤	g  read-text-line¤	g  read-string¤	g  find-string-from-port?¤	 	¤	g  set-current-module¤	 ¤	 ¤	 g  integer->char¤	!g  ascii->char¤	"g  char->integer¤	#g  char->ascii¤	$g  char-newline¤	%g  char-return¤	&g  1+¤	'g  inc¤	(g  1-¤	)g  dec¤	*g  make-syntax-transformer¤	+* ¤	,* ¤	-g  
define-opt¤	.g  macro¤	/g  $sc-dispatch¤	0/ ¤	1/ ¤	2g  _¤	3g  any¤	423¤	5g  syntax->datum¤	65 ¤	75 ¤	8g  datum->syntax¤	98 ¤	:8 ¤	;g  reverse¤	<g  optional¤	=g  define*¤	>g  append¤	?g  optionalS¤	@g  syntax-violation¤	A@ ¤	B@ ¤	Cf  -source expression failed to match any pattern¤	Dg  throw¤	Eg  parser-error¤	Fg  current-input-port¤	Gg  	read-char¤	Hg  	peek-char¤	Ig  memv¤	Jf  Wrong character ¤	Kf   (0x¤	Lg  eof-object?¤	Mf  *eof*¤	Ng  number->string¤	Of  ) ¤	Pf  . ¤	Qf  	 expected¤	Rg  number?¤	Sf  Unexpected EOF while skipping ¤	Tf   characters¤	Ug  memq¤	Vg  *eof*¤	Wf  $Unexpected EOF while skipping until ¤	Xg  make-string¤	Yg  input-parse:init-buffer¤	Zf   ¤	[g  string-length¤	\g  	substring¤	]f  EOF while reading a token ¤	^g  string-append¤	_g  string-set!¤	`g  next-token-old¤	ag  string-concatenate-reverse¤	bg  
procedure?¤	cg  *read-line-breaks*¤	df  reading a line¤	eg  list->string¤	fg  delete¤	gg  read-delimited¤	hg  peek¤	ig  	read-line¤	jg  char=?¤	kg  
string-ref¤C 5h-  g  ]4	
5 4 >  "  G    i!R"i#R
$R%R&i'R(i)R4,-.147:;<=>?       hp   ¶  - 1 3 4 5$  $$  &  "  "  "  $  4455C C  ®      g  bindings
			n g  body			n g  	body-rest				n g  rev-bindings			n g  opt-bindings		>	n  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	J
	
	K			K			M			M	
		M	+		M	$	!	M	
	#	N		&	N		+	M	
	.	O		>	K		F	P		H	Q		I	Q		L	Q		S	Q	5	T	Q	$	V	Q		Z	Q		_	Q		d	T	 			n	
	   C       h   ±   ]	4 5L 4?6©       g  args
		 g  v			  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm		J
 		   CBC        h(   ¯   ]	4 5$   O @ 6 §       g  y
		' g  tmp		'  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	J
 		'   C5-RDE    h   õ   - 1 3  @   í       g  port
			 g  message			 g  rest				  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	V
		W			W	 				
	  g  nameg  parser-error CERFGH      h8   Ú   -  . , 3  #  45  4 >  "  G   6       Ò       g  port
		1  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
	,
		,	,		-		1	.	 		1
  g  nameg  peek-next-char CRFGIEJKLMN"OPQ     hh     - . , 3 #  45 454 5$  C45$  "  4	4
5	5 6
       g  expected-chars
		g g  comment		g g  port			g g  c		 	g  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
	;
		<	.		=		 	=		#	>		/	>		8	?		<	@		=	@		G	@		I	@	&	N	A		Q	A	$	[	A		]	A	;	a	B		e	B	'	g	?	 		g	  g  nameg  assert-curr-char CRFRLGEST)IUVW     hÈ   þ  - . , 3 #  45 4 5$  O"  C
$  94455$  4 >  "  G  "   45"ÿÿ¿C "ÿÿµ"  B4	 5$  C45$  4
 5$  C 645"ÿÿ¾45"ÿÿ±  ö      g  arg
	 Æ g  port	 Æ g  i		(	k g  c		w ¹  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
	Q
		Q	,		S		$	R		(	T		+	U		0	T		1	V		4	V		<	V		@	V	
	A	W		G	W		K	X		P	W		]	T		i	T		w	Z		x	\	
 	[	 	]	
 	[	 	^	 	^	 	^	 	^	 ¨	_	  ¬	_	 ­	`	 ¹	`	 ¹	Z	 º	Z	 Æ	Z	 "	 Æ  g  nameg  
skip-until CRFIGH   h`   ,  - . , 3 #  45 "  /4 5$  4>  "  G  45"ÿÿÔC45"ÿÿÄ       $      g  
skip-chars
		Y g  port		Y g  c			L  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
	k
		k	3		l			m		*	l		+	n		=	l		I	l		M	l	
	Y	l	 		Y  g  nameg  
skip-while CR4Xi 5    h      ] L C          g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
 	 		
   C O  YRZFY[I\LUVE]^X_G'H        hø   Ô  - . , 3 	#  #  45 45 H4J5HJH"  ©45$  	J
645$   4	5$  	J
6
6J$  !4J4J55KJK4J5K"   4J>  "  G  4>  "  G  4545"ÿÿW
4 5"ÿÿE  Ì      g  prefix-skipped-chars
	 ö g  break-chars	 ö g  comment		 ö g  port		 ö g  buffer		( ö g  curr-buf-len		1 ö g  quantum		5 ö g  i		; ä g  c		; ä  	g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
 ¡
	 ¢	+	 ¢	5	$ £		( £		+ ¤		1 £		; ¦		< ¨			H §		Q ¨		R ©			\ §		] ª		a ª		e ª		i ª	
	r «		x ¬		| ¬	  ®	  ®	
  °	  °	1  °	  °	  ±	  ²	! ¢ ²	 § ³	
 ½ ´	
 Ï µ	 Ö µ	 ä µ	
 ä ¦	 æ ¦	 ö ¦	 &	 ö		  g  nameg  next-token-old C`RZF[I\aLUVE]X_G'HY h  ý  - . , 3 
#  #  45 "  Ø45"  À4	5$  (  	
664	5$  04	5$  (  	
66
6$  45	"ÿÿq4	>  "  G  4>  "  G  4545	"ÿÿ@
	"ÿÿ545 4 5"ÿÿõ      g  prefix-skipped-chars
	 g  break-chars	 g  comment		 g  port		 g  buffer		' ÿ g  filled-buffer-l		' ÿ g  c		' ÿ g  curr-buf-len		. ÿ g  i		4 ô g  c			4 ô  
g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
 Ä
	 Å	%	 Å	/	' Æ		( È		. È		4 É		5 Ë		A Ê		G Ì		P Ì	(	Z Í		[ Î		e Ê		f Ï		j Ï		n Ï		r Ï		x Ð	  Ð	*  Ñ	  Ò	!  Ò	  Ó	  Ê	  Ô	 ª Õ	 ¶ Ô	 · ×	 Í Ø	 ß Ù	 æ Ù	 ô Ù	 ô É	 ÿ Æ	  Æ	 Æ	B Ç	 Æ	 (			  g  nameg  
next-token CRFY[bXH_G'\aI    hh  ®  - . , 3 #  45 45 454 5$  "  "  $  45"ÿÿß4 455$  54>  "  G  4>  "  G  4	5"ÿÿ(  	

66
"ÿÿz"ÿÿk"  "  $  45"ÿÿß454 5$  54>  "  G  4>  "  G  4	5"ÿÿ(  	

66
"ÿÿx"ÿÿi¦      g  incl-list/pred
	h g  port	h g  buffer		h g  curr-buf-len		'h g  buffer		8 Â g  filled-buffer-l		8 Â g  i		< » g  c		e » g  buffer	 Ñ] g  filled-buffer-l	 Ñ] g  i	 ÕV g  c	 ùV  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
 ø
	 ù	+	 ú		 ú		! û		' ú		* ü		4 ü		8 ý		< þ		A ÿ		E ÿ		F 		Q 	,	Y 	
	Z		]	#	e		e	
	m		n	 	 	 ¢	 ¨	 ±	* »		 » þ	 Â ý	 Å ý	1 Í ý	 Ñ	 Õ	 Ú	 Þ	 ß	 ê	, ò	
 ó	 ù	
 ü					1	=	C	L	,V	V	]	`	1h	 6	h  g  nameg  next-token-of CR$i%iV cRFLHcdG% 	  h     -  . , 3  #  45  44 55$   64 54 5$  (4 5
$  4 >  "  G  "   "   C            g  port
		z g  line	;	z g  c		D	z  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
*
	*	,	+		+		%+		)+		/+	%	0-		3-		7.		;-		;,		>/		D,		K0		O0		P0	'	Y0	!	]0		^1		 		z
  g  nameg  read-text-line CRFXL\'_GZ 	  h   Ë  - . , 3 #  45  
$  m4 5"  P45$  	
6454>  "  G   $  C45"ÿÿ°
45"ÿÿ C       Ã      g  n
	  g  port	  g  buffer		(  g  i		.	~ g  c		.	~ g  i1		I	~  g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scm
<
	<	+	=		!=		">		(>		.?		/@		9@		B@		CA		IA	
	LB		fC		jC		pD		~D		~?	 ?	 ?	 =	 	   g  nameg  read-string CRZFefVLHUE]gh    h°   Z  - . , 3 #  #  45 44554 (  45"  	4 55$  4	5$  C
6454455$  4	5$  C
6C      R      g  prefix-skipped-chars
	 ª g  break-chars	 ª g  comment		 ª g  port		 ª g  delims		1 ª g  token		{ ª  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	]
		^	+		^	5	$	_		'	_		+	_	&	/	_		1	_		1	_		4	`		<	`		=	a		H	b		R	`		V	`		W	c		[	c		_	c		c	c		e	d		l	e		p	e		q	f		y	f	1	{	f		{	f		~	g	 	g	  	g	 	g	
 	h	 	h	 	h	 	g	 £	i	! §	i	 %	 ª		  g  nameg  
next-token CRFi   h    Î   -  . , 3  #  45   6 Æ       g  port
		  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	l
		l	,		m	 		
  g  nameg  read-text-line CRRHLb  hp   U  ]4M5  $  "  MM $  K4L5 M $   C4 5$  C4M5$  N 4M 5$  C CM $  C CCM      g  t
	
	 g  c
	*	n  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	|			|	!	
	|		
	|			}		#	|		$	~	!	*	~		2			6 	!	@			C 	!	M			P 	!	Q 	%	[ 	!	e 	!	i		 		p
  g  nameg  my-peek-char CGjk[        hp  Ê  - 1 3 
HJ(  "  JKHHO "  G45 $  8K4>  "  G  JK44 
55$  "  
"ÿÿ»C"  b4 5$  JC45 $  B44 55$  #K4>  "  G  JK"ÿÿ«"  C"  c
$  "ÿÿ<"  E	$  "ÿÿr44 	54 	55$  			"ÿÿÈ"ÿÿ¤
	"ÿÿ´Q 
KK"ÿþÉ       Â      g  str
	i g  <input-port>	i g  max-no-char		i g  no-chars-read		i g  peeked?		i g  c		7	y g  pos-to-match		} ß g  c	  ß g  i	 ãF g  matched-substr-len	 ãF g  j	 èF g  k		 ú? g  my-peek-char	Fi g  no-chars-read	Se g  peeked?	Se  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm
	t
		w			w	/		w			x		2 		3 		7 	
	? 		B 		C 	/	W 	/	Y 		Z 		_ 	&	h 		l 		s 		w 		} 	  	  	  	
  	  	  	  	 ¡ 	( « 	 ¯ 	 ² 	 ³ 	/ Ç 	/ É 	 Ì 	3 Ò 	  Ý 	 ã ¦	 è §	 è §	
 î ¨	 ò ¨	 ö ª	 ú «	 ÿ ¬	 ¬	 ­	 ®	 ®	" ¯	"! ¯	2# ¯	"% ®	) ®	, °	 2 °	5 ±	%? ±	? «	F	x	i ³	 >		i	
	  g  nameg  find-string-from-port?g  documentationf  dLooks for @var{str} in @var{<input-port>}, optionally within the
first @var{max-no-char} characters. CRC _      g  m
		, g  buffer
=  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm		6
	3	B
	:	C
	?	D
	D	E
	K	F
	R	G
	V
g  filenamef  a/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/upstream/input-parse.scmÅ	,
ç	;
Ù	Q
	k
 	 	@ 
G ¡
 Ä
Ñ ø
Û(	:Þ(	á(
*
 <
g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/ssax/input-parse.scm#Í	]
$Ñ	l
-	t
 	-
   C6 