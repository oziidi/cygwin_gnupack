GOOF----LE-4-2.0%      ] I 4     hG      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  scripts¤	g  lint¤	 ¤		g  filenameS¤	
f  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm¤	g  importsS¤	g  ice-9¤	g  common-list¤	 ¤	 ¤	g  format¤	 ¤	 ¤	 ¤	g  exportsS¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  %include-in-guild-list¤	f  1Check for bugs and style errors in a Scheme file.¤	g  %summary¤	g  scan-file-for-module-name¤	g  uniq¤	g  scan-file-for-free-variables¤	g  resolve-module¤	 f  Resolved module: ~S
¤	!g  catch¤	"g  eval¤	#f  !Unresolved free variables in ~A:
¤	$g  
write-char¤	%g  write¤	&g  newline¤	'f  #No unresolved free variables in ~A
¤	(g  with-input-from-file¤	)g  eof-object?¤	*g  read¤	+g  define-module¤	,g  append¤	-g  detect-free-variables¤	.g  memq¤	/g  define-generic¤	0g  quote¤	1g  
quasiquote¤	2g  let¤	3g  letrec¤	4g  map¤	5g  car¤	6g  let*¤	7g  and-let*¤	8g  define¤	9g  define-public¤	:g  define-macro¤	;g  lambda¤	<g  lambda*¤	=g  receive¤	>g  define-method¤	?g  define*¤	@g  define-class¤	Ag  detect-free-variables-noncar¤	Bg  case¤	Cg  unquote¤	Dg  unquote-splicing¤	Eg  else¤	Fg  =>¤	Gg  for-each¤	Hg  main¤C 5    hø  ã   ]4	
5 4 >  "  G   RR !"      h      ] LL 6            g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
	x			y			y	 		
   C#$%&   hh   ñ   -  1  3 M$  4L>  "  G  "   4	>  "  G  4L >  "  G  4>   "  G  NC      é       g  args
			b  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
	z			{			|			}			|		)			; 		@ 		E 		N 		` 	 			b


   C'  h   É  ])4 544 5545H4>  "  G  "  8(  "  64O 	 O >  "  G  "ÿÿÈ"ÿÿÀJ$  	
 6C Á      g  filename
	  g  module-name	  g  	free-vars		  g  module		!  g  all-resolved?		!  g  	free-vars		>	v  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
	n
		o		
	p			p			p			o			q		!	q		&	s		+	s		2	s		>	t		D	u		I	w		p 		v 		v	t	  	  	  	
 	   g  nameg  lint CR()*+       hP     ]"  94 5$  C"  45  "ÿÿä $   &   C"ÿÿÜ"ÿÿØ45  "ÿÿ¼            g  x
		?  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 		 		 		 		 		" 		" 		% 		) 		, 		. 	!	2 		5 		? 		@ 		J 	 		J
   C      h   Í   ] 6      Å       g  filename
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 
	
 	 		
  g  nameg  scan-file-for-module-name CR(),*- h@   *  ]"  )4 5$  @45 4 5 "ÿÿ×45  "ÿÿÉ   "      g  x
		/ g  fvlists		/  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 		 		 		 		 		 		 		" 	8	$ 		' 		/ 		/ 		0 		5 	%	= 	 		=
   C   h   Ð   ] 6      È       g  filename
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 
	
 	 		
  g  nameg  scan-file-for-free-variables CR.+/0123,45-  h   ¹   ] L $  L"  M6±       g  binding
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 Ä	"	 Å	;	 Æ	;	 Å	$ 		   C- h   ¨   ] L 6              g  bodyform
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 Ë	"	
 Ì	$ 		
   C67-      h   ¨   ] M 6              g  bodyform
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 Ô		
 Õ	 		
   C-89:;<-      h   ¨   ] L 6              g  bodyform
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 ï		
 ð	 		
   C=-        h   ¨   ] L 6              g  bodyform
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 ø		
 ù	 		
   C>?-      h   ¨   ] L 6              g  bodyform
		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm

		
	 		
   C@A        h   Æ   ] $   "   M 6¾       g  slot/option
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
			<		8		<		 		   CB-  h   ­   ] M 6     ¥       g  case
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
			1		 		   CCDEFA %     hh  ã  ]H $  4 J5$  C  C $ ; $  "  !$  "  $  "  $  C$  "  $  r $       "    4J4	
 5544	O  5?44	O  5?6$  "  $  I (  4	O  5@4 J54   J56$  "  $  "  $  : $   JK  J6 JK J6$  "  $  U"  -(  "  /$  "ÿÿÜ"  J "ÿÿÅ4	O  5@$  .4J 54 J54	O  5@$  "  $  f"  =(  "  @$  !$  "  "ÿÿÌ"  J "ÿÿ´4	O  5@$  4	O  5@$  4 J54	O  5@ $  "  !!$  "  "$  "  #$  	$ J64J54$ J56C  Û      g  x
	f g  locals	f g  key		,d g  letrec?	 ¶ ì g  locals-for-let-body	 ¶ ì g  locals	Æó g  args	Æó g  locals-for-lambda-body	 g  locals-for-receive-body	+N g  locals	i¦ g  args	i¦ g  locals-for-method-body	µË  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
 ¬
		 °			 °		 ±		 ±			 ±		! ±	!	% ³			) °		, ´		, ´			b ·		l ´			} ½		 ½	  ½	  ¾	.  ¾	(  ¾	;  ¾	"  ¿	"  ¾	  ¾	 ¡ Á	  £ Á	( ¤ Á	 ¥ Â	' ª Â	6 ± Â	? ´ Â	6 ¶ Â	' ¶ Á	 ½ Ã	 À Ä	 Ð É	" Ó Ä	 Õ Ã	 Ö Ê	 Ù Ë	 å Î	" è Ë	 ê Ê	 ì Ã	 õ ´		 Ò	 Ò	 Ô	 Ö	 Ô	 Ó	" ×	' ×	5* ×	// ×	0 Ø	4 Ø	/7 Ø	7< Ø	C> Ø	/B Ù	5H Ù	/J Ø	L ×	U ´		t Ü	v Ü	z Ü	} Þ	% Þ	 Þ	 ß	) ß	3 ß	? ß	) ß	 â	% â	  â	¥ ã	)« ã	´ ´		Æ æ	*Ì è	,Õ é	3Ù è	,Ü ê	?ß ê	9â ë	9ê ê	3ï í	3ó æ	*ø ç	; æ	* æ	 ï	 ò	 ï	 î	  ´		! õ	+( õ	:+ õ	++ õ	0 ÷	5 ÷	,; ÷	< ø	H û	L ø	N ö	W ´		i þ	*o 	,x	3| 	,	J	C	?	C	C	9	9	3¢	3¦ þ	*« ÿ	;µ þ	*µ þ	º
	Æ	É
	Ë		Ô ´		×	ã	ç	é	ò ´		õ	ú	*ÿ	 				 ´		I"	*M"	P$	T$	0X$	Y%	^%	7b%	d$	e'	 	f	  g  nameg  detect-free-variables C-R.FA,-    hX     ]
 $  4 5$  C  C $  . $  	 6454 56C         g  x
		U g  locals		U g  key		(	S  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
)
	-				-		
.		.			.		.	!	!0			%-		(1		(1			83	*	<3		?5		C5	0	G5		H6		M6	7	Q6		S5		T8	 		U	  g  nameg  detect-free-variables-noncar CARG    h   ·   -  1  3  6       ¯       g  files
			  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm
:
	;	 			


  g  nameg  main CHRC    Û       g  m
		,  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/lint.scm		f
	0	k
	2	l		5	l
Ü	n
X 
	Ë 
þ ¬
)
ò:
 	ô
   C6 