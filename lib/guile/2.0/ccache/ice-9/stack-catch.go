GOOF----LE-4-2.0�      ]  4       h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  ice-9�	g  stack-catch�	 �		g  filenameS�	
f  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/stack-catch.scm�	g  importsS�	g  
save-stack�	 �	 �	 �	g  exportsS�	 �	g  set-current-module�	 �	 �	g  catch�	g  throw�C 5  hP  �   ]4	
5 4 >  "  G    h(   �   - 1 3 4	>  "  G   @     �       g  key
			# g  args			#  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/stack-catch.scm�
	+		��	
	-	��	#	.	�� 			#
   C    h     ] 6  �      g  key
		 g  thunk		 g  handler			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/stack-catch.scm�
	
��		(	�� 			  g  nameg  stack-catch�g  documentationf Like @code{catch}, invoke @var{thunk} in the dynamic context of
@var{handler} for exceptions matching @var{key}, but also save the
current stack state in the @var{the-last-stack} fluid, for the purpose
of debugging or re-throwing of an error.  If thunk throws to the
symbol @var{key}, then @var{handler} is invoked this way:

@example
 (handler key args ...)
@end example

@var{key} is a symbol or #t.

@var{thunk} takes no arguments.  If @var{thunk} returns normally, that
is the return value of @code{catch}.

Handler is invoked outside the scope of its own @code{catch}.  If
@var{handler} again throws to the same key, a new handler from further
up the call chain is invoked.

If the key is @code{#t}, then a throw to @emph{any} symbol will match
this call to @code{catch}.� CRC�       g  m
		,  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/stack-catch.scm�		
��N	
�� 	P
   C6 