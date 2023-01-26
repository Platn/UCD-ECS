; A list is a 1-D array of numbers.
; A matrix is a 2-D array of numbers, stored in row-major order.

(defun occurs (elem li)
    (cond  ((null li) nil )
           ((equal elem (car li)) t  )
           ( (occurs elem (cdr li) ) )
    )
)

; AreAdjacent returns true iff a and b are adjacent in lst.
(defun are-adjacent (lst a b)
    (cond ((null lst) nil) 
          ((null (cdr lst)) nil)
          ((eql a (car lst)) (if (eql (cadr lst) b) T nil))
          ((eql b (car lst)) (if (eql (cadr lst) a) T nil))
          (t (are-adjacent (cdr lst) a b))
    )
)

; Transpose returns the transpose of the 2D matrix mat.
(defun transpose (matrix)
    (cond ((null matrix) nil) 
           (t (apply #'mapcar #'list matrix))
    )
)

; AreNeighbors returns true iff a and b are neighbors in the 2D
; matrix mat.
(defun are-neighbors (matrix a b)
    (cond ((null matrix) nil)
          ((occurs T (mapcar #'(lambda (l) (are-adjacent l a b)) matrix)) T)
          ((occurs T (mapcar #'(lambda (l) (are-adjacent l a b)) (transpose matrix))) T)
          (t nil)
    )
)