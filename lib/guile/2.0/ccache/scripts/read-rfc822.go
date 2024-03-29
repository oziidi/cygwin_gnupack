GOOF----LE-4-2.0t      ] L 4  h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  scripts¤	g  read-rfc822¤	 ¤		g  filenameS¤	
f  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm¤	g  importsS¤	g  ice-9¤	g  regex¤	 ¤	 ¤	g  rdelim¤	 ¤	 ¤	 ¤	g  exportsS¤	g  read-rfc822-silently¤	 ¤	g  	autoloadsS¤	g  srfi¤	g  srfi-13¤	 ¤	g  string-join¤	 ¤	 ¤	g  set-current-module¤	 ¤	  ¤	!g  %include-in-guild-list¤	"f  Validate an RFC822-style file.¤	#g  %summary¤	$g  make-regexp¤	%f  ^From ¤	&g  from-line-rx¤	'f  ^([^:]+):[ 	]*¤	(g  header-name-rx¤	)f  ^[ 	]+¤	*g  header-cont-rx¤	+g  option¤	,g  eof-object?¤	-g  reverse¤	.g  	read-line¤	/g  regexp-exec¤	0g  for-each¤	1g  unread-char¤	2g  string->list¤	3g  drain-message¤	4g  match:suffix¤	5g  string->symbol¤	6g  match:substring¤	7g  	substring¤	8g  	match:end¤	9f   ¤	:g  string-null?¤	;g  from¤	<g  
body-lines¤	=g  headers¤	>g  body¤	?f  
¤	@g  suffix¤	Ag  error¤	Bf  bad component:¤	Cg  parse-message¤	Dg  format¤	Ef  From ~A
¤	Ff  ~A: ~A
¤	Gf  
~A¤	Hg  display-rfc822¤	Ig  	open-file¤	Jg  	OPEN_READ¤	Kg  main¤C 5  h   J  ]4	
5	 4  >  "  G   !R"#R4$i%5&R4$i'5(R4$i)5*R+R,-.+/&01  h   ¦   ] L 6             g  c
		
  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	A		
	B	 		
   C2 
    h     ]"  p45$  6"  4 5"ÿÿ×$  =45$  ,4 O 
44	55>  "  G  6"ÿÿ©"ÿÿ¥4 5"ÿÿ        g  port
	  g  line		v g  acc			v  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	<
		=			>			>			?			G		'	G	"	/	G		/	>		6	@		B	@		C	A		O	D		R	D	$	Z	D		[	C		`	A		n	E		v	=		w	=		~	=	* 	=	 	   g  nameg  drain-message C3R+4/&.-/(56789 
    hX   Å  ])4 545445544455	5M N C       ½      g  reversed-hlines
		Q g  hlines			Q g  first			Q g  m			Q g  name		(	Q g  data		C	Q  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	P			Q	'			Q			R	&		Q			S	"		Q			T	%		T	5	(	T	%	(	Q		+	U	%	.	V	,	3	V	=	;	V	,	>	W	,	?	V	&	A	X	&	C	U	%	C	Q		J	Y	(	O	Y	 		Q  g  nameg  add-header! C:3*-;<=>?@AB 
    hX   ü   ]	 $  LC $  M C $  MC $  M$  C4M 5NMC	 6    ô       g  	component
		T g  t	1	L  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	g			h		1	l		=	m	&	C	m	>	E	m	C	G	m	&	I	m		P	o		T	o	 		T   C hø   x  ]A$  444 555"  HHHO Q "  45$  )$  4>  "  G  "   4 5K"  m4	5$  4 545"ÿÿ¢$  4>  "  G  "   4 5 "ÿÿr4 5"ÿÿ_4
J5KO C    p      g  port
	 ô g  from	  ô g  
body-lines		# ô g  body		& ô g  headers		) ô g  add-header!		3 ô g  line		< Í g  current-header		< Í g  t		y Í  	g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	I
		J				K			K	"		L	/		K	"		K		 	J		)	O		)	J		<	[		=	\		G	\		M	]		N	]	!	d	^		l	^		q	_		y	\	 	a	 	b	 	b	 	a	  	d	 ¡	d	! ·	e	 Â	e	$ Ê	e	 Í	[	 Î	[	 Ý	[	 Þ	f	 æ	f	 "	 ô  g  nameg  parse-message CCRC       h   É   ] 6Á       g  port
		  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	q
		r	 		  g  nameg  read-rfc822-silently CR;DE0DF   h   Ã   ]  6 »       g  header
		  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	v			w		
	w	$		w	1		w	 		   C=G> 	   hX   R  ]	4 5$  4>  "  G  "   44 5>  "  G  4 56    J      g  parse
		T g  t			-  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	t
		u				u				u					u			u	)		u	4		u	)	.	v		3	x		7	x		9	x		>	v		K	y		L	y		P	y		R	y		T	y	 		T  g  nameg  display-rfc822 CHRIJH       h8     -  1  3 44 554>  "  G  C       ü       g  args
			1 g  parse		/  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm
	{
	
	|			|	%		|	0		|	%		|			|			}	 				1


  g  nameg  read-rfc822 CRiKRCB      g  m
		0  g  filenamef  [/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/read-rfc822.scm		-
	4	3
	6	4		9	4
	:	6		@	6	$	B	6		E	6
	F	7		L	7	$	N	7		Q	7
	R	8		X	8	$	Z	8		]	8
	a	:
Z	<

£	I
	q
=	t
	{
 
 	 
   C6 