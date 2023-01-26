;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "matrix.lisp")

(remove-tests :all)

(setq *print-failures* t)
(setq *print-errors* t)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; matrix test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(setq m1 '(
    (1 1 1 1 1)
    (1 1 7 1 1)
    (1 4 0 6 1)
    (1 1 5 1 1)
    (1 1 1 1 1) ))
(setq m1t '(
    (1 1 1 1 1)
    (1 1 4 1 1)
    (1 7 0 5 1)
    (1 1 6 1 1)
    (1 1 1 1 1) ))
(setq m2 '(
    (1  7  0  9  1)
    (6  1  8  1 15)
    (0  5  1 14  0)
    (4  1 11  1 13)
    (1 10  1 12  1) ))
(setq m2t '(
    (1  6  0  4  1)
    (7  1  5  1 10)
    (0  8  1 11  1)
    (9  1 14  1 12)
    (1 15  0 13  1) ))
(setq m3 '(
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 5 6 0 0 0 0 0 0 0 0 0 0)
    (4 4 4 4 4 4 4 4 4 4 1 1 7 7 7 7 7 7 7 7 7 7)
    (9 9 9 9 9 9 9 9 9 9 1 1 8 8 8 8 8 8 8 8 8 8)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0)
    (1 1 1 1 1 1 1 1 1 1 8 9 0 0 0 0 0 0 0 0 0 0) ))
(setq m3t '(
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (0 0 0 0 0 0 0 0 0 0 4 9 1 1 1 1 1 1 1 1 1 1)
    (5 5 5 5 5 5 5 5 5 5 1 1 8 8 8 8 8 8 8 8 8 8)
    (6 6 6 6 6 6 6 6 6 6 1 1 9 9 9 9 9 9 9 9 9 9)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0)
    (0 0 0 0 0 0 0 0 0 0 7 8 0 0 0 0 0 0 0 0 0 0) ))
(setq m4 '(
    (0 1 2 3)
    (1 2 3 0)
    (2 3 0 1)
    (2 4 6 0)
    (3 0 1 2) ))
(setq m4t '(
    (0 1 2 2 3)
    (1 2 3 4 0)
    (2 3 0 6 1)
    (3 0 1 0 2) ))
(setq m5 '(
    (1 2 3 4 5 6 7 8 9 1 2 3 4 5 6 7 8 9) ))
(setq m5t '(
    (1)
    (2)
    (3)
    (4)
    (5)
    (6)
    (7)
    (8)
    (9)
    (1)
    (2)
    (3)
    (4)
    (5)
    (6)
    (7)
    (8)
    (9) ))

(define-test test-are-adjacent-01
    (assert-equal NIL (are-adjacent '()  12 12)))
(define-test test-are-adjacent-02
    (assert-equal NIL (are-adjacent '(10) 10 10)))
(define-test test-are-adjacent-03
    (assert-equal NIL (are-adjacent '(9) 10 10)))
(define-test test-are-adjacent-04
    (assert-equal T (are-adjacent '(9 9) 9 9)))
(define-test test-are-adjacent-05
    (assert-equal T (are-adjacent '(30 50 70 90 10 20) 50 70)))
(define-test test-are-adjacent-06
    (assert-equal T (are-adjacent '(55 33 11 22 44 66 88 55 99) 33 11)))
(define-test test-are-adjacent-07
    (assert-equal NIL (are-adjacent '(11 33 55 22 44 66 88 99 33 55 44 99 22 11 11 88 22 99 88 33 22 33 77 77 88) 11 55)))
(define-test test-are-adjacent-08
    (assert-equal T (are-adjacent '(11 44 33 22 55 66 77 88 99 11 44 66 55 77 22 88 11 88 99 11 44 22 77 55 66) 44 33)))

(define-test test-transpose-01 (assert-equal '() (transpose '())))
(define-test test-transpose-02 (assert-equal '((42)) (transpose '((42)))))
(define-test test-transpose-03 (assert-equal m1t (transpose m1)))
(define-test test-transpose-04 (assert-equal m2t (transpose m2)))
(define-test test-transpose-05 (assert-equal m3t (transpose m3)))
(define-test test-transpose-06 (assert-equal m4t (transpose m4)))
(define-test test-transpose-07 (assert-equal m5t (transpose m5)))

(define-test test-are-neighbors-01 (assert-equal NIL (are-neighbors '() 0 0)))
(define-test test-are-neighbors-02 (assert-equal NIL (are-neighbors '((0)) 0 0)))
(define-test test-are-neighbors-03 (assert-equal T (are-neighbors m1 0 4)))
(define-test test-are-neighbors-04 (assert-equal T (are-neighbors m1 0 5)))
(define-test test-are-neighbors-05 (assert-equal T (are-neighbors m1 0 6)))
(define-test test-are-neighbors-06 (assert-equal T (are-neighbors m1 0 7)))
(define-test test-are-neighbors-07 (assert-equal T (are-neighbors m1 1 1)))
(define-test test-are-neighbors-08 (assert-equal NIL (are-neighbors m1 5 7)))
(define-test test-are-neighbors-09 (assert-equal T (are-neighbors m2 0 6)))
(define-test test-are-neighbors-10 (assert-equal NIL (are-neighbors m2 15 15)))
(define-test test-are-neighbors-11 (assert-equal T (are-neighbors m3 0 6)))
(define-test test-are-neighbors-12 (assert-equal NIL (are-neighbors m3 4 7)))
(define-test test-are-neighbors-13 (assert-equal NIL (are-neighbors m3 5 8)))
(define-test test-are-neighbors-14 (assert-equal NIL (are-neighbors m5 0 0)))
(define-test test-are-neighbors-15 (assert-equal T (are-neighbors m5 1 9)))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
