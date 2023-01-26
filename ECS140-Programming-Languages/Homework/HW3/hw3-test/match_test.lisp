;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "match.lisp")

(remove-tests :all)

(setq *print-failures* t)
(setq *print-errors* t)
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; match test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
(define-test test-match-01
    (assert-equal T (match '() '())))
(define-test test-match-02
    (assert-equal NIL (match '()  '(AAA))))
(define-test test-match-03
    (assert-equal NIL (match '(AAA) '())))
(define-test test-match-04
    (assert-equal NIL (match '(AAA) '(BBB))))
(define-test test-match-05
    (assert-equal NIL (match '(XX YY ZZ) '(XX YY WW))))
(define-test test-match-06
    (assert-equal NIL (match '(?) '())))
(define-test test-match-07
    (assert-equal T (match '(?) '(XX))))
(define-test test-match-08
    (assert-equal NIL (match '(?) '(NIL YY))))
(define-test test-match-09
    (assert-equal T (match '(? XX) '(YY XX))))
(define-test test-match-10
    (assert-equal NIL (match '(!) '())))
(define-test test-match-11
    (assert-equal T (match '(!) '(XX YY ZZ))))
(define-test test-match-12
    (assert-equal NIL (match '(! WW) '(XX YY ZZ))))
(define-test test-match-13
    (assert-equal NIL (match '(! Z ? W Y) '(X Y Z W Y))))
(define-test test-match-14
    (assert-equal T (match '(! Z ? Z Y Z Y ! ? !) '(Z Z Y Z Y Z Y Z Y Z Y))))
(define-test test-match-15
    (assert-equal T (match '(? !) '(X X Y Z V))))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
