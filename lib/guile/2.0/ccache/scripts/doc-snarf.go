GOOF----LE-4-2.0¤J      ] ¦ 4        he      ] g  foo¤	g  bar¤	g  foo/bar¤	f  0.0.2¤	g  doc-snarf-version¤	g  guile¤	 ¤	g  define-module*¤		 ¤	
 ¤	g  scripts¤	g  	doc-snarf¤	 ¤	g  filenameS¤	f  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm¤	g  importsS¤	g  ice-9¤	g  getopt-long¤	 ¤	 ¤	g  regex¤	 ¤	 ¤	g  
string-fun¤	 ¤	 ¤	g  rdelim¤	 ¤	 ¤	 ¤	g  exportsS¤	  ¤	!g  set-current-module¤	"! ¤	#! ¤	$f  $Snarf out documentation from a file.¤	%g  %summary¤	&g  version¤	'g  single-char¤	('v ¤	)g  value¤	*) ¤	+&(* ¤	,g  help¤	-'h ¤	.,-* ¤	/g  output¤	0'o ¤	1) ¤	2/01 ¤	3g  texinfo¤	4't ¤	534* ¤	6g  lang¤	7'l ¤	8671 ¤	9+.258 ¤	:g  command-synopsis¤	;g  display¤	<f  
doc-snarf ¤	=g  newline¤	>g  display-version¤	?f  (Usage: doc-snarf [options...] inputfile
¤	@f  6  --help, -h              Show this usage information
¤	Af  3  --version, -v           Show version information
¤	Bf  ?  --output=FILE, -o       Specify output file [default=stdout]
¤	Cf  3  --texinfo, -t           Format output as texinfo
¤	Df  5  --lang=[c,scheme], -l   Specify the input language
¤	Eg  display-help¤	Ff  	doc-snarf¤	Gg  
option-ref¤	Hg  string->symbol¤	Ig  string-downcase¤	Jf  scheme¤	Kg  
snarf-file¤	Lg  main¤	Mg  c¤	Nf  ^/\*(.*)¤	Of  ^ \*/¤	Pf  	^ \* (.*)¤	Qf  	^ \*-(.*)¤	Rf  NOTHING AT THIS TIME!!!¤	SMNOPQR ¤	Tg  scheme¤	Uf  ^;; (.*)¤	Vf  ^;;\.¤	Wf  ^;;-(.*)¤	Xf  	^\(define¤	YTUVUWX ¤	ZSY ¤	[g  supported-languages¤	\g  list-ref¤	]g  assq-ref¤	^g  docstring-start¤	_g  docstring-end¤	`g  docstring-prefix¤	ag  option-prefix¤	bg  signature-start¤	cg  std-int-doc?¤	dg  	lang-parm¤	eg  memq¤	fg  map¤	gg  car¤	hg  error¤	if  .doc-snarf: input language must be c or scheme.¤	jg  write-output¤	kg  snarf¤	lg  format-texinfo¤	mg  format-plain¤	ng  unread-string¤	og  read¤	pg  length¤	qg  define¤	rg  lambda¤	sg  string?¤	tg  find-std-int-doc¤	ug  separate-fields-discarding-char¤	vg  string-append¤	wg  split-prefixed¤	xg  open-input-file¤	yg  make-regexp¤	zg  eof-object?¤	{g  close-input-port¤	|g  reverse¤	}g  neutral¤	~g  regexp-exec¤	g  	read-line¤ g  
doc-string¤ g  match:substring¤ g  options¤ f  
internal: ¤ g  append¤ g  parse-entry¤ g  entry¤ g  
make-entry¤ g  entry-symbol¤ g  entry-signature¤ g  entry-docstrings¤ g  entry-options¤ g  entry-filename¤ g  
entry-line¤ g  
get-symbol¤ g  make-prototype¤ f   ¤ g  call-with-input-string¤ g  	read-char¤ g  join-symbols¤ g  symbol->string¤ f  . ¤ f   ¤ g  with-output-to-port¤ g  open-output-file¤ g  current-output-port¤ g  for-each¤ f  
¤ f  @c snarfed from ¤ f  :¤ f  @deffn procedure ¤ g  
write-line¤  f  @c ¤ ¡f  
@end deffn¤ ¢f  Procedure: ¤ £f  ;; ¤ ¤f  Snarfed from ¤ ¥f  ¤C 5      h°;  º  ]      h   Ê   ] $  CC  Â       g  braz
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
	1
		2		
	2			2	 		  g  nameg  foo/bar CRR4
 5 4# >  "  G   $%R9:R;<=      h0   Ï   ] 4>  "  G  4>  "  G  6       Ç       g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
	`
		a			a			a			a		*	a	5 		*
  g  nameg  display-version C>R;?@ABCD    hh   ,  ] 4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  6      $      g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
	e
		f			f			f			g			g			g		'	h		+	h		0	h		9	i		=	j		B	i		K	k		O	k		T	k		`	l		b	l	 		b
  g  nameg  display-help CERF:G,&3HI6J>E/K       h     -  1  3 4 545454544	4
555$  6 $  6 4545$  66          g  args
		  g  options	  g  help-wanted		J  g  version-wanted		J  g  texinfo-wanted		J  g  lang		J  g  input		y  g  output		y   g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
	p
	
	q			q	$		q			q			q			r			r	+	!	r		"	s		(	s	.	+	s		,	t		2	t	.	5	t		6	u		9	v		<	v	"	B	v	6	D	v	<	F	v	"	H	v		J	u		J	r		X	w		\	x		b	w		f	y		g	{		l	{	)	o	{		p	|		v	|	*	y	|		y	{	  	 	}	
  	  	  	 )		 


  g  nameg  	doc-snarf CRiLRZ[R\][^_`abc 
    hh   â   ]4 5$  
"  K$  "  =$  	"  .$  	"  $  	"  	$  	"  6Ú       g  lang
		h g  parm		h  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
  
	 ¡		 ¢		h ¡	 		h	  g  nameg  	lang-parm CdRefg[hijklm   hP   \  ]4455$  "  4>  "  G  4 5$  	"  
6T      g  input
		P g  output		P g  texinfo?			P g  lang			P g  t			3  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 ®
	 ¯		 ¯		 ¯		 ¯		 °		# °		( °		6 ±		F ²		P ±	 		P	  g  nameg  
snarf-file CKRnopqrs hð   m  ]
4 >  "  G  45$  Å"  o	45$  _&  T$  I$  =	45$  *&  45$  	CCCCCCCC	45$  @&  3$  &$  45$  C"ÿÿO"ÿÿK"ÿÿG"ÿÿC"ÿÿ?C   e      g  line
	 í g  
input-port	 í g  form		 í  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 ·
	 º		 »		 »		" ¼		& ¼		- Ä		4 Ä		8 Ã		: Å		= Å		A Ã		D Æ		F Æ		J Ã		M Ç		P Ç		T Ã		W È		\ È		` È		a È		e Ã		g É		j É	"	m É		q Ã		r Ê		w Ê	 	z Ê		~ Ê	  Ã	  Ë	  Ë	  ¼	  ½	 £ ½	 § ¼	 © ¾	 ¬ ¾	 ° ¼	 ³ ¿	 µ ¿	 ¹ ¼	 ¼ À	 ¿ À	 Ã ¼	 Ä Á	 É Á	 Í Á	 Ñ ¼	 Ô Â	 5	 í	  g  nameg  find-std-int-docg  documentationf  Unread @var{line} from @var{input-port}, then read in the entire form and
return the standard internal docstring if found.  Return #f if not. CtRufv    h   ©   ]L  6      ¡       g  line
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 Ô	
	
 Õ	 		
   C    h   ¬   -  1  3 L O  6  ¤       g  lines
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 Ó		 Ô	 			


   C h   Ù   ]
 O 6       Ñ       g  string
		 g  prefix		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 Ð
	 Ñ	 			  g  nameg  split-prefixed CwRxyd^_`abz{|}~ctw  h¸  ¹	  ]4 544554455445544554455" 74	5$  4
>  "  G  6	&  [45$  *4545 
	"ÿÿ45	
	"ÿÿz	& m45454545$  *4545

	"ÿÿ$  +45
45
	"ÿþå$  45$  &45$  45"  "  $  4455"  454
 5
	"ÿþZ$  /454
 5
	"ÿþ%45
	"ÿþ	& 2454545$  +45
45
	"ÿý«$  45$  &45$  45"  "  $  4455"  454
 5
	"ÿý $  /454
 5
	"ÿüë45
	"ÿüËC45

	"ÿü¬     ±	      g  
input-file
	³ g  lang	³ g  i-p			³ g  docstring-start		³ g  docstring-end		)³ g  docstring-prefix		9³ g  option-prefix		I³ g  signature-start		Y³ g  line		_ g  state			_ g  doc-strings	
	_ g  options		_ g  entries		_ g  lno		_ g  m	  å g  m0	Z g  m1	Z g  m2	Z g  m3	Z g  d	¯ g  int-doc	´Ó g  options	Ó g  m1	} g  m2	} g  m3	} g  d	Ðé g  int-doc	î g  options	?  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
 Þ
	 ß			 ß		 à	%	 à	2	 á	(	 à	2	 à	%	 ß		 à	%	 à	2	% â	(	' à	2	) à	%	) ß		, à	%	/ à	2	5 ã	(	7 à	2	9 à	%	9 ß		< à	%	? à	2	E ä	(	G à	2	I à	%	I ß		L à	%	O à	2	U å	(	W à	2	Y à	%	Y ß		_ ï		` ò		j ñ		k ó	  ô	  ø	  ñ	  ù	  ù	  ú	
  û	 ¤ û	  ¥ ü	 ¯ ü	 ° ü	- µ ü	9 Å û	 Æ ý	 Ï ý	& Ð ý	* Õ ý	6 å ý	 é	 í ñ	 î	 ÷	 					
 	(	 )		3		4		:9
	I	O	
P	X	 [	"e	j	z		
 è	  è	0 è	  è	 é	) é	  ê	" ê	)¥ ê	;§ ê	)´ è	¼ ë	½ ì	À ì	Ê ì	Ó	Ö	Þ	"ß	+à	/á	ò	õ			
		 	)	-	'	*	:	;	C	 D	)E	-J	9Z	^	b ñ	c 	l!	u"	} 	#	
%	%	 &	"&	¤'	´%	º#	
» è	 Á è	0Ã è	 Ç è	È é	)Ð é	 Ø ê	"Ù ê	)ß ê	;á ê	)î è	ö ë	÷ ì	ú ì	 ì	)	*	*	"*	+*	/+	,+	/-	?*	E#	
F/	N/	 O/	)P/	-Q0	a0	d2	t/	u4	}4	 ~4	)4	-4	94	 ï	 ï	 ï	+  ï	B¡ ð	¢ ð	$³ ï	 ²	³	  g  nameg  snarf CkR      h   6  ]  C    .      g  symbol
		 g  	signature		 g  
docstrings			 g  options			 g  filename			 g  line			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
6
	7	
	7	 			  g  nameg  
make-entry CR     h   ¾   ] £C ¶       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
8
	9	 		  g  nameg  entry-symbol CR     h   Á   ] 	£C¹       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
:
	;	 		  g  nameg  entry-signature CR  h   Â   ] 	£Cº       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
<
	=	 		  g  nameg  entry-docstrings CR h   ¿   ] 	£C·       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
>
	?	 		  g  nameg  entry-options CR    h   À   ] 	£C¸       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
@
	A	 		  g  nameg  entry-filename CR   h     ] 	£Cþ       g  e
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
B
	D	 		  g  nameg  
entry-lineg  documentationf  4This docstring will not be snarfed, unfortunately... CR|pH 	  h¸     ]$  545454 5454 54564 5
$  =44 554 54 5454 54564 5454 5456      }      g  
docstrings
	 ² g  options	 ² g  def-line		 ² g  filename		 ² g  line-no		 ²  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
I
	K		M		N		N	+	 O		+P		2P		3P	3	:P		;P		=M		>Q		FQ		JK		MR		PR	%	WR	 	YR		ZS		aS		bT		iT		jU		uV		|V		}V	2 V	 V	 R	 X	 X	 X	 X	-  Y	 §Y	 ¨Y	2 ¯Y	 °Y	 ²X	 )	 ²	  g  nameg  parse-entry CRo       hP     ]	4 >  "  G  4 >  "  G  4 5$  6$  6C    
      g  s-p
		L g  tmp	-	L  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
a		b		c		'd		-b		2f			6e		<g			?h			Ce		Ii			Kk		 		L   C   h   Ç   ] 6      ¿       g  def-line
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
^
	
_	 		
  g  nameg  make-prototype CRo hH   	  ]	4 >  "  G  4 >  "  G  4 5$  C$  CC       g  s-p
		G g  tmp	-	G  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
p		q		r		's		-q		2u			6t		9v			=w			At		Fz		 		G   C    h   Ã   ] 6      »       g  def-line
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
m
	
n	 		
  g  nameg  
get-symbol CRv  hH   `  ] (  C $  4 56 (   64 54 56     X      g  s
		C  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
~
			
															"		&		+		-			0		5	(	7		9	1	:	5	?	C	A	5	C		 		C  g  nameg  join-symbols CR        h      ] LL 6             g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
		
	 		

   C       h(     ]$  45"  45  O 6         g  entries
		% g  output-file		% g  writer			%  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm

	
			*		#	%	 		%	  g  nameg  write-output CjR;=    h   ¦   ] 6       g  s
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
			 		   C;       h    Á   ]4>  "  G   6      ¹       g  s
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
¡		¡		¡	!	¡		¡	( 		   C¡      h  	  ]4>  "  G  44 5>  "  G  4>   "  G  4>  "  G  44 5>  "  G  4>  "  G  44 5>  "  G  4>   "  G  4	>  "  G  44
 5>  "  G  4>   "  G  44 5>  "  G  44 5>  "  G  6        g  entry
	  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm

											#		,		<		@		E		N		Q		\		e		i		n		w		z	 	 	 	 ¢	 §	 °	 ³	 ¾	 Ç	 ×	 Ü 	 ç	 ð¡	 õ¢	 ¡	£	£	 %	  g  nameg  format-texinfo ClR;¢=   h   ¦   ] 6       g  s
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
«		«	 		   C;£      h    Á   ]4>  "  G   6      ¹       g  s
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
­		­		­	!	­		­	( 		   C¤¥      hØ   Ã  ]4>  "  G  44 5>  "  G  4>   "  G  44 5>  "  G  44	 5>  "  G  4
>  "  G  44 5>  "  G  4>  "  G  44 5>  "  G  4>   "  G  6   »      g  entry
	 Õ  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm
§
	¨		¨		¨		©		©		#©		,ª		<«		A¬		L«		U­		Z®		e­		n¯		r¯		w¯	 °	 °	 °	 ±	 ±	  ±	 ©²	 ¬²	 ·²	 À³	 Ó´	 Õ´	 	 Õ  g  nameg  format-plain CmRC²      g  m
$  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/doc-snarf.scm õ	1
 ÷	J	 ú	J
 û	N
&	U	)	U
+	X	.	W
J	`
ÿ	e
i	p
p 
r 	u 
å  

· ®
0 ·
Ô Ð
 Þ
 é6
!Á8
":
#u<
$J>
%#@
&AB
) I
,
^
.Vm
0~
2'
7,
;®§
 "	;°
   C6 