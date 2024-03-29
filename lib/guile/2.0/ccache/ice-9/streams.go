GOOF----LE-4-2.0'      ] . 4        h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  ice-9¤	g  streams¤	 ¤		g  filenameS¤	
f  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm¤	g  exportsS¤	g  make-stream¤	g  
stream-car¤	g  
stream-cdr¤	g  stream-null?¤	g  list->stream¤	g  vector->stream¤	g  port->stream¤	g  stream->list¤	g  stream->reversed-list¤	g  stream->list&length¤	g  stream->reversed-list&length¤	g  stream->vector¤	g  stream-fold¤	g  stream-for-each¤	g  
stream-map¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  make-promise¤	  ¤	! ¤	"g  force¤	#g  vector-length¤	$g  reverse!¤	%g  make-vector¤	&g  stream-fold-one¤	'g  stream-fold-many¤	(g  or-map¤	)g  map¤	*g  stream-for-each-one¤	+g  stream-for-each-many¤	,g  apply¤	-g  eof-object?¤C 5    hè!  =  ]4	
5 4 >  "  G   !       h(   è   ]4LL 5  $   4L 5CC     à       g  o
			#  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	V			W				W			X	
		X			Y			Z			Z			Z		 	Y	
	"	[	
 		#
   C     h   Ê   ] O 6 Â       g  m
		 g  state		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	U
		V	 			  g  nameg  make-stream CR"     h     ]4 5C           g  stream
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	]
		_		
	_	 		  g  nameg  
stream-carg  documentationf  BReturns the first element in STREAM.  This is equivalent to `car'. CR"    h      ]4 5C           g  stream
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	a
		c		
	c	 		  g  nameg  
stream-cdrg  documentationf  GReturns the first tail of STREAM. Equivalent to `(force (cdr STREAM))'. CR"       h     ]4 5C     {      g  stream
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	e
		i			
	i	 		  g  nameg  stream-null?g  documentationf  ¨Returns `#t' if STREAM is the end-of-stream marker; otherwise
returns `#f'.  This is equivalent to `null?', but should be used
whenever testing for the end of a stream. CR     h      ] C          g  l
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	o	 		   C     h   =  ] 6      5      g  l
		
  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	k
	
	n	 		
  g  nameg  list->streamg  documentationf  oReturns a newly allocated stream whose elements are the elements of
LIST.  Equivalent to `(apply stream LIST)'. CR# h    Î   ]	 L $  CL £ C    Æ       g  i
		 g  t		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	u			v			v			w			w	"		w	 		   C       h    Ú   ]	4 5 O 
6    Ò       g  v
		 g  len		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	r
		t			t			s	 		  g  nameg  vector->stream CR hH   U  ]"  /45$  D4545"ÿÿÑ 
"ÿÿÃ     M      g  stream
		C g  s		5 g  acc			5 g  len			5  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
	z
		{			|			|			}			~			~	#	(	~		+	~	7	5	~		5	{		8	{		C	{	 		C  g  nameg  stream->reversed-list&length CR  h   ð   ]4 >  G C è       g  stream
		 g  l		 g  len			  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 
	 		 	 		  g  nameg  stream->reversed-list CR$     h       ]4 >  G 45D ø       g  stream
		 g  l		 g  len			  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 
	 		 		 		 	 		  g  nameg  stream->list&length CR$     h   h  ]4 56   `      g  stream
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 
	 		 	 		  g  nameg  stream->listg  documentationf  Returns a newly allocated list whose elements are the elements of STREAM.
If STREAM has infinite length this procedure will not terminate. CR%     hX   y  ])4 >  G 45"  %(  "  &¤"ÿÿÛ
"ÿÿÐC    q      g  stream
		T g  l		T g  len			T g  v			T g  i		!	F g  l		!	F  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 
	 		 		 		 		! 		' 			2 		7 	*	8 		; 		> 		F 		F 	 		T  g  nameg  stream->vector CR&'    h(     - 1 3 (  
 6 6        g  f
			& g  init			& g  stream				& g  rest				&  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 
	 		 		$ 		& 	 			&	
	  g  nameg  stream-fold CR&  h0     ]45$  C 4 455456           g  f
		+ g  r		+ g  stream			+  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
  
	 ¡		 ¡		 £		 £		" £		# £	3	+ £	 			+	  g  nameg  stream-fold-one C&R('      h    í   ] (  L C 4L  5C     å       g  cars
		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 ©	!	 «	#	 ¬	'	 ­	-	 ®	-	 ®	4	 ®	-	 ­	' 			  g  nameg  recur C)  hH   =  ]45$  C 4 O Q 4455?456 5      g  f
		G g  r		G g  streams			G g  recur		"	:  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 ¥
	 ¦		 ¦		 ©		" ©	!	- ª	.	7 ©	!	< ©		= ¯		G ¨	 		G	  g  nameg  stream-fold-many C'R*+        h(     - 1 3 (   6 6      ù       g  f
			" g  stream			" g  rest				"  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 ±
	 ²		 ³		  ´		" ´	 			"	
	  g  nameg  stream-for-each CR*        h8     ]45$  C4 45>  "  G   456           g  f
		3 g  stream		3  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 ¶
	 ·		 ·		 ¹		 ¹		 ¹		+ º		3 º	 			3	  g  nameg  stream-for-each-one C*R(,)+        h@     ]45$  C4 45>  "  G   456           g  f
		; g  streams		;  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 ¼
	 ½		 ½		 ¿		 ¿		$ ¿		1 À	 	; À	 			;	  g  nameg  stream-for-each-many C+R   h0   æ   ]	4 5$  C4L 4 554 5C       Þ       g  s
		) g  t			)  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 Ç		 È			 È		 É		 É	"	  É		! É	2	( É	 			)   C()    h0   ì   ]	4 5$  C4L 4 5?4 5C ä       g  streams
		/ g  t		/  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 Ë		 Ì		 Ì		 Í		 Í	(	$ Í		% Î		. Í	 			/   C h0      - 1 3 (   O 6 O 6          g  f
			, g  stream			, g  rest				,  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 Â
	 Æ		 Ç		* Ï		, Ë	 			,	
	  g  nameg  
stream-mapg  documentationf  Returns a newly allocated stream, each element being the result of
invoking F with the corresponding elements of the STREAMs
as its arguments. CR-    h(   ã   ]4L  545$  C C     Û       g  p
		# g  o			# g  t			#  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 Ò		 Ó			 Ó		 Ô		 Ô		" Õ	 		#   C  h   Ï   ]O  6 Ç       g  port
		 g  read		  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm
 Ñ
	 Ò	 			  g  nameg  port->stream CRC    5      g  m
		(  g  filenamef  U/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/streams.scm		
=	U
~	]
Ã	a
n	e
	k

	r
@	z
[ 
 
# 
 
Z 
¹  
x ¥
¼ ±
 ¶
 ¼
Ó Â
!â Ñ
 	!ä
   C6 