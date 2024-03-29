GOOF----LE-4-2.0T      ] = 4 h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  texinfo�	g  indexing�	 �		g  filenameS�	
f  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/indexing.scm�	g  importsS�	g  sxml�	g  simple�	 �	 �	g  srfi�	g  srfi-13�	 �	 �	 �	g  exportsS�	g  stexi-extract-index�	 �	g  set-current-module�	 �	 �	g  deftp�	g  defcv�	g  defivar�	g  deftypeivar�	g  defop�	 g  	deftypeop�	!g  	defmethod�	"g  deftypemethod�	#g  defopt�	$g  defvr�	%g  defvar�	&g  	deftypevr�	'g  
deftypevar�	(g  deffn�	)g  	deftypefn�	*g  defspec�	+g  defmac�	,g  defun�	-g  
deftypefun�	. !"#$%&'()*+,- �	/g  defines�	0g  cindex�	1g  findex�	2g  vindex�	3g  kindex�	4g  pindex�	5g  tindex�	6012345 �	7g  indices�	8g  anchor�	9g  memq�	:g  assq�	;g  name�	<g  sxml->string�C 5    hP  �   ]4	
5 4 >  "  G   ./R67R89/:;7<       h�   �  ] (  C ��$  �"   �4L  �5 "��� ��$  � ���$  � ��&  {4 ���5$  / ��4 �����5��4 ����5���� "��|4 ���5$  ( ��4 ��54 ����5���� "��D"��U"��Q"��M"��I � "��'           g  in
	 � g  entries	 �  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/indexing.scm�
	5	��		6	��		9	��		9	��		6	��		F	��		F	��	 	F	��	$	F	��	,	F	��	,	:	��	/	;	��	0	;	��	4	:	��	7	;	%��	9	;	��	=	;	��	@	<	��	C	<	��	G	;	��	H	<	%��	M	<	+��	S	<	%��	W	;	��	Z	=	��	\	=	%��	`	=	+��	c	=	6��	g	=	1��	i	=	%��	j	=	��	l	>	%��	p	>	+��	s	>	1��	x	>	%��	y	>	��	{	=	�� �	=	�� �	A	%�� �	A	+�� �	A	%�� �	@	�� �	B	�� �	B	�� �	B	-�� �	B	�� �	C	%�� �	C	+�� �	C	1�� �	C	%�� �	C	�� �	B	�� �	B	�� �	H	�� �	H	�� 8	 �	  g  nameg  loop� C      h   ^  ]O Q  6  V      g  tree
		 g  manual-name		 g  kind			 g  loop		
	  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/indexing.scm�
	*
��	
	5	��		5	 ��		5	�� 			  g  nameg  stexi-extract-index�g  documentationf 6Given an stexi tree @var{tree}, index all of the entries of type
@var{kind}. @var{kind} can be one of the predefined texinfo indices
(@code{concept}, @code{variable}, @code{function}, @code{key},
@code{program}, @code{type}) or one of the special symbols @code{auto} 
or @code{all}. @code{auto} will scan the stext for a @code{(printindex)}
statement, and @code{all} will generate an index from all entries,
regardless of type.

The returned index is a list of pairs, the @sc{car} of which is the
entry (a string) and the @sc{cdr} of which is a node name (a string).� CRC     �       g  m
		,  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/indexing.scm�		
��	.	#	��	1	"
��	3	(	��	6	'
��I	*
�� 	K
   C6 