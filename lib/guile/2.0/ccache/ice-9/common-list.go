GOOF----LE-4-2.0è8      ] 2 4    hy      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  ice-9¤	g  common-list¤	 ¤		g  filenameS¤	
f  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm¤	g  exportsS¤	g  adjoin¤	g  union¤	g  intersection¤	g  set-difference¤	g  reduce-init¤	g  reduce¤	g  some¤	g  every¤	g  notany¤	g  notevery¤	g  count-if¤	g  find-if¤	g  	member-if¤	g  	remove-if¤	g  remove-if-not¤	g  
delete-if!¤	g  delete-if-not!¤	g  butlast¤	g  and?¤	g  or?¤	 g  has-duplicates?¤	!g  pick¤	"g  pick-mappings¤	#g  uniq¤	$ !"# ¤	%g  set-current-module¤	&% ¤	'% ¤	(g  memq¤	)g  reverse!¤	*g  memv¤	+g  map¤	,g  car¤	-g  cdr¤	.g  length¤	/g  error¤	0f  negative argument to butlast¤	1g  member¤C 5  h¨3  n  ]4	
$5 4' >  "  G   (  h   2  ]4 5$  C C*      g  e
		 g  l		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	M
		O			O			O	 			  g  nameg  adjoing  documentationf  GReturn list L, possibly with element E added if it is not already in L. CR   h(   {  ] (  C(   C 4 56   s      g  l1
		% g  l2		%  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	Q
		U			W			W			W	&	#	W		%	W	 		%	  g  nameg  uniong  documentationf  wReturn a new list that is the union of L1 and L2.
Elements that occur in both lists occur only once in
the result list. CR)*  hX      ](  C"  8(  645$  "ÿÿÕ"ÿÿÈ "ÿÿ½      ø      g  l1
		R g  l2		R g  l1			G g  result			G  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	Y
		\			]			^			^			_		!	_		%	_		)	^		,	_	(	/	_	7	2	_	1	:	_	"	=	`		G	`		G	]		J	]	!	R	]	 		R	  g  nameg  intersectiong  documentationf  yReturn a new list that is the intersection of L1 and L2.
Only elements that occur in both lists occur in the result list. CR)*     hP   ¶  ]"  8(  645$  "ÿÿÙ"ÿÿÈ "ÿÿ½       ®      g  l1
		I g  l2		I g  l1			> g  result			>  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	b
		d			e			e			f			f			f		 	e		#	f	$	-	f		0	g		3	g	%	6	g		>	g		>	d		A	d		I	d	 		I	  g  nameg  set-differenceg  documentationf  5Return elements from list L1 that are not in list L2. CR h    _  ](  C 4 56  W      g  p
		 g  init		 g  l			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	i
		k			m			m			m			m	&		m	 			  g  nameg  reduce-initg  documentationf  ESame as `reduce' except it implicitly inserts INIT at the start of L. CR        h(   D  ](  C(  C 6      <      g  p
		" g  l		"  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	o
		u			v			u			v			w		 	w	%	"	w	 			"	  g  nameg  reduceg  documentationf 9Combine all the elements of sequence L using a binary operation P.
The combination is left-associative.  For example, using +, one can
add up all the elements.  `reduce' allows you to apply a function which
accepts only two arguments to more than 2 objects.  Functional
programmers usually refer to this as foldl. CR+,-       h   h  - 1 3 (  0"  $(  C4 5$  C"ÿÿÜ"ÿÿÔ"  8(  C4 45?$  C45"ÿÿÈ"ÿÿ¼ `      g  pred
		  g  l		  g  rest			  g  l			7 g  t		#	7 g  l		C	{ g  rest		C	{ g  t		\	{  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
	y
	 		 			 		 		! 		# 		# 		1 	)	7 	#	7 			C 		I 		L 		Q 	%	R 	-	\ 		\ 		j 		k 	'	{ 		{ 	 		 	
	  g  nameg  someg  documentationf pPRED is a boolean function of as many arguments as there are list
arguments to `some', i.e., L plus any optional arguments.  PRED is
applied to successive elements of the list arguments in order.  As soon
as one of these applications returns a true value, return that value.
If no application returns a true value, return #f.
All the lists should have the same length. CR+,-   h     - 1 3 (  +"  (  C4 5$  	"ÿÿãC"ÿÿÙ"  3(  C4 45?$  45"ÿÿÏC"ÿÿÁ         g  pred
			} g  l			} g  rest				} g  l			2 g  l		>	q g  rest		>	q  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 
	 		 			 		 		 		! 		# 		' 		* 	)	0 	#	2 			> 		D 		E 		G 		L 	%	M 	-	W 		[ 		^ 		_ 	'	o 		q 	 			}	
	  g  nameg  everyg  documentationf  Return #t iff every application of PRED to L, etc., returns #t.
Analogous to `some' except it returns #t if every application of
PRED is #t and #f otherwise. CR        h     - 1 3 4 ?C          g  pred
			 g  ls			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 
	
 		 	 			
  g  nameg  notanyg  documentationf  ®Return #t iff every application of PRED to L, etc., returns #f.
Analogous to some but returns #t if no application of PRED returns a
true value or #f as soon as any one does. CR      h     - 1 3 4 ?C          g  pred
			 g  ls			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 
	
  		  	 			
  g  nameg  noteveryg  documentationf  ªReturn #t iff there is an application of PRED to L, etc., that returns #f.
Analogous to some but returns #t as soon as an application of PRED returns #f,
or #f otherwise. CR    hH   ¬  ]"  0(  C4 5$  "ÿÿÝ"ÿÿÐ
"ÿÿÅ       ¤      g  pred
		A g  l		A g  n			6 g  l			6  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ¢
	 ¤		 ¥		 ¦		 ¦		 ¦		 ¥		 ¦	 	! ¦	(	) ¦		. §		6 §		6 ¤	 		A	  g  nameg  count-ifg  documentationf  IReturn the number of elements in L for which (PRED element) returns true. CR   h(     ](  C4 5$  C 6           g  pred
		# g  l		#  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ©
	 ¬		 ­			 ­		 ­			 ¬		 ­		! ®		# ®	 
		#	  g  nameg  find-ifg  documentationf  |Search for the first element in L for which (PRED element) returns true.
If found, return that element, otherwise return #f. CR   h(   R  ](  C4 5$  C 6      J      g  pred
		" g  l		"  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 °
	 ²		 ³			 ³		 ³			 ²		  ´		" ´	 			"	  g  nameg  	member-ifg  documentationf  9Return the first sublist of L for whose car PRED is true. CR)     hH   ä  ]"  6(  64 5$  "ÿÿÛ"ÿÿÊ"ÿÿ¿ Ü      g  pred
		G g  l		G g  l			< g  result			<  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ¶
	 ¹		 º		 º		 »		 »		 »		 º		! »	 	+ »		. ¼		1 ¼	$	4 ¼		< ¼		< ¹		? ¹		G ¹	 		G	  g  nameg  	remove-ifg  documentationf  WRemove all elements from L where (PRED element) is true.
Return everything that's left. CR)   hH   æ  ]"  6(  64 5$  "ÿÿ×"ÿÿÊ"ÿÿ¿ Þ      g  pred
		G g  l		G g  l			< g  result			<  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ¾
	 Á		 Â		 Â		 Ã		 Ã		 Ã		 Â		! Ä		$ Ä	$	' Ä		/ Ä		2 Ã	&	< Ã	 	< Á		? Á		G Á	 		G	  g  nameg  remove-if-notg  documentationf  URemove all elements from L where (PRED element) is #f.
Return everything that's left. CR    h0     ] (  C4L  5$  	  "ÿÿã 4L 5 C         g  l
		-  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 È		 É			 É		 Ê		 Ê		 Ê		 É		 Ê	%	 Ê		" Ì		' Ì	"	) Ì		* Ì	 		-  g  nameg  	delete-if C      h     ]
O  Q 6      g  pred
		 g  l		 g  	delete-if			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 Æ
	 È	 			  g  nameg  
delete-if!g  documentationf  #Destructive version of `remove-if'. CR    h0   #  ] (  C4L  5$   4L 5 C  "ÿÿÕ         g  l
		-  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 Ñ		 Ò			 Ò		 Ó		 Ó		 Ó		 Ò		 Õ		 Õ	&	  Õ		! Õ		' Ó	/	- Ó	  		-  g  nameg  delete-if-not C  h   *  ]
O  Q 6"      g  pred
		 g  l		 g  delete-if-not			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 Ï
	 Ñ	 			  g  nameg  delete-if-not!g  documentationf  'Destructive version of `remove-if-not'. CRh(     ] (   C
$   4L  5CC         g  lst
		% g  n		%  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 Û		 Ü		 Ý		 Ü		 Þ		 Þ	(	 Þ	,	 Þ	6	! Þ	(	" Þ		$ ß	 		%	  g  nameg  bl C./0        hH     ]HO Q 4 5K 
$  45"  J6      x      g  lst
		B g  n		B g  l			B g  bl			B g  l			%  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 Ø
	 Ú		 Ú		 Ú		, à		1 à		2 á		6 á		: á		B à	 		B	  g  nameg  butlastg  documentationf  *Return all but the last N elements of LST. CR       h(     -  1  3  (  C $   @C             g  args
			!  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ä
	 æ		 ç			 æ		 ç	 	 ç	 			!


  g  nameg  and?g  documentationf  #Return #t iff all of ARGS are true. CR       h(     -  1  3  (  C $  C @             g  args
			!  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ê
	 ì		 í			 ì		 î		! î	 			!


  g  nameg  or?g  documentationf  "Return #t iff any of ARGS is true. CR1        h(   K  ] (  C4  5$  C 6      C      g  lst
		"  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ð
	 ò		 ó			 ó		 ó		 ó			 ò		  ô		" ô	 
		"  g  nameg  has-duplicates?g  documentationf  3Return #t iff 2 members of LST are equal?, else #f. C RhH   Ñ  ]"  3(  C4 5$  "ÿÿÚ"ÿÿÍ"ÿÿÂ    É      g  p
		D g  l		D g  s			9 g  l			9  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 ö
	 ù		 û		 ý		 ý			 ý		 û		 ý	$	! ý		$ ý	/	, ý		1 þ	 	9 þ		9 ù		: ù		D ù	 		D	  g  nameg  pickg  documentationf  ZApply P to each element of L, returning a list of elts
for which P returns a non-#f value. C!R  hH   Ö  ]"  6(  C4 5$  "ÿÿ×"ÿÿÊ"ÿÿ¿ Î      g  p
		G g  l		G g  s			< g  l			< g  t			<  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm
 
														$	0	'	A	/	*	4	 	<		<		=		G	 		G	  g  nameg  pick-mappingsg  documentationf  PApply P to each element of L, returning a list of the
non-#f return values of P. C"R)(       hH      ]"  1(  645$  "  "ÿÿÏ "ÿÿÄ            g  l
		B g  acc		7 g  l			7  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm


													 		)		,		/		7		7		8		B	 		B  g  nameg  uniqg  documentationf  @Return a list containing elements of L, with duplicates removed. C#RC   f      g  m
		(  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/common-list.scm		4
	M
>	Q
«	Y
É	b
	Z	i
ß	o
ë	y

 
Ô 
 
 ¢
w ©
 °
G ¶
  ¾
#1 Æ
%å Ï
) Ø
*k ä
+Á ê
-N ð
/t ö
1¡ 
3£

 	3¥
   C6 